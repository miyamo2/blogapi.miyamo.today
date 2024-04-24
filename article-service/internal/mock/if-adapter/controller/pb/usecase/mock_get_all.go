// Code generated by MockGen. DO NOT EDIT.
// Source: get_all.go
//
// Generated by this command:
//
//	mockgen -source=get_all.go -destination=../../../../mock/if-adapter/controller/pb/usecase/mock_get_all.go -package=mock_usecase
//
// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	usecase "github.com/miyamo2/api.miyamo.today/article-service/internal/if-adapter/controller/pb/usecase"
	gomock "go.uber.org/mock/gomock"
)

// MockGetAll is a mock of GetAll interface.
type MockGetAll[T usecase.Tag, A usecase.Article[T], O usecase.GetAllOutDto[T, A]] struct {
	ctrl     *gomock.Controller
	recorder *MockGetAllMockRecorder[T, A, O]
}

// MockGetAllMockRecorder is the mock recorder for MockGetAll.
type MockGetAllMockRecorder[T usecase.Tag, A usecase.Article[T], O usecase.GetAllOutDto[T, A]] struct {
	mock *MockGetAll[T, A, O]
}

// NewMockGetAll creates a new mock instance.
func NewMockGetAll[T usecase.Tag, A usecase.Article[T], O usecase.GetAllOutDto[T, A]](ctrl *gomock.Controller) *MockGetAll[T, A, O] {
	mock := &MockGetAll[T, A, O]{ctrl: ctrl}
	mock.recorder = &MockGetAllMockRecorder[T, A, O]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetAll[T, A, O]) EXPECT() *MockGetAllMockRecorder[T, A, O] {
	return m.recorder
}

// Execute mocks base method.
func (m *MockGetAll[T, A, O]) Execute(ctx context.Context) (O, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx)
	ret0, _ := ret[0].(O)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockGetAllMockRecorder[T, A, O]) Execute(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockGetAll[T, A, O])(nil).Execute), ctx)
}
