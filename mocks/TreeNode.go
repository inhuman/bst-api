// Code generated by mockery v1.1.2. DO NOT EDIT.

package mocks

import (
	interfaces "github.com/inhuman/bst-api/interfaces"
	mock "github.com/stretchr/testify/mock"
)

// TreeNode is an autogenerated mock type for the TreeNode type
type TreeNode struct {
	mock.Mock
}

// GetKey provides a mock function with given fields:
func (_m *TreeNode) GetKey() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetLeft provides a mock function with given fields:
func (_m *TreeNode) GetLeft() interfaces.TreeNode {
	ret := _m.Called()

	var r0 interfaces.TreeNode
	if rf, ok := ret.Get(0).(func() interfaces.TreeNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.TreeNode)
		}
	}

	return r0
}

// GetRight provides a mock function with given fields:
func (_m *TreeNode) GetRight() interfaces.TreeNode {
	ret := _m.Called()

	var r0 interfaces.TreeNode
	if rf, ok := ret.Get(0).(func() interfaces.TreeNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.TreeNode)
		}
	}

	return r0
}

// GetValue provides a mock function with given fields:
func (_m *TreeNode) GetValue() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// SetKey provides a mock function with given fields: key
func (_m *TreeNode) SetKey(key int) {
	_m.Called(key)
}

// SetLeft provides a mock function with given fields: left
func (_m *TreeNode) SetLeft(left interfaces.TreeNode) {
	_m.Called(left)
}

// SetRight provides a mock function with given fields: right
func (_m *TreeNode) SetRight(right interfaces.TreeNode) {
	_m.Called(right)
}

// SetValue provides a mock function with given fields: value
func (_m *TreeNode) SetValue(value interface{}) {
	_m.Called(value)
}
