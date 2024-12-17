// Code generated by MockGen. DO NOT EDIT.
// Source: blogging_event.go
//
// Generated by this command:
//
//	mockgen -source=blogging_event.go -destination=../../../../../mock/if-adapter/controller/graphql/resolver/usecase/blogging_event.go -package=usecase
//

// Package usecase is a generated GoMock package.
package usecase

import (
	context "context"
	reflect "reflect"

	dto "github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"
	gomock "go.uber.org/mock/gomock"
)

// MockCreateArticle is a mock of CreateArticle interface.
type MockCreateArticle struct {
	ctrl     *gomock.Controller
	recorder *MockCreateArticleMockRecorder
	isgomock struct{}
}

// MockCreateArticleMockRecorder is the mock recorder for MockCreateArticle.
type MockCreateArticleMockRecorder struct {
	mock *MockCreateArticle
}

// NewMockCreateArticle creates a new mock instance.
func NewMockCreateArticle(ctrl *gomock.Controller) *MockCreateArticle {
	mock := &MockCreateArticle{ctrl: ctrl}
	mock.recorder = &MockCreateArticleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreateArticle) EXPECT() *MockCreateArticleMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockCreateArticle) Execute(ctx context.Context, in dto.CreateArticleInDTO) (dto.CreateArticleOutDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, in)
	ret0, _ := ret[0].(dto.CreateArticleOutDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockCreateArticleMockRecorder) Execute(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockCreateArticle)(nil).Execute), ctx, in)
}

// MockUpdateArticleTitle is a mock of UpdateArticleTitle interface.
type MockUpdateArticleTitle struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateArticleTitleMockRecorder
	isgomock struct{}
}

// MockUpdateArticleTitleMockRecorder is the mock recorder for MockUpdateArticleTitle.
type MockUpdateArticleTitleMockRecorder struct {
	mock *MockUpdateArticleTitle
}

// NewMockUpdateArticleTitle creates a new mock instance.
func NewMockUpdateArticleTitle(ctrl *gomock.Controller) *MockUpdateArticleTitle {
	mock := &MockUpdateArticleTitle{ctrl: ctrl}
	mock.recorder = &MockUpdateArticleTitleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpdateArticleTitle) EXPECT() *MockUpdateArticleTitleMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockUpdateArticleTitle) Execute(ctx context.Context, in dto.UpdateArticleTitleInDTO) (dto.UpdateArticleTitleOutDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, in)
	ret0, _ := ret[0].(dto.UpdateArticleTitleOutDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockUpdateArticleTitleMockRecorder) Execute(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockUpdateArticleTitle)(nil).Execute), ctx, in)
}
