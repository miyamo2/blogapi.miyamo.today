// Code generated by MockGen. DO NOT EDIT.
// Source: converter.go
//
// Generated by this command:
//
//	mockgen -source=converter.go -destination=../../../../mock/if-adapter/controller/pb/presenter/mock_converter.go -package=presenter
//
// Package presenter is a generated GoMock package.
package presenter

import (
	context "context"
	reflect "reflect"

	usecase "github.com/miyamo2/api.miyamo.today/article-service/internal/if-adapter/controller/pb/usecase"
	pb "github.com/miyamo2/api.miyamo.today/protogen/article/server/pb"
	gomock "go.uber.org/mock/gomock"
)

// MockToGetNextConverter is a mock of ToGetNextConverter interface.
type MockToGetNextConverter[T usecase.Tag, A usecase.Article[T], O usecase.GetNextOutDto[T, A]] struct {
	ctrl     *gomock.Controller
	recorder *MockToGetNextConverterMockRecorder[T, A, O]
}

// MockToGetNextConverterMockRecorder is the mock recorder for MockToGetNextConverter.
type MockToGetNextConverterMockRecorder[T usecase.Tag, A usecase.Article[T], O usecase.GetNextOutDto[T, A]] struct {
	mock *MockToGetNextConverter[T, A, O]
}

// NewMockToGetNextConverter creates a new mock instance.
func NewMockToGetNextConverter[T usecase.Tag, A usecase.Article[T], O usecase.GetNextOutDto[T, A]](ctrl *gomock.Controller) *MockToGetNextConverter[T, A, O] {
	mock := &MockToGetNextConverter[T, A, O]{ctrl: ctrl}
	mock.recorder = &MockToGetNextConverterMockRecorder[T, A, O]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToGetNextConverter[T, A, O]) EXPECT() *MockToGetNextConverterMockRecorder[T, A, O] {
	return m.recorder
}

// ToGetNextArticlesResponse mocks base method.
func (m *MockToGetNextConverter[T, A, O]) ToGetNextArticlesResponse(ctx context.Context, from O) (*pb.GetNextArticlesResponse, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToGetNextArticlesResponse", ctx, from)
	ret0, _ := ret[0].(*pb.GetNextArticlesResponse)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ToGetNextArticlesResponse indicates an expected call of ToGetNextArticlesResponse.
func (mr *MockToGetNextConverterMockRecorder[T, A, O]) ToGetNextArticlesResponse(ctx, from any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToGetNextArticlesResponse", reflect.TypeOf((*MockToGetNextConverter[T, A, O])(nil).ToGetNextArticlesResponse), ctx, from)
}

// MockToGetAllConverter is a mock of ToGetAllConverter interface.
type MockToGetAllConverter[T usecase.Tag, A usecase.Article[T], O usecase.GetAllOutDto[T, A]] struct {
	ctrl     *gomock.Controller
	recorder *MockToGetAllConverterMockRecorder[T, A, O]
}

// MockToGetAllConverterMockRecorder is the mock recorder for MockToGetAllConverter.
type MockToGetAllConverterMockRecorder[T usecase.Tag, A usecase.Article[T], O usecase.GetAllOutDto[T, A]] struct {
	mock *MockToGetAllConverter[T, A, O]
}

// NewMockToGetAllConverter creates a new mock instance.
func NewMockToGetAllConverter[T usecase.Tag, A usecase.Article[T], O usecase.GetAllOutDto[T, A]](ctrl *gomock.Controller) *MockToGetAllConverter[T, A, O] {
	mock := &MockToGetAllConverter[T, A, O]{ctrl: ctrl}
	mock.recorder = &MockToGetAllConverterMockRecorder[T, A, O]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToGetAllConverter[T, A, O]) EXPECT() *MockToGetAllConverterMockRecorder[T, A, O] {
	return m.recorder
}

// ToGetAllArticlesResponse mocks base method.
func (m *MockToGetAllConverter[T, A, O]) ToGetAllArticlesResponse(ctx context.Context, from O) (*pb.GetAllArticlesResponse, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToGetAllArticlesResponse", ctx, from)
	ret0, _ := ret[0].(*pb.GetAllArticlesResponse)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ToGetAllArticlesResponse indicates an expected call of ToGetAllArticlesResponse.
func (mr *MockToGetAllConverterMockRecorder[T, A, O]) ToGetAllArticlesResponse(ctx, from any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToGetAllArticlesResponse", reflect.TypeOf((*MockToGetAllConverter[T, A, O])(nil).ToGetAllArticlesResponse), ctx, from)
}

// MockToGetByIdConverter is a mock of ToGetByIdConverter interface.
type MockToGetByIdConverter[T usecase.Tag, A usecase.Article[T]] struct {
	ctrl     *gomock.Controller
	recorder *MockToGetByIdConverterMockRecorder[T, A]
}

// MockToGetByIdConverterMockRecorder is the mock recorder for MockToGetByIdConverter.
type MockToGetByIdConverterMockRecorder[T usecase.Tag, A usecase.Article[T]] struct {
	mock *MockToGetByIdConverter[T, A]
}

// NewMockToGetByIdConverter creates a new mock instance.
func NewMockToGetByIdConverter[T usecase.Tag, A usecase.Article[T]](ctrl *gomock.Controller) *MockToGetByIdConverter[T, A] {
	mock := &MockToGetByIdConverter[T, A]{ctrl: ctrl}
	mock.recorder = &MockToGetByIdConverterMockRecorder[T, A]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToGetByIdConverter[T, A]) EXPECT() *MockToGetByIdConverterMockRecorder[T, A] {
	return m.recorder
}

// ToGetByIdArticlesResponse mocks base method.
func (m *MockToGetByIdConverter[T, A]) ToGetByIdArticlesResponse(ctx context.Context, from A) (*pb.GetArticleByIdResponse, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToGetByIdArticlesResponse", ctx, from)
	ret0, _ := ret[0].(*pb.GetArticleByIdResponse)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ToGetByIdArticlesResponse indicates an expected call of ToGetByIdArticlesResponse.
func (mr *MockToGetByIdConverterMockRecorder[T, A]) ToGetByIdArticlesResponse(ctx, from any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToGetByIdArticlesResponse", reflect.TypeOf((*MockToGetByIdConverter[T, A])(nil).ToGetByIdArticlesResponse), ctx, from)
}

// MockToGetPrevConverter is a mock of ToGetPrevConverter interface.
type MockToGetPrevConverter[T usecase.Tag, A usecase.Article[T], O usecase.GetPrevOutDto[T, A]] struct {
	ctrl     *gomock.Controller
	recorder *MockToGetPrevConverterMockRecorder[T, A, O]
}

// MockToGetPrevConverterMockRecorder is the mock recorder for MockToGetPrevConverter.
type MockToGetPrevConverterMockRecorder[T usecase.Tag, A usecase.Article[T], O usecase.GetPrevOutDto[T, A]] struct {
	mock *MockToGetPrevConverter[T, A, O]
}

// NewMockToGetPrevConverter creates a new mock instance.
func NewMockToGetPrevConverter[T usecase.Tag, A usecase.Article[T], O usecase.GetPrevOutDto[T, A]](ctrl *gomock.Controller) *MockToGetPrevConverter[T, A, O] {
	mock := &MockToGetPrevConverter[T, A, O]{ctrl: ctrl}
	mock.recorder = &MockToGetPrevConverterMockRecorder[T, A, O]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToGetPrevConverter[T, A, O]) EXPECT() *MockToGetPrevConverterMockRecorder[T, A, O] {
	return m.recorder
}

// ToGetPrevArticlesResponse mocks base method.
func (m *MockToGetPrevConverter[T, A, O]) ToGetPrevArticlesResponse(ctx context.Context, from O) (*pb.GetPrevArticlesResponse, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToGetPrevArticlesResponse", ctx, from)
	ret0, _ := ret[0].(*pb.GetPrevArticlesResponse)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ToGetPrevArticlesResponse indicates an expected call of ToGetPrevArticlesResponse.
func (mr *MockToGetPrevConverterMockRecorder[T, A, O]) ToGetPrevArticlesResponse(ctx, from any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToGetPrevArticlesResponse", reflect.TypeOf((*MockToGetPrevConverter[T, A, O])(nil).ToGetPrevArticlesResponse), ctx, from)
}
