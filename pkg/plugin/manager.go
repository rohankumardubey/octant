/*
Copyright (c) 2019 the Octant contributors. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package plugin

//go:generate mockgen -destination=./fake/mock_manager.go -package=fake github.com/vmware-tanzu/octant/pkg/plugin ManagerInterface
//go:generate mockgen -destination=./fake/mock_module_registrar.go -package=fake github.com/vmware-tanzu/octant/pkg/plugin ModuleRegistrar
//go:generate mockgen -destination=./fake/mock_action_registrar.go -package=fake github.com/vmware-tanzu/octant/pkg/plugin ActionRegistrar

import (
	"context"
	"fmt"
	"path/filepath"
	"sort"
	"sync"
	"time"

	"github.com/vmware-tanzu/octant/pkg/event"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"

	"github.com/vmware-tanzu/octant/pkg/plugin/javascript"

	"github.com/fsnotify/fsnotify"
	"github.com/hashicorp/go-plugin"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/vmware-tanzu/octant/internal/log"
	"github.com/vmware-tanzu/octant/internal/module"
	"github.com/vmware-tanzu/octant/internal/portforward"
	"github.com/vmware-tanzu/octant/pkg/action"
	"github.com/vmware-tanzu/octant/pkg/plugin/api"
	"github.com/vmware-tanzu/octant/pkg/view/component"
)

// ClientFactory is a factory for creating clients.
type ClientFactory interface {
	// Init initializes a client.
	Init(ctx context.Context, cmd string) Client
}

// DefaultClientFactory is the default client factory
type DefaultClientFactory struct{}

var _ ClientFactory = (*DefaultClientFactory)(nil)

// NewDefaultClientFactory creates an instance of DefaultClientFactory.
func NewDefaultClientFactory() *DefaultClientFactory {
	return &DefaultClientFactory{}
}

// Init creates a new client.
func (f *DefaultClientFactory) Init(ctx context.Context, cmd string) Client {
	loggerAdapter := &zapAdapter{
		dashLogger: log.From(ctx),
	}

	c := pluginCmd(cmd)

	return plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: Handshake,
		Plugins:         pluginMap,
		Cmd:             c,
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolGRPC,
		},
		Logger: loggerAdapter,
	})
}

// Client is an interface that describes a plugin client.
type Client interface {
	Client() (plugin.ClientProtocol, error)
	Kill()
}

// ManagerStore is the data store for Manager.
type ManagerStore interface {
	Store(name string, client Client, metadata *Metadata, cmd string) error
	StoreJS(name string, jspc JSPlugin) error
	Get(name string) (Client, bool)
	GetJS(name string) (JSPlugin, bool)
	Remove(name string)
	RemoveJS(name string)
	NamesJS() []string
	GetMetadata(name string) (*Metadata, error)
	GetModuleService(name string) (ModuleService, error)
	GetService(name string) (Service, error)
	GetCommand(name string) (string, error)
	Clients() map[string]Client
	ClientNames() []string
}

// DefaultStore is the default implement of ManagerStore.
type DefaultStore struct {
	clients  map[string]Client
	metadata map[string]Metadata
	commands map[string]string

	jsPlugins sync.Map
}

var _ ManagerStore = (*DefaultStore)(nil)

// NewDefaultStore creates an instance of DefaultStore.
func NewDefaultStore() *DefaultStore {
	return &DefaultStore{
		clients:  make(map[string]Client),
		metadata: make(map[string]Metadata),
		commands: make(map[string]string),
	}
}

func (s *DefaultStore) NamesJS() []string {
	var names []string
	s.jsPlugins.Range(func(key interface{}, value interface{}) bool {
		name, ok := key.(string)
		if !ok {
			return false
		}
		names = append(names, name)
		return true
	})
	return names
}

func (s *DefaultStore) StoreJS(name string, plugin JSPlugin) error {
	s.jsPlugins.Store(name, plugin)
	return nil
}

func (s *DefaultStore) Get(name string) (Client, bool) {
	client, ok := s.clients[name]
	return client, ok
}

func (s *DefaultStore) GetJS(name string) (JSPlugin, bool) {
	voidStar, ok := s.jsPlugins.Load(name)
	if !ok {
		return nil, false
	}
	jspc, ok := voidStar.(JSPlugin)
	return jspc, ok
}

func (s *DefaultStore) Remove(name string) {
	delete(s.commands, name)
	delete(s.clients, name)
	delete(s.metadata, name)
}

func (s *DefaultStore) RemoveJS(name string) {
	s.jsPlugins.Delete(name)
}

// Store stores information for a plugin.
func (s *DefaultStore) Store(name string, client Client, metadata *Metadata, cmd string) error {
	if metadata == nil {
		return errors.New("metadata is nil")
	}

	s.clients[name] = client
	s.metadata[name] = *metadata
	s.commands[name] = cmd

	return nil
}

// GetModuleService gets the moduleService for a plugin
func (s *DefaultStore) GetModuleService(name string) (ModuleService, error) {
	client, ok := s.clients[name]
	if !ok {
		return nil, errors.Errorf("plugin %q doesn't have a client", name)
	}

	rpcClient, err := client.Client()
	if err != nil {
		return nil, err
	}

	raw, err := rpcClient.Dispense("plugin")
	if err != nil {
		return nil, errors.Wrapf(err, "dispensing plugin for %q", name)
	}

	moduleService, ok := raw.(ModuleService)
	if !ok {
		return nil, errors.Errorf("unknown type for plugin %q: %T", name, raw)
	}

	return moduleService, nil
}

// GetService gets the service for a plugin.
func (s *DefaultStore) GetService(name string) (Service, error) {
	moduleService, err := s.GetModuleService(name)
	if err != nil {
		return nil, err
	}
	service, ok := moduleService.(Service)
	if !ok {
		return nil, fmt.Errorf("failed to GetService, unable to cast ModuleService to Service")
	}
	return service, nil
}

// GetMetadata gets the metadata for a plugin.
func (s *DefaultStore) GetMetadata(name string) (*Metadata, error) {
	metadata, ok := s.metadata[name]
	if !ok {
		return nil, errors.Errorf("plugin %q doesn't have metadata", name)
	}

	return &metadata, nil
}

// GetCommand gets the command for a plugin.
func (s *DefaultStore) GetCommand(name string) (string, error) {
	cmd, ok := s.commands[name]
	if !ok {
		return "", errors.Errorf("plugin %q doesn't have command", name)
	}

	return cmd, nil
}

// Clients returns all the clients in the store.
func (s *DefaultStore) Clients() map[string]Client {
	return s.clients
}

// ClientNames returns the client names in the store.
func (s *DefaultStore) ClientNames() []string {
	var list []string
	for name := range s.Clients() {
		list = append(list, name)
	}
	tsNames := s.NamesJS()
	list = append(list, tsNames...)
	return list
}

type PluginConfig struct {
	Cmd  string
	Name string
}

// ClusterClient defines the cluster client plugin manager has access to.
type ClusterClient interface {
	Resource(gk schema.GroupKind) (schema.GroupVersionResource, bool, error)
	DynamicClient() (dynamic.Interface, error)
}

// ManagerInterface is an interface which represent a plugin manager.
type ManagerInterface interface {
	// Print prints an object.
	Print(ctx context.Context, object runtime.Object) (*PrintResponse, error)

	// Tabs retrieves tabs for an object.
	Tabs(ctx context.Context, object runtime.Object) ([]component.Tab, error)

	// Store returns the manager's storage.
	Store() ManagerStore

	// ObjectStatus returns the object status
	ObjectStatus(ctx context.Context, object runtime.Object) (*ObjectStatusResponse, error)

	// SetOctantClient sets the the Octant client.
	SetOctantClient(octantClient javascript.OctantClient)
}

// ModuleRegistrar is a module registrar.
type ModuleRegistrar interface {
	// Register registers a module.
	Register(mod module.Module) error
	// Unregister unregisters a module.
	Unregister(mod module.Module)
}

// ActionRegistrar is an action registrar.
type ActionRegistrar interface {
	// Register registers an action.
	Register(actionPath string, pluginPath string, actionFunc action.DispatcherFunc) error
	// Unregister unregisters an action.
	Unregister(actionPath string, pluginPath string)
}

// ManagerOption is an option for configuring Manager.
type ManagerOption func(*Manager)

// Manager manages plugins
type Manager struct {
	PortForwarder   portforward.PortForwarder
	API             api.API
	ClientFactory   ClientFactory
	ModuleRegistrar ModuleRegistrar
	ActionRegistrar ActionRegistrar
	WSClient        event.WSClientGetter

	Runners Runners

	octantClient javascript.OctantClient
	configs      []PluginConfig
	store        ManagerStore

	lock sync.Mutex
}

var _ ManagerInterface = (*Manager)(nil)

// NewManager creates an instance of Manager.
func NewManager(apiService api.API, moduleRegistrar ModuleRegistrar, actionRegistrar ActionRegistrar, ws event.WSClientGetter, options ...ManagerOption) *Manager {
	m := &Manager{
		store:           NewDefaultStore(),
		ClientFactory:   NewDefaultClientFactory(),
		Runners:         newDefaultRunners(),
		API:             apiService,
		ModuleRegistrar: moduleRegistrar,
		ActionRegistrar: actionRegistrar,
		WSClient:        ws,
	}

	for _, option := range options {
		option(m)
	}

	return m
}

func (m *Manager) SetOctantClient(client javascript.OctantClient) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.octantClient = client
}

// Store returns the store for the manager.
func (m *Manager) Store() ManagerStore {
	return m.store
}

// SetStore sets the store for the manager.
func (m *Manager) SetStore(store ManagerStore) {
	m.store = store
}

// Load loads a plugin.
func (m *Manager) Load(cmd string) (PluginConfig, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	name := filepath.Base(cmd)

	for _, c := range m.configs {
		if name == c.Name {
			return PluginConfig{}, fmt.Errorf("tried to load plugin %q more than once", name)
		}
	}

	c := PluginConfig{
		Name: name,
		Cmd:  cmd,
	}

	m.configs = append(m.configs, c)

	return c, nil
}

func (m *Manager) Unload(ctx context.Context, cmd string) {
	if IsJavaScriptPlugin(cmd) {
		m.unregisterJSPlugin(ctx, cmd)
	} else {
		m.unregisterGoPlugin(ctx, cmd)
	}
}

func (m *Manager) watchPluginFiles(ctx context.Context) {
	logger := log.From(ctx)

	dirs, err := DefaultConfig.PluginDirs(DefaultConfig.Home())
	if err != nil {
		logger.Errorf("unable to get plugin directories: %w", err)
		return
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.Errorf("initializing plugin watcher: %w", err)
		return
	}
	defer func() {
		if err := watcher.Close(); err != nil {
			logger.Errorf("error closing fsnotify watcher: %w", err)
		}
	}()

	watchedDirs := []string{}
	for _, dir := range dirs {
		if err := watcher.Add(dir); err != nil {
			logger.Warnf("Unable to add %s to the plugin watcher. Error: %s\n", dir, err)
		} else {
			watchedDirs = append(watchedDirs, dir)
		}
	}

	if len(watchedDirs) > 0 {
		logger.Infof("Watching for plugin changes %q\n", dirs)
	}

	writeEvents := make(map[string]bool)
	updatePlugin := func(name string) {
		m.Unload(ctx, name)
		if IsJavaScriptPlugin(name) {
			logger.Infof("reloading: JavaScript plugin: %s", name)
			if err := m.registerJSPlugin(ctx, name); err != nil {
				logger.Errorf("reloading: JavaScript plugin: %w", err)
			}
		} else {
			logger.Infof("reloading: Go plugin: %s", name)
			config, err := m.Load(name)
			if err != nil {
				logger.Errorf("reloading: Go plugin: %w", err)
			} else {
				m.start(ctx, config)
			}
		}
	}

	for {
		select {
		case <-ctx.Done():
			logger.Infof("context cancelled shutting down plugin watcher.")
			return
		case ev, ok := <-watcher.Events:
			if !ok {
				logger.Errorf("bad event returned from plugin watcher")
				return
			}
			if ev.Op&(fsnotify.Chmod|fsnotify.Write|fsnotify.Create) == fsnotify.Chmod {
				continue
			}
			if ev.Op&(fsnotify.Remove|fsnotify.Rename) == fsnotify.Remove || ev.Op&fsnotify.Rename == fsnotify.Rename {
				logger.Infof("removing plugin: %s", ev.Name)
				m.Unload(ctx, ev.Name)
			} else if ev.Op&(fsnotify.Write|ev.Op&fsnotify.Rename|ev.Op&fsnotify.Create) == fsnotify.Create {
				writeEvents[ev.Name] = true
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			logger.Errorf("error:", err)
		case <-time.After(1 * time.Second):
			for k := range writeEvents {
				updatePlugin(k)
			}
			writeEvents = make(map[string]bool)
			continue
		}

	}
}

func (m *Manager) unregisterGoPlugin(ctx context.Context, cmd string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	logger := log.From(ctx)

	name := filepath.Base(cmd)

	var match bool
	var matchIdx int

	for i, c := range m.configs {
		if name == c.Name {
			match = true
			matchIdx = i
			break
		}
	}

	if match {
		client, ok := m.store.Get(name)
		if ok {
			client.Kill()
		}
		if metadata, err := m.store.GetMetadata(name); err == nil {
			moduleService, err := m.store.GetModuleService(name)
			if err != nil {
				logger.Errorf("failed unregister service (go): %w", err)
			}
			if err := m.unregisterMetadata(ctx, cmd, metadata, moduleService); err != nil {
				logger.Errorf("failed unregister metadata (go): %w", err)
			}
		}

		m.configs[matchIdx] = m.configs[len(m.configs)-1]
		m.configs = m.configs[:len(m.configs)-1]

		m.store.Remove(name)
	}
}

func (m *Manager) unregisterJSPlugin(ctx context.Context, cmd string) error {
	jsPlugin, ok := m.store.GetJS(cmd)
	if ok {
		m.store.RemoveJS(cmd)
		jsPlugin.Close()
		if err := m.unregisterMetadata(ctx, jsPlugin.PluginPath(), jsPlugin.Metadata(), jsPlugin); err != nil {
			logger := log.From(ctx)
			logger.Errorf("failed unregister metadata (js): %w", err)
		}
	}
	return nil
}

func (m *Manager) unregisterMetadata(ctx context.Context, path string, metadata *Metadata, moduleService ModuleService) error {
	if metadata.Capabilities.IsModule {
		mp, err := NewModuleProxy(metadata.Name, metadata, moduleService)
		if err != nil {
			return fmt.Errorf("unregister: creating module proxy: %w", err)
		}
		m.ModuleRegistrar.Unregister(mp)
	}

	for _, actionName := range metadata.Capabilities.ActionNames {
		actionPath := actionName
		m.ActionRegistrar.Unregister(actionPath, path)
	}
	return nil
}

func (m *Manager) registerJSPlugin(ctx context.Context, pluginPath string) error {
	dashboardClientFactory := javascript.NewModularDashboardClientFactory(javascript.DefaultFunctions(m.octantClient, m.WSClient))

	jsPlugin, err := NewJSPlugin(ctx, pluginPath, dashboardClientFactory)
	if err != nil {
		return err
	}
	if err := m.store.StoreJS(pluginPath, jsPlugin); err != nil {
		return err
	}

	metadata := jsPlugin.Metadata()

	pluginLogger := log.From(ctx).With("plugin-name", pluginPath)
	pluginLogger.With(
		"cmd", pluginPath,
		"metadata", metadata,
	).Infof("registered plugin %q", metadata.Name)

	for _, actionName := range metadata.Capabilities.ActionNames {
		actionPath := actionName
		pluginLogger.With("action-path", actionPath).Infof("registering plugin action")
		err := m.ActionRegistrar.Register(actionPath, pluginPath, func(ctx context.Context, alerter action.Alerter, payload action.Payload) error {
			return jsPlugin.HandleAction(ctx, actionPath, payload)
		})

		if err != nil {
			return fmt.Errorf("configuring plugin action: %w", err)
		}
	}

	if metadata.Capabilities.IsModule {
		pluginLogger.Infof("plugin supports navigation")

		mp, err := NewModuleProxy(metadata.Name, metadata, jsPlugin)
		if err != nil {
			return fmt.Errorf("creating module proxy: %w", err)
		}

		if err := m.ModuleRegistrar.Register(mp); err != nil {
			return fmt.Errorf("register module %s: %w", metadata.Name, err)
		}
	}

	return nil
}

func (m *Manager) startJS(ctx context.Context, pluginPath string) error {
	if err := m.registerJSPlugin(ctx, pluginPath); err != nil {
		return fmt.Errorf("javascript plugin: %w", err)
	}

	return nil
}

// Start starts all plugins.
func (m *Manager) Start(ctx context.Context) error {
	if m.store == nil {
		return errors.New("manager store is nil")
	}

	if m.ClientFactory == nil {
		return errors.New("manager client factory is nil")
	}

	if err := m.API.Start(ctx); err != nil {
		return errors.Wrap(err, "start api service")
	}

	logger := log.From(ctx)
	logger.With("addr", m.API.Addr()).Debugf("starting plugin api service")

	m.lock.Lock()
	defer m.lock.Unlock()

	pluginList, err := AvailablePlugins(DefaultConfig)
	if err != nil {
		return err
	}
	for _, pluginPath := range pluginList {
		if IsJavaScriptPlugin(pluginPath) {
			logger := log.From(ctx)
			logger.Debugf("creating ts plugin client")

			if err := m.startJS(ctx, pluginPath); err != nil {
				logger.Warnf("start JS plugin: %s\n", err)
			}
		}
	}

	for i := range m.configs {
		c := m.configs[i]

		if err := m.start(ctx, c); err != nil {
			logger.Warnf("start plugin: %s\n", err)
		}
	}

	go m.watchPluginFiles(ctx)
	go m.goPluginPingPong(ctx)

	return nil
}

// goPluginPingPong will attempt to restart a Go plugin process if the Ping method for a client returns a non-nil error.
func (m *Manager) goPluginPingPong(ctx context.Context) {
	logger := log.From(ctx)

	timer := time.NewTimer(5 * time.Second)
	running := true

	for running {
		select {
		case <-ctx.Done():
			logger.Infof("shutting down plugin watcher")
			running = false
			return
		case <-timer.C:
			for clientName, client := range m.store.Clients() {
				rpcClient, err := client.Client()
				if err != nil {
					logger.WithErr(err).Errorf("retrieve plugin client for ping")
				}

				if err := rpcClient.Ping(); err != nil {
					logger.With("plugin-name", clientName).Infof("restarting plugin")

					cmd, err := m.store.GetCommand(clientName)
					if err != nil {
						logger.WithErr(err).Errorf("unable to find command for plugin")
						continue
					}

					m.Unload(ctx, cmd)

					c := PluginConfig{
						Name: clientName,
						Cmd:  cmd,
					}

					if err := m.start(ctx, c); err != nil {
						logger.WithErr(err).Errorf("unable to restart plugin")
						continue
					}
				}
			}

			timer.Reset(5 * time.Second)
		}
	}

}

func (m *Manager) start(ctx context.Context, c PluginConfig) error {
	client := m.ClientFactory.Init(ctx, c.Cmd)

	rpcClient, err := client.Client()
	if err != nil {
		return errors.Wrapf(err, "get rpc client for %q", c.Name)
	}

	pluginLogger := log.From(ctx).With("plugin-name", c.Name)

	raw, err := rpcClient.Dispense("plugin")
	if err != nil {
		return errors.Wrapf(err, "dispensing plugin for %q", c.Name)
	}

	service, ok := raw.(Service)
	if !ok {
		return errors.Errorf("unknown type for plugin %q: %T", c.Name, raw)
	}

	metadata, err := service.Register(ctx, m.API.Addr())
	if err != nil {
		return errors.Wrapf(err, "register plugin %q", c.Name)
	}

	if err := m.store.Store(c.Name, client, &metadata, c.Cmd); err != nil {
		return errors.Wrapf(err, "storing plugin")
	}

	for _, actionName := range metadata.Capabilities.ActionNames {
		actionPath := actionName
		pluginLogger.With("action-path", actionPath).Infof("registering plugin action")
		err := m.ActionRegistrar.Register(actionPath, c.Name, func(ctx context.Context, alerter action.Alerter, payload action.Payload) error {
			return service.HandleAction(ctx, actionPath, payload)
		})

		if err != nil {
			return errors.Wrap(err, "configuring plugin action")
		}
	}

	pluginLogger.With(
		"cmd", c.Cmd,
		"metadata", metadata,
	).Infof("registered plugin %q", metadata.Name)

	if metadata.Capabilities.IsModule {
		service, ok := raw.(ModuleService)
		if !ok {
			return errors.Errorf("plugin type %T is a not a module", raw)
		}

		pluginLogger.Infof("plugin supports navigation")

		mp, err := NewModuleProxy(c.Name, &metadata, service)
		if err != nil {
			return errors.Wrap(err, "creating module proxy")
		}

		if err := m.ModuleRegistrar.Register(mp); err != nil {
			return errors.Wrapf(err, "register module %s", metadata.Name)
		}
	}

	return nil
}

// Stop stops all plugins.
func (m *Manager) Stop(ctx context.Context) {
	logger := log.From(ctx)

	m.lock.Lock()
	defer m.lock.Unlock()

	for name, client := range m.store.Clients() {
		logger.With("plugin-name", name).Debugf("stopping plugin")
		client.Kill()
	}
}

// Print prints an object with plugins which are configured to print the objects's
// GVK.
func (m *Manager) Print(ctx context.Context, object runtime.Object) (*PrintResponse, error) {
	if m.Runners == nil {
		return nil, errors.New("runners is nil")
	}

	runner, ch := m.Runners.Print(m.store)
	done := make(chan bool)

	var pr PrintResponse

	go func() {
		for resp := range ch {
			pr.Config = append(pr.Config, resp.Config...)
			pr.Status = append(pr.Status, resp.Status...)
			pr.Items = append(pr.Items, resp.Items...)
		}

		done <- true
	}()

	if err := runner.Run(ctx, object, m.store.ClientNames()); err != nil {
		return nil, fmt.Errorf("print runner failed: %w", err)
	}
	close(ch)

	<-done

	// Attempt to eliminate whitespace before fallback
	sort.Slice(pr.Items, func(i, j int) bool {
		if a, b := pr.Items[i].Width, pr.Items[j].Width; a != b {
			return a < b
		}

		a, _ := component.TitleFromTitleComponent(pr.Items[i].View.GetMetadata().Title)
		b, _ := component.TitleFromTitleComponent(pr.Items[j].View.GetMetadata().Title)

		return a < b
	})

	return &pr, nil
}

// Tabs queries plugins for tabs for an object.
func (m *Manager) Tabs(ctx context.Context, object runtime.Object) ([]component.Tab, error) {
	if m.Runners == nil {
		return nil, errors.New("runners is nil")
	}

	runner, ch := m.Runners.Tab(m.store)
	done := make(chan bool)

	var tabs []component.Tab

	go func() {
		for t := range ch {
			tabs = append(tabs, t...)
		}

		done <- true
	}()

	if err := runner.Run(ctx, object, m.store.ClientNames()); err != nil {
		return nil, err
	}

	close(ch)
	<-done

	sort.Slice(tabs, func(i, j int) bool {
		return tabs[i].Name < tabs[j].Name
	})

	return tabs, nil
}

// ObjectStatus updates the object status of an object configured from a plugin
func (m *Manager) ObjectStatus(ctx context.Context, object runtime.Object) (*ObjectStatusResponse, error) {
	if m.Runners == nil {
		return nil, errors.New("runners is nil")
	}

	runner, ch := m.Runners.ObjectStatus(m.store)
	done := make(chan bool)

	var osr ObjectStatusResponse

	go func() {
		for resp := range ch {
			osr.ObjectStatus = resp.ObjectStatus
		}

		done <- true
	}()

	if err := runner.Run(ctx, object, m.store.ClientNames()); err != nil {
		return nil, err
	}
	close(ch)

	<-done
	return &osr, nil
}
