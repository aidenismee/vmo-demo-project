// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nekizz/vmo-demo-project/pkg/filer (interfaces: Filer)
//
// Generated by this command:
//
//	mockgen -destination pkg/filer/mock.go -package=filer github.com/nekizz/vmo-demo-project/pkg/filer Filer
//

// Package filer is a generated GoMock package.
package filer

import (
	os "os"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockFiler is a mock of Filer interface.
type MockFiler struct {
	ctrl     *gomock.Controller
	recorder *MockFilerMockRecorder
}

// MockFilerMockRecorder is the mock recorder for MockFiler.
type MockFilerMockRecorder struct {
	mock *MockFiler
}

// NewMockFiler creates a new mock instance.
func NewMockFiler(ctrl *gomock.Controller) *MockFiler {
	mock := &MockFiler{ctrl: ctrl}
	mock.recorder = &MockFilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFiler) EXPECT() *MockFilerMockRecorder {
	return m.recorder
}

// CountLine mocks base method.
func (m *MockFiler) CountLine(arg0 *os.File) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountLine", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountLine indicates an expected call of CountLine.
func (mr *MockFilerMockRecorder) CountLine(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountLine", reflect.TypeOf((*MockFiler)(nil).CountLine), arg0)
}

// OpenFile mocks base method.
func (m *MockFiler) OpenFile(arg0 string) (*os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenFile", arg0)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenFile indicates an expected call of OpenFile.
func (mr *MockFilerMockRecorder) OpenFile(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenFile", reflect.TypeOf((*MockFiler)(nil).OpenFile), arg0)
}
