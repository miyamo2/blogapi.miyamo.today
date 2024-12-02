// Code generated by MockGen. DO NOT EDIT.
// Source: get_prev.go
//
// Generated by this command:
//
//	mockgen -source=get_prev.go -destination=../../../../mock/if-adapter/controller/pb/usecase/mock_get_prev.go -package=mock_usecase
//

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	dto "github.com/miyamo2/blogapi.miyamo.today/article-service/internal/app/usecase/dto"
	gomock "go.uber.org/mock/gomock"
)

// MockGetPrev is a mock of GetPrev interface.
type MockGetPrev struct {
	ctrl     *gomock.Controller
	recorder *MockGetPrevMockRecorder
	isgomock struct{}
}

// MockGetPrevMockRecorder is the mock recorder for MockGetPrev.
type MockGetPrevMockRecorder struct {
	mock *MockGetPrev
}

// NewMockGetPrev creates a new mock instance.
func NewMockGetPrev(ctrl *gomock.Controller) *MockGetPrev {
	mock := &MockGetPrev{ctrl: ctrl}
	mock.recorder = &MockGetPrevMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetPrev) EXPECT() *MockGetPrevMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockGetPrev) Execute(ctx context.Context, in dto.GetPrevInDto) (*dto.GetPrevOutDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, in)
	ret0, _ := ret[0].(*dto.GetPrevOutDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockGetPrevMockRecorder) Execute(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockGetPrev)(nil).Execute), ctx, in)
}
