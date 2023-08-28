// Code generated by mockery v2.33.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "avito-backend/internal/app/model"
)

// IService is an autogenerated mock type for the IService type
type IService struct {
	mock.Mock
}

// AddDeleteUserSegment provides a mock function with given fields: ctx, userId, toAdd, toDelete
func (_m *IService) AddDeleteUserSegment(ctx context.Context, userId int, toAdd []model.SegmentWithExpires, toDelete []string) ([]int, error) {
	ret := _m.Called(ctx, userId, toAdd, toDelete)

	var r0 []int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, []model.SegmentWithExpires, []string) ([]int, error)); ok {
		return rf(ctx, userId, toAdd, toDelete)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, []model.SegmentWithExpires, []string) []int); ok {
		r0 = rf(ctx, userId, toAdd, toDelete)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, []model.SegmentWithExpires, []string) error); ok {
		r1 = rf(ctx, userId, toAdd, toDelete)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddSegment provides a mock function with given fields: ctx, name
func (_m *IService) AddSegment(ctx context.Context, name string) (int, error) {
	ret := _m.Called(ctx, name)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSegment provides a mock function with given fields: ctx, name
func (_m *IService) DeleteSegment(ctx context.Context, name string) error {
	ret := _m.Called(ctx, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FlushExpired provides a mock function with given fields: ctx
func (_m *IService) FlushExpired(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetSegmentsOfUser provides a mock function with given fields: ctx, userID
func (_m *IService) GetSegmentsOfUser(ctx context.Context, userID int) ([]model.SegmentWithExpires, error) {
	ret := _m.Called(ctx, userID)

	var r0 []model.SegmentWithExpires
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) ([]model.SegmentWithExpires, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) []model.SegmentWithExpires); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.SegmentWithExpires)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIService creates a new instance of IService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IService {
	mock := &IService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
