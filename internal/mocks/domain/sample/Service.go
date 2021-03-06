// Code generated by mockery v2.12.1. DO NOT EDIT.

package sample

import (
	context "context"

	domainsample "github.com/ftec-project/internal/domain/sample"
	entity "github.com/ftec-project/internal/infra/database/entity"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateSample provides a mock function with given fields: ctx, createDTO
func (_m *Service) CreateSample(ctx context.Context, createDTO domainsample.CreateDTO) (entity.Sample, error) {
	ret := _m.Called(ctx, createDTO)

	var r0 entity.Sample
	if rf, ok := ret.Get(0).(func(context.Context, domainsample.CreateDTO) entity.Sample); ok {
		r0 = rf(ctx, createDTO)
	} else {
		r0 = ret.Get(0).(entity.Sample)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domainsample.CreateDTO) error); ok {
		r1 = rf(ctx, createDTO)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByReferenceUUID provides a mock function with given fields: ctx, referenceUUID
func (_m *Service) GetByReferenceUUID(ctx context.Context, referenceUUID string) (entity.Sample, error) {
	ret := _m.Called(ctx, referenceUUID)

	var r0 entity.Sample
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.Sample); ok {
		r0 = rf(ctx, referenceUUID)
	} else {
		r0 = ret.Get(0).(entity.Sample)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, referenceUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewService creates a new instance of Service. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t testing.TB) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
