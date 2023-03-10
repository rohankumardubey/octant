// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware-tanzu/octant/pkg/plugin (interfaces: Runners,ManagerStore,ClientFactory,ModuleService,Service,Broker)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	runtime "k8s.io/apimachinery/pkg/runtime"

	action "github.com/vmware-tanzu/octant/pkg/action"
	navigation "github.com/vmware-tanzu/octant/pkg/navigation"
	plugin "github.com/vmware-tanzu/octant/pkg/plugin"
	component "github.com/vmware-tanzu/octant/pkg/view/component"
)

// MockRunners is a mock of Runners interface.
type MockRunners struct {
	ctrl     *gomock.Controller
	recorder *MockRunnersMockRecorder
}

// MockRunnersMockRecorder is the mock recorder for MockRunners.
type MockRunnersMockRecorder struct {
	mock *MockRunners
}

// NewMockRunners creates a new mock instance.
func NewMockRunners(ctrl *gomock.Controller) *MockRunners {
	mock := &MockRunners{ctrl: ctrl}
	mock.recorder = &MockRunnersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRunners) EXPECT() *MockRunnersMockRecorder {
	return m.recorder
}

// ObjectStatus mocks base method.
func (m *MockRunners) ObjectStatus(arg0 plugin.ManagerStore) (plugin.DefaultRunner, chan plugin.ObjectStatusResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectStatus", arg0)
	ret0, _ := ret[0].(plugin.DefaultRunner)
	ret1, _ := ret[1].(chan plugin.ObjectStatusResponse)
	return ret0, ret1
}

// ObjectStatus indicates an expected call of ObjectStatus.
func (mr *MockRunnersMockRecorder) ObjectStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectStatus", reflect.TypeOf((*MockRunners)(nil).ObjectStatus), arg0)
}

// Print mocks base method.
func (m *MockRunners) Print(arg0 plugin.ManagerStore) (plugin.DefaultRunner, chan plugin.PrintResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Print", arg0)
	ret0, _ := ret[0].(plugin.DefaultRunner)
	ret1, _ := ret[1].(chan plugin.PrintResponse)
	return ret0, ret1
}

// Print indicates an expected call of Print.
func (mr *MockRunnersMockRecorder) Print(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Print", reflect.TypeOf((*MockRunners)(nil).Print), arg0)
}

// Tab mocks base method.
func (m *MockRunners) Tab(arg0 plugin.ManagerStore) (plugin.DefaultRunner, chan []component.Tab) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tab", arg0)
	ret0, _ := ret[0].(plugin.DefaultRunner)
	ret1, _ := ret[1].(chan []component.Tab)
	return ret0, ret1
}

// Tab indicates an expected call of Tab.
func (mr *MockRunnersMockRecorder) Tab(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tab", reflect.TypeOf((*MockRunners)(nil).Tab), arg0)
}

// MockManagerStore is a mock of ManagerStore interface.
type MockManagerStore struct {
	ctrl     *gomock.Controller
	recorder *MockManagerStoreMockRecorder
}

// MockManagerStoreMockRecorder is the mock recorder for MockManagerStore.
type MockManagerStoreMockRecorder struct {
	mock *MockManagerStore
}

// NewMockManagerStore creates a new mock instance.
func NewMockManagerStore(ctrl *gomock.Controller) *MockManagerStore {
	mock := &MockManagerStore{ctrl: ctrl}
	mock.recorder = &MockManagerStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManagerStore) EXPECT() *MockManagerStoreMockRecorder {
	return m.recorder
}

// ClientNames mocks base method.
func (m *MockManagerStore) ClientNames() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientNames")
	ret0, _ := ret[0].([]string)
	return ret0
}

// ClientNames indicates an expected call of ClientNames.
func (mr *MockManagerStoreMockRecorder) ClientNames() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientNames", reflect.TypeOf((*MockManagerStore)(nil).ClientNames))
}

// Clients mocks base method.
func (m *MockManagerStore) Clients() map[string]plugin.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clients")
	ret0, _ := ret[0].(map[string]plugin.Client)
	return ret0
}

// Clients indicates an expected call of Clients.
func (mr *MockManagerStoreMockRecorder) Clients() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clients", reflect.TypeOf((*MockManagerStore)(nil).Clients))
}

// Get mocks base method.
func (m *MockManagerStore) Get(arg0 string) (plugin.Client, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(plugin.Client)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockManagerStoreMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockManagerStore)(nil).Get), arg0)
}

// GetCommand mocks base method.
func (m *MockManagerStore) GetCommand(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommand", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommand indicates an expected call of GetCommand.
func (mr *MockManagerStoreMockRecorder) GetCommand(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommand", reflect.TypeOf((*MockManagerStore)(nil).GetCommand), arg0)
}

// GetJS mocks base method.
func (m *MockManagerStore) GetJS(arg0 string) (plugin.JSPlugin, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJS", arg0)
	ret0, _ := ret[0].(plugin.JSPlugin)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetJS indicates an expected call of GetJS.
func (mr *MockManagerStoreMockRecorder) GetJS(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJS", reflect.TypeOf((*MockManagerStore)(nil).GetJS), arg0)
}

// GetMetadata mocks base method.
func (m *MockManagerStore) GetMetadata(arg0 string) (*plugin.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetadata", arg0)
	ret0, _ := ret[0].(*plugin.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetadata indicates an expected call of GetMetadata.
func (mr *MockManagerStoreMockRecorder) GetMetadata(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadata", reflect.TypeOf((*MockManagerStore)(nil).GetMetadata), arg0)
}

// GetModuleService mocks base method.
func (m *MockManagerStore) GetModuleService(arg0 string) (plugin.ModuleService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleService", arg0)
	ret0, _ := ret[0].(plugin.ModuleService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModuleService indicates an expected call of GetModuleService.
func (mr *MockManagerStoreMockRecorder) GetModuleService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleService", reflect.TypeOf((*MockManagerStore)(nil).GetModuleService), arg0)
}

// GetService mocks base method.
func (m *MockManagerStore) GetService(arg0 string) (plugin.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", arg0)
	ret0, _ := ret[0].(plugin.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService.
func (mr *MockManagerStoreMockRecorder) GetService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockManagerStore)(nil).GetService), arg0)
}

// NamesJS mocks base method.
func (m *MockManagerStore) NamesJS() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NamesJS")
	ret0, _ := ret[0].([]string)
	return ret0
}

// NamesJS indicates an expected call of NamesJS.
func (mr *MockManagerStoreMockRecorder) NamesJS() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NamesJS", reflect.TypeOf((*MockManagerStore)(nil).NamesJS))
}

// Remove mocks base method.
func (m *MockManagerStore) Remove(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Remove", arg0)
}

// Remove indicates an expected call of Remove.
func (mr *MockManagerStoreMockRecorder) Remove(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockManagerStore)(nil).Remove), arg0)
}

// RemoveJS mocks base method.
func (m *MockManagerStore) RemoveJS(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveJS", arg0)
}

// RemoveJS indicates an expected call of RemoveJS.
func (mr *MockManagerStoreMockRecorder) RemoveJS(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveJS", reflect.TypeOf((*MockManagerStore)(nil).RemoveJS), arg0)
}

// Store mocks base method.
func (m *MockManagerStore) Store(arg0 string, arg1 plugin.Client, arg2 *plugin.Metadata, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockManagerStoreMockRecorder) Store(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockManagerStore)(nil).Store), arg0, arg1, arg2, arg3)
}

// StoreJS mocks base method.
func (m *MockManagerStore) StoreJS(arg0 string, arg1 plugin.JSPlugin) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreJS", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreJS indicates an expected call of StoreJS.
func (mr *MockManagerStoreMockRecorder) StoreJS(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreJS", reflect.TypeOf((*MockManagerStore)(nil).StoreJS), arg0, arg1)
}

// MockClientFactory is a mock of ClientFactory interface.
type MockClientFactory struct {
	ctrl     *gomock.Controller
	recorder *MockClientFactoryMockRecorder
}

// MockClientFactoryMockRecorder is the mock recorder for MockClientFactory.
type MockClientFactoryMockRecorder struct {
	mock *MockClientFactory
}

// NewMockClientFactory creates a new mock instance.
func NewMockClientFactory(ctrl *gomock.Controller) *MockClientFactory {
	mock := &MockClientFactory{ctrl: ctrl}
	mock.recorder = &MockClientFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientFactory) EXPECT() *MockClientFactoryMockRecorder {
	return m.recorder
}

// Init mocks base method.
func (m *MockClientFactory) Init(arg0 context.Context, arg1 string) plugin.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0, arg1)
	ret0, _ := ret[0].(plugin.Client)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockClientFactoryMockRecorder) Init(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockClientFactory)(nil).Init), arg0, arg1)
}

// MockModuleService is a mock of ModuleService interface.
type MockModuleService struct {
	ctrl     *gomock.Controller
	recorder *MockModuleServiceMockRecorder
}

// MockModuleServiceMockRecorder is the mock recorder for MockModuleService.
type MockModuleServiceMockRecorder struct {
	mock *MockModuleService
}

// NewMockModuleService creates a new mock instance.
func NewMockModuleService(ctrl *gomock.Controller) *MockModuleService {
	mock := &MockModuleService{ctrl: ctrl}
	mock.recorder = &MockModuleServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModuleService) EXPECT() *MockModuleServiceMockRecorder {
	return m.recorder
}

// Content mocks base method.
func (m *MockModuleService) Content(arg0 context.Context, arg1 string) (component.ContentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Content", arg0, arg1)
	ret0, _ := ret[0].(component.ContentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Content indicates an expected call of Content.
func (mr *MockModuleServiceMockRecorder) Content(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Content", reflect.TypeOf((*MockModuleService)(nil).Content), arg0, arg1)
}

// HandleAction mocks base method.
func (m *MockModuleService) HandleAction(arg0 context.Context, arg1 string, arg2 action.Payload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleAction", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleAction indicates an expected call of HandleAction.
func (mr *MockModuleServiceMockRecorder) HandleAction(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleAction", reflect.TypeOf((*MockModuleService)(nil).HandleAction), arg0, arg1, arg2)
}

// Navigation mocks base method.
func (m *MockModuleService) Navigation(arg0 context.Context) (navigation.Navigation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Navigation", arg0)
	ret0, _ := ret[0].(navigation.Navigation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Navigation indicates an expected call of Navigation.
func (mr *MockModuleServiceMockRecorder) Navigation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Navigation", reflect.TypeOf((*MockModuleService)(nil).Navigation), arg0)
}

// ObjectStatus mocks base method.
func (m *MockModuleService) ObjectStatus(arg0 context.Context, arg1 runtime.Object) (plugin.ObjectStatusResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectStatus", arg0, arg1)
	ret0, _ := ret[0].(plugin.ObjectStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObjectStatus indicates an expected call of ObjectStatus.
func (mr *MockModuleServiceMockRecorder) ObjectStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectStatus", reflect.TypeOf((*MockModuleService)(nil).ObjectStatus), arg0, arg1)
}

// Print mocks base method.
func (m *MockModuleService) Print(arg0 context.Context, arg1 runtime.Object) (plugin.PrintResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Print", arg0, arg1)
	ret0, _ := ret[0].(plugin.PrintResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Print indicates an expected call of Print.
func (mr *MockModuleServiceMockRecorder) Print(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Print", reflect.TypeOf((*MockModuleService)(nil).Print), arg0, arg1)
}

// PrintTabs mocks base method.
func (m *MockModuleService) PrintTabs(arg0 context.Context, arg1 runtime.Object) ([]plugin.TabResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrintTabs", arg0, arg1)
	ret0, _ := ret[0].([]plugin.TabResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrintTabs indicates an expected call of PrintTabs.
func (mr *MockModuleServiceMockRecorder) PrintTabs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrintTabs", reflect.TypeOf((*MockModuleService)(nil).PrintTabs), arg0, arg1)
}

// Register mocks base method.
func (m *MockModuleService) Register(arg0 context.Context, arg1 string) (plugin.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(plugin.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockModuleServiceMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockModuleService)(nil).Register), arg0, arg1)
}

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// HandleAction mocks base method.
func (m *MockService) HandleAction(arg0 context.Context, arg1 string, arg2 action.Payload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleAction", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleAction indicates an expected call of HandleAction.
func (mr *MockServiceMockRecorder) HandleAction(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleAction", reflect.TypeOf((*MockService)(nil).HandleAction), arg0, arg1, arg2)
}

// ObjectStatus mocks base method.
func (m *MockService) ObjectStatus(arg0 context.Context, arg1 runtime.Object) (plugin.ObjectStatusResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectStatus", arg0, arg1)
	ret0, _ := ret[0].(plugin.ObjectStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObjectStatus indicates an expected call of ObjectStatus.
func (mr *MockServiceMockRecorder) ObjectStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectStatus", reflect.TypeOf((*MockService)(nil).ObjectStatus), arg0, arg1)
}

// Print mocks base method.
func (m *MockService) Print(arg0 context.Context, arg1 runtime.Object) (plugin.PrintResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Print", arg0, arg1)
	ret0, _ := ret[0].(plugin.PrintResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Print indicates an expected call of Print.
func (mr *MockServiceMockRecorder) Print(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Print", reflect.TypeOf((*MockService)(nil).Print), arg0, arg1)
}

// PrintTabs mocks base method.
func (m *MockService) PrintTabs(arg0 context.Context, arg1 runtime.Object) ([]plugin.TabResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrintTabs", arg0, arg1)
	ret0, _ := ret[0].([]plugin.TabResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrintTabs indicates an expected call of PrintTabs.
func (mr *MockServiceMockRecorder) PrintTabs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrintTabs", reflect.TypeOf((*MockService)(nil).PrintTabs), arg0, arg1)
}

// Register mocks base method.
func (m *MockService) Register(arg0 context.Context, arg1 string) (plugin.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(plugin.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockServiceMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockService)(nil).Register), arg0, arg1)
}

// MockBroker is a mock of Broker interface.
type MockBroker struct {
	ctrl     *gomock.Controller
	recorder *MockBrokerMockRecorder
}

// MockBrokerMockRecorder is the mock recorder for MockBroker.
type MockBrokerMockRecorder struct {
	mock *MockBroker
}

// NewMockBroker creates a new mock instance.
func NewMockBroker(ctrl *gomock.Controller) *MockBroker {
	mock := &MockBroker{ctrl: ctrl}
	mock.recorder = &MockBrokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBroker) EXPECT() *MockBrokerMockRecorder {
	return m.recorder
}

// AcceptAndServe mocks base method.
func (m *MockBroker) AcceptAndServe(arg0 uint32, arg1 func([]grpc.ServerOption) *grpc.Server) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AcceptAndServe", arg0, arg1)
}

// AcceptAndServe indicates an expected call of AcceptAndServe.
func (mr *MockBrokerMockRecorder) AcceptAndServe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptAndServe", reflect.TypeOf((*MockBroker)(nil).AcceptAndServe), arg0, arg1)
}

// NextId mocks base method.
func (m *MockBroker) NextId() uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextId")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// NextId indicates an expected call of NextId.
func (mr *MockBrokerMockRecorder) NextId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextId", reflect.TypeOf((*MockBroker)(nil).NextId))
}
