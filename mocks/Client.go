// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mongo "github.com/royw835/mongo"
	mock "github.com/stretchr/testify/mock"

	mongo_drivermongo "go.mongodb.org/mongo-driver/mongo"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Connect provides a mock function with given fields: _a0
func (_m *Client) Connect(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Database provides a mock function with given fields: _a0
func (_m *Client) Database(_a0 string) mongo.Database {
	ret := _m.Called(_a0)

	var r0 mongo.Database
	if rf, ok := ret.Get(0).(func(string) mongo.Database); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongo.Database)
		}
	}

	return r0
}

// Disconnect provides a mock function with given fields: _a0
func (_m *Client) Disconnect(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Ping provides a mock function with given fields: _a0
func (_m *Client) Ping(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartSession provides a mock function with given fields:
func (_m *Client) StartSession() (mongo_drivermongo.Session, error) {
	ret := _m.Called()

	var r0 mongo_drivermongo.Session
	if rf, ok := ret.Get(0).(func() mongo_drivermongo.Session); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongo_drivermongo.Session)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UseSession provides a mock function with given fields: ctx, fn
func (_m *Client) UseSession(ctx context.Context, fn func(mongo_drivermongo.SessionContext) error) error {
	ret := _m.Called(ctx, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(mongo_drivermongo.SessionContext) error) error); ok {
		r0 = rf(ctx, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
