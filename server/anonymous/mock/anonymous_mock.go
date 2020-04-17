// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bakurits/mattermost-plugin-anonymous/server/anonymous (interfaces: Anonymous)

// Package mock is a generated GoMock package.
package mock

import (
	config "github.com/bakurits/mattermost-plugin-anonymous/server/config"
	crypto "github.com/bakurits/mattermost-plugin-anonymous/server/crypto"
	store "github.com/bakurits/mattermost-plugin-anonymous/server/store"
	gomock "github.com/golang/mock/gomock"
	model "github.com/mattermost/mattermost-server/v5/model"
	reflect "reflect"
)

// MockAnonymous is a mock of Anonymous interface
type MockAnonymous struct {
	ctrl     *gomock.Controller
	recorder *MockAnonymousMockRecorder
}

// MockAnonymousMockRecorder is the mock recorder for MockAnonymous
type MockAnonymousMockRecorder struct {
	mock *MockAnonymous
}

// NewMockAnonymous creates a new mock instance
func NewMockAnonymous(ctrl *gomock.Controller) *MockAnonymous {
	mock := &MockAnonymous{ctrl: ctrl}
	mock.recorder = &MockAnonymousMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAnonymous) EXPECT() *MockAnonymousMockRecorder {
	return m.recorder
}

// DeleteUser mocks base method
func (m *MockAnonymous) DeleteUser(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *MockAnonymousMockRecorder) DeleteUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockAnonymous)(nil).DeleteUser), arg0)
}

// GetConfiguration mocks base method
func (m *MockAnonymous) GetConfiguration() *config.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfiguration")
	ret0, _ := ret[0].(*config.Config)
	return ret0
}

// GetConfiguration indicates an expected call of GetConfiguration
func (mr *MockAnonymousMockRecorder) GetConfiguration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfiguration", reflect.TypeOf((*MockAnonymous)(nil).GetConfiguration))
}

// GetPublicKey mocks base method
func (m *MockAnonymous) GetPublicKey(arg0 string) (crypto.PublicKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicKey", arg0)
	ret0, _ := ret[0].(crypto.PublicKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicKey indicates an expected call of GetPublicKey
func (mr *MockAnonymousMockRecorder) GetPublicKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicKey", reflect.TypeOf((*MockAnonymous)(nil).GetPublicKey), arg0)
}

// LoadUser mocks base method
func (m *MockAnonymous) LoadUser(arg0 string) (*store.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadUser", arg0)
	ret0, _ := ret[0].(*store.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadUser indicates an expected call of LoadUser
func (mr *MockAnonymousMockRecorder) LoadUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadUser", reflect.TypeOf((*MockAnonymous)(nil).LoadUser), arg0)
}

// SendEphemeralPost mocks base method
func (m *MockAnonymous) SendEphemeralPost(arg0 string, arg1 *model.Post) *model.Post {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendEphemeralPost", arg0, arg1)
	ret0, _ := ret[0].(*model.Post)
	return ret0
}

// SendEphemeralPost indicates an expected call of SendEphemeralPost
func (mr *MockAnonymousMockRecorder) SendEphemeralPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEphemeralPost", reflect.TypeOf((*MockAnonymous)(nil).SendEphemeralPost), arg0, arg1)
}

// SetConfiguration mocks base method
func (m *MockAnonymous) SetConfiguration(arg0 *config.Config) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetConfiguration", arg0)
}

// SetConfiguration indicates an expected call of SetConfiguration
func (mr *MockAnonymousMockRecorder) SetConfiguration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfiguration", reflect.TypeOf((*MockAnonymous)(nil).SetConfiguration), arg0)
}

// StorePublicKey mocks base method
func (m *MockAnonymous) StorePublicKey(arg0 string, arg1 crypto.PublicKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorePublicKey", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StorePublicKey indicates an expected call of StorePublicKey
func (mr *MockAnonymousMockRecorder) StorePublicKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorePublicKey", reflect.TypeOf((*MockAnonymous)(nil).StorePublicKey), arg0, arg1)
}

// StoreUser mocks base method
func (m *MockAnonymous) StoreUser(arg0 *store.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreUser indicates an expected call of StoreUser
func (mr *MockAnonymousMockRecorder) StoreUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreUser", reflect.TypeOf((*MockAnonymous)(nil).StoreUser), arg0)
}
