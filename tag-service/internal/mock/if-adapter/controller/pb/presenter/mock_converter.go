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

	dto "blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
	grpc "blogapi.miyamo.today/tag-service/internal/infra/grpc"
	connect "connectrpc.com/connect"
	gomock "go.uber.org/mock/gomock"
)

// MockToGetByIdConverter is a mock of ToGetByIdConverter interface.
type MockToGetByIdConverter struct {
	ctrl     *gomock.Controller
	recorder *MockToGetByIdConverterMockRecorder
	isgomock struct{}
}

// MockToGetByIdConverterMockRecorder is the mock recorder for MockToGetByIdConverter.
type MockToGetByIdConverterMockRecorder struct {
	mock *MockToGetByIdConverter
}

// NewMockToGetByIdConverter creates a new mock instance.
func NewMockToGetByIdConverter(ctrl *gomock.Controller) *MockToGetByIdConverter {
	mock := &MockToGetByIdConverter{ctrl: ctrl}
	mock.recorder = &MockToGetByIdConverterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToGetByIdConverter) EXPECT() *MockToGetByIdConverterMockRecorder {
	return m.recorder
}

// ToGetByIdTagResponse mocks base method.
func (m *MockToGetByIdConverter) ToGetByIdTagResponse(ctx context.Context, from *dto.GetByIdOutDto) (*connect.Response[grpc.GetTagByIdResponse], bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToGetByIdTagResponse", ctx, from)
	ret0, _ := ret[0].(*connect.Response[grpc.GetTagByIdResponse])
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ToGetByIdTagResponse indicates an expected call of ToGetByIdTagResponse.
func (mr *MockToGetByIdConverterMockRecorder) ToGetByIdTagResponse(ctx, from any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToGetByIdTagResponse", reflect.TypeOf((*MockToGetByIdConverter)(nil).ToGetByIdTagResponse), ctx, from)
}

// MockToGetAllConverter is a mock of ToGetAllConverter interface.
type MockToGetAllConverter struct {
	ctrl     *gomock.Controller
	recorder *MockToGetAllConverterMockRecorder
	isgomock struct{}
}

// MockToGetAllConverterMockRecorder is the mock recorder for MockToGetAllConverter.
type MockToGetAllConverterMockRecorder struct {
	mock *MockToGetAllConverter
}

// NewMockToGetAllConverter creates a new mock instance.
func NewMockToGetAllConverter(ctrl *gomock.Controller) *MockToGetAllConverter {
	mock := &MockToGetAllConverter{ctrl: ctrl}
	mock.recorder = &MockToGetAllConverterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToGetAllConverter) EXPECT() *MockToGetAllConverterMockRecorder {
	return m.recorder
}

// ToGetAllTagsResponse mocks base method.
func (m *MockToGetAllConverter) ToGetAllTagsResponse(ctx context.Context, from *dto.GetAllOutDto) (*connect.Response[grpc.GetAllTagsResponse], bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToGetAllTagsResponse", ctx, from)
	ret0, _ := ret[0].(*connect.Response[grpc.GetAllTagsResponse])
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ToGetAllTagsResponse indicates an expected call of ToGetAllTagsResponse.
func (mr *MockToGetAllConverterMockRecorder) ToGetAllTagsResponse(ctx, from any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToGetAllTagsResponse", reflect.TypeOf((*MockToGetAllConverter)(nil).ToGetAllTagsResponse), ctx, from)
}

// MockToGetNextConverter is a mock of ToGetNextConverter interface.
type MockToGetNextConverter struct {
	ctrl     *gomock.Controller
	recorder *MockToGetNextConverterMockRecorder
	isgomock struct{}
}

// MockToGetNextConverterMockRecorder is the mock recorder for MockToGetNextConverter.
type MockToGetNextConverterMockRecorder struct {
	mock *MockToGetNextConverter
}

// NewMockToGetNextConverter creates a new mock instance.
func NewMockToGetNextConverter(ctrl *gomock.Controller) *MockToGetNextConverter {
	mock := &MockToGetNextConverter{ctrl: ctrl}
	mock.recorder = &MockToGetNextConverterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToGetNextConverter) EXPECT() *MockToGetNextConverterMockRecorder {
	return m.recorder
}

// ToGetNextTagsResponse mocks base method.
func (m *MockToGetNextConverter) ToGetNextTagsResponse(ctx context.Context, from *dto.GetNextOutDto) (*connect.Response[grpc.GetNextTagResponse], bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToGetNextTagsResponse", ctx, from)
	ret0, _ := ret[0].(*connect.Response[grpc.GetNextTagResponse])
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ToGetNextTagsResponse indicates an expected call of ToGetNextTagsResponse.
func (mr *MockToGetNextConverterMockRecorder) ToGetNextTagsResponse(ctx, from any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToGetNextTagsResponse", reflect.TypeOf((*MockToGetNextConverter)(nil).ToGetNextTagsResponse), ctx, from)
}

// MockToGetPrevConverter is a mock of ToGetPrevConverter interface.
type MockToGetPrevConverter struct {
	ctrl     *gomock.Controller
	recorder *MockToGetPrevConverterMockRecorder
	isgomock struct{}
}

// MockToGetPrevConverterMockRecorder is the mock recorder for MockToGetPrevConverter.
type MockToGetPrevConverterMockRecorder struct {
	mock *MockToGetPrevConverter
}

// NewMockToGetPrevConverter creates a new mock instance.
func NewMockToGetPrevConverter(ctrl *gomock.Controller) *MockToGetPrevConverter {
	mock := &MockToGetPrevConverter{ctrl: ctrl}
	mock.recorder = &MockToGetPrevConverterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToGetPrevConverter) EXPECT() *MockToGetPrevConverterMockRecorder {
	return m.recorder
}

// ToGetPrevTagsResponse mocks base method.
func (m *MockToGetPrevConverter) ToGetPrevTagsResponse(ctx context.Context, from *dto.GetPrevOutDto) (*connect.Response[grpc.GetPrevTagResponse], bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToGetPrevTagsResponse", ctx, from)
	ret0, _ := ret[0].(*connect.Response[grpc.GetPrevTagResponse])
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ToGetPrevTagsResponse indicates an expected call of ToGetPrevTagsResponse.
func (mr *MockToGetPrevConverterMockRecorder) ToGetPrevTagsResponse(ctx, from any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToGetPrevTagsResponse", reflect.TypeOf((*MockToGetPrevConverter)(nil).ToGetPrevTagsResponse), ctx, from)
}
