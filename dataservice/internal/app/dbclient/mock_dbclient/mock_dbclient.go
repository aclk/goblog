// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pkg/dbclient/cockroachdb.go

// Package mock_dbclient is a generated GoMock package.
package mock_dbclient

import (
	context "context"
	model "github.com/aclk/goblog/common/model"
	dbclient "github.com/aclk/goblog/dataservice/internal/app/dbclient"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIGormClient is a mock of IGormClient interface
type MockIGormClient struct {
	ctrl     *gomock.Controller
	recorder *MockIGormClientMockRecorder
}

// MockIGormClientMockRecorder is the mock recorder for MockIGormClient
type MockIGormClientMockRecorder struct {
	mock *MockIGormClient
}

// NewMockIGormClient creates a new mock instance
func NewMockIGormClient(ctrl *gomock.Controller) *MockIGormClient {
	mock := &MockIGormClient{ctrl: ctrl}
	mock.recorder = &MockIGormClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIGormClient) EXPECT() *MockIGormClientMockRecorder {
	return m.recorder
}

// UpdateAccount mocks base method
func (m *MockIGormClient) UpdateAccount(ctx context.Context, accountData model.AccountData) (model.AccountData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", ctx, accountData)
	ret0, _ := ret[0].(model.AccountData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccount indicates an expected call of UpdateAccount
func (mr *MockIGormClientMockRecorder) UpdateAccount(ctx, accountData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockIGormClient)(nil).UpdateAccount), ctx, accountData)
}

// StoreAccount mocks base method
func (m *MockIGormClient) StoreAccount(ctx context.Context, accountData model.AccountData) (model.AccountData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreAccount", ctx, accountData)
	ret0, _ := ret[0].(model.AccountData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreAccount indicates an expected call of StoreAccount
func (mr *MockIGormClientMockRecorder) StoreAccount(ctx, accountData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreAccount", reflect.TypeOf((*MockIGormClient)(nil).StoreAccount), ctx, accountData)
}

// QueryAccount mocks base method
func (m *MockIGormClient) QueryAccount(ctx context.Context, accountId string) (model.AccountData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryAccount", ctx, accountId)
	ret0, _ := ret[0].(model.AccountData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryAccount indicates an expected call of QueryAccount
func (mr *MockIGormClientMockRecorder) QueryAccount(ctx, accountId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryAccount", reflect.TypeOf((*MockIGormClient)(nil).QueryAccount), ctx, accountId)
}

// GetRandomAccount mocks base method
func (m *MockIGormClient) GetRandomAccount(ctx context.Context) (model.AccountData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRandomAccount", ctx)
	ret0, _ := ret[0].(model.AccountData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRandomAccount indicates an expected call of GetRandomAccount
func (mr *MockIGormClientMockRecorder) GetRandomAccount(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRandomAccount", reflect.TypeOf((*MockIGormClient)(nil).GetRandomAccount), ctx)
}

// QueryAccountByNameWithCount mocks base method
func (m *MockIGormClient) QueryAccountByNameWithCount(ctx context.Context, name string) ([]dbclient.Pair, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryAccountByNameWithCount", ctx, name)
	ret0, _ := ret[0].([]dbclient.Pair)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryAccountByNameWithCount indicates an expected call of QueryAccountByNameWithCount
func (mr *MockIGormClientMockRecorder) QueryAccountByNameWithCount(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryAccountByNameWithCount", reflect.TypeOf((*MockIGormClient)(nil).QueryAccountByNameWithCount), ctx, name)
}

// SetupDB mocks base method
func (m *MockIGormClient) SetupDB(addr string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetupDB", addr)
}

// SetupDB indicates an expected call of SetupDB
func (mr *MockIGormClientMockRecorder) SetupDB(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupDB", reflect.TypeOf((*MockIGormClient)(nil).SetupDB), addr)
}

// SeedAccounts mocks base method
func (m *MockIGormClient) SeedAccounts() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeedAccounts")
	ret0, _ := ret[0].(error)
	return ret0
}

// SeedAccounts indicates an expected call of SeedAccounts
func (mr *MockIGormClientMockRecorder) SeedAccounts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeedAccounts", reflect.TypeOf((*MockIGormClient)(nil).SeedAccounts))
}

// Check mocks base method
func (m *MockIGormClient) Check() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Check indicates an expected call of Check
func (mr *MockIGormClientMockRecorder) Check() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockIGormClient)(nil).Check))
}

// Close mocks base method
func (m *MockIGormClient) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockIGormClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIGormClient)(nil).Close))
}
