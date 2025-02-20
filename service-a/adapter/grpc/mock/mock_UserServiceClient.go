// Code generated by MockGen. DO NOT EDIT.
// Source: github/herochi/orbi/service-a/adapter/grpc/user (interfaces: UserServiceClient)

// Package UserServiceClient_mock is a generated GoMock package.
package UserServiceClient_mock

import (
	context "context"
	user "github/herochi/orbi/service-a/adapter/grpc/user"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockUserServiceClient is a mock of UserServiceClient interface.
type MockUserServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceClientMockRecorder
}

// MockUserServiceClientMockRecorder is the mock recorder for MockUserServiceClient.
type MockUserServiceClientMockRecorder struct {
	mock *MockUserServiceClient
}

// NewMockUserServiceClient creates a new mock instance.
func NewMockUserServiceClient(ctrl *gomock.Controller) *MockUserServiceClient {
	mock := &MockUserServiceClient{ctrl: ctrl}
	mock.recorder = &MockUserServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceClient) EXPECT() *MockUserServiceClientMockRecorder {
	return m.recorder
}

// NotifyUser mocks base method.
func (m *MockUserServiceClient) NotifyUser(arg0 context.Context, arg1 *user.Request, arg2 ...grpc.CallOption) (*user.NotifyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NotifyUser", varargs...)
	ret0, _ := ret[0].(*user.NotifyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotifyUser indicates an expected call of NotifyUser.
func (mr *MockUserServiceClientMockRecorder) NotifyUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyUser", reflect.TypeOf((*MockUserServiceClient)(nil).NotifyUser), varargs...)
}
