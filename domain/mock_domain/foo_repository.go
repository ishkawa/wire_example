// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ishkawa/wire_example/domain (interfaces: FooRepository)

package mock_domain

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	domain "github.com/ishkawa/wire_example/domain"
	reflect "reflect"
)

// MockFooRepository is a mock of FooRepository interface
type MockFooRepository struct {
	ctrl     *gomock.Controller
	recorder *MockFooRepositoryMockRecorder
}

// MockFooRepositoryMockRecorder is the mock recorder for MockFooRepository
type MockFooRepositoryMockRecorder struct {
	mock *MockFooRepository
}

// NewMockFooRepository creates a new mock instance
func NewMockFooRepository(ctrl *gomock.Controller) *MockFooRepository {
	mock := &MockFooRepository{ctrl: ctrl}
	mock.recorder = &MockFooRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockFooRepository) EXPECT() *MockFooRepositoryMockRecorder {
	return _m.recorder
}

// Get mocks base method
func (_m *MockFooRepository) Get(_param0 context.Context, _param1 int64) (*domain.Foo, error) {
	ret := _m.ctrl.Call(_m, "Get", _param0, _param1)
	ret0, _ := ret[0].(*domain.Foo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (_mr *MockFooRepositoryMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Get", reflect.TypeOf((*MockFooRepository)(nil).Get), arg0, arg1)
}

// Put mocks base method
func (_m *MockFooRepository) Put(_param0 context.Context, _param1 *domain.Foo) error {
	ret := _m.ctrl.Call(_m, "Put", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put
func (_mr *MockFooRepositoryMockRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Put", reflect.TypeOf((*MockFooRepository)(nil).Put), arg0, arg1)
}
