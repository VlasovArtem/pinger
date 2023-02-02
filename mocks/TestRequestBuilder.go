// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	resthelper "github.com/VlasovArtem/pinger/src/test/resthelper"
	mock "github.com/stretchr/testify/mock"
)

// TestRequestBuilder is an autogenerated mock type for the TestRequestBuilder type
type TestRequestBuilder struct {
	mock.Mock
}

// Build provides a mock function with given fields:
func (_m *TestRequestBuilder) Build() *resthelper.TestRequest {
	ret := _m.Called()

	var r0 *resthelper.TestRequest
	if rf, ok := ret.Get(0).(func() *resthelper.TestRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resthelper.TestRequest)
		}
	}

	return r0
}

// WithBody provides a mock function with given fields: body
func (_m *TestRequestBuilder) WithBody(body interface{}) *resthelper.TestRequest {
	ret := _m.Called(body)

	var r0 *resthelper.TestRequest
	if rf, ok := ret.Get(0).(func(interface{}) *resthelper.TestRequest); ok {
		r0 = rf(body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resthelper.TestRequest)
		}
	}

	return r0
}

// WithHandler provides a mock function with given fields: handler
func (_m *TestRequestBuilder) WithHandler(handler http.HandlerFunc) *resthelper.TestRequest {
	ret := _m.Called(handler)

	var r0 *resthelper.TestRequest
	if rf, ok := ret.Get(0).(func(http.HandlerFunc) *resthelper.TestRequest); ok {
		r0 = rf(handler)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resthelper.TestRequest)
		}
	}

	return r0
}

// WithMethod provides a mock function with given fields: method
func (_m *TestRequestBuilder) WithMethod(method string) *resthelper.TestRequest {
	ret := _m.Called(method)

	var r0 *resthelper.TestRequest
	if rf, ok := ret.Get(0).(func(string) *resthelper.TestRequest); ok {
		r0 = rf(method)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resthelper.TestRequest)
		}
	}

	return r0
}

// WithParameter provides a mock function with given fields: key, value
func (_m *TestRequestBuilder) WithParameter(key string, value string) *resthelper.TestRequest {
	ret := _m.Called(key, value)

	var r0 *resthelper.TestRequest
	if rf, ok := ret.Get(0).(func(string, string) *resthelper.TestRequest); ok {
		r0 = rf(key, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resthelper.TestRequest)
		}
	}

	return r0
}

// WithURL provides a mock function with given fields: target
func (_m *TestRequestBuilder) WithURL(target string) *resthelper.TestRequest {
	ret := _m.Called(target)

	var r0 *resthelper.TestRequest
	if rf, ok := ret.Get(0).(func(string) *resthelper.TestRequest); ok {
		r0 = rf(target)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resthelper.TestRequest)
		}
	}

	return r0
}

// WithVar provides a mock function with given fields: key, value
func (_m *TestRequestBuilder) WithVar(key string, value string) *resthelper.TestRequest {
	ret := _m.Called(key, value)

	var r0 *resthelper.TestRequest
	if rf, ok := ret.Get(0).(func(string, string) *resthelper.TestRequest); ok {
		r0 = rf(key, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resthelper.TestRequest)
		}
	}

	return r0
}

type mockConstructorTestingTNewTestRequestBuilder interface {
	mock.TestingT
	Cleanup(func())
}

// NewTestRequestBuilder creates a new instance of TestRequestBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTestRequestBuilder(t mockConstructorTestingTNewTestRequestBuilder) *TestRequestBuilder {
	mock := &TestRequestBuilder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}