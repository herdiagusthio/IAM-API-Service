// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	user "iam-api-service/service/user"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: _a0
func (_m *Service) CreateUser(_a0 user.CreateUserData) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(user.CreateUserData) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoginUser provides a mock function with given fields: email, password
func (_m *Service) LoginUser(email string, password string) (string, error) {
	ret := _m.Called(email, password)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
