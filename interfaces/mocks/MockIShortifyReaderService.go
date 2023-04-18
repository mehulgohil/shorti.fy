// Code generated by MockGen. DO NOT EDIT.
// Source: .\IShortifyReaderService.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIShortifyReaderService is a mock of IShortifyReaderService interface.
type MockIShortifyReaderService struct {
	ctrl     *gomock.Controller
	recorder *MockIShortifyReaderServiceMockRecorder
}

// MockIShortifyReaderServiceMockRecorder is the mock recorder for MockIShortifyReaderService.
type MockIShortifyReaderServiceMockRecorder struct {
	mock *MockIShortifyReaderService
}

// NewMockIShortifyReaderService creates a new mock instance.
func NewMockIShortifyReaderService(ctrl *gomock.Controller) *MockIShortifyReaderService {
	mock := &MockIShortifyReaderService{ctrl: ctrl}
	mock.recorder = &MockIShortifyReaderServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIShortifyReaderService) EXPECT() *MockIShortifyReaderServiceMockRecorder {
	return m.recorder
}

// Reader mocks base method.
func (m *MockIShortifyReaderService) Reader(shortURLHash string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reader", shortURLHash)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reader indicates an expected call of Reader.
func (mr *MockIShortifyReaderServiceMockRecorder) Reader(shortURLHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reader", reflect.TypeOf((*MockIShortifyReaderService)(nil).Reader), shortURLHash)
}