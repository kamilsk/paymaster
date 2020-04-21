// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kamilsk/guard/pkg/transport/grpc/rpc (interfaces: Storage)

// Package rpc_test is a generated GoMock package.
package rpc_test

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	types "github.com/kamilsk/guard/pkg/service/types"
	query "github.com/kamilsk/guard/pkg/storage/query"
	types0 "github.com/kamilsk/guard/pkg/storage/types"
)

// MockStorage is a mock of Storage interface
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// AddEmployee mocks base method
func (m *MockStorage) AddEmployee(arg0 context.Context, arg1 types.Token, arg2 query.LicenseEmployee) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEmployee", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEmployee indicates an expected call of AddEmployee
func (mr *MockStorageMockRecorder) AddEmployee(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEmployee", reflect.TypeOf((*MockStorage)(nil).AddEmployee), arg0, arg1, arg2)
}

// AddWorkplace mocks base method
func (m *MockStorage) AddWorkplace(arg0 context.Context, arg1 types.Token, arg2 query.LicenseWorkplace) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddWorkplace", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddWorkplace indicates an expected call of AddWorkplace
func (mr *MockStorageMockRecorder) AddWorkplace(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWorkplace", reflect.TypeOf((*MockStorage)(nil).AddWorkplace), arg0, arg1, arg2)
}

// CreateLicense mocks base method
func (m *MockStorage) CreateLicense(arg0 context.Context, arg1 types.Token, arg2 query.CreateLicense) (types0.License, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLicense", arg0, arg1, arg2)
	ret0, _ := ret[0].(types0.License)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLicense indicates an expected call of CreateLicense
func (mr *MockStorageMockRecorder) CreateLicense(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLicense", reflect.TypeOf((*MockStorage)(nil).CreateLicense), arg0, arg1, arg2)
}

// DeleteEmployee mocks base method
func (m *MockStorage) DeleteEmployee(arg0 context.Context, arg1 types.Token, arg2 query.LicenseEmployee) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEmployee", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEmployee indicates an expected call of DeleteEmployee
func (mr *MockStorageMockRecorder) DeleteEmployee(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEmployee", reflect.TypeOf((*MockStorage)(nil).DeleteEmployee), arg0, arg1, arg2)
}

// DeleteLicense mocks base method
func (m *MockStorage) DeleteLicense(arg0 context.Context, arg1 types.Token, arg2 query.DeleteLicense) (types0.License, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLicense", arg0, arg1, arg2)
	ret0, _ := ret[0].(types0.License)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteLicense indicates an expected call of DeleteLicense
func (mr *MockStorageMockRecorder) DeleteLicense(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLicense", reflect.TypeOf((*MockStorage)(nil).DeleteLicense), arg0, arg1, arg2)
}

// DeleteWorkplace mocks base method
func (m *MockStorage) DeleteWorkplace(arg0 context.Context, arg1 types.Token, arg2 query.LicenseWorkplace) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWorkplace", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWorkplace indicates an expected call of DeleteWorkplace
func (mr *MockStorageMockRecorder) DeleteWorkplace(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWorkplace", reflect.TypeOf((*MockStorage)(nil).DeleteWorkplace), arg0, arg1, arg2)
}

// LicenseEmployees mocks base method
func (m *MockStorage) LicenseEmployees(arg0 context.Context, arg1 types.Token, arg2 query.EmployeeList) ([]types0.Employee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LicenseEmployees", arg0, arg1, arg2)
	ret0, _ := ret[0].([]types0.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LicenseEmployees indicates an expected call of LicenseEmployees
func (mr *MockStorageMockRecorder) LicenseEmployees(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LicenseEmployees", reflect.TypeOf((*MockStorage)(nil).LicenseEmployees), arg0, arg1, arg2)
}

// LicenseWorkplaces mocks base method
func (m *MockStorage) LicenseWorkplaces(arg0 context.Context, arg1 types.Token, arg2 query.WorkplaceList) ([]types0.Workplace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LicenseWorkplaces", arg0, arg1, arg2)
	ret0, _ := ret[0].([]types0.Workplace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LicenseWorkplaces indicates an expected call of LicenseWorkplaces
func (mr *MockStorageMockRecorder) LicenseWorkplaces(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LicenseWorkplaces", reflect.TypeOf((*MockStorage)(nil).LicenseWorkplaces), arg0, arg1, arg2)
}

// PushWorkplace mocks base method
func (m *MockStorage) PushWorkplace(arg0 context.Context, arg1 types.Token, arg2 query.LicenseWorkplace) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushWorkplace", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushWorkplace indicates an expected call of PushWorkplace
func (mr *MockStorageMockRecorder) PushWorkplace(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushWorkplace", reflect.TypeOf((*MockStorage)(nil).PushWorkplace), arg0, arg1, arg2)
}

// ReadLicense mocks base method
func (m *MockStorage) ReadLicense(arg0 context.Context, arg1 types.Token, arg2 query.ReadLicense) (types0.License, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadLicense", arg0, arg1, arg2)
	ret0, _ := ret[0].(types0.License)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadLicense indicates an expected call of ReadLicense
func (mr *MockStorageMockRecorder) ReadLicense(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadLicense", reflect.TypeOf((*MockStorage)(nil).ReadLicense), arg0, arg1, arg2)
}

// RegisterLicense mocks base method
func (m *MockStorage) RegisterLicense(arg0 context.Context, arg1 types.Token, arg2 query.RegisterLicense) (types0.License, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterLicense", arg0, arg1, arg2)
	ret0, _ := ret[0].(types0.License)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterLicense indicates an expected call of RegisterLicense
func (mr *MockStorageMockRecorder) RegisterLicense(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterLicense", reflect.TypeOf((*MockStorage)(nil).RegisterLicense), arg0, arg1, arg2)
}

// RestoreLicense mocks base method
func (m *MockStorage) RestoreLicense(arg0 context.Context, arg1 types.Token, arg2 query.RestoreLicense) (types0.License, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreLicense", arg0, arg1, arg2)
	ret0, _ := ret[0].(types0.License)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RestoreLicense indicates an expected call of RestoreLicense
func (mr *MockStorageMockRecorder) RestoreLicense(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreLicense", reflect.TypeOf((*MockStorage)(nil).RestoreLicense), arg0, arg1, arg2)
}

// UpdateLicense mocks base method
func (m *MockStorage) UpdateLicense(arg0 context.Context, arg1 types.Token, arg2 query.UpdateLicense) (types0.License, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLicense", arg0, arg1, arg2)
	ret0, _ := ret[0].(types0.License)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateLicense indicates an expected call of UpdateLicense
func (mr *MockStorageMockRecorder) UpdateLicense(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLicense", reflect.TypeOf((*MockStorage)(nil).UpdateLicense), arg0, arg1, arg2)
}
