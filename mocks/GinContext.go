// Code generated by mockery v1.1.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// GinContext is an autogenerated mock type for the GinContext type
type GinContext struct {
	mock.Mock
}

// AbortWithStatus provides a mock function with given fields: code
func (_m *GinContext) AbortWithStatus(code int) {
	_m.Called(code)
}

// AbortWithStatusJSON provides a mock function with given fields: code, jsonObj
func (_m *GinContext) AbortWithStatusJSON(code int, jsonObj interface{}) {
	_m.Called(code, jsonObj)
}

// BindJSON provides a mock function with given fields: obj
func (_m *GinContext) BindJSON(obj interface{}) error {
	ret := _m.Called(obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Query provides a mock function with given fields: key
func (_m *GinContext) Query(key string) string {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}