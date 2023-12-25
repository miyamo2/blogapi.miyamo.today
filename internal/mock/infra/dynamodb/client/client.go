// Code generated by MockGen. DO NOT EDIT.
// Source: client.go
//
// Generated by this command:
//
//	mockgen -source=client.go -destination=../../../internal/mock/infra/dynamodb/client/client.go
//
// Package mock_client is a generated GoMock package.
package mock_client

import (
	context "context"
	reflect "reflect"

	dynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// BatchExecuteStatement mocks base method.
func (m *MockClient) BatchExecuteStatement(ctx context.Context, params *dynamodb.BatchExecuteStatementInput, optFns ...func(*dynamodb.Options)) (*dynamodb.BatchExecuteStatementOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchExecuteStatement", varargs...)
	ret0, _ := ret[0].(*dynamodb.BatchExecuteStatementOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchExecuteStatement indicates an expected call of BatchExecuteStatement.
func (mr *MockClientMockRecorder) BatchExecuteStatement(ctx, params any, optFns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchExecuteStatement", reflect.TypeOf((*MockClient)(nil).BatchExecuteStatement), varargs...)
}

// BatchGetItem mocks base method.
func (m *MockClient) BatchGetItem(ctx context.Context, params *dynamodb.BatchGetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.BatchGetItemOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchGetItem", varargs...)
	ret0, _ := ret[0].(*dynamodb.BatchGetItemOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetItem indicates an expected call of BatchGetItem.
func (mr *MockClientMockRecorder) BatchGetItem(ctx, params any, optFns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetItem", reflect.TypeOf((*MockClient)(nil).BatchGetItem), varargs...)
}

// BatchWriteItem mocks base method.
func (m *MockClient) BatchWriteItem(ctx context.Context, params *dynamodb.BatchWriteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.BatchWriteItemOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchWriteItem", varargs...)
	ret0, _ := ret[0].(*dynamodb.BatchWriteItemOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchWriteItem indicates an expected call of BatchWriteItem.
func (mr *MockClientMockRecorder) BatchWriteItem(ctx, params any, optFns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchWriteItem", reflect.TypeOf((*MockClient)(nil).BatchWriteItem), varargs...)
}

// DeleteItem mocks base method.
func (m *MockClient) DeleteItem(ctx context.Context, params *dynamodb.DeleteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteItem", varargs...)
	ret0, _ := ret[0].(*dynamodb.DeleteItemOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteItem indicates an expected call of DeleteItem.
func (mr *MockClientMockRecorder) DeleteItem(ctx, params any, optFns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteItem", reflect.TypeOf((*MockClient)(nil).DeleteItem), varargs...)
}

// ExecuteStatement mocks base method.
func (m *MockClient) ExecuteStatement(ctx context.Context, params *dynamodb.ExecuteStatementInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ExecuteStatementOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecuteStatement", varargs...)
	ret0, _ := ret[0].(*dynamodb.ExecuteStatementOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteStatement indicates an expected call of ExecuteStatement.
func (mr *MockClientMockRecorder) ExecuteStatement(ctx, params any, optFns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteStatement", reflect.TypeOf((*MockClient)(nil).ExecuteStatement), varargs...)
}

// ExecuteTransaction mocks base method.
func (m *MockClient) ExecuteTransaction(ctx context.Context, params *dynamodb.ExecuteTransactionInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ExecuteTransactionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecuteTransaction", varargs...)
	ret0, _ := ret[0].(*dynamodb.ExecuteTransactionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteTransaction indicates an expected call of ExecuteTransaction.
func (mr *MockClientMockRecorder) ExecuteTransaction(ctx, params any, optFns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteTransaction", reflect.TypeOf((*MockClient)(nil).ExecuteTransaction), varargs...)
}

// GetItem mocks base method.
func (m *MockClient) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetItem", varargs...)
	ret0, _ := ret[0].(*dynamodb.GetItemOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItem indicates an expected call of GetItem.
func (mr *MockClientMockRecorder) GetItem(ctx, params any, optFns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItem", reflect.TypeOf((*MockClient)(nil).GetItem), varargs...)
}

// PutItem mocks base method.
func (m *MockClient) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutItem", varargs...)
	ret0, _ := ret[0].(*dynamodb.PutItemOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutItem indicates an expected call of PutItem.
func (mr *MockClientMockRecorder) PutItem(ctx, params any, optFns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutItem", reflect.TypeOf((*MockClient)(nil).PutItem), varargs...)
}

// Query mocks base method.
func (m *MockClient) Query(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(*dynamodb.QueryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockClientMockRecorder) Query(ctx, params any, optFns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockClient)(nil).Query), varargs...)
}

// Scan mocks base method.
func (m *MockClient) Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Scan", varargs...)
	ret0, _ := ret[0].(*dynamodb.ScanOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Scan indicates an expected call of Scan.
func (mr *MockClientMockRecorder) Scan(ctx, params any, optFns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockClient)(nil).Scan), varargs...)
}

// TransactGetItems mocks base method.
func (m *MockClient) TransactGetItems(ctx context.Context, params *dynamodb.TransactGetItemsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.TransactGetItemsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TransactGetItems", varargs...)
	ret0, _ := ret[0].(*dynamodb.TransactGetItemsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactGetItems indicates an expected call of TransactGetItems.
func (mr *MockClientMockRecorder) TransactGetItems(ctx, params any, optFns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactGetItems", reflect.TypeOf((*MockClient)(nil).TransactGetItems), varargs...)
}

// TransactWriteItems mocks base method.
func (m *MockClient) TransactWriteItems(ctx context.Context, params *dynamodb.TransactWriteItemsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.TransactWriteItemsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TransactWriteItems", varargs...)
	ret0, _ := ret[0].(*dynamodb.TransactWriteItemsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactWriteItems indicates an expected call of TransactWriteItems.
func (mr *MockClientMockRecorder) TransactWriteItems(ctx, params any, optFns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactWriteItems", reflect.TypeOf((*MockClient)(nil).TransactWriteItems), varargs...)
}

// UpdateItem mocks base method.
func (m *MockClient) UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateItem", varargs...)
	ret0, _ := ret[0].(*dynamodb.UpdateItemOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateItem indicates an expected call of UpdateItem.
func (mr *MockClientMockRecorder) UpdateItem(ctx, params any, optFns ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItem", reflect.TypeOf((*MockClient)(nil).UpdateItem), varargs...)
}
