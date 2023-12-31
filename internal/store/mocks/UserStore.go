// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/merge/shopping-card/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// UserStore is an autogenerated mock type for the UserStore type
type UserStore struct {
	mock.Mock
}

// FindByEmail provides a mock function with given fields: ctx, email
func (_m *UserStore) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	ret := _m.Called(ctx, email)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByUserID provides a mock function with given fields: ctx, userID
func (_m *UserStore) FindByUserID(ctx context.Context, userID int) (*model.User, error) {
	ret := _m.Called(ctx, userID)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*model.User, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *model.User); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, u
func (_m *UserStore) Save(ctx context.Context, u *model.User) error {
	ret := _m.Called(ctx, u)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) error); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserStore creates a new instance of UserStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserStore {
	mock := &UserStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
