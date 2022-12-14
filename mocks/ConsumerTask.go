// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ConsumerTask is an autogenerated mock type for the ConsumerTask type
type ConsumerTask struct {
	mock.Mock
}

// ConsumeNews provides a mock function with given fields: ctx
func (_m *ConsumerTask) ConsumeNews(ctx context.Context) {
	_m.Called(ctx)
}

type mockConstructorTestingTNewConsumerTask interface {
	mock.TestingT
	Cleanup(func())
}

// NewConsumerTask creates a new instance of ConsumerTask. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConsumerTask(t mockConstructorTestingTNewConsumerTask) *ConsumerTask {
	mock := &ConsumerTask{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
