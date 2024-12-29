package pb

import (
	"blogapi.miyamo.today/blogging-event-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/blogging-event-service/internal/infra/grpc"
	mpresenter "blogapi.miyamo.today/blogging-event-service/internal/mock/if-adapter/controller/pb/presenter"
	musecase "blogapi.miyamo.today/blogging-event-service/internal/mock/if-adapter/controller/pb/usecase"
	"blogapi.miyamo.today/blogging-event-service/internal/pkg"
	"context"
	"github.com/cockroachdb/errors"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/testing/protocmp"
	"testing"
)

func TestBloggingEventServiceServer_CreateArticle(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *grpc.CreateArticleRequest
	}
	type want struct {
		response *grpc.BloggingEventResponse
		err      error
	}
	type testCase struct {
		outDto         dto.CreateArticleOutDto
		setupUsecase   func(out dto.CreateArticleOutDto, u *musecase.MockCreateArticle)
		setupConverter func(from dto.CreateArticleOutDto, res *grpc.BloggingEventResponse, conv *mpresenter.MockToCreateArticleResponse)
		args           args
		want           want
	}
	errInUsecase := errors.New("error in usecase")
	errInConverter := errors.New("error in converter")
	tests := map[string]testCase{
		"happy_path": {
			outDto: dto.NewCreateArticleOutDto("eventID", "articleID"),
			setupUsecase: func(out dto.CreateArticleOutDto, u *musecase.MockCreateArticle) {
				in := dto.NewCreateArticleInDto("title", "body", "https://example.com/example.jpg", []string{"tag1", "tag2"})
				u.EXPECT().
					Execute(gomock.Any(), &in).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.CreateArticleOutDto, res *grpc.BloggingEventResponse, conv *mpresenter.MockToCreateArticleResponse) {
				conv.EXPECT().
					ToCreateArticleArticleResponse(gomock.Any(), &from).
					Return(res, nil).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.CreateArticleRequest{
					Title:        "title",
					Body:         "body",
					ThumbnailUrl: "https://example.com/example.jpg",
					TagNames:     []string{"tag1", "tag2"},
				},
			},
			want: want{
				response: &grpc.BloggingEventResponse{EventId: "eventID", ArticleId: "articleID"},
			},
		},
		"unhappy_path/usecase-returns-error": {
			outDto: dto.NewCreateArticleOutDto("", ""),
			setupUsecase: func(out dto.CreateArticleOutDto, u *musecase.MockCreateArticle) {
				in := dto.NewCreateArticleInDto("title", "body", "https://example.com/example.jpg", []string{"tag1", "tag2"})
				u.EXPECT().
					Execute(gomock.Any(), &in).
					Return(&out, errInUsecase).
					Times(1)
			},
			setupConverter: func(from dto.CreateArticleOutDto, res *grpc.BloggingEventResponse, conv *mpresenter.MockToCreateArticleResponse) {
				conv.EXPECT().
					ToCreateArticleArticleResponse(gomock.Any(), gomock.Any()).
					Times(0)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.CreateArticleRequest{
					Title:        "title",
					Body:         "body",
					ThumbnailUrl: "https://example.com/example.jpg",
					TagNames:     []string{"tag1", "tag2"},
				},
			},
			want: want{
				err: errInUsecase,
			},
		},
		"unhappy_path/converter-returns-error": {
			outDto: dto.NewCreateArticleOutDto("eventID", "articleID"),
			setupUsecase: func(out dto.CreateArticleOutDto, u *musecase.MockCreateArticle) {
				in := dto.NewCreateArticleInDto("title", "body", "https://example.com/example.jpg", []string{"tag1", "tag2"})
				u.EXPECT().
					Execute(gomock.Any(), &in).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.CreateArticleOutDto, res *grpc.BloggingEventResponse, conv *mpresenter.MockToCreateArticleResponse) {
				conv.EXPECT().
					ToCreateArticleArticleResponse(gomock.Any(), &from).
					Return(nil, errInConverter).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.CreateArticleRequest{
					Title:        "title",
					Body:         "body",
					ThumbnailUrl: "https://example.com/example.jpg",
					TagNames:     []string{"tag1", "tag2"},
				},
			},
			want: want{
				err: errInConverter,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			out := tt.outDto
			u := musecase.NewMockCreateArticle(ctrl)
			tt.setupUsecase(out, u)
			response := tt.want.response
			conv := mpresenter.NewMockToCreateArticleResponse(ctrl)
			tt.setupConverter(out, response, conv)
			s := NewBloggingEventServiceServer(WithCreateArticleUsecase(u), WithCreateArticleConverter(conv))
			got, err := s.CreateArticle(tt.args.ctx, tt.args.in)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("CreateArticle() error = %v, wantErr %v", err, tt.want.err)
			}
			if diff := cmp.Diff(got, tt.want.response, protocmp.Transform()); diff != "" {
				t.Errorf("GetArticleById() got = %v, want %v", got, tt.want.response)
			}
		})
	}
}

func TestBloggingEventServiceServer_UpdateArticleTitle(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *grpc.UpdateArticleTitleRequest
	}
	type want struct {
		response *grpc.BloggingEventResponse
		err      error
	}
	type testCase struct {
		outDto         dto.UpdateArticleTitleOutDto
		setupUsecase   func(out dto.UpdateArticleTitleOutDto, u *musecase.MockUpdateArticleTitle)
		setupConverter func(from dto.UpdateArticleTitleOutDto, res *grpc.BloggingEventResponse, conv *mpresenter.MockToUpdateArticleTitleResponse)
		args           args
		want           want
	}

	errInUsecase := errors.New("error in usecase")
	errInConverter := errors.New("error in converter")

	tests := map[string]testCase{
		"happy_path": {
			outDto: dto.NewUpdateArticleTitleOutDto("eventID", "articleID"),
			setupUsecase: func(out dto.UpdateArticleTitleOutDto, u *musecase.MockUpdateArticleTitle) {
				in := dto.NewUpdateArticleTitleInDto("articleID", "title")
				u.EXPECT().
					Execute(gomock.Any(), &in).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.UpdateArticleTitleOutDto, res *grpc.BloggingEventResponse, conv *mpresenter.MockToUpdateArticleTitleResponse) {
				conv.EXPECT().ToUpdateArticleTitleResponse(gomock.Any(), &from).
					Return(res, nil).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.UpdateArticleTitleRequest{
					Id:    "articleID",
					Title: "title",
				},
			},
			want: want{
				response: &grpc.BloggingEventResponse{EventId: "eventID", ArticleId: "articleID"},
			},
		},
		"unhappy_path/usecase-returns-error": {
			outDto: dto.NewUpdateArticleTitleOutDto("", ""),
			setupUsecase: func(out dto.UpdateArticleTitleOutDto, u *musecase.MockUpdateArticleTitle) {
				in := dto.NewUpdateArticleTitleInDto("articleID", "title")
				u.EXPECT().
					Execute(gomock.Any(), &in).
					Return(&out, errInUsecase).
					Times(1)
			},
			setupConverter: func(from dto.UpdateArticleTitleOutDto, res *grpc.BloggingEventResponse, conv *mpresenter.MockToUpdateArticleTitleResponse) {
				conv.EXPECT().
					ToUpdateArticleTitleResponse(gomock.Any(), gomock.Any()).
					Times(0)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.UpdateArticleTitleRequest{
					Id:    "articleID",
					Title: "title",
				},
			},
			want: want{
				err: errInUsecase,
			},
		},
		"unhappy_path/converter-returns-error": {
			outDto: dto.NewUpdateArticleTitleOutDto("eventID", "articleID"),
			setupUsecase: func(out dto.UpdateArticleTitleOutDto, u *musecase.MockUpdateArticleTitle) {
				in := dto.NewUpdateArticleTitleInDto("articleID", "title")
				u.EXPECT().
					Execute(gomock.Any(), &in).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.UpdateArticleTitleOutDto, res *grpc.BloggingEventResponse, conv *mpresenter.MockToUpdateArticleTitleResponse) {
				conv.EXPECT().
					ToUpdateArticleTitleResponse(gomock.Any(), &from).
					Return(nil, errInConverter).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.UpdateArticleTitleRequest{
					Id:    "articleID",
					Title: "title",
				},
			},
			want: want{
				err: errInConverter,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			out := tt.outDto
			u := musecase.NewMockUpdateArticleTitle(ctrl)
			tt.setupUsecase(out, u)
			response := tt.want.response
			conv := mpresenter.NewMockToUpdateArticleTitleResponse(ctrl)
			tt.setupConverter(out, response, conv)
			s := NewBloggingEventServiceServer(WithUpdateArticleTitleUsecase(u), WithUpdateArticleTitleConverter(conv))
			got, err := s.UpdateArticleTitle(tt.args.ctx, tt.args.in)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("CreateArticle() error = %v, wantErr %v", err, tt.want.err)
			}
			if diff := cmp.Diff(got, tt.want.response, protocmp.Transform()); diff != "" {
				t.Errorf("GetArticleById() got = %v, want %v", got, tt.want.response)
			}
		})
	}
}

func TestBloggingEventServiceServer_UpdateArticleBody(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *grpc.UpdateArticleBodyRequest
	}
	type want struct {
		response *grpc.BloggingEventResponse
		err      error
	}
	type testCase struct {
		outDto         dto.UpdateArticleBodyOutDto
		setupUsecase   func(out dto.UpdateArticleBodyOutDto, u *musecase.MockUpdateArticleBody)
		setupConverter func(from dto.UpdateArticleBodyOutDto, res *grpc.BloggingEventResponse, conv *mpresenter.MockToUpdateArticleBodyResponse)
		args           args
		want           want
	}

	errInUsecase := errors.New("error in usecase")
	errInConverter := errors.New("error in converter")

	tests := map[string]testCase{
		"happy_path": {
			outDto: dto.NewUpdateArticleBodyOutDto("eventID", "articleID"),
			setupUsecase: func(out dto.UpdateArticleBodyOutDto, u *musecase.MockUpdateArticleBody) {
				in := dto.NewUpdateArticleBodyInDto("articleID", "body")
				u.EXPECT().
					Execute(gomock.Any(), &in).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.UpdateArticleBodyOutDto, res *grpc.BloggingEventResponse, conv *mpresenter.MockToUpdateArticleBodyResponse) {
				conv.EXPECT().ToUpdateArticleBodyResponse(gomock.Any(), &from).
					Return(res, nil).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.UpdateArticleBodyRequest{
					Id:   "articleID",
					Body: "body",
				},
			},
			want: want{
				response: &grpc.BloggingEventResponse{EventId: "eventID", ArticleId: "articleID"},
			},
		},
		"unhappy_path/usecase-returns-error": {
			outDto: dto.NewUpdateArticleBodyOutDto("", ""),
			setupUsecase: func(out dto.UpdateArticleBodyOutDto, u *musecase.MockUpdateArticleBody) {
				in := dto.NewUpdateArticleBodyInDto("articleID", "body")
				u.EXPECT().
					Execute(gomock.Any(), &in).
					Return(&out, errInUsecase).
					Times(1)
			},
			setupConverter: func(from dto.UpdateArticleBodyOutDto, res *grpc.BloggingEventResponse, conv *mpresenter.MockToUpdateArticleBodyResponse) {
				conv.EXPECT().
					ToUpdateArticleBodyResponse(gomock.Any(), gomock.Any()).
					Times(0)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.UpdateArticleBodyRequest{
					Id:   "articleID",
					Body: "body",
				},
			},
			want: want{
				err: errInUsecase,
			},
		},
		"unhappy_path/converter-returns-error": {
			outDto: dto.NewUpdateArticleBodyOutDto("eventID", "articleID"),
			setupUsecase: func(out dto.UpdateArticleBodyOutDto, u *musecase.MockUpdateArticleBody) {
				in := dto.NewUpdateArticleBodyInDto("articleID", "body")
				u.EXPECT().
					Execute(gomock.Any(), &in).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.UpdateArticleBodyOutDto, res *grpc.BloggingEventResponse, conv *mpresenter.MockToUpdateArticleBodyResponse) {
				conv.EXPECT().
					ToUpdateArticleBodyResponse(gomock.Any(), &from).
					Return(nil, errInConverter).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.UpdateArticleBodyRequest{
					Id:   "articleID",
					Body: "body",
				},
			},
			want: want{
				err: errInConverter,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			out := tt.outDto
			u := musecase.NewMockUpdateArticleBody(ctrl)
			tt.setupUsecase(out, u)
			response := tt.want.response
			conv := mpresenter.NewMockToUpdateArticleBodyResponse(ctrl)
			tt.setupConverter(out, response, conv)
			s := NewBloggingEventServiceServer(WithUpdateArticleBodyUsecase(u), WithUpdateArticleBodyConverter(conv))
			got, err := s.UpdateArticleBody(tt.args.ctx, tt.args.in)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("CreateArticle() error = %v, wantErr %v", err, tt.want.err)
			}
			if diff := cmp.Diff(got, tt.want.response, protocmp.Transform()); diff != "" {
				t.Errorf("GetArticleById() got = %v, want %v", got, tt.want.response)
			}
		})
	}
}

func TestBloggingEventServiceServer_UpdateArticleThumbnail(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *grpc.UpdateArticleThumbnailRequest
	}
	type want struct {
		response *grpc.BloggingEventResponse
		err      error
	}
	type testCase struct {
		outDto         dto.UpdateArticleThumbnailOutDto
		setupUsecase   func(out dto.UpdateArticleThumbnailOutDto, u *musecase.MockUpdateArticleThumbnail)
		setupConverter func(from dto.UpdateArticleThumbnailOutDto, res *grpc.BloggingEventResponse, conv *mpresenter.MockToUpdateArticleThumbnailResponse)
		args           args
		want           want
	}

	errInUsecase := errors.New("error in usecase")
	errInConverter := errors.New("error in converter")

	tests := map[string]testCase{
		"happy_path": {
			outDto: dto.NewUpdateArticleThumbnailOutDto("eventID", "articleID"),
			setupUsecase: func(out dto.UpdateArticleThumbnailOutDto, u *musecase.MockUpdateArticleThumbnail) {
				in := dto.NewUpdateArticleThumbnailInDto("articleID", *pkg.MustParseURL("https://example.com/example.jpg"))
				u.EXPECT().
					Execute(gomock.Any(), &in).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.UpdateArticleThumbnailOutDto, res *grpc.BloggingEventResponse, conv *mpresenter.MockToUpdateArticleThumbnailResponse) {
				conv.EXPECT().ToUpdateArticleThumbnailResponse(gomock.Any(), &from).
					Return(res, nil).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.UpdateArticleThumbnailRequest{
					Id:           "articleID",
					ThumbnailUrl: "https://example.com/example.jpg",
				},
			},
			want: want{
				response: &grpc.BloggingEventResponse{EventId: "eventID", ArticleId: "articleID"},
			},
		},
		"unhappy_path/usecase-returns-error": {
			outDto: dto.NewUpdateArticleThumbnailOutDto("", ""),
			setupUsecase: func(out dto.UpdateArticleThumbnailOutDto, u *musecase.MockUpdateArticleThumbnail) {
				in := dto.NewUpdateArticleThumbnailInDto("articleID", *pkg.MustParseURL("https://example.com/example.jpg"))
				u.EXPECT().
					Execute(gomock.Any(), &in).
					Return(&out, errInUsecase).
					Times(1)
			},
			setupConverter: func(from dto.UpdateArticleThumbnailOutDto, res *grpc.BloggingEventResponse, conv *mpresenter.MockToUpdateArticleThumbnailResponse) {
				conv.EXPECT().
					ToUpdateArticleThumbnailResponse(gomock.Any(), gomock.Any()).
					Times(0)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.UpdateArticleThumbnailRequest{
					Id:           "articleID",
					ThumbnailUrl: "https://example.com/example.jpg",
				},
			},
			want: want{
				err: errInUsecase,
			},
		},
		"unhappy_path/converter-returns-error": {
			outDto: dto.NewUpdateArticleThumbnailOutDto("eventID", "articleID"),
			setupUsecase: func(out dto.UpdateArticleThumbnailOutDto, u *musecase.MockUpdateArticleThumbnail) {
				in := dto.NewUpdateArticleThumbnailInDto("articleID", *pkg.MustParseURL("https://example.com/example.jpg"))
				u.EXPECT().
					Execute(gomock.Any(), &in).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.UpdateArticleThumbnailOutDto, res *grpc.BloggingEventResponse, conv *mpresenter.MockToUpdateArticleThumbnailResponse) {
				conv.EXPECT().
					ToUpdateArticleThumbnailResponse(gomock.Any(), &from).
					Return(nil, errInConverter).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.UpdateArticleThumbnailRequest{
					Id:           "articleID",
					ThumbnailUrl: "https://example.com/example.jpg",
				},
			},
			want: want{
				err: errInConverter,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			out := tt.outDto
			u := musecase.NewMockUpdateArticleThumbnail(ctrl)
			tt.setupUsecase(out, u)
			response := tt.want.response
			conv := mpresenter.NewMockToUpdateArticleThumbnailResponse(ctrl)
			tt.setupConverter(out, response, conv)
			s := NewBloggingEventServiceServer(WithUpdateArticleThumbnailUsecase(u), WithUpdateArticleThumbnailConverter(conv))
			got, err := s.UpdateArticleThumbnail(tt.args.ctx, tt.args.in)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("CreateArticle() error = %v, wantErr %v", err, tt.want.err)
			}
			if diff := cmp.Diff(got, tt.want.response, protocmp.Transform()); diff != "" {
				t.Errorf("GetArticleById() got = %v, want %v", got, tt.want.response)
			}
		})
	}
}

func TestBloggingEventServiceServer_AttachTags(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *grpc.AttachTagsRequest
	}
	type want struct {
		response *grpc.BloggingEventResponse
		err      error
	}
	type testCase struct {
		outDto         dto.AttachTagsOutDto
		setupUsecase   func(out dto.AttachTagsOutDto, u *musecase.MockAttachTags)
		setupConverter func(from dto.AttachTagsOutDto, res *grpc.BloggingEventResponse, conv *mpresenter.MockToAttachTagsResponse)
		args           args
		want           want
	}

	errInUsecase := errors.New("error in usecase")
	errInConverter := errors.New("error in converter")

	tests := map[string]testCase{
		"happy_path": {
			outDto: dto.NewAttachTagsOutDto("eventID", "articleID"),
			setupUsecase: func(out dto.AttachTagsOutDto, u *musecase.MockAttachTags) {
				in := dto.NewAttachTagsInDto("articleID", []string{"tag1", "tag2"})
				u.EXPECT().
					Execute(gomock.Any(), &in).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.AttachTagsOutDto, res *grpc.BloggingEventResponse, conv *mpresenter.MockToAttachTagsResponse) {
				conv.EXPECT().ToAttachTagsResponse(gomock.Any(), &from).
					Return(res, nil).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.AttachTagsRequest{
					Id:       "articleID",
					TagNames: []string{"tag1", "tag2"},
				},
			},
			want: want{
				response: &grpc.BloggingEventResponse{EventId: "eventID", ArticleId: "articleID"},
			},
		},
		"unhappy_path/usecase-returns-error": {
			outDto: dto.NewAttachTagsOutDto("", ""),
			setupUsecase: func(out dto.AttachTagsOutDto, u *musecase.MockAttachTags) {
				in := dto.NewAttachTagsInDto("articleID", []string{"tag1", "tag2"})
				u.EXPECT().
					Execute(gomock.Any(), &in).
					Return(&out, errInUsecase).
					Times(1)
			},
			setupConverter: func(from dto.AttachTagsOutDto, res *grpc.BloggingEventResponse, conv *mpresenter.MockToAttachTagsResponse) {
				conv.EXPECT().
					ToAttachTagsResponse(gomock.Any(), gomock.Any()).
					Times(0)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.AttachTagsRequest{
					Id:       "articleID",
					TagNames: []string{"tag1", "tag2"},
				},
			},
			want: want{
				err: errInUsecase,
			},
		},
		"unhappy_path/converter-returns-error": {
			outDto: dto.NewAttachTagsOutDto("eventID", "articleID"),
			setupUsecase: func(out dto.AttachTagsOutDto, u *musecase.MockAttachTags) {
				in := dto.NewAttachTagsInDto("articleID", []string{"tag1", "tag2"})
				u.EXPECT().
					Execute(gomock.Any(), &in).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.AttachTagsOutDto, res *grpc.BloggingEventResponse, conv *mpresenter.MockToAttachTagsResponse) {
				conv.EXPECT().
					ToAttachTagsResponse(gomock.Any(), &from).
					Return(nil, errInConverter).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.AttachTagsRequest{
					Id:       "articleID",
					TagNames: []string{"tag1", "tag2"},
				},
			},
			want: want{
				err: errInConverter,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			out := tt.outDto
			u := musecase.NewMockAttachTags(ctrl)
			tt.setupUsecase(out, u)
			response := tt.want.response
			conv := mpresenter.NewMockToAttachTagsResponse(ctrl)
			tt.setupConverter(out, response, conv)
			s := NewBloggingEventServiceServer(WithAttachTagsUsecase(u), WithAttachTagsConverter(conv))
			got, err := s.AttachTags(tt.args.ctx, tt.args.in)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("CreateArticle() error = %v, wantErr %v", err, tt.want.err)
			}
			if diff := cmp.Diff(got, tt.want.response, protocmp.Transform()); diff != "" {
				t.Errorf("GetArticleById() got = %v, want %v", got, tt.want.response)
			}
		})
	}
}

func TestBloggingEventServiceServer_DetachTags(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *grpc.DetachTagsRequest
	}
	type want struct {
		response *grpc.BloggingEventResponse
		err      error
	}
	type testCase struct {
		outDto         dto.DetachTagsOutDto
		setupUsecase   func(out dto.DetachTagsOutDto, u *musecase.MockDetachTags)
		setupConverter func(from dto.DetachTagsOutDto, res *grpc.BloggingEventResponse, conv *mpresenter.MockToDetachTagsResponse)
		args           args
		want           want
	}

	errInUsecase := errors.New("error in usecase")
	errInConverter := errors.New("error in converter")

	tests := map[string]testCase{
		"happy_path": {
			outDto: dto.NewDetachTagsOutDto("eventID", "articleID"),
			setupUsecase: func(out dto.DetachTagsOutDto, u *musecase.MockDetachTags) {
				in := dto.NewDetachTagsInDto("articleID", []string{"tag1", "tag2"})
				u.EXPECT().
					Execute(gomock.Any(), &in).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.DetachTagsOutDto, res *grpc.BloggingEventResponse, conv *mpresenter.MockToDetachTagsResponse) {
				conv.EXPECT().ToDetachTagsResponse(gomock.Any(), &from).
					Return(res, nil).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.DetachTagsRequest{
					Id:       "articleID",
					TagNames: []string{"tag1", "tag2"},
				},
			},
			want: want{
				response: &grpc.BloggingEventResponse{EventId: "eventID", ArticleId: "articleID"},
			},
		},
		"unhappy_path/usecase-returns-error": {
			outDto: dto.NewDetachTagsOutDto("", ""),
			setupUsecase: func(out dto.DetachTagsOutDto, u *musecase.MockDetachTags) {
				in := dto.NewDetachTagsInDto("articleID", []string{"tag1", "tag2"})
				u.EXPECT().
					Execute(gomock.Any(), &in).
					Return(&out, errInUsecase).
					Times(1)
			},
			setupConverter: func(from dto.DetachTagsOutDto, res *grpc.BloggingEventResponse, conv *mpresenter.MockToDetachTagsResponse) {
				conv.EXPECT().
					ToDetachTagsResponse(gomock.Any(), gomock.Any()).
					Times(0)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.DetachTagsRequest{
					Id:       "articleID",
					TagNames: []string{"tag1", "tag2"},
				},
			},
			want: want{
				err: errInUsecase,
			},
		},
		"unhappy_path/converter-returns-error": {
			outDto: dto.NewDetachTagsOutDto("eventID", "articleID"),
			setupUsecase: func(out dto.DetachTagsOutDto, u *musecase.MockDetachTags) {
				in := dto.NewDetachTagsInDto("articleID", []string{"tag1", "tag2"})
				u.EXPECT().
					Execute(gomock.Any(), &in).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.DetachTagsOutDto, res *grpc.BloggingEventResponse, conv *mpresenter.MockToDetachTagsResponse) {
				conv.EXPECT().
					ToDetachTagsResponse(gomock.Any(), &from).
					Return(nil, errInConverter).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.DetachTagsRequest{
					Id:       "articleID",
					TagNames: []string{"tag1", "tag2"},
				},
			},
			want: want{
				err: errInConverter,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			out := tt.outDto
			u := musecase.NewMockDetachTags(ctrl)
			tt.setupUsecase(out, u)
			response := tt.want.response
			conv := mpresenter.NewMockToDetachTagsResponse(ctrl)
			tt.setupConverter(out, response, conv)
			s := NewBloggingEventServiceServer(WithDetachTagsUsecase(u), WithDetachTagsConverter(conv))
			got, err := s.DetachTags(tt.args.ctx, tt.args.in)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("CreateArticle() error = %v, wantErr %v", err, tt.want.err)
			}
			if diff := cmp.Diff(got, tt.want.response, protocmp.Transform()); diff != "" {
				t.Errorf("GetArticleById() got = %v, want %v", got, tt.want.response)
			}
		})
	}
}
