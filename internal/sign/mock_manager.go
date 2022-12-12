// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go

// Package sign is a generated GoMock package.
package sign

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "github.com/kubernetes-sigs/kernel-module-management/api/v1beta1"
	utils "github.com/kubernetes-sigs/kernel-module-management/internal/utils"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MockSignManager is a mock of SignManager interface.
type MockSignManager struct {
	ctrl     *gomock.Controller
	recorder *MockSignManagerMockRecorder
}

// MockSignManagerMockRecorder is the mock recorder for MockSignManager.
type MockSignManagerMockRecorder struct {
	mock *MockSignManager
}

// NewMockSignManager creates a new mock instance.
func NewMockSignManager(ctrl *gomock.Controller) *MockSignManager {
	mock := &MockSignManager{ctrl: ctrl}
	mock.recorder = &MockSignManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSignManager) EXPECT() *MockSignManagerMockRecorder {
	return m.recorder
}

// ShouldSync mocks base method.
func (m_2 *MockSignManager) ShouldSync(ctx context.Context, mod v1beta1.Module, m v1beta1.KernelMapping) (bool, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "ShouldSync", ctx, mod, m)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShouldSync indicates an expected call of ShouldSync.
func (mr *MockSignManagerMockRecorder) ShouldSync(ctx, mod, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldSync", reflect.TypeOf((*MockSignManager)(nil).ShouldSync), ctx, mod, m)
}

// Sync mocks base method.
func (m_2 *MockSignManager) Sync(ctx context.Context, mod v1beta1.Module, m v1beta1.KernelMapping, targetKernel, imageToSign string, pushImage bool, owner v1.Object) (utils.Result, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Sync", ctx, mod, m, targetKernel, imageToSign, pushImage, owner)
	ret0, _ := ret[0].(utils.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sync indicates an expected call of Sync.
func (mr *MockSignManagerMockRecorder) Sync(ctx, mod, m, targetKernel, imageToSign, pushImage, owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockSignManager)(nil).Sync), ctx, mod, m, targetKernel, imageToSign, pushImage, owner)
}