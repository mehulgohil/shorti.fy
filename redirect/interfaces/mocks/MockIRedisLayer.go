// Code generated by MockGen. DO NOT EDIT.
// Source: .\IRedisLayer.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockIRedisLayer is a mock of IRedisLayer interface.
type MockIRedisLayer struct {
	ctrl     *gomock.Controller
	recorder *MockIRedisLayerMockRecorder
}

// MockIRedisLayerMockRecorder is the mock recorder for MockIRedisLayer.
type MockIRedisLayerMockRecorder struct {
	mock *MockIRedisLayer
}

// NewMockIRedisLayer creates a new mock instance.
func NewMockIRedisLayer(ctrl *gomock.Controller) *MockIRedisLayer {
	mock := &MockIRedisLayer{ctrl: ctrl}
	mock.recorder = &MockIRedisLayerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRedisLayer) EXPECT() *MockIRedisLayerMockRecorder {
	return m.recorder
}

// GetKeyValue mocks base method.
func (m *MockIRedisLayer) GetKeyValue(key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKeyValue", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeyValue indicates an expected call of GetKeyValue.
func (mr *MockIRedisLayerMockRecorder) GetKeyValue(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeyValue", reflect.TypeOf((*MockIRedisLayer)(nil).GetKeyValue), key)
}

// SetKeyValue mocks base method.
func (m *MockIRedisLayer) SetKeyValue(key, value string, exp time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetKeyValue", key, value, exp)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetKeyValue indicates an expected call of SetKeyValue.
func (mr *MockIRedisLayerMockRecorder) SetKeyValue(key, value, exp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetKeyValue", reflect.TypeOf((*MockIRedisLayer)(nil).SetKeyValue), key, value, exp)
}