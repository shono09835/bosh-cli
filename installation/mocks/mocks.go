// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/shono09835/bosh-cli/v7/installation (interfaces: Installation,Installer,InstallerFactory,Uninstaller,JobResolver,PackageCompiler,JobRenderer)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	installation "github.com/shono09835/bosh-cli/v7/installation"
	manifest "github.com/shono09835/bosh-cli/v7/installation/manifest"
	job "github.com/shono09835/bosh-cli/v7/release/job"
	ui "github.com/shono09835/bosh-cli/v7/ui"
	gomock "github.com/golang/mock/gomock"
)

// MockInstallation is a mock of Installation interface.
type MockInstallation struct {
	ctrl     *gomock.Controller
	recorder *MockInstallationMockRecorder
}

// MockInstallationMockRecorder is the mock recorder for MockInstallation.
type MockInstallationMockRecorder struct {
	mock *MockInstallation
}

// NewMockInstallation creates a new mock instance.
func NewMockInstallation(ctrl *gomock.Controller) *MockInstallation {
	mock := &MockInstallation{ctrl: ctrl}
	mock.recorder = &MockInstallationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstallation) EXPECT() *MockInstallationMockRecorder {
	return m.recorder
}

// Job mocks base method.
func (m *MockInstallation) Job() installation.InstalledJob {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Job")
	ret0, _ := ret[0].(installation.InstalledJob)
	return ret0
}

// Job indicates an expected call of Job.
func (mr *MockInstallationMockRecorder) Job() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Job", reflect.TypeOf((*MockInstallation)(nil).Job))
}

// Target mocks base method.
func (m *MockInstallation) Target() installation.Target {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Target")
	ret0, _ := ret[0].(installation.Target)
	return ret0
}

// Target indicates an expected call of Target.
func (mr *MockInstallationMockRecorder) Target() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Target", reflect.TypeOf((*MockInstallation)(nil).Target))
}

// MockInstaller is a mock of Installer interface.
type MockInstaller struct {
	ctrl     *gomock.Controller
	recorder *MockInstallerMockRecorder
}

// MockInstallerMockRecorder is the mock recorder for MockInstaller.
type MockInstallerMockRecorder struct {
	mock *MockInstaller
}

// NewMockInstaller creates a new mock instance.
func NewMockInstaller(ctrl *gomock.Controller) *MockInstaller {
	mock := &MockInstaller{ctrl: ctrl}
	mock.recorder = &MockInstallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstaller) EXPECT() *MockInstallerMockRecorder {
	return m.recorder
}

// Cleanup mocks base method.
func (m *MockInstaller) Cleanup(arg0 installation.Installation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cleanup", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Cleanup indicates an expected call of Cleanup.
func (mr *MockInstallerMockRecorder) Cleanup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cleanup", reflect.TypeOf((*MockInstaller)(nil).Cleanup), arg0)
}

// Install mocks base method.
func (m *MockInstaller) Install(arg0 manifest.Manifest, arg1 ui.Stage) (installation.Installation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install", arg0, arg1)
	ret0, _ := ret[0].(installation.Installation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Install indicates an expected call of Install.
func (mr *MockInstallerMockRecorder) Install(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockInstaller)(nil).Install), arg0, arg1)
}

// MockInstallerFactory is a mock of InstallerFactory interface.
type MockInstallerFactory struct {
	ctrl     *gomock.Controller
	recorder *MockInstallerFactoryMockRecorder
}

// MockInstallerFactoryMockRecorder is the mock recorder for MockInstallerFactory.
type MockInstallerFactoryMockRecorder struct {
	mock *MockInstallerFactory
}

// NewMockInstallerFactory creates a new mock instance.
func NewMockInstallerFactory(ctrl *gomock.Controller) *MockInstallerFactory {
	mock := &MockInstallerFactory{ctrl: ctrl}
	mock.recorder = &MockInstallerFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstallerFactory) EXPECT() *MockInstallerFactoryMockRecorder {
	return m.recorder
}

// NewInstaller mocks base method.
func (m *MockInstallerFactory) NewInstaller(arg0 installation.Target) installation.Installer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewInstaller", arg0)
	ret0, _ := ret[0].(installation.Installer)
	return ret0
}

// NewInstaller indicates an expected call of NewInstaller.
func (mr *MockInstallerFactoryMockRecorder) NewInstaller(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewInstaller", reflect.TypeOf((*MockInstallerFactory)(nil).NewInstaller), arg0)
}

// MockUninstaller is a mock of Uninstaller interface.
type MockUninstaller struct {
	ctrl     *gomock.Controller
	recorder *MockUninstallerMockRecorder
}

// MockUninstallerMockRecorder is the mock recorder for MockUninstaller.
type MockUninstallerMockRecorder struct {
	mock *MockUninstaller
}

// NewMockUninstaller creates a new mock instance.
func NewMockUninstaller(ctrl *gomock.Controller) *MockUninstaller {
	mock := &MockUninstaller{ctrl: ctrl}
	mock.recorder = &MockUninstallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUninstaller) EXPECT() *MockUninstallerMockRecorder {
	return m.recorder
}

// Uninstall mocks base method.
func (m *MockUninstaller) Uninstall(arg0 installation.Target) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uninstall", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Uninstall indicates an expected call of Uninstall.
func (mr *MockUninstallerMockRecorder) Uninstall(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uninstall", reflect.TypeOf((*MockUninstaller)(nil).Uninstall), arg0)
}

// MockJobResolver is a mock of JobResolver interface.
type MockJobResolver struct {
	ctrl     *gomock.Controller
	recorder *MockJobResolverMockRecorder
}

// MockJobResolverMockRecorder is the mock recorder for MockJobResolver.
type MockJobResolverMockRecorder struct {
	mock *MockJobResolver
}

// NewMockJobResolver creates a new mock instance.
func NewMockJobResolver(ctrl *gomock.Controller) *MockJobResolver {
	mock := &MockJobResolver{ctrl: ctrl}
	mock.recorder = &MockJobResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobResolver) EXPECT() *MockJobResolverMockRecorder {
	return m.recorder
}

// From mocks base method.
func (m *MockJobResolver) From(arg0 manifest.Manifest) ([]job.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "From", arg0)
	ret0, _ := ret[0].([]job.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// From indicates an expected call of From.
func (mr *MockJobResolverMockRecorder) From(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "From", reflect.TypeOf((*MockJobResolver)(nil).From), arg0)
}

// MockPackageCompiler is a mock of PackageCompiler interface.
type MockPackageCompiler struct {
	ctrl     *gomock.Controller
	recorder *MockPackageCompilerMockRecorder
}

// MockPackageCompilerMockRecorder is the mock recorder for MockPackageCompiler.
type MockPackageCompilerMockRecorder struct {
	mock *MockPackageCompiler
}

// NewMockPackageCompiler creates a new mock instance.
func NewMockPackageCompiler(ctrl *gomock.Controller) *MockPackageCompiler {
	mock := &MockPackageCompiler{ctrl: ctrl}
	mock.recorder = &MockPackageCompilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPackageCompiler) EXPECT() *MockPackageCompilerMockRecorder {
	return m.recorder
}

// For mocks base method.
func (m *MockPackageCompiler) For(arg0 []job.Job, arg1 ui.Stage) ([]installation.CompiledPackageRef, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "For", arg0, arg1)
	ret0, _ := ret[0].([]installation.CompiledPackageRef)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// For indicates an expected call of For.
func (mr *MockPackageCompilerMockRecorder) For(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "For", reflect.TypeOf((*MockPackageCompiler)(nil).For), arg0, arg1)
}

// MockJobRenderer is a mock of JobRenderer interface.
type MockJobRenderer struct {
	ctrl     *gomock.Controller
	recorder *MockJobRendererMockRecorder
}

// MockJobRendererMockRecorder is the mock recorder for MockJobRenderer.
type MockJobRendererMockRecorder struct {
	mock *MockJobRenderer
}

// NewMockJobRenderer creates a new mock instance.
func NewMockJobRenderer(ctrl *gomock.Controller) *MockJobRenderer {
	mock := &MockJobRenderer{ctrl: ctrl}
	mock.recorder = &MockJobRendererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobRenderer) EXPECT() *MockJobRendererMockRecorder {
	return m.recorder
}

// RenderAndUploadFrom mocks base method.
func (m *MockJobRenderer) RenderAndUploadFrom(arg0 manifest.Manifest, arg1 []job.Job, arg2 ui.Stage) ([]installation.RenderedJobRef, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenderAndUploadFrom", arg0, arg1, arg2)
	ret0, _ := ret[0].([]installation.RenderedJobRef)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenderAndUploadFrom indicates an expected call of RenderAndUploadFrom.
func (mr *MockJobRendererMockRecorder) RenderAndUploadFrom(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenderAndUploadFrom", reflect.TypeOf((*MockJobRenderer)(nil).RenderAndUploadFrom), arg0, arg1, arg2)
}
