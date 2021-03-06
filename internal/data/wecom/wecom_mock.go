// Code generated by MockGen. DO NOT EDIT.
// Source: wecom.go

// Package wecom is a generated GoMock package.
package wecom

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockWeCom is a mock of WeCom interface.
type MockWeCom struct {
	ctrl     *gomock.Controller
	recorder *MockWeComMockRecorder
}

// MockWeComMockRecorder is the mock recorder for MockWeCom.
type MockWeComMockRecorder struct {
	mock *MockWeCom
}

// NewMockWeCom creates a new mock instance.
func NewMockWeCom(ctrl *gomock.Controller) *MockWeCom {
	mock := &MockWeCom{ctrl: ctrl}
	mock.recorder = &MockWeComMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWeCom) EXPECT() *MockWeComMockRecorder {
	return m.recorder
}

// LoginUrl mocks base method.
func (m *MockWeCom) LoginUrl(redirectUrl string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUrl", redirectUrl)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginUrl indicates an expected call of LoginUrl.
func (mr *MockWeComMockRecorder) LoginUrl(redirectUrl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUrl", reflect.TypeOf((*MockWeCom)(nil).LoginUrl), redirectUrl)
}
