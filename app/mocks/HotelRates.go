// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	handlers "nuitee-task/handlers"

	mock "github.com/stretchr/testify/mock"
)

// HotelRates is an autogenerated mock type for the HotelRates type
type HotelRates struct {
	mock.Mock
}

type HotelRates_Expecter struct {
	mock *mock.Mock
}

func (_m *HotelRates) EXPECT() *HotelRates_Expecter {
	return &HotelRates_Expecter{mock: &_m.Mock}
}

// GetHotelPrices provides a mock function with given fields: _a0
func (_m *HotelRates) GetHotelPrices(_a0 *handlers.GetHotelsRequest) ([]*handlers.HotelPrice, string, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetHotelPrices")
	}

	var r0 []*handlers.HotelPrice
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(*handlers.GetHotelsRequest) ([]*handlers.HotelPrice, string, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*handlers.GetHotelsRequest) []*handlers.HotelPrice); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*handlers.HotelPrice)
		}
	}

	if rf, ok := ret.Get(1).(func(*handlers.GetHotelsRequest) string); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(*handlers.GetHotelsRequest) error); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// HotelRates_GetHotelPrices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHotelPrices'
type HotelRates_GetHotelPrices_Call struct {
	*mock.Call
}

// GetHotelPrices is a helper method to define mock.On call
//   - _a0 *handlers.GetHotelsRequest
func (_e *HotelRates_Expecter) GetHotelPrices(_a0 interface{}) *HotelRates_GetHotelPrices_Call {
	return &HotelRates_GetHotelPrices_Call{Call: _e.mock.On("GetHotelPrices", _a0)}
}

func (_c *HotelRates_GetHotelPrices_Call) Run(run func(_a0 *handlers.GetHotelsRequest)) *HotelRates_GetHotelPrices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*handlers.GetHotelsRequest))
	})
	return _c
}

func (_c *HotelRates_GetHotelPrices_Call) Return(_a0 []*handlers.HotelPrice, _a1 string, _a2 error) *HotelRates_GetHotelPrices_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *HotelRates_GetHotelPrices_Call) RunAndReturn(run func(*handlers.GetHotelsRequest) ([]*handlers.HotelPrice, string, error)) *HotelRates_GetHotelPrices_Call {
	_c.Call.Return(run)
	return _c
}

// NewHotelRates creates a new instance of HotelRates. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHotelRates(t interface {
	mock.TestingT
	Cleanup(func())
}) *HotelRates {
	mock := &HotelRates{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
