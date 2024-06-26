// Code generated by mockery v2.42.2. DO NOT EDIT.

package service

import (
	dtos "codeberg.org/Birkenfunk/SQS/dtos"
	mock "github.com/stretchr/testify/mock"
)

// MockIWeatherService is an autogenerated mock type for the IWeatherService type
type MockIWeatherService struct {
	mock.Mock
}

type MockIWeatherService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIWeatherService) EXPECT() *MockIWeatherService_Expecter {
	return &MockIWeatherService_Expecter{mock: &_m.Mock}
}

// GetHealth provides a mock function with given fields:
func (_m *MockIWeatherService) GetHealth() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetHealth")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIWeatherService_GetHealth_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHealth'
type MockIWeatherService_GetHealth_Call struct {
	*mock.Call
}

// GetHealth is a helper method to define mock.On call
func (_e *MockIWeatherService_Expecter) GetHealth() *MockIWeatherService_GetHealth_Call {
	return &MockIWeatherService_GetHealth_Call{Call: _e.mock.On("GetHealth")}
}

func (_c *MockIWeatherService_GetHealth_Call) Run(run func()) *MockIWeatherService_GetHealth_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockIWeatherService_GetHealth_Call) Return(_a0 error) *MockIWeatherService_GetHealth_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIWeatherService_GetHealth_Call) RunAndReturn(run func() error) *MockIWeatherService_GetHealth_Call {
	_c.Call.Return(run)
	return _c
}

// GetWeather provides a mock function with given fields: location
func (_m *MockIWeatherService) GetWeather(location string) (*dtos.WeatherDto, error) {
	ret := _m.Called(location)

	if len(ret) == 0 {
		panic("no return value specified for GetWeather")
	}

	var r0 *dtos.WeatherDto
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*dtos.WeatherDto, error)); ok {
		return rf(location)
	}
	if rf, ok := ret.Get(0).(func(string) *dtos.WeatherDto); ok {
		r0 = rf(location)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dtos.WeatherDto)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(location)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIWeatherService_GetWeather_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWeather'
type MockIWeatherService_GetWeather_Call struct {
	*mock.Call
}

// GetWeather is a helper method to define mock.On call
//   - location string
func (_e *MockIWeatherService_Expecter) GetWeather(location interface{}) *MockIWeatherService_GetWeather_Call {
	return &MockIWeatherService_GetWeather_Call{Call: _e.mock.On("GetWeather", location)}
}

func (_c *MockIWeatherService_GetWeather_Call) Run(run func(location string)) *MockIWeatherService_GetWeather_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockIWeatherService_GetWeather_Call) Return(_a0 *dtos.WeatherDto, _a1 error) *MockIWeatherService_GetWeather_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIWeatherService_GetWeather_Call) RunAndReturn(run func(string) (*dtos.WeatherDto, error)) *MockIWeatherService_GetWeather_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIWeatherService creates a new instance of MockIWeatherService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIWeatherService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIWeatherService {
	mock := &MockIWeatherService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
