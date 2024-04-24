// Code generated by MockGen. DO NOT EDIT.
// Source: get_by_id.go
//
// Generated by this command:
//
//	mockgen -source=get_by_id.go -destination=../../../../mock/if-adapter/controller/pb/usecase/mock_get_by_id.go -package=mock_usecase
//
// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	usecase "github.com/miyamo2/blogapi-tag-service/internal/if-adapter/controller/pb/usecase"
	gomock "go.uber.org/mock/gomock"
)

// MockGetById is a mock of GetById interface.
type MockGetById[I usecase.GetByIdInDto, A usecase.Article, T usecase.Tag[A]] struct {
	ctrl     *gomock.Controller
	recorder *MockGetByIdMockRecorder[I, A, T]
}

// MockGetByIdMockRecorder is the mock recorder for MockGetById.
type MockGetByIdMockRecorder[I usecase.GetByIdInDto, A usecase.Article, T usecase.Tag[A]] struct {
	mock *MockGetById[I, A, T]
}

// NewMockGetById creates a new mock instance.
func NewMockGetById[I usecase.GetByIdInDto, A usecase.Article, T usecase.Tag[A]](ctrl *gomock.Controller) *MockGetById[I, A, T] {
	mock := &MockGetById[I, A, T]{ctrl: ctrl}
	mock.recorder = &MockGetByIdMockRecorder[I, A, T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetById[I, A, T]) EXPECT() *MockGetByIdMockRecorder[I, A, T] {
	return m.recorder
}

// Execute mocks base method.
func (m *MockGetById[I, A, T]) Execute(ctx context.Context, in I) (T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, in)
	ret0, _ := ret[0].(T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockGetByIdMockRecorder[I, A, T]) Execute(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockGetById[I, A, T])(nil).Execute), ctx, in)
}
