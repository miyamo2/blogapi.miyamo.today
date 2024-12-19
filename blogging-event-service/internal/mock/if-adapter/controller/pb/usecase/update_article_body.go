// Code generated by MockGen. DO NOT EDIT.
// Source: update_article_body.go
//
// Generated by this command:
//
//	mockgen -source=update_article_body.go -destination=../../../../mock/if-adapter/controller/pb/usecase/update_article_body.go -package=usecase
//

// Package usecase is a generated GoMock package.
package usecase

import (
	context "context"
	reflect "reflect"

	dto "github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/app/usecase/dto"
	gomock "go.uber.org/mock/gomock"
)

// MockUpdateArticleBody is a mock of UpdateArticleBody interface.
type MockUpdateArticleBody struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateArticleBodyMockRecorder
	isgomock struct{}
}

// MockUpdateArticleBodyMockRecorder is the mock recorder for MockUpdateArticleBody.
type MockUpdateArticleBodyMockRecorder struct {
	mock *MockUpdateArticleBody
}

// NewMockUpdateArticleBody creates a new mock instance.
func NewMockUpdateArticleBody(ctrl *gomock.Controller) *MockUpdateArticleBody {
	mock := &MockUpdateArticleBody{ctrl: ctrl}
	mock.recorder = &MockUpdateArticleBodyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpdateArticleBody) EXPECT() *MockUpdateArticleBodyMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockUpdateArticleBody) Execute(ctx context.Context, in *dto.UpdateArticleBodyInDto) (*dto.UpdateArticleBodyOutDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, in)
	ret0, _ := ret[0].(*dto.UpdateArticleBodyOutDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockUpdateArticleBodyMockRecorder) Execute(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockUpdateArticleBody)(nil).Execute), ctx, in)
}