// Code generated by mockery v2.42.1. DO NOT EDIT.

package websocketsmocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// WebSocketsNamespaced is an autogenerated mock type for the WebSocketsNamespaced type
type WebSocketsNamespaced struct {
	mock.Mock
}

// ServeHTTPNamespaced provides a mock function with given fields: namespace, res, req
func (_m *WebSocketsNamespaced) ServeHTTPNamespaced(namespace string, res http.ResponseWriter, req *http.Request) {
	_m.Called(namespace, res, req)
}

// NewWebSocketsNamespaced creates a new instance of WebSocketsNamespaced. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWebSocketsNamespaced(t interface {
	mock.TestingT
	Cleanup(func())
}) *WebSocketsNamespaced {
	mock := &WebSocketsNamespaced{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}