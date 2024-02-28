// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// RistrettoCache is an autogenerated mock type for the RistrettoCache type
type RistrettoCache struct {
	mock.Mock
}

// Del provides a mock function with given fields: key
func (_m *RistrettoCache) Del(key ...string) {
	_va := make([]interface{}, len(key))
	for _i := range key {
		_va[_i] = key[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Get provides a mock function with given fields: key
func (_m *RistrettoCache) Get(key string) interface{} {
	ret := _m.Called(key)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// Set provides a mock function with given fields: key, value
func (_m *RistrettoCache) Set(key string, value interface{}) {
	_m.Called(key, value)
}

type mockConstructorTestingTNewRistrettoCache interface {
	mock.TestingT
	Cleanup(func())
}

// NewRistrettoCache creates a new instance of RistrettoCache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRistrettoCache(t mockConstructorTestingTNewRistrettoCache) *RistrettoCache {
	mock := &RistrettoCache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
