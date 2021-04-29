// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-plugin-incident-collaboration/server/bot (interfaces: Poster)

// Package mock_bot is a generated GoMock package.
package mock_bot

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/mattermost/mattermost-server/v5/model"
	reflect "reflect"
)

// MockPoster is a mock of Poster interface
type MockPoster struct {
	ctrl     *gomock.Controller
	recorder *MockPosterMockRecorder
}

// MockPosterMockRecorder is the mock recorder for MockPoster
type MockPosterMockRecorder struct {
	mock *MockPoster
}

// NewMockPoster creates a new mock instance
func NewMockPoster(ctrl *gomock.Controller) *MockPoster {
	mock := &MockPoster{ctrl: ctrl}
	mock.recorder = &MockPosterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPoster) EXPECT() *MockPosterMockRecorder {
	return m.recorder
}

// DM mocks base method
func (m *MockPoster) DM(arg0, arg1 string, arg2 ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DM", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DM indicates an expected call of DM
func (mr *MockPosterMockRecorder) DM(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DM", reflect.TypeOf((*MockPoster)(nil).DM), varargs...)
}

// DMWithAttachments mocks base method
func (m *MockPoster) DMWithAttachments(arg0 string, arg1 ...*model.SlackAttachment) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DMWithAttachments", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DMWithAttachments indicates an expected call of DMWithAttachments
func (mr *MockPosterMockRecorder) DMWithAttachments(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DMWithAttachments", reflect.TypeOf((*MockPoster)(nil).DMWithAttachments), varargs...)
}

// EphemeralPost mocks base method
func (m *MockPoster) EphemeralPost(arg0, arg1 string, arg2 *model.Post) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EphemeralPost", arg0, arg1, arg2)
}

// EphemeralPost indicates an expected call of EphemeralPost
func (mr *MockPosterMockRecorder) EphemeralPost(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EphemeralPost", reflect.TypeOf((*MockPoster)(nil).EphemeralPost), arg0, arg1, arg2)
}

// NotifyAdmins mocks base method
func (m *MockPoster) NotifyAdmins(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyAdmins", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotifyAdmins indicates an expected call of NotifyAdmins
func (mr *MockPosterMockRecorder) NotifyAdmins(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyAdmins", reflect.TypeOf((*MockPoster)(nil).NotifyAdmins), arg0)
}

// PostMessage mocks base method
func (m *MockPoster) PostMessage(arg0, arg1 string, arg2 ...interface{}) (*model.Post, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostMessage", varargs...)
	ret0, _ := ret[0].(*model.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostMessage indicates an expected call of PostMessage
func (mr *MockPosterMockRecorder) PostMessage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostMessage", reflect.TypeOf((*MockPoster)(nil).PostMessage), varargs...)
}

// PostMessageWithAttachments mocks base method
func (m *MockPoster) PostMessageWithAttachments(arg0 string, arg1 []*model.SlackAttachment, arg2 string, arg3 ...interface{}) (*model.Post, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PostMessageWithAttachments", varargs...)
	ret0, _ := ret[0].(*model.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostMessageWithAttachments indicates an expected call of PostMessageWithAttachments
func (mr *MockPosterMockRecorder) PostMessageWithAttachments(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostMessageWithAttachments", reflect.TypeOf((*MockPoster)(nil).PostMessageWithAttachments), varargs...)
}

// PublishWebsocketEventToChannel mocks base method
func (m *MockPoster) PublishWebsocketEventToChannel(arg0 string, arg1 interface{}, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PublishWebsocketEventToChannel", arg0, arg1, arg2)
}

// PublishWebsocketEventToChannel indicates an expected call of PublishWebsocketEventToChannel
func (mr *MockPosterMockRecorder) PublishWebsocketEventToChannel(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishWebsocketEventToChannel", reflect.TypeOf((*MockPoster)(nil).PublishWebsocketEventToChannel), arg0, arg1, arg2)
}

// PublishWebsocketEventToTeam mocks base method
func (m *MockPoster) PublishWebsocketEventToTeam(arg0 string, arg1 interface{}, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PublishWebsocketEventToTeam", arg0, arg1, arg2)
}

// PublishWebsocketEventToTeam indicates an expected call of PublishWebsocketEventToTeam
func (mr *MockPosterMockRecorder) PublishWebsocketEventToTeam(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishWebsocketEventToTeam", reflect.TypeOf((*MockPoster)(nil).PublishWebsocketEventToTeam), arg0, arg1, arg2)
}

// PublishWebsocketEventToUser mocks base method
func (m *MockPoster) PublishWebsocketEventToUser(arg0 string, arg1 interface{}, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PublishWebsocketEventToUser", arg0, arg1, arg2)
}

// PublishWebsocketEventToUser indicates an expected call of PublishWebsocketEventToUser
func (mr *MockPosterMockRecorder) PublishWebsocketEventToUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishWebsocketEventToUser", reflect.TypeOf((*MockPoster)(nil).PublishWebsocketEventToUser), arg0, arg1, arg2)
}
