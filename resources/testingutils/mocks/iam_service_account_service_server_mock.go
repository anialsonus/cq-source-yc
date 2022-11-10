// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1 (interfaces: ServiceAccountServiceServer)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	access "github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	iam "github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	reflect "reflect"
)

// MockServiceAccountServiceServer is a mock of ServiceAccountServiceServer interface
type MockServiceAccountServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockServiceAccountServiceServerMockRecorder
}

// MockServiceAccountServiceServerMockRecorder is the mock recorder for MockServiceAccountServiceServer
type MockServiceAccountServiceServerMockRecorder struct {
	mock *MockServiceAccountServiceServer
}

// NewMockServiceAccountServiceServer creates a new mock instance
func NewMockServiceAccountServiceServer(ctrl *gomock.Controller) *MockServiceAccountServiceServer {
	mock := &MockServiceAccountServiceServer{ctrl: ctrl}
	mock.recorder = &MockServiceAccountServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceAccountServiceServer) EXPECT() *MockServiceAccountServiceServerMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockServiceAccountServiceServer) Create(arg0 context.Context, arg1 *iam.CreateServiceAccountRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockServiceAccountServiceServerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockServiceAccountServiceServer)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockServiceAccountServiceServer) Delete(arg0 context.Context, arg1 *iam.DeleteServiceAccountRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockServiceAccountServiceServerMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServiceAccountServiceServer)(nil).Delete), arg0, arg1)
}

// Get mocks base method
func (m *MockServiceAccountServiceServer) Get(arg0 context.Context, arg1 *iam.GetServiceAccountRequest) (*iam.ServiceAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*iam.ServiceAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockServiceAccountServiceServerMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockServiceAccountServiceServer)(nil).Get), arg0, arg1)
}

// List mocks base method
func (m *MockServiceAccountServiceServer) List(arg0 context.Context, arg1 *iam.ListServiceAccountsRequest) (*iam.ListServiceAccountsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*iam.ListServiceAccountsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockServiceAccountServiceServerMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockServiceAccountServiceServer)(nil).List), arg0, arg1)
}

// ListAccessBindings mocks base method
func (m *MockServiceAccountServiceServer) ListAccessBindings(arg0 context.Context, arg1 *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccessBindings", arg0, arg1)
	ret0, _ := ret[0].(*access.ListAccessBindingsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccessBindings indicates an expected call of ListAccessBindings
func (mr *MockServiceAccountServiceServerMockRecorder) ListAccessBindings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccessBindings", reflect.TypeOf((*MockServiceAccountServiceServer)(nil).ListAccessBindings), arg0, arg1)
}

// ListOperations mocks base method
func (m *MockServiceAccountServiceServer) ListOperations(arg0 context.Context, arg1 *iam.ListServiceAccountOperationsRequest) (*iam.ListServiceAccountOperationsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOperations", arg0, arg1)
	ret0, _ := ret[0].(*iam.ListServiceAccountOperationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOperations indicates an expected call of ListOperations
func (mr *MockServiceAccountServiceServerMockRecorder) ListOperations(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOperations", reflect.TypeOf((*MockServiceAccountServiceServer)(nil).ListOperations), arg0, arg1)
}

// SetAccessBindings mocks base method
func (m *MockServiceAccountServiceServer) SetAccessBindings(arg0 context.Context, arg1 *access.SetAccessBindingsRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAccessBindings", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetAccessBindings indicates an expected call of SetAccessBindings
func (mr *MockServiceAccountServiceServerMockRecorder) SetAccessBindings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccessBindings", reflect.TypeOf((*MockServiceAccountServiceServer)(nil).SetAccessBindings), arg0, arg1)
}

// Update mocks base method
func (m *MockServiceAccountServiceServer) Update(arg0 context.Context, arg1 *iam.UpdateServiceAccountRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockServiceAccountServiceServerMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockServiceAccountServiceServer)(nil).Update), arg0, arg1)
}

// UpdateAccessBindings mocks base method
func (m *MockServiceAccountServiceServer) UpdateAccessBindings(arg0 context.Context, arg1 *access.UpdateAccessBindingsRequest) (*operation.Operation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccessBindings", arg0, arg1)
	ret0, _ := ret[0].(*operation.Operation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccessBindings indicates an expected call of UpdateAccessBindings
func (mr *MockServiceAccountServiceServerMockRecorder) UpdateAccessBindings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccessBindings", reflect.TypeOf((*MockServiceAccountServiceServer)(nil).UpdateAccessBindings), arg0, arg1)
}
