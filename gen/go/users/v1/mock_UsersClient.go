// Code generated by mockery v1.0.0. DO NOT EDIT.

package v1pb

import context "context"
import grpc "google.golang.org/grpc"
import mock "github.com/stretchr/testify/mock"

// MockUsersClient is an autogenerated mock type for the UsersClient type
type MockUsersClient struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, in, opts
func (_m *MockUsersClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*User, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *User
	if rf, ok := ret.Get(0).(func(context.Context, *CreateRequest, ...grpc.CallOption) *User); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *CreateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, in, opts
func (_m *MockUsersClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*User, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *User
	if rf, ok := ret.Get(0).(func(context.Context, *GetRequest, ...grpc.CallOption) *User); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
