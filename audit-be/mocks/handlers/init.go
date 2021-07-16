// Code generated by MockGen. DO NOT EDIT.
// Source: handlers/init.go

// Package mock_handlers is a generated GoMock package.
package mock_handlers

import (
	reflect "reflect"

	datatransfers "github.com/daystram/audit/audit-be/datatransfers"
	gomock "github.com/golang/mock/gomock"
)

// MockHandlerFunc is a mock of HandlerFunc interface.
type MockHandlerFunc struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerFuncMockRecorder
}

// MockHandlerFuncMockRecorder is the mock recorder for MockHandlerFunc.
type MockHandlerFuncMockRecorder struct {
	mock *MockHandlerFunc
}

// NewMockHandlerFunc creates a new mock instance.
func NewMockHandlerFunc(ctrl *gomock.Controller) *MockHandlerFunc {
	mock := &MockHandlerFunc{ctrl: ctrl}
	mock.recorder = &MockHandlerFuncMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandlerFunc) EXPECT() *MockHandlerFuncMockRecorder {
	return m.recorder
}

// ApplicationCreate mocks base method.
func (m *MockHandlerFunc) ApplicationCreate(applicationInfo datatransfers.ApplicationInfo) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationCreate", applicationInfo)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationCreate indicates an expected call of ApplicationCreate.
func (mr *MockHandlerFuncMockRecorder) ApplicationCreate(applicationInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationCreate", reflect.TypeOf((*MockHandlerFunc)(nil).ApplicationCreate), applicationInfo)
}

// ApplicationDelete mocks base method.
func (m *MockHandlerFunc) ApplicationDelete(applicationID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationDelete", applicationID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplicationDelete indicates an expected call of ApplicationDelete.
func (mr *MockHandlerFuncMockRecorder) ApplicationDelete(applicationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationDelete", reflect.TypeOf((*MockHandlerFunc)(nil).ApplicationDelete), applicationID)
}

// ApplicationGetAll mocks base method.
func (m *MockHandlerFunc) ApplicationGetAll() ([]datatransfers.ApplicationInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationGetAll")
	ret0, _ := ret[0].([]datatransfers.ApplicationInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationGetAll indicates an expected call of ApplicationGetAll.
func (mr *MockHandlerFuncMockRecorder) ApplicationGetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationGetAll", reflect.TypeOf((*MockHandlerFunc)(nil).ApplicationGetAll))
}

// ApplicationGetOne mocks base method.
func (m *MockHandlerFunc) ApplicationGetOne(applicationID string) (datatransfers.ApplicationInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationGetOne", applicationID)
	ret0, _ := ret[0].(datatransfers.ApplicationInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationGetOne indicates an expected call of ApplicationGetOne.
func (mr *MockHandlerFuncMockRecorder) ApplicationGetOne(applicationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationGetOne", reflect.TypeOf((*MockHandlerFunc)(nil).ApplicationGetOne), applicationID)
}

// ApplicationUpdate mocks base method.
func (m *MockHandlerFunc) ApplicationUpdate(applicationInfo datatransfers.ApplicationInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationUpdate", applicationInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplicationUpdate indicates an expected call of ApplicationUpdate.
func (mr *MockHandlerFuncMockRecorder) ApplicationUpdate(applicationInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationUpdate", reflect.TypeOf((*MockHandlerFunc)(nil).ApplicationUpdate), applicationInfo)
}

// ServiceCreate mocks base method.
func (m *MockHandlerFunc) ServiceCreate(serviceInfo datatransfers.ServiceInfo) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceCreate", serviceInfo)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServiceCreate indicates an expected call of ServiceCreate.
func (mr *MockHandlerFuncMockRecorder) ServiceCreate(serviceInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceCreate", reflect.TypeOf((*MockHandlerFunc)(nil).ServiceCreate), serviceInfo)
}

// ServiceDelete mocks base method.
func (m *MockHandlerFunc) ServiceDelete(serviceID, applicationID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceDelete", serviceID, applicationID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ServiceDelete indicates an expected call of ServiceDelete.
func (mr *MockHandlerFuncMockRecorder) ServiceDelete(serviceID, applicationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceDelete", reflect.TypeOf((*MockHandlerFunc)(nil).ServiceDelete), serviceID, applicationID)
}

// ServiceGetAll mocks base method.
func (m *MockHandlerFunc) ServiceGetAll(applicationID string) ([]datatransfers.ServiceInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceGetAll", applicationID)
	ret0, _ := ret[0].([]datatransfers.ServiceInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServiceGetAll indicates an expected call of ServiceGetAll.
func (mr *MockHandlerFuncMockRecorder) ServiceGetAll(applicationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceGetAll", reflect.TypeOf((*MockHandlerFunc)(nil).ServiceGetAll), applicationID)
}

// ServiceGetOne mocks base method.
func (m *MockHandlerFunc) ServiceGetOne(serviceID, applicationID string) (datatransfers.ServiceInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceGetOne", serviceID, applicationID)
	ret0, _ := ret[0].(datatransfers.ServiceInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServiceGetOne indicates an expected call of ServiceGetOne.
func (mr *MockHandlerFuncMockRecorder) ServiceGetOne(serviceID, applicationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceGetOne", reflect.TypeOf((*MockHandlerFunc)(nil).ServiceGetOne), serviceID, applicationID)
}

// ServiceUpdate mocks base method.
func (m *MockHandlerFunc) ServiceUpdate(serviceInfo datatransfers.ServiceInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceUpdate", serviceInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// ServiceUpdate indicates an expected call of ServiceUpdate.
func (mr *MockHandlerFuncMockRecorder) ServiceUpdate(serviceInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceUpdate", reflect.TypeOf((*MockHandlerFunc)(nil).ServiceUpdate), serviceInfo)
}
