package pb

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/google/go-cmp/cmp"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/infra/grpc"
	mpresenter "github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/mock/if-adapter/controller/pb/presenter"
	musecase "github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/mock/if-adapter/controller/pb/usecase"
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
			s := NewBloggingEventServiceServer(u, conv, nil, nil)
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
			s := NewBloggingEventServiceServer(nil, nil, u, conv)
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
