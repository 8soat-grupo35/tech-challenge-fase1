// Code generated by MockGen. DO NOT EDIT.
// Source: item.go
//
// Generated by this command:
//
//	mockgen -source=item.go -destination=../../../test/gateways/mock/item_mock.go
//

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	entities "github.com/8soat-grupo35/tech-challenge-fase1/src/entities"
	gomock "go.uber.org/mock/gomock"
)

// MockItemRepository is a mock of ItemRepository interface.
type MockItemRepository struct {
	ctrl     *gomock.Controller
	recorder *MockItemRepositoryMockRecorder
}

// MockItemRepositoryMockRecorder is the mock recorder for MockItemRepository.
type MockItemRepositoryMockRecorder struct {
	mock *MockItemRepository
}

// NewMockItemRepository creates a new mock instance.
func NewMockItemRepository(ctrl *gomock.Controller) *MockItemRepository {
	mock := &MockItemRepository{ctrl: ctrl}
	mock.recorder = &MockItemRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemRepository) EXPECT() *MockItemRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockItemRepository) Create(item entities.Item) (*entities.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", item)
	ret0, _ := ret[0].(*entities.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockItemRepositoryMockRecorder) Create(item any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockItemRepository)(nil).Create), item)
}

// Delete mocks base method.
func (m *MockItemRepository) Delete(itemId uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", itemId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockItemRepositoryMockRecorder) Delete(itemId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockItemRepository)(nil).Delete), itemId)
}

// GetAll mocks base method.
func (m *MockItemRepository) GetAll(arg0 entities.Item) ([]entities.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].([]entities.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockItemRepositoryMockRecorder) GetAll(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockItemRepository)(nil).GetAll), arg0)
}

// GetOne mocks base method.
func (m *MockItemRepository) GetOne(arg0 entities.Item) (*entities.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOne", arg0)
	ret0, _ := ret[0].(*entities.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOne indicates an expected call of GetOne.
func (mr *MockItemRepositoryMockRecorder) GetOne(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOne", reflect.TypeOf((*MockItemRepository)(nil).GetOne), arg0)
}

// Update mocks base method.
func (m *MockItemRepository) Update(itemId uint32, item entities.Item) (*entities.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", itemId, item)
	ret0, _ := ret[0].(*entities.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockItemRepositoryMockRecorder) Update(itemId, item any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockItemRepository)(nil).Update), itemId, item)
}
