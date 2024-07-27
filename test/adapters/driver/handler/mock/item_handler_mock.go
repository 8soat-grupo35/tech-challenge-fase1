// Code generated by MockGen. DO NOT EDIT.
// Source: item_handler.go
//
// Generated by this command:
//
//	mockgen -source=item_handler.go -destination=../../../../test/adapters/driver/handler/mock/item_handler_mock.go
//

// Package mock_handler is a generated GoMock package.
package mock_handler

import (
	reflect "reflect"

	echo "github.com/labstack/echo"
	gomock "go.uber.org/mock/gomock"
)

// MockItemHandler is a mock of ItemHandler interface.
type MockItemHandler struct {
	ctrl     *gomock.Controller
	recorder *MockItemHandlerMockRecorder
}

// MockItemHandlerMockRecorder is the mock recorder for MockItemHandler.
type MockItemHandlerMockRecorder struct {
	mock *MockItemHandler
}

// NewMockItemHandler creates a new mock instance.
func NewMockItemHandler(ctrl *gomock.Controller) *MockItemHandler {
	mock := &MockItemHandler{ctrl: ctrl}
	mock.recorder = &MockItemHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemHandler) EXPECT() *MockItemHandlerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockItemHandler) Create(echo echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", echo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockItemHandlerMockRecorder) Create(echo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockItemHandler)(nil).Create), echo)
}

// Delete mocks base method.
func (m *MockItemHandler) Delete(echo echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", echo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockItemHandlerMockRecorder) Delete(echo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockItemHandler)(nil).Delete), echo)
}

// GetAll mocks base method.
func (m *MockItemHandler) GetAll(echo echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", echo)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAll indicates an expected call of GetAll.
func (mr *MockItemHandlerMockRecorder) GetAll(echo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockItemHandler)(nil).GetAll), echo)
}

// Update mocks base method.
func (m *MockItemHandler) Update(echo echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", echo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockItemHandlerMockRecorder) Update(echo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockItemHandler)(nil).Update), echo)
}
