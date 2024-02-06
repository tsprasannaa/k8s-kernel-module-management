// Code generated by MockGen. DO NOT EDIT.
// Source: module_nmc_reconciler.go
//
// Generated by this command:
//
//	mockgen -source=module_nmc_reconciler.go -package=controllers -destination=mock_module_nmc_reconciler.go moduleNMCReconcilerHelperAPI,namespaceLabeler
//
// Package controllers is a generated GoMock package.
package controllers

import (
	context "context"
	reflect "reflect"

	v1beta1 "github.com/kubernetes-sigs/kernel-module-management/api/v1beta1"
	api "github.com/kubernetes-sigs/kernel-module-management/internal/api"
	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/core/v1"
	types "k8s.io/apimachinery/pkg/types"
	sets "k8s.io/apimachinery/pkg/util/sets"
)

// MockmoduleNMCReconcilerHelperAPI is a mock of moduleNMCReconcilerHelperAPI interface.
type MockmoduleNMCReconcilerHelperAPI struct {
	ctrl     *gomock.Controller
	recorder *MockmoduleNMCReconcilerHelperAPIMockRecorder
}

// MockmoduleNMCReconcilerHelperAPIMockRecorder is the mock recorder for MockmoduleNMCReconcilerHelperAPI.
type MockmoduleNMCReconcilerHelperAPIMockRecorder struct {
	mock *MockmoduleNMCReconcilerHelperAPI
}

// NewMockmoduleNMCReconcilerHelperAPI creates a new mock instance.
func NewMockmoduleNMCReconcilerHelperAPI(ctrl *gomock.Controller) *MockmoduleNMCReconcilerHelperAPI {
	mock := &MockmoduleNMCReconcilerHelperAPI{ctrl: ctrl}
	mock.recorder = &MockmoduleNMCReconcilerHelperAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockmoduleNMCReconcilerHelperAPI) EXPECT() *MockmoduleNMCReconcilerHelperAPIMockRecorder {
	return m.recorder
}

// disableModuleOnNode mocks base method.
func (m *MockmoduleNMCReconcilerHelperAPI) disableModuleOnNode(ctx context.Context, modNamespace, modName, nodeName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "disableModuleOnNode", ctx, modNamespace, modName, nodeName)
	ret0, _ := ret[0].(error)
	return ret0
}

// disableModuleOnNode indicates an expected call of disableModuleOnNode.
func (mr *MockmoduleNMCReconcilerHelperAPIMockRecorder) disableModuleOnNode(ctx, modNamespace, modName, nodeName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "disableModuleOnNode", reflect.TypeOf((*MockmoduleNMCReconcilerHelperAPI)(nil).disableModuleOnNode), ctx, modNamespace, modName, nodeName)
}

// enableModuleOnNode mocks base method.
func (m *MockmoduleNMCReconcilerHelperAPI) enableModuleOnNode(ctx context.Context, mld *api.ModuleLoaderData, node *v1.Node) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "enableModuleOnNode", ctx, mld, node)
	ret0, _ := ret[0].(error)
	return ret0
}

// enableModuleOnNode indicates an expected call of enableModuleOnNode.
func (mr *MockmoduleNMCReconcilerHelperAPIMockRecorder) enableModuleOnNode(ctx, mld, node any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "enableModuleOnNode", reflect.TypeOf((*MockmoduleNMCReconcilerHelperAPI)(nil).enableModuleOnNode), ctx, mld, node)
}

// finalizeModule mocks base method.
func (m *MockmoduleNMCReconcilerHelperAPI) finalizeModule(ctx context.Context, mod *v1beta1.Module) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "finalizeModule", ctx, mod)
	ret0, _ := ret[0].(error)
	return ret0
}

// finalizeModule indicates an expected call of finalizeModule.
func (mr *MockmoduleNMCReconcilerHelperAPIMockRecorder) finalizeModule(ctx, mod any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "finalizeModule", reflect.TypeOf((*MockmoduleNMCReconcilerHelperAPI)(nil).finalizeModule), ctx, mod)
}

// getNMCsByModuleSet mocks base method.
func (m *MockmoduleNMCReconcilerHelperAPI) getNMCsByModuleSet(ctx context.Context, mod *v1beta1.Module) (sets.Set[string], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getNMCsByModuleSet", ctx, mod)
	ret0, _ := ret[0].(sets.Set[string])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getNMCsByModuleSet indicates an expected call of getNMCsByModuleSet.
func (mr *MockmoduleNMCReconcilerHelperAPIMockRecorder) getNMCsByModuleSet(ctx, mod any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getNMCsByModuleSet", reflect.TypeOf((*MockmoduleNMCReconcilerHelperAPI)(nil).getNMCsByModuleSet), ctx, mod)
}

// getNodesListBySelector mocks base method.
func (m *MockmoduleNMCReconcilerHelperAPI) getNodesListBySelector(ctx context.Context, mod *v1beta1.Module) ([]v1.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getNodesListBySelector", ctx, mod)
	ret0, _ := ret[0].([]v1.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getNodesListBySelector indicates an expected call of getNodesListBySelector.
func (mr *MockmoduleNMCReconcilerHelperAPIMockRecorder) getNodesListBySelector(ctx, mod any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getNodesListBySelector", reflect.TypeOf((*MockmoduleNMCReconcilerHelperAPI)(nil).getNodesListBySelector), ctx, mod)
}

// getRequestedModule mocks base method.
func (m *MockmoduleNMCReconcilerHelperAPI) getRequestedModule(ctx context.Context, namespacedName types.NamespacedName) (*v1beta1.Module, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getRequestedModule", ctx, namespacedName)
	ret0, _ := ret[0].(*v1beta1.Module)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getRequestedModule indicates an expected call of getRequestedModule.
func (mr *MockmoduleNMCReconcilerHelperAPIMockRecorder) getRequestedModule(ctx, namespacedName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getRequestedModule", reflect.TypeOf((*MockmoduleNMCReconcilerHelperAPI)(nil).getRequestedModule), ctx, namespacedName)
}

// moduleUpdateWorkerPodsStatus mocks base method.
func (m *MockmoduleNMCReconcilerHelperAPI) moduleUpdateWorkerPodsStatus(ctx context.Context, mod *v1beta1.Module, targetedNodes []v1.Node) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "moduleUpdateWorkerPodsStatus", ctx, mod, targetedNodes)
	ret0, _ := ret[0].(error)
	return ret0
}

// moduleUpdateWorkerPodsStatus indicates an expected call of moduleUpdateWorkerPodsStatus.
func (mr *MockmoduleNMCReconcilerHelperAPIMockRecorder) moduleUpdateWorkerPodsStatus(ctx, mod, targetedNodes any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "moduleUpdateWorkerPodsStatus", reflect.TypeOf((*MockmoduleNMCReconcilerHelperAPI)(nil).moduleUpdateWorkerPodsStatus), ctx, mod, targetedNodes)
}

// prepareSchedulingData mocks base method.
func (m *MockmoduleNMCReconcilerHelperAPI) prepareSchedulingData(ctx context.Context, mod *v1beta1.Module, targetedNodes []v1.Node, currentNMCs sets.Set[string]) (map[string]schedulingData, []error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "prepareSchedulingData", ctx, mod, targetedNodes, currentNMCs)
	ret0, _ := ret[0].(map[string]schedulingData)
	ret1, _ := ret[1].([]error)
	return ret0, ret1
}

// prepareSchedulingData indicates an expected call of prepareSchedulingData.
func (mr *MockmoduleNMCReconcilerHelperAPIMockRecorder) prepareSchedulingData(ctx, mod, targetedNodes, currentNMCs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "prepareSchedulingData", reflect.TypeOf((*MockmoduleNMCReconcilerHelperAPI)(nil).prepareSchedulingData), ctx, mod, targetedNodes, currentNMCs)
}

// setFinalizerAndStatus mocks base method.
func (m *MockmoduleNMCReconcilerHelperAPI) setFinalizerAndStatus(ctx context.Context, mod *v1beta1.Module) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "setFinalizerAndStatus", ctx, mod)
	ret0, _ := ret[0].(error)
	return ret0
}

// setFinalizerAndStatus indicates an expected call of setFinalizerAndStatus.
func (mr *MockmoduleNMCReconcilerHelperAPIMockRecorder) setFinalizerAndStatus(ctx, mod any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "setFinalizerAndStatus", reflect.TypeOf((*MockmoduleNMCReconcilerHelperAPI)(nil).setFinalizerAndStatus), ctx, mod)
}

// MocknamespaceLabeler is a mock of namespaceLabeler interface.
type MocknamespaceLabeler struct {
	ctrl     *gomock.Controller
	recorder *MocknamespaceLabelerMockRecorder
}

// MocknamespaceLabelerMockRecorder is the mock recorder for MocknamespaceLabeler.
type MocknamespaceLabelerMockRecorder struct {
	mock *MocknamespaceLabeler
}

// NewMocknamespaceLabeler creates a new mock instance.
func NewMocknamespaceLabeler(ctrl *gomock.Controller) *MocknamespaceLabeler {
	mock := &MocknamespaceLabeler{ctrl: ctrl}
	mock.recorder = &MocknamespaceLabelerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocknamespaceLabeler) EXPECT() *MocknamespaceLabelerMockRecorder {
	return m.recorder
}

// setLabel mocks base method.
func (m *MocknamespaceLabeler) setLabel(ctx context.Context, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "setLabel", ctx, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// setLabel indicates an expected call of setLabel.
func (mr *MocknamespaceLabelerMockRecorder) setLabel(ctx, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "setLabel", reflect.TypeOf((*MocknamespaceLabeler)(nil).setLabel), ctx, name)
}

// tryRemovingLabel mocks base method.
func (m *MocknamespaceLabeler) tryRemovingLabel(ctx context.Context, name, moduleName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "tryRemovingLabel", ctx, name, moduleName)
	ret0, _ := ret[0].(error)
	return ret0
}

// tryRemovingLabel indicates an expected call of tryRemovingLabel.
func (mr *MocknamespaceLabelerMockRecorder) tryRemovingLabel(ctx, name, moduleName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "tryRemovingLabel", reflect.TypeOf((*MocknamespaceLabeler)(nil).tryRemovingLabel), ctx, name, moduleName)
}
