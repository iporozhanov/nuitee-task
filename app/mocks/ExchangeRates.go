// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	big "math/big"

	mock "github.com/stretchr/testify/mock"
)

// ExchangeRates is an autogenerated mock type for the ExchangeRates type
type ExchangeRates struct {
	mock.Mock
}

type ExchangeRates_Expecter struct {
	mock *mock.Mock
}

func (_m *ExchangeRates) EXPECT() *ExchangeRates_Expecter {
	return &ExchangeRates_Expecter{mock: &_m.Mock}
}

// ConvertCurrency provides a mock function with given fields: amount, from, to
func (_m *ExchangeRates) ConvertCurrency(amount *big.Float, from string, to string) (*big.Float, error) {
	ret := _m.Called(amount, from, to)

	if len(ret) == 0 {
		panic("no return value specified for ConvertCurrency")
	}

	var r0 *big.Float
	var r1 error
	if rf, ok := ret.Get(0).(func(*big.Float, string, string) (*big.Float, error)); ok {
		return rf(amount, from, to)
	}
	if rf, ok := ret.Get(0).(func(*big.Float, string, string) *big.Float); ok {
		r0 = rf(amount, from, to)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Float)
		}
	}

	if rf, ok := ret.Get(1).(func(*big.Float, string, string) error); ok {
		r1 = rf(amount, from, to)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExchangeRates_ConvertCurrency_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConvertCurrency'
type ExchangeRates_ConvertCurrency_Call struct {
	*mock.Call
}

// ConvertCurrency is a helper method to define mock.On call
//   - amount *big.Float
//   - from string
//   - to string
func (_e *ExchangeRates_Expecter) ConvertCurrency(amount interface{}, from interface{}, to interface{}) *ExchangeRates_ConvertCurrency_Call {
	return &ExchangeRates_ConvertCurrency_Call{Call: _e.mock.On("ConvertCurrency", amount, from, to)}
}

func (_c *ExchangeRates_ConvertCurrency_Call) Run(run func(amount *big.Float, from string, to string)) *ExchangeRates_ConvertCurrency_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*big.Float), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ExchangeRates_ConvertCurrency_Call) Return(_a0 *big.Float, _a1 error) *ExchangeRates_ConvertCurrency_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ExchangeRates_ConvertCurrency_Call) RunAndReturn(run func(*big.Float, string, string) (*big.Float, error)) *ExchangeRates_ConvertCurrency_Call {
	_c.Call.Return(run)
	return _c
}

// NewExchangeRates creates a new instance of ExchangeRates. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewExchangeRates(t interface {
	mock.TestingT
	Cleanup(func())
}) *ExchangeRates {
	mock := &ExchangeRates{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
