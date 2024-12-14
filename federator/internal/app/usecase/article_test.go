package usecase

import (
	"context"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	grpc "github.com/miyamo2/blogapi.miyamo.today/federator/internal/infra/grpc/article"
	mgrpc "github.com/miyamo2/blogapi.miyamo.today/federator/internal/mock/infra/grpc/article"
	"reflect"
	"testing"

	"github.com/cockroachdb/errors"
	blogapictx "github.com/miyamo2/blogapi.miyamo.today/core/context"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"
	"go.uber.org/mock/gomock"
)

func TestArticle_Execute(t *testing.T) {
	type args struct {
		ctx context.Context
		in  dto.ArticleInDto
	}
	type want struct {
		out dto.ArticleOutDto
		err error
	}
	type testCase struct {
		articleServiceClient func(ctrl *gomock.Controller) grpc.ArticleServiceClient
		args                 args
		want                 want
		wantErr              bool
	}
	errTestArticle := errors.New("test error")
	mockBlogAPIContext := func() context.Context {
		return blogapictx.StoreToContext(
			context.Background(),
			blogapictx.New(
				"1234567890",
				"0987654321",
				blogapictx.RequestTypeGRPC,
				nil,
				nil))
	}
	tests := map[string]testCase{
		"happy_path/single_tag": {
			articleServiceClient: func(ctrl *gomock.Controller) grpc.ArticleServiceClient {
				aSvcClt := mgrpc.NewMockArticleServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetArticleById(gomock.Any(), gomock.Any()).
					Return(&grpc.GetArticleByIdResponse{
						Article: &grpc.Article{
							Id:           "Article1",
							Title:        "happy_path/single_tag",
							Body:         "## happy_path/single_tag",
							ThumbnailUrl: "example.test",
							CreatedAt:    "2020-01-01T00:00:00.000000Z",
							UpdatedAt:    "2020-01-01T00:00:00.000000Z",
							Tags: []*grpc.Tag{
								{
									Id:   "Tag1",
									Name: "Tag1",
								},
							},
						},
					}, nil).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.NewArticleInDto("Article1"),
			},
			want: want{
				out: dto.NewArticleOutDto(
					dto.NewArticleTag(
						"Article1",
						"happy_path/single_tag",
						"## happy_path/single_tag",
						"example.test",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						[]dto.Tag{
							dto.NewTag("Tag1", "Tag1"),
						}),
				),
			},
		},
		"happy_path/multiple_tags": {
			articleServiceClient: func(ctrl *gomock.Controller) grpc.ArticleServiceClient {
				aSvcClt := mgrpc.NewMockArticleServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetArticleById(gomock.Any(), gomock.Any()).
					Return(&grpc.GetArticleByIdResponse{
						Article: &grpc.Article{
							Id:           "Article1",
							Title:        "happy_path/multiple_tags",
							Body:         "## happy_path/multiple_tags",
							ThumbnailUrl: "example.test",
							CreatedAt:    "2020-01-01T00:00:00.000000Z",
							UpdatedAt:    "2020-01-01T00:00:00.000000Z",
							Tags: []*grpc.Tag{
								{
									Id:   "Tag1",
									Name: "Tag1",
								},
								{
									Id:   "Tag2",
									Name: "Tag2",
								},
								{
									Id:   "Tag3",
									Name: "Tag3",
								},
							},
						},
					}, nil).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.NewArticleInDto("Article1"),
			},
			want: want{
				out: dto.NewArticleOutDto(
					dto.NewArticleTag(
						"Article1",
						"happy_path/multiple_tags",
						"## happy_path/multiple_tags",
						"example.test",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						[]dto.Tag{
							dto.NewTag("Tag1", "Tag1"),
							dto.NewTag("Tag2", "Tag2"),
							dto.NewTag("Tag3", "Tag3"),
						},
					),
				),
			},
		},
		"happy_path/no_tags": {
			articleServiceClient: func(ctrl *gomock.Controller) grpc.ArticleServiceClient {
				aSvcClt := mgrpc.NewMockArticleServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetArticleById(gomock.Any(), gomock.Any()).
					Return(&grpc.GetArticleByIdResponse{
						Article: &grpc.Article{
							Id:           "Article1",
							Title:        "happy_path/no_tags",
							Body:         "## happy_path/no_tags",
							ThumbnailUrl: "example.test",
							CreatedAt:    "2020-01-01T00:00:00.000000Z",
							UpdatedAt:    "2020-01-01T00:00:00.000000Z",
						},
					}, nil).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.NewArticleInDto("Article1"),
			},
			want: want{
				out: dto.NewArticleOutDto(
					dto.NewArticleTag(
						"Article1",
						"happy_path/no_tags",
						"## happy_path/no_tags",
						"example.test",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						[]dto.Tag{})),
			},
		},
		"unhappy_path/grpc_returns_error": {
			articleServiceClient: func(ctrl *gomock.Controller) grpc.ArticleServiceClient {
				aSvcClt := mgrpc.NewMockArticleServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetArticleById(gomock.Any(), gomock.Any()).
					Return(&grpc.GetArticleByIdResponse{}, errTestArticle).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.NewArticleInDto("Article1"),
			},
			want: want{
				out: dto.ArticleOutDto{},
				err: errTestArticle,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			aSvcClt := tt.articleServiceClient(ctrl)
			u := NewArticle(aSvcClt)
			got, err := u.Execute(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Execute() expected error but got nil")
					return
				}
				if !errors.Is(err, tt.want.err) {
					t.Errorf("Execute() error = %v, want %v", err, tt.wantErr)
					return
				}
			} else if err != nil {
				t.Errorf("Execute() expected nil but got error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want.out) {
				t.Errorf("Execute() got = %v, want %v", got, tt.want.out)
			}
		})
	}
}
