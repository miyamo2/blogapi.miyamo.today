// Code generated by MockGen. DO NOT EDIT.
// Source: converter.go
//
// Generated by this command:
//
//	mockgen -source=converter.go -destination=../../../../mock/if-adapter/controller/pb/presenter/converter.go -package=presenters
//

// Package presenters is a generated GoMock package.
package presenters

import (
	context "context"
	reflect "reflect"

	dto "github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/app/usecase/dto"
	grpc "github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/infra/grpc"
	gomock "go.uber.org/mock/gomock"
)

// MockToCreateArticleResponse is a mock of ToCreateArticleResponse interface.
type MockToCreateArticleResponse struct {
	ctrl     *gomock.Controller
	recorder *MockToCreateArticleResponseMockRecorder
	isgomock struct{}
}

// MockToCreateArticleResponseMockRecorder is the mock recorder for MockToCreateArticleResponse.
type MockToCreateArticleResponseMockRecorder struct {
	mock *MockToCreateArticleResponse
}

// NewMockToCreateArticleResponse creates a new mock instance.
func NewMockToCreateArticleResponse(ctrl *gomock.Controller) *MockToCreateArticleResponse {
	mock := &MockToCreateArticleResponse{ctrl: ctrl}
	mock.recorder = &MockToCreateArticleResponseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToCreateArticleResponse) EXPECT() *MockToCreateArticleResponseMockRecorder {
	return m.recorder
}

// ToCreateArticleArticleResponse mocks base method.
func (m *MockToCreateArticleResponse) ToCreateArticleArticleResponse(ctx context.Context, from *dto.CreateArticleOutDto) (*grpc.BloggingEventResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToCreateArticleArticleResponse", ctx, from)
	ret0, _ := ret[0].(*grpc.BloggingEventResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToCreateArticleArticleResponse indicates an expected call of ToCreateArticleArticleResponse.
func (mr *MockToCreateArticleResponseMockRecorder) ToCreateArticleArticleResponse(ctx, from any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToCreateArticleArticleResponse", reflect.TypeOf((*MockToCreateArticleResponse)(nil).ToCreateArticleArticleResponse), ctx, from)
}

// MockToUpdateArticleTitleResponse is a mock of ToUpdateArticleTitleResponse interface.
type MockToUpdateArticleTitleResponse struct {
	ctrl     *gomock.Controller
	recorder *MockToUpdateArticleTitleResponseMockRecorder
	isgomock struct{}
}

// MockToUpdateArticleTitleResponseMockRecorder is the mock recorder for MockToUpdateArticleTitleResponse.
type MockToUpdateArticleTitleResponseMockRecorder struct {
	mock *MockToUpdateArticleTitleResponse
}

// NewMockToUpdateArticleTitleResponse creates a new mock instance.
func NewMockToUpdateArticleTitleResponse(ctrl *gomock.Controller) *MockToUpdateArticleTitleResponse {
	mock := &MockToUpdateArticleTitleResponse{ctrl: ctrl}
	mock.recorder = &MockToUpdateArticleTitleResponseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToUpdateArticleTitleResponse) EXPECT() *MockToUpdateArticleTitleResponseMockRecorder {
	return m.recorder
}

// ToUpdateArticleTitleResponse mocks base method.
func (m *MockToUpdateArticleTitleResponse) ToUpdateArticleTitleResponse(ctx context.Context, from *dto.UpdateArticleTitleOutDto) (*grpc.BloggingEventResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToUpdateArticleTitleResponse", ctx, from)
	ret0, _ := ret[0].(*grpc.BloggingEventResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToUpdateArticleTitleResponse indicates an expected call of ToUpdateArticleTitleResponse.
func (mr *MockToUpdateArticleTitleResponseMockRecorder) ToUpdateArticleTitleResponse(ctx, from any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToUpdateArticleTitleResponse", reflect.TypeOf((*MockToUpdateArticleTitleResponse)(nil).ToUpdateArticleTitleResponse), ctx, from)
}

// MockToUpdateArticleBodyResponse is a mock of ToUpdateArticleBodyResponse interface.
type MockToUpdateArticleBodyResponse struct {
	ctrl     *gomock.Controller
	recorder *MockToUpdateArticleBodyResponseMockRecorder
	isgomock struct{}
}

// MockToUpdateArticleBodyResponseMockRecorder is the mock recorder for MockToUpdateArticleBodyResponse.
type MockToUpdateArticleBodyResponseMockRecorder struct {
	mock *MockToUpdateArticleBodyResponse
}

// NewMockToUpdateArticleBodyResponse creates a new mock instance.
func NewMockToUpdateArticleBodyResponse(ctrl *gomock.Controller) *MockToUpdateArticleBodyResponse {
	mock := &MockToUpdateArticleBodyResponse{ctrl: ctrl}
	mock.recorder = &MockToUpdateArticleBodyResponseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToUpdateArticleBodyResponse) EXPECT() *MockToUpdateArticleBodyResponseMockRecorder {
	return m.recorder
}

// ToUpdateArticleBodyResponse mocks base method.
func (m *MockToUpdateArticleBodyResponse) ToUpdateArticleBodyResponse(ctx context.Context, from *dto.UpdateArticleBodyOutDto) (*grpc.BloggingEventResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToUpdateArticleBodyResponse", ctx, from)
	ret0, _ := ret[0].(*grpc.BloggingEventResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToUpdateArticleBodyResponse indicates an expected call of ToUpdateArticleBodyResponse.
func (mr *MockToUpdateArticleBodyResponseMockRecorder) ToUpdateArticleBodyResponse(ctx, from any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToUpdateArticleBodyResponse", reflect.TypeOf((*MockToUpdateArticleBodyResponse)(nil).ToUpdateArticleBodyResponse), ctx, from)
}

// MockToUpdateArticleThumbnailResponse is a mock of ToUpdateArticleThumbnailResponse interface.
type MockToUpdateArticleThumbnailResponse struct {
	ctrl     *gomock.Controller
	recorder *MockToUpdateArticleThumbnailResponseMockRecorder
	isgomock struct{}
}

// MockToUpdateArticleThumbnailResponseMockRecorder is the mock recorder for MockToUpdateArticleThumbnailResponse.
type MockToUpdateArticleThumbnailResponseMockRecorder struct {
	mock *MockToUpdateArticleThumbnailResponse
}

// NewMockToUpdateArticleThumbnailResponse creates a new mock instance.
func NewMockToUpdateArticleThumbnailResponse(ctrl *gomock.Controller) *MockToUpdateArticleThumbnailResponse {
	mock := &MockToUpdateArticleThumbnailResponse{ctrl: ctrl}
	mock.recorder = &MockToUpdateArticleThumbnailResponseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToUpdateArticleThumbnailResponse) EXPECT() *MockToUpdateArticleThumbnailResponseMockRecorder {
	return m.recorder
}

// ToUpdateArticleThumbnailResponse mocks base method.
func (m *MockToUpdateArticleThumbnailResponse) ToUpdateArticleThumbnailResponse(ctx context.Context, from *dto.UpdateArticleThumbnailOutDto) (*grpc.BloggingEventResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToUpdateArticleThumbnailResponse", ctx, from)
	ret0, _ := ret[0].(*grpc.BloggingEventResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToUpdateArticleThumbnailResponse indicates an expected call of ToUpdateArticleThumbnailResponse.
func (mr *MockToUpdateArticleThumbnailResponseMockRecorder) ToUpdateArticleThumbnailResponse(ctx, from any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToUpdateArticleThumbnailResponse", reflect.TypeOf((*MockToUpdateArticleThumbnailResponse)(nil).ToUpdateArticleThumbnailResponse), ctx, from)
}

// MockToAttachTagsResponse is a mock of ToAttachTagsResponse interface.
type MockToAttachTagsResponse struct {
	ctrl     *gomock.Controller
	recorder *MockToAttachTagsResponseMockRecorder
	isgomock struct{}
}

// MockToAttachTagsResponseMockRecorder is the mock recorder for MockToAttachTagsResponse.
type MockToAttachTagsResponseMockRecorder struct {
	mock *MockToAttachTagsResponse
}

// NewMockToAttachTagsResponse creates a new mock instance.
func NewMockToAttachTagsResponse(ctrl *gomock.Controller) *MockToAttachTagsResponse {
	mock := &MockToAttachTagsResponse{ctrl: ctrl}
	mock.recorder = &MockToAttachTagsResponseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToAttachTagsResponse) EXPECT() *MockToAttachTagsResponseMockRecorder {
	return m.recorder
}

// ToAttachTagsResponse mocks base method.
func (m *MockToAttachTagsResponse) ToAttachTagsResponse(ctx context.Context, from *dto.AttachTagsOutDto) (*grpc.BloggingEventResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToAttachTagsResponse", ctx, from)
	ret0, _ := ret[0].(*grpc.BloggingEventResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToAttachTagsResponse indicates an expected call of ToAttachTagsResponse.
func (mr *MockToAttachTagsResponseMockRecorder) ToAttachTagsResponse(ctx, from any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToAttachTagsResponse", reflect.TypeOf((*MockToAttachTagsResponse)(nil).ToAttachTagsResponse), ctx, from)
}

// MockToDetachTagsResponse is a mock of ToDetachTagsResponse interface.
type MockToDetachTagsResponse struct {
	ctrl     *gomock.Controller
	recorder *MockToDetachTagsResponseMockRecorder
	isgomock struct{}
}

// MockToDetachTagsResponseMockRecorder is the mock recorder for MockToDetachTagsResponse.
type MockToDetachTagsResponseMockRecorder struct {
	mock *MockToDetachTagsResponse
}

// NewMockToDetachTagsResponse creates a new mock instance.
func NewMockToDetachTagsResponse(ctrl *gomock.Controller) *MockToDetachTagsResponse {
	mock := &MockToDetachTagsResponse{ctrl: ctrl}
	mock.recorder = &MockToDetachTagsResponseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockToDetachTagsResponse) EXPECT() *MockToDetachTagsResponseMockRecorder {
	return m.recorder
}

// ToDetachTagsResponse mocks base method.
func (m *MockToDetachTagsResponse) ToDetachTagsResponse(ctx context.Context, from *dto.DetachTagsOutDto) (*grpc.BloggingEventResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToDetachTagsResponse", ctx, from)
	ret0, _ := ret[0].(*grpc.BloggingEventResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ToDetachTagsResponse indicates an expected call of ToDetachTagsResponse.
func (mr *MockToDetachTagsResponseMockRecorder) ToDetachTagsResponse(ctx, from any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToDetachTagsResponse", reflect.TypeOf((*MockToDetachTagsResponse)(nil).ToDetachTagsResponse), ctx, from)
}
