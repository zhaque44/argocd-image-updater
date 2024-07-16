// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	distribution "github.com/distribution/distribution/v3"
	mock "github.com/stretchr/testify/mock"
)

// Manifest is an autogenerated mock type for the Manifest type
type Manifest struct {
	mock.Mock
}

// Payload provides a mock function with given fields:
func (_m *Manifest) Payload() (string, []byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Payload")
	}

	var r0 string
	var r1 []byte
	var r2 error
	if rf, ok := ret.Get(0).(func() (string, []byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() []byte); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// References provides a mock function with given fields:
func (_m *Manifest) References() []distribution.Descriptor {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for References")
	}

	var r0 []distribution.Descriptor
	if rf, ok := ret.Get(0).(func() []distribution.Descriptor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]distribution.Descriptor)
		}
	}

	return r0
}

// NewManifest creates a new instance of Manifest. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewManifest(t interface {
	mock.TestingT
	Cleanup(func())
}) *Manifest {
	mock := &Manifest{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}