// Code generated by MockGen. DO NOT EDIT.
// Source: article_service.go
//
// Generated by this command:
//
//	mockgen -source=article_service.go -destination=../../../mock/app/usecase/query/mock_article_service.go -package=mock_query
//

// Package mock_query is a generated GoMock package.
package mock_query

import (
	context "context"
	reflect "reflect"

	query "github.com/miyamo2/blogapi.miyamo.today/article-service/internal/infra/rdb/query"
	db "github.com/miyamo2/blogapi.miyamo.today/core/db"
	gomock "go.uber.org/mock/gomock"
)

// MockArticleService is a mock of ArticleService interface.
type MockArticleService struct {
	ctrl     *gomock.Controller
	recorder *MockArticleServiceMockRecorder
	isgomock struct{}
}

// MockArticleServiceMockRecorder is the mock recorder for MockArticleService.
type MockArticleServiceMockRecorder struct {
	mock *MockArticleService
}

// NewMockArticleService creates a new mock instance.
func NewMockArticleService(ctrl *gomock.Controller) *MockArticleService {
	mock := &MockArticleService{ctrl: ctrl}
	mock.recorder = &MockArticleServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleService) EXPECT() *MockArticleServiceMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockArticleService) GetAll(ctx context.Context, out *db.MultipleStatementResult[query.Article], paginationOption ...db.PaginationOption) db.Statement {
	m.ctrl.T.Helper()
	varargs := []any{ctx, out}
	for _, a := range paginationOption {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAll", varargs...)
	ret0, _ := ret[0].(db.Statement)
	return ret0
}

// GetAll indicates an expected call of GetAll.
func (mr *MockArticleServiceMockRecorder) GetAll(ctx, out any, paginationOption ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, out}, paginationOption...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockArticleService)(nil).GetAll), varargs...)
}

// GetById mocks base method.
func (m *MockArticleService) GetById(ctx context.Context, id string, out *db.SingleStatementResult[query.Article]) db.Statement {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id, out)
	ret0, _ := ret[0].(db.Statement)
	return ret0
}

// GetById indicates an expected call of GetById.
func (mr *MockArticleServiceMockRecorder) GetById(ctx, id, out any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockArticleService)(nil).GetById), ctx, id, out)
}
