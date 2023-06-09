// Code generated by MockGen. DO NOT EDIT.
// Source: role.go

// Package mbiz is a generated GoMock package.
package mbiz

import (
	context "context"
	reflect "reflect"

	v1 "github.com/869413421/micro-chat/api/user/service/v1"
	biz "github.com/869413421/micro-chat/app/user/service/internal/biz"
	gomock "github.com/golang/mock/gomock"
)

// MockRoleRepo is a mock of RoleRepo interface.
type MockRoleRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRoleRepoMockRecorder
}

// MockRoleRepoMockRecorder is the mock recorder for MockRoleRepo.
type MockRoleRepoMockRecorder struct {
	mock *MockRoleRepo
}

// NewMockRoleRepo creates a new mock instance.
func NewMockRoleRepo(ctrl *gomock.Controller) *MockRoleRepo {
	mock := &MockRoleRepo{ctrl: ctrl}
	mock.recorder = &MockRoleRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleRepo) EXPECT() *MockRoleRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRoleRepo) Create(arg0 context.Context, arg1 *biz.Role) (*biz.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*biz.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRoleRepoMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRoleRepo)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockRoleRepo) Delete(arg0 context.Context, arg1 uint64) (*biz.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*biz.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockRoleRepoMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRoleRepo)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockRoleRepo) Get(arg0 context.Context, arg1 map[string]interface{}) (*biz.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*biz.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRoleRepoMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRoleRepo)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockRoleRepo) List(arg0 context.Context, arg1 map[string]interface{}, arg2 map[string]bool, arg3, arg4 int64) ([]*biz.Role, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*biz.Role)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockRoleRepoMockRecorder) List(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRoleRepo)(nil).List), arg0, arg1, arg2, arg3, arg4)
}

// Update mocks base method.
func (m *MockRoleRepo) Update(arg0 context.Context, arg1 *biz.Role) (*biz.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*biz.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRoleRepoMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRoleRepo)(nil).Update), arg0, arg1)
}

// MockRoleUsecase is a mock of RoleUsecase interface.
type MockRoleUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockRoleUsecaseMockRecorder
}

// MockRoleUsecaseMockRecorder is the mock recorder for MockRoleUsecase.
type MockRoleUsecaseMockRecorder struct {
	mock *MockRoleUsecase
}

// NewMockRoleUsecase creates a new mock instance.
func NewMockRoleUsecase(ctrl *gomock.Controller) *MockRoleUsecase {
	mock := &MockRoleUsecase{ctrl: ctrl}
	mock.recorder = &MockRoleUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoleUsecase) EXPECT() *MockRoleUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRoleUsecase) Create(arg0 context.Context, arg1 *v1.CreateRoleRequest) (*biz.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*biz.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRoleUsecaseMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRoleUsecase)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockRoleUsecase) Delete(arg0 context.Context, arg1 uint64) (*biz.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*biz.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockRoleUsecaseMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRoleUsecase)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockRoleUsecase) Get(arg0 context.Context, arg1 *v1.GetRoleRequest) (*biz.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*biz.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRoleUsecaseMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRoleUsecase)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockRoleUsecase) List(arg0 context.Context, arg1 *v1.ListRoleRequest) ([]*biz.Role, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*biz.Role)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockRoleUsecaseMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRoleUsecase)(nil).List), arg0, arg1)
}

// Update mocks base method.
func (m *MockRoleUsecase) Update(arg0 context.Context, arg1 *v1.UpdateRoleRequest) (*biz.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*biz.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockRoleUsecaseMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRoleUsecase)(nil).Update), arg0, arg1)
}
