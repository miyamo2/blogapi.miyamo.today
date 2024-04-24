// Code generated by MockGen. DO NOT EDIT.
// Source: article.go
//
// Generated by this command:
//
//	mockgen -source=article.go -destination=../../../../../mock/if-adapter/controller/graphql/resolver/usecase/mock_article.go -package=usecase
//
// Package usecase is a generated GoMock package.
package usecase

import (
	context "context"
	reflect "reflect"

	dto "github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/controller/graphql/resolver/usecase/dto"
	gomock "go.uber.org/mock/gomock"
)

// MockArticle is a mock of Article interface.
type MockArticle[I dto.ArticleInDto, T dto.Tag, AT dto.ArticleTag[T], O dto.ArticleOutDto[T, AT]] struct {
	ctrl     *gomock.Controller
	recorder *MockArticleMockRecorder[I, T, AT, O]
}

// MockArticleMockRecorder is the mock recorder for MockArticle.
type MockArticleMockRecorder[I dto.ArticleInDto, T dto.Tag, AT dto.ArticleTag[T], O dto.ArticleOutDto[T, AT]] struct {
	mock *MockArticle[I, T, AT, O]
}

// NewMockArticle creates a new mock instance.
func NewMockArticle[I dto.ArticleInDto, T dto.Tag, AT dto.ArticleTag[T], O dto.ArticleOutDto[T, AT]](ctrl *gomock.Controller) *MockArticle[I, T, AT, O] {
	mock := &MockArticle[I, T, AT, O]{ctrl: ctrl}
	mock.recorder = &MockArticleMockRecorder[I, T, AT, O]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticle[I, T, AT, O]) EXPECT() *MockArticleMockRecorder[I, T, AT, O] {
	return m.recorder
}

// Execute mocks base method.
func (m *MockArticle[I, T, AT, O]) Execute(ctx context.Context, in I) (O, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, in)
	ret0, _ := ret[0].(O)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockArticleMockRecorder[I, T, AT, O]) Execute(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockArticle[I, T, AT, O])(nil).Execute), ctx, in)
}

// MockArticles is a mock of Articles interface.
type MockArticles[I dto.ArticlesInDto, T dto.Tag, AT dto.ArticleTag[T], O dto.ArticlesOutDto[T, AT]] struct {
	ctrl     *gomock.Controller
	recorder *MockArticlesMockRecorder[I, T, AT, O]
}

// MockArticlesMockRecorder is the mock recorder for MockArticles.
type MockArticlesMockRecorder[I dto.ArticlesInDto, T dto.Tag, AT dto.ArticleTag[T], O dto.ArticlesOutDto[T, AT]] struct {
	mock *MockArticles[I, T, AT, O]
}

// NewMockArticles creates a new mock instance.
func NewMockArticles[I dto.ArticlesInDto, T dto.Tag, AT dto.ArticleTag[T], O dto.ArticlesOutDto[T, AT]](ctrl *gomock.Controller) *MockArticles[I, T, AT, O] {
	mock := &MockArticles[I, T, AT, O]{ctrl: ctrl}
	mock.recorder = &MockArticlesMockRecorder[I, T, AT, O]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticles[I, T, AT, O]) EXPECT() *MockArticlesMockRecorder[I, T, AT, O] {
	return m.recorder
}

// Execute mocks base method.
func (m *MockArticles[I, T, AT, O]) Execute(ctx context.Context, in I) (O, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, in)
	ret0, _ := ret[0].(O)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockArticlesMockRecorder[I, T, AT, O]) Execute(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockArticles[I, T, AT, O])(nil).Execute), ctx, in)
}
