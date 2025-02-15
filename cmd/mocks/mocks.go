// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/shono09835/bosh-cli/v7/cmd (interfaces: DeploymentDeleter,DeploymentStateManager)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	ui "github.com/shono09835/bosh-cli/v7/ui"
	gomock "github.com/golang/mock/gomock"
)

// MockDeploymentDeleter is a mock of DeploymentDeleter interface.
type MockDeploymentDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentDeleterMockRecorder
}

// MockDeploymentDeleterMockRecorder is the mock recorder for MockDeploymentDeleter.
type MockDeploymentDeleterMockRecorder struct {
	mock *MockDeploymentDeleter
}

// NewMockDeploymentDeleter creates a new mock instance.
func NewMockDeploymentDeleter(ctrl *gomock.Controller) *MockDeploymentDeleter {
	mock := &MockDeploymentDeleter{ctrl: ctrl}
	mock.recorder = &MockDeploymentDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentDeleter) EXPECT() *MockDeploymentDeleterMockRecorder {
	return m.recorder
}

// DeleteDeployment mocks base method.
func (m *MockDeploymentDeleter) DeleteDeployment(arg0 bool, arg1 ui.Stage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDeployment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDeployment indicates an expected call of DeleteDeployment.
func (mr *MockDeploymentDeleterMockRecorder) DeleteDeployment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDeployment", reflect.TypeOf((*MockDeploymentDeleter)(nil).DeleteDeployment), arg0, arg1)
}

// MockDeploymentStateManager is a mock of DeploymentStateManager interface.
type MockDeploymentStateManager struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentStateManagerMockRecorder
}

// MockDeploymentStateManagerMockRecorder is the mock recorder for MockDeploymentStateManager.
type MockDeploymentStateManagerMockRecorder struct {
	mock *MockDeploymentStateManager
}

// NewMockDeploymentStateManager creates a new mock instance.
func NewMockDeploymentStateManager(ctrl *gomock.Controller) *MockDeploymentStateManager {
	mock := &MockDeploymentStateManager{ctrl: ctrl}
	mock.recorder = &MockDeploymentStateManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentStateManager) EXPECT() *MockDeploymentStateManagerMockRecorder {
	return m.recorder
}

// StartDeployment mocks base method.
func (m *MockDeploymentStateManager) StartDeployment(arg0 ui.Stage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartDeployment", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartDeployment indicates an expected call of StartDeployment.
func (mr *MockDeploymentStateManagerMockRecorder) StartDeployment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartDeployment", reflect.TypeOf((*MockDeploymentStateManager)(nil).StartDeployment), arg0)
}

// StopDeployment mocks base method.
func (m *MockDeploymentStateManager) StopDeployment(arg0 bool, arg1 ui.Stage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopDeployment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopDeployment indicates an expected call of StopDeployment.
func (mr *MockDeploymentStateManagerMockRecorder) StopDeployment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopDeployment", reflect.TypeOf((*MockDeploymentStateManager)(nil).StopDeployment), arg0, arg1)
}
