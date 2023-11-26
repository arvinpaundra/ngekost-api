// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/arvinpaundra/ngekost-api/internal/entity"

	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	request "github.com/arvinpaundra/ngekost-api/internal/adapter/request"
)

// KostRepository is an autogenerated mock type for the KostRepository type
type KostRepository struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, query
func (_m *KostRepository) Count(ctx context.Context, query *request.Common) (int, error) {
	ret := _m.Called(ctx, query)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *request.Common) (int, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *request.Common) int); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *request.Common) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountByOwnerId provides a mock function with given fields: ctx, ownerId, query
func (_m *KostRepository) CountByOwnerId(ctx context.Context, ownerId string, query *request.Common) (int, error) {
	ret := _m.Called(ctx, ownerId, query)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *request.Common) (int, error)); ok {
		return rf(ctx, ownerId, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *request.Common) int); ok {
		r0 = rf(ctx, ownerId, query)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *request.Common) error); ok {
		r1 = rf(ctx, ownerId, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, kostId
func (_m *KostRepository) Delete(ctx context.Context, kostId string) error {
	ret := _m.Called(ctx, kostId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, kostId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteWithTx provides a mock function with given fields: ctx, tx, kostId
func (_m *KostRepository) DeleteWithTx(ctx context.Context, tx *gorm.DB, kostId string) error {
	ret := _m.Called(ctx, tx, kostId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, string) error); ok {
		r0 = rf(ctx, tx, kostId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: ctx, query
func (_m *KostRepository) Find(ctx context.Context, query *request.Common) ([]*entity.Kost, error) {
	ret := _m.Called(ctx, query)

	var r0 []*entity.Kost
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *request.Common) ([]*entity.Kost, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *request.Common) []*entity.Kost); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Kost)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *request.Common) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: ctx, kostId
func (_m *KostRepository) FindById(ctx context.Context, kostId string) (*entity.Kost, error) {
	ret := _m.Called(ctx, kostId)

	var r0 *entity.Kost
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Kost, error)); ok {
		return rf(ctx, kostId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Kost); ok {
		r0 = rf(ctx, kostId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Kost)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, kostId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByOwnerId provides a mock function with given fields: ctx, ownerId, query
func (_m *KostRepository) FindByOwnerId(ctx context.Context, ownerId string, query *request.Common) ([]*entity.Kost, error) {
	ret := _m.Called(ctx, ownerId, query)

	var r0 []*entity.Kost
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *request.Common) ([]*entity.Kost, error)); ok {
		return rf(ctx, ownerId, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *request.Common) []*entity.Kost); ok {
		r0 = rf(ctx, ownerId, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Kost)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *request.Common) error); ok {
		r1 = rf(ctx, ownerId, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, kost
func (_m *KostRepository) Save(ctx context.Context, kost *entity.Kost) error {
	ret := _m.Called(ctx, kost)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Kost) error); ok {
		r0 = rf(ctx, kost)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveWithTx provides a mock function with given fields: ctx, tx, kost
func (_m *KostRepository) SaveWithTx(ctx context.Context, tx *gorm.DB, kost *entity.Kost) error {
	ret := _m.Called(ctx, tx, kost)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, *entity.Kost) error); ok {
		r0 = rf(ctx, tx, kost)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, kost, kostId
func (_m *KostRepository) Update(ctx context.Context, kost *entity.Kost, kostId string) error {
	ret := _m.Called(ctx, kost, kostId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Kost, string) error); ok {
		r0 = rf(ctx, kost, kostId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateWithTx provides a mock function with given fields: ctx, tx, kost, kostId
func (_m *KostRepository) UpdateWithTx(ctx context.Context, tx *gorm.DB, kost *entity.Kost, kostId string) error {
	ret := _m.Called(ctx, tx, kost, kostId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, *entity.Kost, string) error); ok {
		r0 = rf(ctx, tx, kost, kostId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewKostRepository creates a new instance of KostRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewKostRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *KostRepository {
	mock := &KostRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
