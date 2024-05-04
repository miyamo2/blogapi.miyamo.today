package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/cockroachdb/errors"
	blogapictx "github.com/miyamo2/blogapi.miyamo.today/core/context"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"
	mpb "github.com/miyamo2/blogapi.miyamo.today/federator/internal/mock/protogen/article/client/pb"
	"github.com/miyamo2/blogapi.miyamo.today/protogen/article/client/pb"
	"go.uber.org/mock/gomock"
)

func TestArticles_Execute(t *testing.T) {
	type args struct {
		ctx context.Context
		in  dto.ArticlesInDto
	}
	type want struct {
		out dto.ArticlesOutDto
		err error
	}
	type testCase struct {
		articleServiceClient func(ctrl *gomock.Controller) pb.ArticleServiceClient
		args                 args
		want                 want
		wantErr              bool
	}
	errTestArticles := errors.New("test error")
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
		"happy_path/next_paging": {
			articleServiceClient: func(ctrl *gomock.Controller) pb.ArticleServiceClient {
				aSvcClt := mpb.NewMockArticleServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetNextArticles(gomock.Any(), gomock.Any()).
					Return(&pb.GetNextArticlesResponse{
						Articles: []*pb.Article{
							{
								Id:           "Article1",
								Title:        "happy_path/next_paging",
								Body:         "## happy_path/next_paging",
								ThumbnailUrl: "example.test",
								CreatedAt:    "2020-01-01T00:00:00Z",
								UpdatedAt:    "2020-01-01T00:00:00Z",
								Tags: []*pb.Tag{
									{
										Id:   "Tag1",
										Name: "Tag1",
									},
								},
							},
						},
						StillExists: true,
					}, nil).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.ArticlesInDto {
					in, _ := dto.NewArticlesInDto(dto.ArticlesInWithFirst(1), dto.ArticlesInWithAfter("Article0"))
					return in
				}(),
			},
			want: want{
				out: dto.NewArticlesOutDto(
					[]dto.ArticleTag{
						dto.NewArticleTag(
							"Article1",
							"happy_path/next_paging",
							"## happy_path/next_paging",
							"example.test",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z",
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
							}),
					},
					dto.ArticlesOutDtoWithHasNext(true),
				),
			},
		},
		"unhappy_path/next_paging_returns_error": {
			articleServiceClient: func(ctrl *gomock.Controller) pb.ArticleServiceClient {
				aSvcClt := mpb.NewMockArticleServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetNextArticles(gomock.Any(), gomock.Any()).
					Return(&pb.GetNextArticlesResponse{}, errTestArticles).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.ArticlesInDto {
					in, _ := dto.NewArticlesInDto(dto.ArticlesInWithFirst(1), dto.ArticlesInWithAfter("Article0"))
					return in
				}(),
			},
			want: want{
				out: dto.ArticlesOutDto{},
				err: errTestArticles,
			},
			wantErr: true,
		},
		"happy_path/prev_paging": {
			articleServiceClient: func(ctrl *gomock.Controller) pb.ArticleServiceClient {
				aSvcClt := mpb.NewMockArticleServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetPrevArticles(gomock.Any(), gomock.Any()).
					Return(&pb.GetPrevArticlesResponse{
						Articles: []*pb.Article{
							{
								Id:           "Article1",
								Title:        "happy_path/prev_paging",
								Body:         "## happy_path/prev_paging",
								ThumbnailUrl: "example.test",
								CreatedAt:    "2020-01-01T00:00:00Z",
								UpdatedAt:    "2020-01-01T00:00:00Z",
								Tags: []*pb.Tag{
									{
										Id:   "Tag1",
										Name: "Tag1",
									},
								},
							},
						},
						StillExists: true,
					}, nil).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.ArticlesInDto {
					in, _ := dto.NewArticlesInDto(dto.ArticlesInWithLast(1), dto.ArticlesInWithBefore("Article2"))
					return in
				}(),
			},
			want: want{
				out: dto.NewArticlesOutDto(
					[]dto.ArticleTag{
						dto.NewArticleTag(
							"Article1",
							"happy_path/prev_paging",
							"## happy_path/prev_paging",
							"example.test",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z",
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
							}),
					},
					dto.ArticlesOutDtoWithHasPrev(true),
				),
			},
		},
		"unhappy_path/prev_paging_returns_error": {
			articleServiceClient: func(ctrl *gomock.Controller) pb.ArticleServiceClient {
				aSvcClt := mpb.NewMockArticleServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetPrevArticles(gomock.Any(), gomock.Any()).
					Return(&pb.GetPrevArticlesResponse{}, errTestArticles).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.ArticlesInDto {
					in, _ := dto.NewArticlesInDto(dto.ArticlesInWithLast(1), dto.ArticlesInWithBefore("Article2"))
					return in
				}(),
			},
			want: want{
				out: dto.ArticlesOutDto{},
				err: errTestArticles,
			},
			wantErr: true,
		},
		"happy_path/execute": {
			articleServiceClient: func(ctrl *gomock.Controller) pb.ArticleServiceClient {
				aSvcClt := mpb.NewMockArticleServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetAllArticles(gomock.Any(), gomock.Any()).
					Return(&pb.GetAllArticlesResponse{
						Articles: []*pb.Article{
							{
								Id:           "Article1",
								Title:        "happy_path/execute",
								Body:         "## happy_path/execute",
								ThumbnailUrl: "example.test",
								CreatedAt:    "2020-01-01T00:00:00Z",
								UpdatedAt:    "2020-01-01T00:00:00Z",
								Tags: []*pb.Tag{
									{
										Id:   "Tag1",
										Name: "Tag1",
									},
								},
							},
						},
					}, nil).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.ArticlesInDto{},
			},
			want: want{
				out: dto.NewArticlesOutDto(
					[]dto.ArticleTag{
						dto.NewArticleTag(
							"Article1",
							"happy_path/execute",
							"## happy_path/execute",
							"example.test",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z",
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
							}),
					},
				),
			},
		},
		"unhappy_path/execute_returns_error": {
			articleServiceClient: func(ctrl *gomock.Controller) pb.ArticleServiceClient {
				aSvcClt := mpb.NewMockArticleServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetAllArticles(gomock.Any(), gomock.Any()).
					Return(&pb.GetAllArticlesResponse{}, errTestArticles).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.ArticlesInDto{},
			},
			want: want{
				out: dto.ArticlesOutDto{},
				err: errTestArticles,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			aSvcClt := tt.articleServiceClient(ctrl)
			u := NewArticles(aSvcClt)
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
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArticles_executeNextPaging(t *testing.T) {
	type args struct {
		ctx context.Context
		in  dto.ArticlesInDto
	}
	type want struct {
		out dto.ArticlesOutDto
		err error
	}
	type testCase struct {
		articleServiceClient func(ctrl *gomock.Controller) pb.ArticleServiceClient
		args                 args
		want                 want
		wantErr              bool
	}
	errTestArticles := errors.New("test error")
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
		"happy_path/next_paging": {
			articleServiceClient: func(ctrl *gomock.Controller) pb.ArticleServiceClient {
				aSvcClt := mpb.NewMockArticleServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetNextArticles(gomock.Any(), gomock.Any()).
					Return(&pb.GetNextArticlesResponse{
						Articles: []*pb.Article{
							{
								Id:           "Article1",
								Title:        "happy_path/next_paging",
								Body:         "## happy_path/next_paging",
								ThumbnailUrl: "example.test",
								CreatedAt:    "2020-01-01T00:00:00Z",
								UpdatedAt:    "2020-01-01T00:00:00Z",
								Tags: []*pb.Tag{
									{
										Id:   "Tag1",
										Name: "Tag1",
									},
								},
							},
						},
						StillExists: true,
					}, nil).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.ArticlesInDto {
					in, _ := dto.NewArticlesInDto(dto.ArticlesInWithFirst(1), dto.ArticlesInWithAfter("Article0"))
					return in
				}(),
			},
			want: want{
				out: dto.NewArticlesOutDto(
					[]dto.ArticleTag{
						dto.NewArticleTag(
							"Article1",
							"happy_path/next_paging",
							"## happy_path/next_paging",
							"example.test",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z",
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
							}),
					},
					dto.ArticlesOutDtoWithHasNext(true),
				),
			},
		},
		"unhappy_path/next_paging_returns_error": {
			articleServiceClient: func(ctrl *gomock.Controller) pb.ArticleServiceClient {
				aSvcClt := mpb.NewMockArticleServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetNextArticles(gomock.Any(), gomock.Any()).
					Return(&pb.GetNextArticlesResponse{}, errTestArticles).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.ArticlesInDto {
					in, _ := dto.NewArticlesInDto(dto.ArticlesInWithFirst(1), dto.ArticlesInWithAfter("Article0"))
					return in
				}(),
			},
			want: want{
				out: dto.ArticlesOutDto{},
				err: errTestArticles,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			aSvcClt := tt.articleServiceClient(ctrl)
			u := NewArticles(aSvcClt)
			got, err := u.executeNextPaging(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				if err == nil {
					t.Errorf("executeNextPaging() expected error but got nil")
					return
				}
				if !errors.Is(err, tt.want.err) {
					t.Errorf("executeNextPaging() error = %v, want %v", err, tt.wantErr)
					return
				}
			} else if err != nil {
				t.Errorf("executeNextPaging() expected nil but got error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want.out) {
				t.Errorf("executeNextPaging() got = %v, want %v", got, tt.want.out)
			}
		})
	}
}

func TestArticles_executePrevPaging(t *testing.T) {
	type args struct {
		ctx context.Context
		in  dto.ArticlesInDto
	}
	type want struct {
		out dto.ArticlesOutDto
		err error
	}
	type testCase struct {
		articleServiceClient func(ctrl *gomock.Controller) pb.ArticleServiceClient
		args                 args
		want                 want
		wantErr              bool
	}
	errTestArticles := errors.New("test error")
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
		"happy_path/prev_paging": {
			articleServiceClient: func(ctrl *gomock.Controller) pb.ArticleServiceClient {
				aSvcClt := mpb.NewMockArticleServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetPrevArticles(gomock.Any(), gomock.Any()).
					Return(&pb.GetPrevArticlesResponse{
						Articles: []*pb.Article{
							{
								Id:           "Article1",
								Title:        "happy_path/prev_paging",
								Body:         "## happy_path/prev_paging",
								ThumbnailUrl: "example.test",
								CreatedAt:    "2020-01-01T00:00:00Z",
								UpdatedAt:    "2020-01-01T00:00:00Z",
								Tags: []*pb.Tag{
									{
										Id:   "Tag1",
										Name: "Tag1",
									},
								},
							},
						},
						StillExists: true,
					}, nil).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.ArticlesInDto {
					in, _ := dto.NewArticlesInDto(dto.ArticlesInWithLast(2), dto.ArticlesInWithBefore("Article2"))
					return in
				}(),
			},
			want: want{
				out: dto.NewArticlesOutDto(
					[]dto.ArticleTag{
						dto.NewArticleTag(
							"Article1",
							"happy_path/prev_paging",
							"## happy_path/prev_paging",
							"example.test",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z",
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
							}),
					},
					dto.ArticlesOutDtoWithHasPrev(true),
				),
			},
		},
		"unhappy_path/prev_paging_returns_error": {
			articleServiceClient: func(ctrl *gomock.Controller) pb.ArticleServiceClient {
				aSvcClt := mpb.NewMockArticleServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetPrevArticles(gomock.Any(), gomock.Any()).
					Return(&pb.GetPrevArticlesResponse{}, errTestArticles).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.ArticlesInDto {
					in, _ := dto.NewArticlesInDto(dto.ArticlesInWithLast(2), dto.ArticlesInWithBefore("Article2"))
					return in
				}(),
			},
			want: want{
				out: dto.ArticlesOutDto{},
				err: errTestArticles,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			aSvcClt := tt.articleServiceClient(ctrl)
			u := NewArticles(aSvcClt)
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
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArticles_execute(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	type want struct {
		out dto.ArticlesOutDto
		err error
	}
	type testCase struct {
		articleServiceClient func(ctrl *gomock.Controller) pb.ArticleServiceClient
		args                 args
		want                 want
		wantErr              bool
	}
	errTestArticles := errors.New("test error")
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
		"happy_path/all_articles": {
			articleServiceClient: func(ctrl *gomock.Controller) pb.ArticleServiceClient {
				aSvcClt := mpb.NewMockArticleServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetAllArticles(gomock.Any(), gomock.Any()).
					Return(&pb.GetAllArticlesResponse{
						Articles: []*pb.Article{
							{
								Id:           "Article1",
								Title:        "happy_path/all_articles",
								Body:         "## happy_path/all_articles",
								ThumbnailUrl: "example.test",
								CreatedAt:    "2020-01-01T00:00:00Z",
								UpdatedAt:    "2020-01-01T00:00:00Z",
								Tags: []*pb.Tag{
									{
										Id:   "Tag1",
										Name: "Tag1",
									},
								},
							},
						},
					}, nil).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
			},
			want: want{
				out: dto.NewArticlesOutDto(
					[]dto.ArticleTag{
						dto.NewArticleTag(
							"Article1",
							"happy_path/all_articles",
							"## happy_path/all_articles",
							"example.test",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z",
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
							}),
					},
				),
			},
		},
		"unhappy_path/all_articles_returns_error": {
			articleServiceClient: func(ctrl *gomock.Controller) pb.ArticleServiceClient {
				aSvcClt := mpb.NewMockArticleServiceClient(ctrl)
				aSvcClt.EXPECT().
					GetAllArticles(gomock.Any(), gomock.Any()).
					Return(&pb.GetAllArticlesResponse{}, errTestArticles).
					Times(1)
				return aSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
			},
			want: want{
				out: dto.ArticlesOutDto{},
				err: errTestArticles,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			aSvcClt := tt.articleServiceClient(ctrl)
			u := NewArticles(aSvcClt)
			got, err := u.execute(tt.args.ctx)
			if tt.wantErr {
				if err == nil {
					t.Errorf("execute() expected error but got nil")
					return
				}
				if !errors.Is(err, tt.want.err) {
					t.Errorf("execute() error = %v, want %v", err, tt.wantErr)
					return
				}
			} else if err != nil {
				t.Errorf("execute() expected nil but got error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want.out) {
				t.Errorf("execute() got = %v, want %v", got, tt.want.out)
			}
		})
	}
}
