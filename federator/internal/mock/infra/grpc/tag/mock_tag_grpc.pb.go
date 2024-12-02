// Code generated by MockGen. DO NOT EDIT.
// Source: tag_grpc.pb.go
//
// Generated by this command:
//
//	mockgen -source=tag_grpc.pb.go -destination=../../../mock/infra/grpc/tag/mock_tag_grpc.pb.go -package=mock_tag
//

// Package mock_tag is a generated GoMock package.
package mock_tag

import (
	context "context"
	reflect "reflect"

	tag "github.com/miyamo2/blogapi.miyamo.today/federator/internal/infra/grpc/tag"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockTagServiceClient is a mock of TagServiceClient interface.
type MockTagServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockTagServiceClientMockRecorder
	isgomock struct{}
}

// MockTagServiceClientMockRecorder is the mock recorder for MockTagServiceClient.
type MockTagServiceClientMockRecorder struct {
	mock *MockTagServiceClient
}

// NewMockTagServiceClient creates a new mock instance.
func NewMockTagServiceClient(ctrl *gomock.Controller) *MockTagServiceClient {
	mock := &MockTagServiceClient{ctrl: ctrl}
	mock.recorder = &MockTagServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTagServiceClient) EXPECT() *MockTagServiceClientMockRecorder {
	return m.recorder
}

// GetAllTags mocks base method.
func (m *MockTagServiceClient) GetAllTags(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*tag.GetAllTagsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllTags", varargs...)
	ret0, _ := ret[0].(*tag.GetAllTagsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTags indicates an expected call of GetAllTags.
func (mr *MockTagServiceClientMockRecorder) GetAllTags(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTags", reflect.TypeOf((*MockTagServiceClient)(nil).GetAllTags), varargs...)
}

// GetNextTags mocks base method.
func (m *MockTagServiceClient) GetNextTags(ctx context.Context, in *tag.GetNextTagsRequest, opts ...grpc.CallOption) (*tag.GetNextTagResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetNextTags", varargs...)
	ret0, _ := ret[0].(*tag.GetNextTagResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextTags indicates an expected call of GetNextTags.
func (mr *MockTagServiceClientMockRecorder) GetNextTags(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextTags", reflect.TypeOf((*MockTagServiceClient)(nil).GetNextTags), varargs...)
}

// GetPrevTags mocks base method.
func (m *MockTagServiceClient) GetPrevTags(ctx context.Context, in *tag.GetPrevTagsRequest, opts ...grpc.CallOption) (*tag.GetPrevTagResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPrevTags", varargs...)
	ret0, _ := ret[0].(*tag.GetPrevTagResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrevTags indicates an expected call of GetPrevTags.
func (mr *MockTagServiceClientMockRecorder) GetPrevTags(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrevTags", reflect.TypeOf((*MockTagServiceClient)(nil).GetPrevTags), varargs...)
}

// GetTagById mocks base method.
func (m *MockTagServiceClient) GetTagById(ctx context.Context, in *tag.GetTagByIdRequest, opts ...grpc.CallOption) (*tag.GetTagByIdResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTagById", varargs...)
	ret0, _ := ret[0].(*tag.GetTagByIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagById indicates an expected call of GetTagById.
func (mr *MockTagServiceClientMockRecorder) GetTagById(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagById", reflect.TypeOf((*MockTagServiceClient)(nil).GetTagById), varargs...)
}

// MockTagServiceServer is a mock of TagServiceServer interface.
type MockTagServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockTagServiceServerMockRecorder
	isgomock struct{}
}

// MockTagServiceServerMockRecorder is the mock recorder for MockTagServiceServer.
type MockTagServiceServerMockRecorder struct {
	mock *MockTagServiceServer
}

// NewMockTagServiceServer creates a new mock instance.
func NewMockTagServiceServer(ctrl *gomock.Controller) *MockTagServiceServer {
	mock := &MockTagServiceServer{ctrl: ctrl}
	mock.recorder = &MockTagServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTagServiceServer) EXPECT() *MockTagServiceServerMockRecorder {
	return m.recorder
}

// GetAllTags mocks base method.
func (m *MockTagServiceServer) GetAllTags(arg0 context.Context, arg1 *emptypb.Empty) (*tag.GetAllTagsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTags", arg0, arg1)
	ret0, _ := ret[0].(*tag.GetAllTagsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTags indicates an expected call of GetAllTags.
func (mr *MockTagServiceServerMockRecorder) GetAllTags(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTags", reflect.TypeOf((*MockTagServiceServer)(nil).GetAllTags), arg0, arg1)
}

// GetNextTags mocks base method.
func (m *MockTagServiceServer) GetNextTags(arg0 context.Context, arg1 *tag.GetNextTagsRequest) (*tag.GetNextTagResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextTags", arg0, arg1)
	ret0, _ := ret[0].(*tag.GetNextTagResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextTags indicates an expected call of GetNextTags.
func (mr *MockTagServiceServerMockRecorder) GetNextTags(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextTags", reflect.TypeOf((*MockTagServiceServer)(nil).GetNextTags), arg0, arg1)
}

// GetPrevTags mocks base method.
func (m *MockTagServiceServer) GetPrevTags(arg0 context.Context, arg1 *tag.GetPrevTagsRequest) (*tag.GetPrevTagResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrevTags", arg0, arg1)
	ret0, _ := ret[0].(*tag.GetPrevTagResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrevTags indicates an expected call of GetPrevTags.
func (mr *MockTagServiceServerMockRecorder) GetPrevTags(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrevTags", reflect.TypeOf((*MockTagServiceServer)(nil).GetPrevTags), arg0, arg1)
}

// GetTagById mocks base method.
func (m *MockTagServiceServer) GetTagById(arg0 context.Context, arg1 *tag.GetTagByIdRequest) (*tag.GetTagByIdResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTagById", arg0, arg1)
	ret0, _ := ret[0].(*tag.GetTagByIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTagById indicates an expected call of GetTagById.
func (mr *MockTagServiceServerMockRecorder) GetTagById(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTagById", reflect.TypeOf((*MockTagServiceServer)(nil).GetTagById), arg0, arg1)
}

// mustEmbedUnimplementedTagServiceServer mocks base method.
func (m *MockTagServiceServer) mustEmbedUnimplementedTagServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedTagServiceServer")
}

// mustEmbedUnimplementedTagServiceServer indicates an expected call of mustEmbedUnimplementedTagServiceServer.
func (mr *MockTagServiceServerMockRecorder) mustEmbedUnimplementedTagServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedTagServiceServer", reflect.TypeOf((*MockTagServiceServer)(nil).mustEmbedUnimplementedTagServiceServer))
}

// MockUnsafeTagServiceServer is a mock of UnsafeTagServiceServer interface.
type MockUnsafeTagServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeTagServiceServerMockRecorder
	isgomock struct{}
}

// MockUnsafeTagServiceServerMockRecorder is the mock recorder for MockUnsafeTagServiceServer.
type MockUnsafeTagServiceServerMockRecorder struct {
	mock *MockUnsafeTagServiceServer
}

// NewMockUnsafeTagServiceServer creates a new mock instance.
func NewMockUnsafeTagServiceServer(ctrl *gomock.Controller) *MockUnsafeTagServiceServer {
	mock := &MockUnsafeTagServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeTagServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeTagServiceServer) EXPECT() *MockUnsafeTagServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedTagServiceServer mocks base method.
func (m *MockUnsafeTagServiceServer) mustEmbedUnimplementedTagServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedTagServiceServer")
}

// mustEmbedUnimplementedTagServiceServer indicates an expected call of mustEmbedUnimplementedTagServiceServer.
func (mr *MockUnsafeTagServiceServerMockRecorder) mustEmbedUnimplementedTagServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedTagServiceServer", reflect.TypeOf((*MockUnsafeTagServiceServer)(nil).mustEmbedUnimplementedTagServiceServer))
}
