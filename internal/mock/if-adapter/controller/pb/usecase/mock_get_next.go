// Code generated by MockGen. DO NOT EDIT.
// Source: get_next.go
//
// Generated by this command:
//
//	mockgen -source=get_next.go -destination=../../../../mock/if-adapter/controller/pb/usecase/mock_get_next.go -package=mock_usecase
//
// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	usecase "github.com/miyamo2/blogapi-tag-service/internal/if-adapter/controller/pb/usecase"
	gomock "go.uber.org/mock/gomock"
)

// MockGetNext is a mock of GetNext interface.
type MockGetNext[I usecase.GetNextInDto, A usecase.Article, T usecase.Tag[A], O usecase.GetNextOutDto[A, T]] struct {
	ctrl     *gomock.Controller
	recorder *MockGetNextMockRecorder[I, A, T, O]
}

// MockGetNextMockRecorder is the mock recorder for MockGetNext.
type MockGetNextMockRecorder[I usecase.GetNextInDto, A usecase.Article, T usecase.Tag[A], O usecase.GetNextOutDto[A, T]] struct {
	mock *MockGetNext[I, A, T, O]
}

// NewMockGetNext creates a new mock instance.
func NewMockGetNext[I usecase.GetNextInDto, A usecase.Article, T usecase.Tag[A], O usecase.GetNextOutDto[A, T]](ctrl *gomock.Controller) *MockGetNext[I, A, T, O] {
	mock := &MockGetNext[I, A, T, O]{ctrl: ctrl}
	mock.recorder = &MockGetNextMockRecorder[I, A, T, O]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetNext[I, A, T, O]) EXPECT() *MockGetNextMockRecorder[I, A, T, O] {
	return m.recorder
}

// Execute mocks base method.
func (m *MockGetNext[I, A, T, O]) Execute(ctx context.Context, in I) (O, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, in)
	ret0, _ := ret[0].(O)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockGetNextMockRecorder[I, A, T, O]) Execute(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockGetNext[I, A, T, O])(nil).Execute), ctx, in)
}
