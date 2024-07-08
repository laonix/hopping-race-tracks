// Code generated by mockery v2.36.1. DO NOT EDIT.

package dispatcher

import mock "github.com/stretchr/testify/mock"

// MockLogger is an autogenerated mock type for the Logger type
type MockLogger struct {
	mock.Mock
}

type MockLogger_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLogger) EXPECT() *MockLogger_Expecter {
	return &MockLogger_Expecter{mock: &_m.Mock}
}

// Debug provides a mock function with given fields: msg, fields
func (_m *MockLogger) Debug(msg string, fields ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, fields...)
	_m.Called(_ca...)
}

// MockLogger_Debug_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Debug'
type MockLogger_Debug_Call struct {
	*mock.Call
}

// Debug is a helper method to define mock.On call
//   - msg string
//   - fields ...interface{}
func (_e *MockLogger_Expecter) Debug(msg interface{}, fields ...interface{}) *MockLogger_Debug_Call {
	return &MockLogger_Debug_Call{Call: _e.mock.On("Debug",
		append([]interface{}{msg}, fields...)...)}
}

func (_c *MockLogger_Debug_Call) Run(run func(msg string, fields ...interface{})) *MockLogger_Debug_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Debug_Call) Return() *MockLogger_Debug_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Debug_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Debug_Call {
	_c.Call.Return(run)
	return _c
}

// Error provides a mock function with given fields: err, msg, fields
func (_m *MockLogger) Error(err error, msg string, fields ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, err, msg)
	_ca = append(_ca, fields...)
	_m.Called(_ca...)
}

// MockLogger_Error_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Error'
type MockLogger_Error_Call struct {
	*mock.Call
}

// Error is a helper method to define mock.On call
//   - err error
//   - msg string
//   - fields ...interface{}
func (_e *MockLogger_Expecter) Error(err interface{}, msg interface{}, fields ...interface{}) *MockLogger_Error_Call {
	return &MockLogger_Error_Call{Call: _e.mock.On("Error",
		append([]interface{}{err, msg}, fields...)...)}
}

func (_c *MockLogger_Error_Call) Run(run func(err error, msg string, fields ...interface{})) *MockLogger_Error_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(error), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Error_Call) Return() *MockLogger_Error_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Error_Call) RunAndReturn(run func(error, string, ...interface{})) *MockLogger_Error_Call {
	_c.Call.Return(run)
	return _c
}

// Fatal provides a mock function with given fields: err, msg, fields
func (_m *MockLogger) Fatal(err error, msg string, fields ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, err, msg)
	_ca = append(_ca, fields...)
	_m.Called(_ca...)
}

// MockLogger_Fatal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fatal'
type MockLogger_Fatal_Call struct {
	*mock.Call
}

// Fatal is a helper method to define mock.On call
//   - err error
//   - msg string
//   - fields ...interface{}
func (_e *MockLogger_Expecter) Fatal(err interface{}, msg interface{}, fields ...interface{}) *MockLogger_Fatal_Call {
	return &MockLogger_Fatal_Call{Call: _e.mock.On("Fatal",
		append([]interface{}{err, msg}, fields...)...)}
}

func (_c *MockLogger_Fatal_Call) Run(run func(err error, msg string, fields ...interface{})) *MockLogger_Fatal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(error), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Fatal_Call) Return() *MockLogger_Fatal_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Fatal_Call) RunAndReturn(run func(error, string, ...interface{})) *MockLogger_Fatal_Call {
	_c.Call.Return(run)
	return _c
}

// Info provides a mock function with given fields: msg, fields
func (_m *MockLogger) Info(msg string, fields ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, fields...)
	_m.Called(_ca...)
}

// MockLogger_Info_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Info'
type MockLogger_Info_Call struct {
	*mock.Call
}

// Info is a helper method to define mock.On call
//   - msg string
//   - fields ...interface{}
func (_e *MockLogger_Expecter) Info(msg interface{}, fields ...interface{}) *MockLogger_Info_Call {
	return &MockLogger_Info_Call{Call: _e.mock.On("Info",
		append([]interface{}{msg}, fields...)...)}
}

func (_c *MockLogger_Info_Call) Run(run func(msg string, fields ...interface{})) *MockLogger_Info_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Info_Call) Return() *MockLogger_Info_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Info_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Info_Call {
	_c.Call.Return(run)
	return _c
}

// Warn provides a mock function with given fields: msg, fields
func (_m *MockLogger) Warn(msg string, fields ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, fields...)
	_m.Called(_ca...)
}

// MockLogger_Warn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Warn'
type MockLogger_Warn_Call struct {
	*mock.Call
}

// Warn is a helper method to define mock.On call
//   - msg string
//   - fields ...interface{}
func (_e *MockLogger_Expecter) Warn(msg interface{}, fields ...interface{}) *MockLogger_Warn_Call {
	return &MockLogger_Warn_Call{Call: _e.mock.On("Warn",
		append([]interface{}{msg}, fields...)...)}
}

func (_c *MockLogger_Warn_Call) Run(run func(msg string, fields ...interface{})) *MockLogger_Warn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockLogger_Warn_Call) Return() *MockLogger_Warn_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockLogger_Warn_Call) RunAndReturn(run func(string, ...interface{})) *MockLogger_Warn_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockLogger creates a new instance of MockLogger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLogger(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLogger {
	mock := &MockLogger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}