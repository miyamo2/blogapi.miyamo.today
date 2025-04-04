// Code generated by MockGen. DO NOT EDIT.
// Source: detach_tags.go
//
// Generated by this command:
//
//	mockgen -source=detach_tags.go -destination=../../../../mock/if-adapter/controller/pb/usecase/detach_tags.go -package=usecase
//

// Package usecase is a generated GoMock package.
package usecase

import (
	context "context"
	reflect "reflect"

	dto "blogapi.miyamo.today/blogging-event-service/internal/app/usecase/dto"
	gomock "go.uber.org/mock/gomock"
)

// MockDetachTags is a mock of DetachTags interface.
type MockDetachTags struct {
	ctrl     *gomock.Controller
	recorder *MockDetachTagsMockRecorder
	isgomock struct{}
}

// MockDetachTagsMockRecorder is the mock recorder for MockDetachTags.
type MockDetachTagsMockRecorder struct {
	mock *MockDetachTags
}

// NewMockDetachTags creates a new mock instance.
func NewMockDetachTags(ctrl *gomock.Controller) *MockDetachTags {
	mock := &MockDetachTags{ctrl: ctrl}
	mock.recorder = &MockDetachTagsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDetachTags) EXPECT() *MockDetachTagsMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockDetachTags) Execute(ctx context.Context, in *dto.DetachTagsInDto) (*dto.DetachTagsOutDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, in)
	ret0, _ := ret[0].(*dto.DetachTagsOutDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockDetachTagsMockRecorder) Execute(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockDetachTags)(nil).Execute), ctx, in)
}
