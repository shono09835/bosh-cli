// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/shono09835/bosh-cli/v7/templatescompiler (interfaces: JobRenderer,JobListRenderer,RenderedJob,RenderedJobList,RenderedJobListArchive,RenderedJobListCompressor)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	job "github.com/shono09835/bosh-cli/v7/release/job"
	templatescompiler "github.com/shono09835/bosh-cli/v7/templatescompiler"
	property "github.com/cloudfoundry/bosh-utils/property"
	gomock "github.com/golang/mock/gomock"
)

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

// Render mocks base method.
func (m *MockJobRenderer) Render(arg0 job.Job, arg1 *property.Map, arg2, arg3 property.Map, arg4, arg5 string) (templatescompiler.RenderedJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Render", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(templatescompiler.RenderedJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Render indicates an expected call of Render.
func (mr *MockJobRendererMockRecorder) Render(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Render", reflect.TypeOf((*MockJobRenderer)(nil).Render), arg0, arg1, arg2, arg3, arg4, arg5)
}

// MockJobListRenderer is a mock of JobListRenderer interface.
type MockJobListRenderer struct {
	ctrl     *gomock.Controller
	recorder *MockJobListRendererMockRecorder
}

// MockJobListRendererMockRecorder is the mock recorder for MockJobListRenderer.
type MockJobListRendererMockRecorder struct {
	mock *MockJobListRenderer
}

// NewMockJobListRenderer creates a new mock instance.
func NewMockJobListRenderer(ctrl *gomock.Controller) *MockJobListRenderer {
	mock := &MockJobListRenderer{ctrl: ctrl}
	mock.recorder = &MockJobListRendererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobListRenderer) EXPECT() *MockJobListRendererMockRecorder {
	return m.recorder
}

// Render mocks base method.
func (m *MockJobListRenderer) Render(arg0 []job.Job, arg1 map[string]*property.Map, arg2, arg3 property.Map, arg4, arg5 string) (templatescompiler.RenderedJobList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Render", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(templatescompiler.RenderedJobList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Render indicates an expected call of Render.
func (mr *MockJobListRendererMockRecorder) Render(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Render", reflect.TypeOf((*MockJobListRenderer)(nil).Render), arg0, arg1, arg2, arg3, arg4, arg5)
}

// MockRenderedJob is a mock of RenderedJob interface.
type MockRenderedJob struct {
	ctrl     *gomock.Controller
	recorder *MockRenderedJobMockRecorder
}

// MockRenderedJobMockRecorder is the mock recorder for MockRenderedJob.
type MockRenderedJobMockRecorder struct {
	mock *MockRenderedJob
}

// NewMockRenderedJob creates a new mock instance.
func NewMockRenderedJob(ctrl *gomock.Controller) *MockRenderedJob {
	mock := &MockRenderedJob{ctrl: ctrl}
	mock.recorder = &MockRenderedJobMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRenderedJob) EXPECT() *MockRenderedJobMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockRenderedJob) Delete() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete")
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRenderedJobMockRecorder) Delete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRenderedJob)(nil).Delete))
}

// DeleteSilently mocks base method.
func (m *MockRenderedJob) DeleteSilently() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteSilently")
}

// DeleteSilently indicates an expected call of DeleteSilently.
func (mr *MockRenderedJobMockRecorder) DeleteSilently() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSilently", reflect.TypeOf((*MockRenderedJob)(nil).DeleteSilently))
}

// Job mocks base method.
func (m *MockRenderedJob) Job() job.Job {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Job")
	ret0, _ := ret[0].(job.Job)
	return ret0
}

// Job indicates an expected call of Job.
func (mr *MockRenderedJobMockRecorder) Job() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Job", reflect.TypeOf((*MockRenderedJob)(nil).Job))
}

// Path mocks base method.
func (m *MockRenderedJob) Path() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Path")
	ret0, _ := ret[0].(string)
	return ret0
}

// Path indicates an expected call of Path.
func (mr *MockRenderedJobMockRecorder) Path() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Path", reflect.TypeOf((*MockRenderedJob)(nil).Path))
}

// MockRenderedJobList is a mock of RenderedJobList interface.
type MockRenderedJobList struct {
	ctrl     *gomock.Controller
	recorder *MockRenderedJobListMockRecorder
}

// MockRenderedJobListMockRecorder is the mock recorder for MockRenderedJobList.
type MockRenderedJobListMockRecorder struct {
	mock *MockRenderedJobList
}

// NewMockRenderedJobList creates a new mock instance.
func NewMockRenderedJobList(ctrl *gomock.Controller) *MockRenderedJobList {
	mock := &MockRenderedJobList{ctrl: ctrl}
	mock.recorder = &MockRenderedJobListMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRenderedJobList) EXPECT() *MockRenderedJobListMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockRenderedJobList) Add(arg0 templatescompiler.RenderedJob) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", arg0)
}

// Add indicates an expected call of Add.
func (mr *MockRenderedJobListMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockRenderedJobList)(nil).Add), arg0)
}

// All mocks base method.
func (m *MockRenderedJobList) All() []templatescompiler.RenderedJob {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All")
	ret0, _ := ret[0].([]templatescompiler.RenderedJob)
	return ret0
}

// All indicates an expected call of All.
func (mr *MockRenderedJobListMockRecorder) All() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockRenderedJobList)(nil).All))
}

// Delete mocks base method.
func (m *MockRenderedJobList) Delete() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete")
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRenderedJobListMockRecorder) Delete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRenderedJobList)(nil).Delete))
}

// DeleteSilently mocks base method.
func (m *MockRenderedJobList) DeleteSilently() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteSilently")
}

// DeleteSilently indicates an expected call of DeleteSilently.
func (mr *MockRenderedJobListMockRecorder) DeleteSilently() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSilently", reflect.TypeOf((*MockRenderedJobList)(nil).DeleteSilently))
}

// MockRenderedJobListArchive is a mock of RenderedJobListArchive interface.
type MockRenderedJobListArchive struct {
	ctrl     *gomock.Controller
	recorder *MockRenderedJobListArchiveMockRecorder
}

// MockRenderedJobListArchiveMockRecorder is the mock recorder for MockRenderedJobListArchive.
type MockRenderedJobListArchiveMockRecorder struct {
	mock *MockRenderedJobListArchive
}

// NewMockRenderedJobListArchive creates a new mock instance.
func NewMockRenderedJobListArchive(ctrl *gomock.Controller) *MockRenderedJobListArchive {
	mock := &MockRenderedJobListArchive{ctrl: ctrl}
	mock.recorder = &MockRenderedJobListArchiveMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRenderedJobListArchive) EXPECT() *MockRenderedJobListArchiveMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockRenderedJobListArchive) Delete() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete")
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRenderedJobListArchiveMockRecorder) Delete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRenderedJobListArchive)(nil).Delete))
}

// DeleteSilently mocks base method.
func (m *MockRenderedJobListArchive) DeleteSilently() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteSilently")
}

// DeleteSilently indicates an expected call of DeleteSilently.
func (mr *MockRenderedJobListArchiveMockRecorder) DeleteSilently() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSilently", reflect.TypeOf((*MockRenderedJobListArchive)(nil).DeleteSilently))
}

// List mocks base method.
func (m *MockRenderedJobListArchive) List() templatescompiler.RenderedJobList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(templatescompiler.RenderedJobList)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockRenderedJobListArchiveMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRenderedJobListArchive)(nil).List))
}

// Path mocks base method.
func (m *MockRenderedJobListArchive) Path() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Path")
	ret0, _ := ret[0].(string)
	return ret0
}

// Path indicates an expected call of Path.
func (mr *MockRenderedJobListArchiveMockRecorder) Path() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Path", reflect.TypeOf((*MockRenderedJobListArchive)(nil).Path))
}

// SHA1 mocks base method.
func (m *MockRenderedJobListArchive) SHA1() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SHA1")
	ret0, _ := ret[0].(string)
	return ret0
}

// SHA1 indicates an expected call of SHA1.
func (mr *MockRenderedJobListArchiveMockRecorder) SHA1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SHA1", reflect.TypeOf((*MockRenderedJobListArchive)(nil).SHA1))
}

// MockRenderedJobListCompressor is a mock of RenderedJobListCompressor interface.
type MockRenderedJobListCompressor struct {
	ctrl     *gomock.Controller
	recorder *MockRenderedJobListCompressorMockRecorder
}

// MockRenderedJobListCompressorMockRecorder is the mock recorder for MockRenderedJobListCompressor.
type MockRenderedJobListCompressorMockRecorder struct {
	mock *MockRenderedJobListCompressor
}

// NewMockRenderedJobListCompressor creates a new mock instance.
func NewMockRenderedJobListCompressor(ctrl *gomock.Controller) *MockRenderedJobListCompressor {
	mock := &MockRenderedJobListCompressor{ctrl: ctrl}
	mock.recorder = &MockRenderedJobListCompressorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRenderedJobListCompressor) EXPECT() *MockRenderedJobListCompressorMockRecorder {
	return m.recorder
}

// Compress mocks base method.
func (m *MockRenderedJobListCompressor) Compress(arg0 templatescompiler.RenderedJobList) (templatescompiler.RenderedJobListArchive, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compress", arg0)
	ret0, _ := ret[0].(templatescompiler.RenderedJobListArchive)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Compress indicates an expected call of Compress.
func (mr *MockRenderedJobListCompressorMockRecorder) Compress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compress", reflect.TypeOf((*MockRenderedJobListCompressor)(nil).Compress), arg0)
}
