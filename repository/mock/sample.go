// Code generated by MockGen. DO NOT EDIT.
// Source: repository/sample.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/hapoon/go-api-template/entity"
)

// MockSampleRepository is a mock of SampleRepository interface.
type MockSampleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSampleRepositoryMockRecorder
}

// MockSampleRepositoryMockRecorder is the mock recorder for MockSampleRepository.
type MockSampleRepositoryMockRecorder struct {
	mock *MockSampleRepository
}

// NewMockSampleRepository creates a new mock instance.
func NewMockSampleRepository(ctrl *gomock.Controller) *MockSampleRepository {
	mock := &MockSampleRepository{ctrl: ctrl}
	mock.recorder = &MockSampleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSampleRepository) EXPECT() *MockSampleRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSampleRepository) Create(data string) (*entity.SampleEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", data)
	ret0, _ := ret[0].(*entity.SampleEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockSampleRepositoryMockRecorder) Create(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSampleRepository)(nil).Create), data)
}

// GetById mocks base method.
func (m *MockSampleRepository) GetById(id string) (entity.SampleEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(entity.SampleEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockSampleRepositoryMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockSampleRepository)(nil).GetById), id)
}

// List mocks base method.
func (m *MockSampleRepository) List() (entity.SampleEntities, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(entity.SampleEntities)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockSampleRepositoryMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSampleRepository)(nil).List))
}
