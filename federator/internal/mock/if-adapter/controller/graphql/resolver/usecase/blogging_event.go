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
func (m *MockUpdateArticleBody) Execute(ctx context.Context, in dto.UpdateArticleBodyInDTO) (dto.UpdateArticleBodyOutDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, in)
	ret0, _ := ret[0].(dto.UpdateArticleBodyOutDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockUpdateArticleBodyMockRecorder) Execute(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockUpdateArticleBody)(nil).Execute), ctx, in)
}

// MockUpdateArticleThumbnail is a mock of UpdateArticleThumbnail interface.
type MockUpdateArticleThumbnail struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateArticleThumbnailMockRecorder
	isgomock struct{}
}

// MockUpdateArticleThumbnailMockRecorder is the mock recorder for MockUpdateArticleThumbnail.
type MockUpdateArticleThumbnailMockRecorder struct {
	mock *MockUpdateArticleThumbnail
}

// NewMockUpdateArticleThumbnail creates a new mock instance.
func NewMockUpdateArticleThumbnail(ctrl *gomock.Controller) *MockUpdateArticleThumbnail {
	mock := &MockUpdateArticleThumbnail{ctrl: ctrl}
	mock.recorder = &MockUpdateArticleThumbnailMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpdateArticleThumbnail) EXPECT() *MockUpdateArticleThumbnailMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockUpdateArticleThumbnail) Execute(ctx context.Context, in dto.UpdateArticleThumbnailInDTO) (dto.UpdateArticleThumbnailOutDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, in)
	ret0, _ := ret[0].(dto.UpdateArticleThumbnailOutDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockUpdateArticleThumbnailMockRecorder) Execute(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockUpdateArticleThumbnail)(nil).Execute), ctx, in)
}

// MockAttachTags is a mock of AttachTags interface.
type MockAttachTags struct {
	ctrl     *gomock.Controller
	recorder *MockAttachTagsMockRecorder
	isgomock struct{}
}

// MockAttachTagsMockRecorder is the mock recorder for MockAttachTags.
type MockAttachTagsMockRecorder struct {
	mock *MockAttachTags
}

// NewMockAttachTags creates a new mock instance.
func NewMockAttachTags(ctrl *gomock.Controller) *MockAttachTags {
	mock := &MockAttachTags{ctrl: ctrl}
	mock.recorder = &MockAttachTagsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAttachTags) EXPECT() *MockAttachTagsMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockAttachTags) Execute(ctx context.Context, in dto.AttachTagsInDTO) (dto.AttachTagsOutDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, in)
	ret0, _ := ret[0].(dto.AttachTagsOutDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockAttachTagsMockRecorder) Execute(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockAttachTags)(nil).Execute), ctx, in)
}

// MockDetachTags is a mock of DetachTags interface.
type MockDetachTags struct {
	ctrl     *gomock.Controller
	recorder *MockDetachTagsMockRecorder
	isgomock struct{}
}

// MockDetachTagsMockRecorder is the mock recorder for MockDetachTags.
type MockDetachTagsMockRecorder struct {
	mock *MockDetachTags
}

// NewMockDetachTags creates a new mock instance.
func NewMockDetachTags(ctrl *gomock.Controller) *MockDetachTags {
	mock := &MockDetachTags{ctrl: ctrl}
	mock.recorder = &MockDetachTagsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDetachTags) EXPECT() *MockDetachTagsMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockDetachTags) Execute(ctx context.Context, in dto.DetachTagsInDTO) (dto.DetachTagsOutDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, in)
	ret0, _ := ret[0].(dto.DetachTagsOutDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockDetachTagsMockRecorder) Execute(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockDetachTags)(nil).Execute), ctx, in)
}

// MockUploadImage is a mock of UploadImage interface.
type MockUploadImage struct {
	ctrl     *gomock.Controller
	recorder *MockUploadImageMockRecorder
	isgomock struct{}
}

// MockUploadImageMockRecorder is the mock recorder for MockUploadImage.
type MockUploadImageMockRecorder struct {
	mock *MockUploadImage
}

// NewMockUploadImage creates a new mock instance.
func NewMockUploadImage(ctrl *gomock.Controller) *MockUploadImage {
	mock := &MockUploadImage{ctrl: ctrl}
	mock.recorder = &MockUploadImageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUploadImage) EXPECT() *MockUploadImageMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockUploadImage) Execute(ctx context.Context, in dto.UploadImageInDTO) (dto.UploadImageOutDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, in)
	ret0, _ := ret[0].(dto.UploadImageOutDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockUploadImageMockRecorder) Execute(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockUploadImage)(nil).Execute), ctx, in)
}
