// Code generated by MockGen. DO NOT EDIT.
// Source: article_grpc.pb.go
//
// Generated by this command:
//
//	mockgen -source=article_grpc.pb.go -destination=../../../mock/infra/grpc/article/mock_article_grpc.pb.go -package=article
//

// Package article is a generated GoMock package.
package article

import (
	context "context"
	reflect "reflect"

	article "github.com/miyamo2/blogapi.miyamo.today/federator/internal/infra/grpc/article"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockArticleServiceClient is a mock of ArticleServiceClient interface.
type MockArticleServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockArticleServiceClientMockRecorder
	isgomock struct{}
}

// MockArticleServiceClientMockRecorder is the mock recorder for MockArticleServiceClient.
type MockArticleServiceClientMockRecorder struct {
	mock *MockArticleServiceClient
}

// NewMockArticleServiceClient creates a new mock instance.
func NewMockArticleServiceClient(ctrl *gomock.Controller) *MockArticleServiceClient {
	mock := &MockArticleServiceClient{ctrl: ctrl}
	mock.recorder = &MockArticleServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleServiceClient) EXPECT() *MockArticleServiceClientMockRecorder {
	return m.recorder
}

// GetAllArticles mocks base method.
func (m *MockArticleServiceClient) GetAllArticles(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*article.GetAllArticlesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllArticles", varargs...)
	ret0, _ := ret[0].(*article.GetAllArticlesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllArticles indicates an expected call of GetAllArticles.
func (mr *MockArticleServiceClientMockRecorder) GetAllArticles(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllArticles", reflect.TypeOf((*MockArticleServiceClient)(nil).GetAllArticles), varargs...)
}

// GetArticleById mocks base method.
func (m *MockArticleServiceClient) GetArticleById(ctx context.Context, in *article.GetArticleByIdRequest, opts ...grpc.CallOption) (*article.GetArticleByIdResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetArticleById", varargs...)
	ret0, _ := ret[0].(*article.GetArticleByIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticleById indicates an expected call of GetArticleById.
func (mr *MockArticleServiceClientMockRecorder) GetArticleById(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticleById", reflect.TypeOf((*MockArticleServiceClient)(nil).GetArticleById), varargs...)
}

// GetNextArticles mocks base method.
func (m *MockArticleServiceClient) GetNextArticles(ctx context.Context, in *article.GetNextArticlesRequest, opts ...grpc.CallOption) (*article.GetNextArticlesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetNextArticles", varargs...)
	ret0, _ := ret[0].(*article.GetNextArticlesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextArticles indicates an expected call of GetNextArticles.
func (mr *MockArticleServiceClientMockRecorder) GetNextArticles(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextArticles", reflect.TypeOf((*MockArticleServiceClient)(nil).GetNextArticles), varargs...)
}

// GetPrevArticles mocks base method.
func (m *MockArticleServiceClient) GetPrevArticles(ctx context.Context, in *article.GetPrevArticlesRequest, opts ...grpc.CallOption) (*article.GetPrevArticlesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPrevArticles", varargs...)
	ret0, _ := ret[0].(*article.GetPrevArticlesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrevArticles indicates an expected call of GetPrevArticles.
func (mr *MockArticleServiceClientMockRecorder) GetPrevArticles(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrevArticles", reflect.TypeOf((*MockArticleServiceClient)(nil).GetPrevArticles), varargs...)
}

// MockArticleServiceServer is a mock of ArticleServiceServer interface.
type MockArticleServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockArticleServiceServerMockRecorder
	isgomock struct{}
}

// MockArticleServiceServerMockRecorder is the mock recorder for MockArticleServiceServer.
type MockArticleServiceServerMockRecorder struct {
	mock *MockArticleServiceServer
}

// NewMockArticleServiceServer creates a new mock instance.
func NewMockArticleServiceServer(ctrl *gomock.Controller) *MockArticleServiceServer {
	mock := &MockArticleServiceServer{ctrl: ctrl}
	mock.recorder = &MockArticleServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleServiceServer) EXPECT() *MockArticleServiceServerMockRecorder {
	return m.recorder
}

// GetAllArticles mocks base method.
func (m *MockArticleServiceServer) GetAllArticles(arg0 context.Context, arg1 *emptypb.Empty) (*article.GetAllArticlesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllArticles", arg0, arg1)
	ret0, _ := ret[0].(*article.GetAllArticlesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllArticles indicates an expected call of GetAllArticles.
func (mr *MockArticleServiceServerMockRecorder) GetAllArticles(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllArticles", reflect.TypeOf((*MockArticleServiceServer)(nil).GetAllArticles), arg0, arg1)
}

// GetArticleById mocks base method.
func (m *MockArticleServiceServer) GetArticleById(arg0 context.Context, arg1 *article.GetArticleByIdRequest) (*article.GetArticleByIdResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticleById", arg0, arg1)
	ret0, _ := ret[0].(*article.GetArticleByIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticleById indicates an expected call of GetArticleById.
func (mr *MockArticleServiceServerMockRecorder) GetArticleById(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticleById", reflect.TypeOf((*MockArticleServiceServer)(nil).GetArticleById), arg0, arg1)
}

// GetNextArticles mocks base method.
func (m *MockArticleServiceServer) GetNextArticles(arg0 context.Context, arg1 *article.GetNextArticlesRequest) (*article.GetNextArticlesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextArticles", arg0, arg1)
	ret0, _ := ret[0].(*article.GetNextArticlesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextArticles indicates an expected call of GetNextArticles.
func (mr *MockArticleServiceServerMockRecorder) GetNextArticles(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextArticles", reflect.TypeOf((*MockArticleServiceServer)(nil).GetNextArticles), arg0, arg1)
}

// GetPrevArticles mocks base method.
func (m *MockArticleServiceServer) GetPrevArticles(arg0 context.Context, arg1 *article.GetPrevArticlesRequest) (*article.GetPrevArticlesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrevArticles", arg0, arg1)
	ret0, _ := ret[0].(*article.GetPrevArticlesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrevArticles indicates an expected call of GetPrevArticles.
func (mr *MockArticleServiceServerMockRecorder) GetPrevArticles(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrevArticles", reflect.TypeOf((*MockArticleServiceServer)(nil).GetPrevArticles), arg0, arg1)
}

// mustEmbedUnimplementedArticleServiceServer mocks base method.
func (m *MockArticleServiceServer) mustEmbedUnimplementedArticleServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedArticleServiceServer")
}

// mustEmbedUnimplementedArticleServiceServer indicates an expected call of mustEmbedUnimplementedArticleServiceServer.
func (mr *MockArticleServiceServerMockRecorder) mustEmbedUnimplementedArticleServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedArticleServiceServer", reflect.TypeOf((*MockArticleServiceServer)(nil).mustEmbedUnimplementedArticleServiceServer))
}

// MockUnsafeArticleServiceServer is a mock of UnsafeArticleServiceServer interface.
type MockUnsafeArticleServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeArticleServiceServerMockRecorder
	isgomock struct{}
}

// MockUnsafeArticleServiceServerMockRecorder is the mock recorder for MockUnsafeArticleServiceServer.
type MockUnsafeArticleServiceServerMockRecorder struct {
	mock *MockUnsafeArticleServiceServer
}

// NewMockUnsafeArticleServiceServer creates a new mock instance.
func NewMockUnsafeArticleServiceServer(ctrl *gomock.Controller) *MockUnsafeArticleServiceServer {
	mock := &MockUnsafeArticleServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeArticleServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeArticleServiceServer) EXPECT() *MockUnsafeArticleServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedArticleServiceServer mocks base method.
func (m *MockUnsafeArticleServiceServer) mustEmbedUnimplementedArticleServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedArticleServiceServer")
}

// mustEmbedUnimplementedArticleServiceServer indicates an expected call of mustEmbedUnimplementedArticleServiceServer.
func (mr *MockUnsafeArticleServiceServerMockRecorder) mustEmbedUnimplementedArticleServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedArticleServiceServer", reflect.TypeOf((*MockUnsafeArticleServiceServer)(nil).mustEmbedUnimplementedArticleServiceServer))
}
