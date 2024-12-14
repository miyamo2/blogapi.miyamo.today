// Code generated by MockGen. DO NOT EDIT.
// Source: blogging_event.go
//
// Generated by this command:
//
//	mockgen -source=blogging_event.go -destination=../../../mock/app/usecase/command/blogging_event.go -package=command
//

// Package command is a generated GoMock package.
package command

import (
	context "context"
	reflect "reflect"

	model "github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/domain/model"
	db "github.com/miyamo2/blogapi.miyamo.today/core/db"
	gomock "go.uber.org/mock/gomock"
)

// MockBloggingEventService is a mock of BloggingEventService interface.
type MockBloggingEventService struct {
	ctrl     *gomock.Controller
	recorder *MockBloggingEventServiceMockRecorder
	isgomock struct{}
}

// MockBloggingEventServiceMockRecorder is the mock recorder for MockBloggingEventService.
type MockBloggingEventServiceMockRecorder struct {
	mock *MockBloggingEventService
}

// NewMockBloggingEventService creates a new mock instance.
func NewMockBloggingEventService(ctrl *gomock.Controller) *MockBloggingEventService {
	mock := &MockBloggingEventService{ctrl: ctrl}
	mock.recorder = &MockBloggingEventServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBloggingEventService) EXPECT() *MockBloggingEventServiceMockRecorder {
	return m.recorder
}

// CreateArticle mocks base method.
func (m *MockBloggingEventService) CreateArticle(ctx context.Context, in model.CreateArticleEvent, out *db.SingleStatementResult[*model.BloggingEventKey]) db.Statement {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateArticle", ctx, in, out)
	ret0, _ := ret[0].(db.Statement)
	return ret0
}

// CreateArticle indicates an expected call of CreateArticle.
func (mr *MockBloggingEventServiceMockRecorder) CreateArticle(ctx, in, out any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateArticle", reflect.TypeOf((*MockBloggingEventService)(nil).CreateArticle), ctx, in, out)
}
