// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/router/router.go

// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	middleware "github.com/hasemeneh/poc-material/poc-pricing/pkg/middleware"
	router "github.com/hasemeneh/poc-material/poc-pricing/pkg/router"
)

// MockHTTPRouter is a mock of HTTPRouter interface.
type MockHTTPRouter struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPRouterMockRecorder
}

// MockHTTPRouterMockRecorder is the mock recorder for MockHTTPRouter.
type MockHTTPRouterMockRecorder struct {
	mock *MockHTTPRouter
}

// NewMockHTTPRouter creates a new mock instance.
func NewMockHTTPRouter(ctrl *gomock.Controller) *MockHTTPRouter {
	mock := &MockHTTPRouter{ctrl: ctrl}
	mock.recorder = &MockHTTPRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPRouter) EXPECT() *MockHTTPRouterMockRecorder {
	return m.recorder
}

// GET mocks base method.
func (m *MockHTTPRouter) GET(path string, handler router.HandlerFunc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GET", path, handler)
}

// GET indicates an expected call of GET.
func (mr *MockHTTPRouterMockRecorder) GET(path, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GET", reflect.TypeOf((*MockHTTPRouter)(nil).GET), path, handler)
}

// OPTIONS mocks base method.
func (m *MockHTTPRouter) OPTIONS(path string, handler router.HandlerFunc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OPTIONS", path, handler)
}

// OPTIONS indicates an expected call of OPTIONS.
func (mr *MockHTTPRouterMockRecorder) OPTIONS(path, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OPTIONS", reflect.TypeOf((*MockHTTPRouter)(nil).OPTIONS), path, handler)
}

// POST mocks base method.
func (m *MockHTTPRouter) POST(path string, handler router.HandlerFunc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "POST", path, handler)
}

// POST indicates an expected call of POST.
func (mr *MockHTTPRouterMockRecorder) POST(path, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "POST", reflect.TypeOf((*MockHTTPRouter)(nil).POST), path, handler)
}

// PUT mocks base method.
func (m *MockHTTPRouter) PUT(path string, handler router.HandlerFunc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PUT", path, handler)
}

// PUT indicates an expected call of PUT.
func (mr *MockHTTPRouterMockRecorder) PUT(path, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PUT", reflect.TypeOf((*MockHTTPRouter)(nil).PUT), path, handler)
}

// MockRegistrator is a mock of Registrator interface.
type MockRegistrator struct {
	ctrl     *gomock.Controller
	recorder *MockRegistratorMockRecorder
}

// MockRegistratorMockRecorder is the mock recorder for MockRegistrator.
type MockRegistratorMockRecorder struct {
	mock *MockRegistrator
}

// NewMockRegistrator creates a new mock instance.
func NewMockRegistrator(ctrl *gomock.Controller) *MockRegistrator {
	mock := &MockRegistrator{ctrl: ctrl}
	mock.recorder = &MockRegistratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistrator) EXPECT() *MockRegistratorMockRecorder {
	return m.recorder
}

// AddMiddlewareWrapper mocks base method.
func (m *MockRegistrator) AddMiddlewareWrapper(wrapper ...middleware.Wrapper) router.Registrator {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range wrapper {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddMiddlewareWrapper", varargs...)
	ret0, _ := ret[0].(router.Registrator)
	return ret0
}

// AddMiddlewareWrapper indicates an expected call of AddMiddlewareWrapper.
func (mr *MockRegistratorMockRecorder) AddMiddlewareWrapper(wrapper ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMiddlewareWrapper", reflect.TypeOf((*MockRegistrator)(nil).AddMiddlewareWrapper), wrapper...)
}

// Register mocks base method.
func (m *MockRegistrator) Register(method, path string, handler router.HandlerFunc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Register", method, path, handler)
}

// Register indicates an expected call of Register.
func (mr *MockRegistratorMockRecorder) Register(method, path, handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockRegistrator)(nil).Register), method, path, handler)
}

// ServeHTTP mocks base method.
func (m *MockRegistrator) ServeHTTP(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ServeHTTP", arg0, arg1)
}

// ServeHTTP indicates an expected call of ServeHTTP.
func (mr *MockRegistratorMockRecorder) ServeHTTP(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServeHTTP", reflect.TypeOf((*MockRegistrator)(nil).ServeHTTP), arg0, arg1)
}
