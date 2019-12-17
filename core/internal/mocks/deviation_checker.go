// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import eth "chainlink/core/eth"
import mock "github.com/stretchr/testify/mock"

// DeviationChecker is an autogenerated mock type for the DeviationChecker type
type DeviationChecker struct {
	mock.Mock
}

// Initialize provides a mock function with given fields: _a0
func (_m *DeviationChecker) Initialize(_a0 eth.Client) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(eth.Client) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields:
func (_m *DeviationChecker) Start() {
	_m.Called()
}

// Stop provides a mock function with given fields:
func (_m *DeviationChecker) Stop() {
	_m.Called()
}