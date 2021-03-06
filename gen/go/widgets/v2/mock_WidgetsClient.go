// Code generated by mockery v1.0.0. DO NOT EDIT.

package v2pb

import context "context"
import empty "github.com/golang/protobuf/ptypes/empty"
import grpc "google.golang.org/grpc"
import mock "github.com/stretchr/testify/mock"

// MockWidgetsClient is an autogenerated mock type for the WidgetsClient type
type MockWidgetsClient struct {
	mock.Mock
}

// BatchGet provides a mock function with given fields: ctx, in, opts
func (_m *MockWidgetsClient) BatchGet(ctx context.Context, in *BatchGetRequest, opts ...grpc.CallOption) (*BatchGetResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *BatchGetResponse
	if rf, ok := ret.Get(0).(func(context.Context, *BatchGetRequest, ...grpc.CallOption) *BatchGetResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*BatchGetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *BatchGetRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, in, opts
func (_m *MockWidgetsClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Widget, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *Widget
	if rf, ok := ret.Get(0).(func(context.Context, *CreateRequest, ...grpc.CallOption) *Widget); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Widget)
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

// Delete provides a mock function with given fields: ctx, in, opts
func (_m *MockWidgetsClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *empty.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *DeleteRequest, ...grpc.CallOption) *empty.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*empty.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *DeleteRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, in, opts
func (_m *MockWidgetsClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Widget, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *Widget
	if rf, ok := ret.Get(0).(func(context.Context, *GetRequest, ...grpc.CallOption) *Widget); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Widget)
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

// List provides a mock function with given fields: ctx, in, opts
func (_m *MockWidgetsClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ListResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ListRequest, ...grpc.CallOption) *ListResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ListResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ListRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, in, opts
func (_m *MockWidgetsClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Widget, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *Widget
	if rf, ok := ret.Get(0).(func(context.Context, *UpdateRequest, ...grpc.CallOption) *Widget); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Widget)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *UpdateRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
