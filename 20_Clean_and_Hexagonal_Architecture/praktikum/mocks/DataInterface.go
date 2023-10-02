// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	entity "clean/api/entity"

	mock "github.com/stretchr/testify/mock"
)

// DataInterface is an autogenerated mock type for the DataInterface type
type DataInterface struct {
	mock.Mock
}

// CheckLogin provides a mock function with given fields: email, password
func (_m *DataInterface) CheckLogin(email string, password string) (entity.Main, string, error) {
	ret := _m.Called(email, password)

	var r0 entity.Main
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) (entity.Main, string, error)); ok {
		return rf(email, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) entity.Main); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Get(0).(entity.Main)
	}

	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(email, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Create provides a mock function with given fields: data
func (_m *DataInterface) Create(data entity.Main) (int, error) {
	ret := _m.Called(data)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.Main) (int, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(entity.Main) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(entity.Main) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllUsers provides a mock function with given fields:
func (_m *DataInterface) FindAllUsers() ([]entity.Main, error) {
	ret := _m.Called()

	var r0 []entity.Main
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]entity.Main, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []entity.Main); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Main)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDataInterface creates a new instance of DataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDataInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *DataInterface {
	mock := &DataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}