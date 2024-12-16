// Code generated by MockGen. DO NOT EDIT.
// Source: converter.go
//
// Generated by this command:
//
//	mockgen -source=converter.go -destination=../../../../../../mock/if-adapter/controller/graphql/resolver/presenter/converter/converter.go -package=converters
//

// Package converters is a generated GoMock package.
package converters

import (
	context "context"
	reflect "reflect"

	dto "github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"
	model "github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/presenters/graphql/model"
	gomock "go.uber.org/mock/gomock"
)

// MockArticleConverter is a mock of ArticleConverter interface.
type MockArticleConverter struct {
	ctrl     *gomock.Controller
	recorder *MockArticleConverterMockRecorder
	isgomock struct{}
}

// MockArticleConverterMockRecorder is the mock recorder for MockArticleConverter.
type MockArticleConverterMockRecorder struct {
	mock *MockArticleConverter
}

// NewMockArticleConverter creates a new mock instance.
func NewMockArticleConverter(ctrl *gomock.Controller) *MockArticleConverter {
	mock := &MockArticleConverter{ctrl: ctrl}
	mock.recorder = &MockArticleConverterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleConverter) EXPECT() *MockArticleConverterMockRecorder {
	return m.recorder
}

// ToArticle mocks base method.
func (m *MockArticleConverter) ToArticle(ctx context.Context, from dto.ArticleOutDTO) (*model.ArticleNode, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToArticle", ctx, from)
	ret0, _ := ret[0].(*model.ArticleNode)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ToArticle indicates an expected call of ToArticle.
func (mr *MockArticleConverterMockRecorder) ToArticle(ctx, from any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToArticle", reflect.TypeOf((*MockArticleConverter)(nil).ToArticle), ctx, from)
}

// MockArticlesConverter is a mock of ArticlesConverter interface.
type MockArticlesConverter struct {
	ctrl     *gomock.Controller
	recorder *MockArticlesConverterMockRecorder
	isgomock struct{}
}

// MockArticlesConverterMockRecorder is the mock recorder for MockArticlesConverter.
type MockArticlesConverterMockRecorder struct {
	mock *MockArticlesConverter
}

// NewMockArticlesConverter creates a new mock instance.
func NewMockArticlesConverter(ctrl *gomock.Controller) *MockArticlesConverter {
	mock := &MockArticlesConverter{ctrl: ctrl}
	mock.recorder = &MockArticlesConverterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticlesConverter) EXPECT() *MockArticlesConverterMockRecorder {
	return m.recorder
}

// ToArticles mocks base method.
func (m *MockArticlesConverter) ToArticles(ctx context.Context, from dto.ArticlesOutDTO) (*model.ArticleConnection, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToArticles", ctx, from)
	ret0, _ := ret[0].(*model.ArticleConnection)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ToArticles indicates an expected call of ToArticles.
func (mr *MockArticlesConverterMockRecorder) ToArticles(ctx, from any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToArticles", reflect.TypeOf((*MockArticlesConverter)(nil).ToArticles), ctx, from)
}

// MockTagConverter is a mock of TagConverter interface.
type MockTagConverter struct {
	ctrl     *gomock.Controller
	recorder *MockTagConverterMockRecorder
	isgomock struct{}
}

// MockTagConverterMockRecorder is the mock recorder for MockTagConverter.
type MockTagConverterMockRecorder struct {
	mock *MockTagConverter
}

// NewMockTagConverter creates a new mock instance.
func NewMockTagConverter(ctrl *gomock.Controller) *MockTagConverter {
	mock := &MockTagConverter{ctrl: ctrl}
	mock.recorder = &MockTagConverterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTagConverter) EXPECT() *MockTagConverterMockRecorder {
	return m.recorder
}

// ToTag mocks base method.
func (m *MockTagConverter) ToTag(ctx context.Context, from dto.TagOutDTO) (*model.TagNode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToTag", ctx, from)
	ret0, _ := ret[0].(*model.TagNode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToTag indicates an expected call of ToTag.
func (mr *MockTagConverterMockRecorder) ToTag(ctx, from any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToTag", reflect.TypeOf((*MockTagConverter)(nil).ToTag), ctx, from)
}

// MockTagsConverter is a mock of TagsConverter interface.
type MockTagsConverter struct {
	ctrl     *gomock.Controller
	recorder *MockTagsConverterMockRecorder
	isgomock struct{}
}

// MockTagsConverterMockRecorder is the mock recorder for MockTagsConverter.
type MockTagsConverterMockRecorder struct {
	mock *MockTagsConverter
}

// NewMockTagsConverter creates a new mock instance.
func NewMockTagsConverter(ctrl *gomock.Controller) *MockTagsConverter {
	mock := &MockTagsConverter{ctrl: ctrl}
	mock.recorder = &MockTagsConverterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTagsConverter) EXPECT() *MockTagsConverterMockRecorder {
	return m.recorder
}

// ToTags mocks base method.
func (m *MockTagsConverter) ToTags(ctx context.Context, from dto.TagsOutDTO) (*model.TagConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToTags", ctx, from)
	ret0, _ := ret[0].(*model.TagConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToTags indicates an expected call of ToTags.
func (mr *MockTagsConverterMockRecorder) ToTags(ctx, from any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToTags", reflect.TypeOf((*MockTagsConverter)(nil).ToTags), ctx, from)
}

// MockCreateArticleConverter is a mock of CreateArticleConverter interface.
type MockCreateArticleConverter struct {
	ctrl     *gomock.Controller
	recorder *MockCreateArticleConverterMockRecorder
	isgomock struct{}
}

// MockCreateArticleConverterMockRecorder is the mock recorder for MockCreateArticleConverter.
type MockCreateArticleConverterMockRecorder struct {
	mock *MockCreateArticleConverter
}

// NewMockCreateArticleConverter creates a new mock instance.
func NewMockCreateArticleConverter(ctrl *gomock.Controller) *MockCreateArticleConverter {
	mock := &MockCreateArticleConverter{ctrl: ctrl}
	mock.recorder = &MockCreateArticleConverterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreateArticleConverter) EXPECT() *MockCreateArticleConverterMockRecorder {
	return m.recorder
}

// ToCreateArticle mocks base method.
func (m *MockCreateArticleConverter) ToCreateArticle(ctx context.Context, from dto.CreateArticleOutDTO) (*model.CreateArticlePayload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToCreateArticle", ctx, from)
	ret0, _ := ret[0].(*model.CreateArticlePayload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToCreateArticle indicates an expected call of ToCreateArticle.
func (mr *MockCreateArticleConverterMockRecorder) ToCreateArticle(ctx, from any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToCreateArticle", reflect.TypeOf((*MockCreateArticleConverter)(nil).ToCreateArticle), ctx, from)
}