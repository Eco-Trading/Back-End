package mocked

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type GenericMock[T any] struct {
	mock.Mock
}

func (a *GenericMock[T]) AssertNumberOfCallsTime(t *testing.T, method string, calls int) {
	a.AssertCalled(t, method)
	a.AssertNumberOfCalls(t, method, calls)
}
