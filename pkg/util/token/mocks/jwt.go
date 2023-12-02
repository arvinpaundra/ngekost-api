// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	token "github.com/arvinpaundra/ngekost-api/pkg/util/token"
	mock "github.com/stretchr/testify/mock"
)

// JSONWebToken is an autogenerated mock type for the JSONWebToken type
type JSONWebToken struct {
	mock.Mock
}

// Decode provides a mock function with given fields: _a0
func (_m *JSONWebToken) Decode(_a0 string) (*token.JWTCustomClaim, error) {
	ret := _m.Called(_a0)

	var r0 *token.JWTCustomClaim
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*token.JWTCustomClaim, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) *token.JWTCustomClaim); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*token.JWTCustomClaim)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Encode provides a mock function with given fields: claims
func (_m *JSONWebToken) Encode(claims *token.JWTCustomClaim) (string, error) {
	ret := _m.Called(claims)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*token.JWTCustomClaim) (string, error)); ok {
		return rf(claims)
	}
	if rf, ok := ret.Get(0).(func(*token.JWTCustomClaim) string); ok {
		r0 = rf(claims)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*token.JWTCustomClaim) error); ok {
		r1 = rf(claims)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewJSONWebToken creates a new instance of JSONWebToken. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJSONWebToken(t interface {
	mock.TestingT
	Cleanup(func())
}) *JSONWebToken {
	mock := &JSONWebToken{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}