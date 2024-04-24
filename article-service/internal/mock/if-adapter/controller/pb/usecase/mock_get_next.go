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

	usecase "github.com/miyamo2/blogapi-article-service/internal/if-adapter/controller/pb/usecase"
	gomock "go.uber.org/mock/gomock"
)

// MockGetNext is a mock of GetNext interface.
type MockGetNext[I usecase.GetNextInDto, T usecase.Tag, A usecase.Article[T], O usecase.GetNextOutDto[T, A]] struct {
	ctrl     *gomock.Controller
	recorder *MockGetNextMockRecorder[I, T, A, O]
}

// MockGetNextMockRecorder is the mock recorder for MockGetNext.
type MockGetNextMockRecorder[I usecase.GetNextInDto, T usecase.Tag, A usecase.Article[T], O usecase.GetNextOutDto[T, A]] struct {
	mock *MockGetNext[I, T, A, O]
}

// NewMockGetNext creates a new mock instance.
func NewMockGetNext[I usecase.GetNextInDto, T usecase.Tag, A usecase.Article[T], O usecase.GetNextOutDto[T, A]](ctrl *gomock.Controller) *MockGetNext[I, T, A, O] {
	mock := &MockGetNext[I, T, A, O]{ctrl: ctrl}
	mock.recorder = &MockGetNextMockRecorder[I, T, A, O]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetNext[I, T, A, O]) EXPECT() *MockGetNextMockRecorder[I, T, A, O] {
	return m.recorder
}

// Execute mocks base method.
func (m *MockGetNext[I, T, A, O]) Execute(ctx context.Context, in I) (O, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, in)
	ret0, _ := ret[0].(O)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockGetNextMockRecorder[I, T, A, O]) Execute(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockGetNext[I, T, A, O])(nil).Execute), ctx, in)
}
