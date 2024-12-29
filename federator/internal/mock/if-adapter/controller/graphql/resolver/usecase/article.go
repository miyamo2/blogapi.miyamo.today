// Code generated by MockGen. DO NOT EDIT.
// Source: article.go
//
// Generated by this command:
//
//	mockgen -source=article.go -destination=../../../../../mock/if-adapter/controller/graphql/resolver/usecase/article.go -package=usecase
//

// Package usecase is a generated GoMock package.
package usecase

import (
	context "context"
	reflect "reflect"

	dto "blogapi.miyamo.today/federator/internal/app/usecase/dto"
	gomock "go.uber.org/mock/gomock"
)

// MockArticle is a mock of Article interface.
type MockArticle struct {
	ctrl     *gomock.Controller
	recorder *MockArticleMockRecorder
	isgomock struct{}
}

// MockArticleMockRecorder is the mock recorder for MockArticle.
type MockArticleMockRecorder struct {
	mock *MockArticle
}

// NewMockArticle creates a new mock instance.
func NewMockArticle(ctrl *gomock.Controller) *MockArticle {
	mock := &MockArticle{ctrl: ctrl}
	mock.recorder = &MockArticleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticle) EXPECT() *MockArticleMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockArticle) Execute(ctx context.Context, in dto.ArticleInDTO) (dto.ArticleOutDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, in)
	ret0, _ := ret[0].(dto.ArticleOutDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockArticleMockRecorder) Execute(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockArticle)(nil).Execute), ctx, in)
}

// MockArticles is a mock of Articles interface.
type MockArticles struct {
	ctrl     *gomock.Controller
	recorder *MockArticlesMockRecorder
	isgomock struct{}
}

// MockArticlesMockRecorder is the mock recorder for MockArticles.
type MockArticlesMockRecorder struct {
	mock *MockArticles
}

// NewMockArticles creates a new mock instance.
func NewMockArticles(ctrl *gomock.Controller) *MockArticles {
	mock := &MockArticles{ctrl: ctrl}
	mock.recorder = &MockArticlesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticles) EXPECT() *MockArticlesMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockArticles) Execute(ctx context.Context, in dto.ArticlesInDTO) (dto.ArticlesOutDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, in)
	ret0, _ := ret[0].(dto.ArticlesOutDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockArticlesMockRecorder) Execute(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockArticles)(nil).Execute), ctx, in)
}
