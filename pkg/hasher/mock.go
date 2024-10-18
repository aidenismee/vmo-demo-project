// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nekizz/vmo-demo-project/pkg/hasher (interfaces: Service)
//
// Generated by this command:
//
//	mockgen -destination pkg/hasher/mock.go -package=hasher github.com/nekizz/vmo-demo-project/pkg/hasher Service
//

// Package hasher is a generated GoMock package.
package hasher

import (
	io "io"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

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

// HashFile mocks base method.
func (m *MockService) HashFile(arg0 io.Reader) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashFile", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashFile indicates an expected call of HashFile.
func (mr *MockServiceMockRecorder) HashFile(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashFile", reflect.TypeOf((*MockService)(nil).HashFile), arg0)
}
