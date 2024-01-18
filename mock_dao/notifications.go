// Code generated by MockGen. DO NOT EDIT.
// Source: notifications.go

// Package mock_dao is a generated GoMock package.
package mock_dao

import (
	reflect "reflect"

	model "github.com/TUM-Dev/gocast/model"
	gomock "github.com/golang/mock/gomock"
)

// MockNotificationsDao is a mock of NotificationsDao interface.
type MockNotificationsDao struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationsDaoMockRecorder
}

// MockNotificationsDaoMockRecorder is the mock recorder for MockNotificationsDao.
type MockNotificationsDaoMockRecorder struct {
	mock *MockNotificationsDao
}

// NewMockNotificationsDao creates a new mock instance.
func NewMockNotificationsDao(ctrl *gomock.Controller) *MockNotificationsDao {
	mock := &MockNotificationsDao{ctrl: ctrl}
	mock.recorder = &MockNotificationsDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotificationsDao) EXPECT() *MockNotificationsDaoMockRecorder {
	return m.recorder
}

// AddNotification mocks base method.
func (m *MockNotificationsDao) AddNotification(notification *model.Notification) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNotification", notification)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNotification indicates an expected call of AddNotification.
func (mr *MockNotificationsDaoMockRecorder) AddNotification(notification interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNotification", reflect.TypeOf((*MockNotificationsDao)(nil).AddNotification), notification)
}

// DeleteNotification mocks base method.
func (m *MockNotificationsDao) DeleteNotification(id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNotification", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNotification indicates an expected call of DeleteNotification.
func (mr *MockNotificationsDaoMockRecorder) DeleteNotification(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNotification", reflect.TypeOf((*MockNotificationsDao)(nil).DeleteNotification), id)
}

// GetAllNotifications mocks base method.
func (m *MockNotificationsDao) GetAllNotifications() ([]model.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllNotifications")
	ret0, _ := ret[0].([]model.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllNotifications indicates an expected call of GetAllNotifications.
func (mr *MockNotificationsDaoMockRecorder) GetAllNotifications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllNotifications", reflect.TypeOf((*MockNotificationsDao)(nil).GetAllNotifications))
}

// GetNotifications mocks base method.
func (m *MockNotificationsDao) GetNotifications(target ...model.NotificationTarget) ([]model.Notification, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range target {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetNotifications", varargs...)
	ret0, _ := ret[0].([]model.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotifications indicates an expected call of GetNotifications.
func (mr *MockNotificationsDaoMockRecorder) GetNotifications(target ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotifications", reflect.TypeOf((*MockNotificationsDao)(nil).GetNotifications), target...)
}
