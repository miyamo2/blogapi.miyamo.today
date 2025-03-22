// Code generated by MockGen. DO NOT EDIT.
// Source: article.connect.go
//
// Generated by this command:
//
//	mockgen -source=article.connect.go -destination=../../../../mock/infra/grpc/article/articleconnect/mock_article.connect.go -package=articleconnect
//

// Package articleconnect is a generated GoMock package.
package articleconnect

import (
	context "context"
	reflect "reflect"

	article "blogapi.miyamo.today/federator/internal/infra/grpc/article"
	connect "connectrpc.com/connect"
	gomock "go.uber.org/mock/gomock"
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
func (m *MockArticleServiceClient) GetAllArticles(arg0 context.Context, arg1 *connect.Request[emptypb.Empty]) (*connect.Response[article.GetAllArticlesResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllArticles", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[article.GetAllArticlesResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllArticles indicates an expected call of GetAllArticles.
func (mr *MockArticleServiceClientMockRecorder) GetAllArticles(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllArticles", reflect.TypeOf((*MockArticleServiceClient)(nil).GetAllArticles), arg0, arg1)
}

// GetArticleById mocks base method.
func (m *MockArticleServiceClient) GetArticleById(arg0 context.Context, arg1 *connect.Request[article.GetArticleByIdRequest]) (*connect.Response[article.GetArticleByIdResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticleById", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[article.GetArticleByIdResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticleById indicates an expected call of GetArticleById.
func (mr *MockArticleServiceClientMockRecorder) GetArticleById(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticleById", reflect.TypeOf((*MockArticleServiceClient)(nil).GetArticleById), arg0, arg1)
}

// GetNextArticles mocks base method.
func (m *MockArticleServiceClient) GetNextArticles(arg0 context.Context, arg1 *connect.Request[article.GetNextArticlesRequest]) (*connect.Response[article.GetNextArticlesResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextArticles", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[article.GetNextArticlesResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextArticles indicates an expected call of GetNextArticles.
func (mr *MockArticleServiceClientMockRecorder) GetNextArticles(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextArticles", reflect.TypeOf((*MockArticleServiceClient)(nil).GetNextArticles), arg0, arg1)
}

// GetPrevArticles mocks base method.
func (m *MockArticleServiceClient) GetPrevArticles(arg0 context.Context, arg1 *connect.Request[article.GetPrevArticlesRequest]) (*connect.Response[article.GetPrevArticlesResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrevArticles", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[article.GetPrevArticlesResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrevArticles indicates an expected call of GetPrevArticles.
func (mr *MockArticleServiceClientMockRecorder) GetPrevArticles(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrevArticles", reflect.TypeOf((*MockArticleServiceClient)(nil).GetPrevArticles), arg0, arg1)
}

// MockArticleServiceHandler is a mock of ArticleServiceHandler interface.
type MockArticleServiceHandler struct {
	ctrl     *gomock.Controller
	recorder *MockArticleServiceHandlerMockRecorder
	isgomock struct{}
}

// MockArticleServiceHandlerMockRecorder is the mock recorder for MockArticleServiceHandler.
type MockArticleServiceHandlerMockRecorder struct {
	mock *MockArticleServiceHandler
}

// NewMockArticleServiceHandler creates a new mock instance.
func NewMockArticleServiceHandler(ctrl *gomock.Controller) *MockArticleServiceHandler {
	mock := &MockArticleServiceHandler{ctrl: ctrl}
	mock.recorder = &MockArticleServiceHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleServiceHandler) EXPECT() *MockArticleServiceHandlerMockRecorder {
	return m.recorder
}

// GetAllArticles mocks base method.
func (m *MockArticleServiceHandler) GetAllArticles(arg0 context.Context, arg1 *connect.Request[emptypb.Empty]) (*connect.Response[article.GetAllArticlesResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllArticles", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[article.GetAllArticlesResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllArticles indicates an expected call of GetAllArticles.
func (mr *MockArticleServiceHandlerMockRecorder) GetAllArticles(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllArticles", reflect.TypeOf((*MockArticleServiceHandler)(nil).GetAllArticles), arg0, arg1)
}

// GetArticleById mocks base method.
func (m *MockArticleServiceHandler) GetArticleById(arg0 context.Context, arg1 *connect.Request[article.GetArticleByIdRequest]) (*connect.Response[article.GetArticleByIdResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticleById", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[article.GetArticleByIdResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticleById indicates an expected call of GetArticleById.
func (mr *MockArticleServiceHandlerMockRecorder) GetArticleById(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticleById", reflect.TypeOf((*MockArticleServiceHandler)(nil).GetArticleById), arg0, arg1)
}

// GetNextArticles mocks base method.
func (m *MockArticleServiceHandler) GetNextArticles(arg0 context.Context, arg1 *connect.Request[article.GetNextArticlesRequest]) (*connect.Response[article.GetNextArticlesResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextArticles", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[article.GetNextArticlesResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextArticles indicates an expected call of GetNextArticles.
func (mr *MockArticleServiceHandlerMockRecorder) GetNextArticles(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextArticles", reflect.TypeOf((*MockArticleServiceHandler)(nil).GetNextArticles), arg0, arg1)
}

// GetPrevArticles mocks base method.
func (m *MockArticleServiceHandler) GetPrevArticles(arg0 context.Context, arg1 *connect.Request[article.GetPrevArticlesRequest]) (*connect.Response[article.GetPrevArticlesResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrevArticles", arg0, arg1)
	ret0, _ := ret[0].(*connect.Response[article.GetPrevArticlesResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrevArticles indicates an expected call of GetPrevArticles.
func (mr *MockArticleServiceHandlerMockRecorder) GetPrevArticles(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrevArticles", reflect.TypeOf((*MockArticleServiceHandler)(nil).GetPrevArticles), arg0, arg1)
}
