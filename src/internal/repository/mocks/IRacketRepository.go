// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"
	dto "src/internal/dto"

	mock "github.com/stretchr/testify/mock"

	model "src/internal/model"
)

// IRacketRepository is an autogenerated mock type for the IRacketRepository type
type IRacketRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, racket
func (_m *IRacketRepository) Create(ctx context.Context, racket *model.Racket) error {
	ret := _m.Called(ctx, racket)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Racket) error); ok {
		r0 = rf(ctx, racket)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *IRacketRepository) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllRackets provides a mock function with given fields: ctx, req
func (_m *IRacketRepository) GetAllRackets(ctx context.Context, req *dto.ListRacketsReq) ([]*model.Racket, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for GetAllRackets")
	}

	var r0 []*model.Racket
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.ListRacketsReq) ([]*model.Racket, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.ListRacketsReq) []*model.Racket); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Racket)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.ListRacketsReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRacketByID provides a mock function with given fields: ctx, id
func (_m *IRacketRepository) GetRacketByID(ctx context.Context, id int) (*model.Racket, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetRacketByID")
	}

	var r0 *model.Racket
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*model.Racket, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *model.Racket); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Racket)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, racket
func (_m *IRacketRepository) Update(ctx context.Context, racket *model.Racket) error {
	ret := _m.Called(ctx, racket)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Racket) error); ok {
		r0 = rf(ctx, racket)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIRacketRepository creates a new instance of IRacketRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIRacketRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IRacketRepository {
	mock := &IRacketRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
