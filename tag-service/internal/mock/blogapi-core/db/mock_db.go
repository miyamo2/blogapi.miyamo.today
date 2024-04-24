// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/miyamo2/blogapi.miyamo.today/core/db (interfaces: TransactionManager,Transaction,Statement)
//
// Generated by this command:
//
//	mockgen.exe github.com/miyamo2/blogapi.miyamo.today/core/db TransactionManager,Transaction,Statement
//
// Package mock_db is a generated GoMock package.
package mock_db

import (
	context "context"
	reflect "reflect"

	db "github.com/miyamo2/blogapi.miyamo.today/core/db"
	gomock "go.uber.org/mock/gomock"
)

// MockTransactionManager is a mock of TransactionManager interface.
type MockTransactionManager struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionManagerMockRecorder
}

// MockTransactionManagerMockRecorder is the mock recorder for MockTransactionManager.
type MockTransactionManagerMockRecorder struct {
	mock *MockTransactionManager
}

// NewMockTransactionManager creates a new mock instance.
func NewMockTransactionManager(ctrl *gomock.Controller) *MockTransactionManager {
	mock := &MockTransactionManager{ctrl: ctrl}
	mock.recorder = &MockTransactionManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionManager) EXPECT() *MockTransactionManagerMockRecorder {
	return m.recorder
}

// GetAndStart mocks base method.
func (m *MockTransactionManager) GetAndStart(arg0 context.Context) (db.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAndStart", arg0)
	ret0, _ := ret[0].(db.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAndStart indicates an expected call of GetAndStart.
func (mr *MockTransactionManagerMockRecorder) GetAndStart(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAndStart", reflect.TypeOf((*MockTransactionManager)(nil).GetAndStart), arg0)
}

// MockTransaction is a mock of Transaction interface.
type MockTransaction struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionMockRecorder
}

// MockTransactionMockRecorder is the mock recorder for MockTransaction.
type MockTransactionMockRecorder struct {
	mock *MockTransaction
}

// NewMockTransaction creates a new mock instance.
func NewMockTransaction(ctrl *gomock.Controller) *MockTransaction {
	mock := &MockTransaction{ctrl: ctrl}
	mock.recorder = &MockTransactionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransaction) EXPECT() *MockTransactionMockRecorder {
	return m.recorder
}

// Commit mocks base method.
func (m *MockTransaction) Commit(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockTransactionMockRecorder) Commit(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockTransaction)(nil).Commit), arg0)
}

// ExecuteStatement mocks base method.
func (m *MockTransaction) ExecuteStatement(arg0 context.Context, arg1 db.Statement) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteStatement", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecuteStatement indicates an expected call of ExecuteStatement.
func (mr *MockTransactionMockRecorder) ExecuteStatement(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteStatement", reflect.TypeOf((*MockTransaction)(nil).ExecuteStatement), arg0, arg1)
}

// Rollback mocks base method.
func (m *MockTransaction) Rollback(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockTransactionMockRecorder) Rollback(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockTransaction)(nil).Rollback), arg0)
}

// SubscribeError mocks base method.
func (m *MockTransaction) SubscribeError() <-chan error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeError")
	ret0, _ := ret[0].(<-chan error)
	return ret0
}

// SubscribeError indicates an expected call of SubscribeError.
func (mr *MockTransactionMockRecorder) SubscribeError() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeError", reflect.TypeOf((*MockTransaction)(nil).SubscribeError))
}

// MockStatement is a mock of Statement interface.
type MockStatement struct {
	ctrl     *gomock.Controller
	recorder *MockStatementMockRecorder
}

// MockStatementMockRecorder is the mock recorder for MockStatement.
type MockStatementMockRecorder struct {
	mock *MockStatement
}

// NewMockStatement creates a new mock instance.
func NewMockStatement(ctrl *gomock.Controller) *MockStatement {
	mock := &MockStatement{ctrl: ctrl}
	mock.recorder = &MockStatementMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatement) EXPECT() *MockStatementMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockStatement) Execute(arg0 context.Context, arg1 ...db.ExecuteOption) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Execute", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockStatementMockRecorder) Execute(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockStatement)(nil).Execute), varargs...)
}

// Result mocks base method.
func (m *MockStatement) Result() db.StatementResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Result")
	ret0, _ := ret[0].(db.StatementResult)
	return ret0
}

// Result indicates an expected call of Result.
func (mr *MockStatementMockRecorder) Result() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Result", reflect.TypeOf((*MockStatement)(nil).Result))
}
