package usecase

import (
	grpc "blogapi.miyamo.today/federator/internal/infra/grpc/article"
	"blogapi.miyamo.today/federator/internal/infra/grpc/article/articleconnect"
	"blogapi.miyamo.today/federator/internal/utils"
	"connectrpc.com/connect"
	"context"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"google.golang.org/protobuf/types/known/timestamppb"
	"reflect"
	"testing"

	blogapictx "blogapi.miyamo.today/core/context"
	"blogapi.miyamo.today/federator/internal/app/usecase/dto"
	marticleconnect "blogapi.miyamo.today/federator/internal/mock/infra/grpc/article/articleconnect"
	"github.com/cockroachdb/errors"
	"go.uber.org/mock/gomock"
)

func TestArticles_Execute(t *testing.T) {
	type args struct {
		ctx context.Context
		in  dto.ArticlesInDTO
	}
	type want struct {
		out dto.ArticlesOutDTO
		err error
	}
	type testCase struct {
		articleServiceClient func(ctrl *gomock.Controller) articleconnect.ArticleServiceClient
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
			articleServiceClient: func(ctrl *gomock.Controller) articleconnect.ArticleServiceClient {
				articleServiceClient := marticleconnect.NewMockArticleServiceClient(ctrl)
				articleServiceClient.EXPECT().
					GetNextArticles(gomock.Any(), gomock.Any()).
					Return(connect.NewResponse(&grpc.GetNextArticlesResponse{
						Articles: []*grpc.Article{
							{
								Id:           "Article1",
								Title:        "happy_path/next_paging",
								Body:         "## happy_path/next_paging",
								ThumbnailUrl: "example.com/example.png",
								CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
								UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
								Tags: []*grpc.Tag{
									{
										Id:   "Tag1",
										Name: "Tag1",
									},
								},
							},
						},
						StillExists: true,
					}), nil).
					Times(1)
				return articleServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.ArticlesInDTO {
					in, _ := dto.NewArticlesInDTO(dto.ArticlesInWithFirst(1), dto.ArticlesInWithAfter("Article0"))
					return in
				}(),
			},
			want: want{
				out: dto.NewArticlesOutDTO(
					[]dto.ArticleTag{
						dto.NewArticleTag(
							"Article1",
							"happy_path/next_paging",
							"## happy_path/next_paging",
							utils.MustURLParse("example.com/example.png"),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
							}),
					},
					dto.ArticlesOutDTOWithHasNext(true),
				),
			},
		},
		"unhappy_path/next_paging_returns_error": {
			articleServiceClient: func(ctrl *gomock.Controller) articleconnect.ArticleServiceClient {
				articleServiceClient := marticleconnect.NewMockArticleServiceClient(ctrl)
				articleServiceClient.EXPECT().
					GetNextArticles(gomock.Any(), gomock.Any()).
					Return(connect.NewResponse(&grpc.GetNextArticlesResponse{}), errTestArticles).
					Times(1)
				return articleServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.ArticlesInDTO {
					in, _ := dto.NewArticlesInDTO(dto.ArticlesInWithFirst(1), dto.ArticlesInWithAfter("Article0"))
					return in
				}(),
			},
			want: want{
				out: dto.ArticlesOutDTO{},
				err: errTestArticles,
			},
			wantErr: true,
		},
		"happy_path/prev_paging": {
			articleServiceClient: func(ctrl *gomock.Controller) articleconnect.ArticleServiceClient {
				articleServiceClient := marticleconnect.NewMockArticleServiceClient(ctrl)
				articleServiceClient.EXPECT().
					GetPrevArticles(gomock.Any(), gomock.Any()).
					Return(connect.NewResponse(&grpc.GetPrevArticlesResponse{
						Articles: []*grpc.Article{
							{
								Id:           "Article1",
								Title:        "happy_path/prev_paging",
								Body:         "## happy_path/prev_paging",
								ThumbnailUrl: "example.com/example.png",
								CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
								UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
								Tags: []*grpc.Tag{
									{
										Id:   "Tag1",
										Name: "Tag1",
									},
								},
							},
						},
						StillExists: true,
					}), nil).
					Times(1)
				return articleServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.ArticlesInDTO {
					in, _ := dto.NewArticlesInDTO(dto.ArticlesInWithLast(1), dto.ArticlesInWithBefore("Article2"))
					return in
				}(),
			},
			want: want{
				out: dto.NewArticlesOutDTO(
					[]dto.ArticleTag{
						dto.NewArticleTag(
							"Article1",
							"happy_path/prev_paging",
							"## happy_path/prev_paging",
							utils.MustURLParse("example.com/example.png"),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
							}),
					},
					dto.ArticlesOutDTOWithHasPrev(true),
				),
			},
		},
		"unhappy_path/prev_paging_returns_error": {
			articleServiceClient: func(ctrl *gomock.Controller) articleconnect.ArticleServiceClient {
				articleServiceClient := marticleconnect.NewMockArticleServiceClient(ctrl)
				articleServiceClient.EXPECT().
					GetPrevArticles(gomock.Any(), gomock.Any()).
					Return(connect.NewResponse(&grpc.GetPrevArticlesResponse{}), errTestArticles).
					Times(1)
				return articleServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.ArticlesInDTO {
					in, _ := dto.NewArticlesInDTO(dto.ArticlesInWithLast(1), dto.ArticlesInWithBefore("Article2"))
					return in
				}(),
			},
			want: want{
				out: dto.ArticlesOutDTO{},
				err: errTestArticles,
			},
			wantErr: true,
		},
		"happy_path/execute": {
			articleServiceClient: func(ctrl *gomock.Controller) articleconnect.ArticleServiceClient {
				articleServiceClient := marticleconnect.NewMockArticleServiceClient(ctrl)
				articleServiceClient.EXPECT().
					GetAllArticles(gomock.Any(), gomock.Any()).
					Return(connect.NewResponse(&grpc.GetAllArticlesResponse{
						Articles: []*grpc.Article{
							{
								Id:           "Article1",
								Title:        "happy_path/execute",
								Body:         "## happy_path/execute",
								ThumbnailUrl: "example.com/example.png",
								CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
								UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
								Tags: []*grpc.Tag{
									{
										Id:   "Tag1",
										Name: "Tag1",
									},
								},
							},
						},
					}), nil).
					Times(1)
				return articleServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.ArticlesInDTO{},
			},
			want: want{
				out: dto.NewArticlesOutDTO(
					[]dto.ArticleTag{
						dto.NewArticleTag(
							"Article1",
							"happy_path/execute",
							"## happy_path/execute",
							utils.MustURLParse("example.com/example.png"),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
							}),
					},
				),
			},
		},
		"unhappy_path/execute_returns_error": {
			articleServiceClient: func(ctrl *gomock.Controller) articleconnect.ArticleServiceClient {
				articleServiceClient := marticleconnect.NewMockArticleServiceClient(ctrl)
				articleServiceClient.EXPECT().
					GetAllArticles(gomock.Any(), gomock.Any()).
					Return(connect.NewResponse(&grpc.GetAllArticlesResponse{}), errTestArticles).
					Times(1)
				return articleServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.ArticlesInDTO{},
			},
			want: want{
				out: dto.ArticlesOutDTO{},
				err: errTestArticles,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			articleServiceClient := tt.articleServiceClient(ctrl)
			u := NewArticles(articleServiceClient)
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
		in  dto.ArticlesInDTO
	}
	type want struct {
		out dto.ArticlesOutDTO
		err error
	}
	type testCase struct {
		articleServiceClient func(ctrl *gomock.Controller) articleconnect.ArticleServiceClient
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
			articleServiceClient: func(ctrl *gomock.Controller) articleconnect.ArticleServiceClient {
				articleServiceClient := marticleconnect.NewMockArticleServiceClient(ctrl)
				articleServiceClient.EXPECT().
					GetNextArticles(gomock.Any(), gomock.Any()).
					Return(connect.NewResponse(&grpc.GetNextArticlesResponse{
						Articles: []*grpc.Article{
							{
								Id:           "Article1",
								Title:        "happy_path/next_paging",
								Body:         "## happy_path/next_paging",
								ThumbnailUrl: "example.com/example.png",
								CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
								UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
								Tags: []*grpc.Tag{
									{
										Id:   "Tag1",
										Name: "Tag1",
									},
								},
							},
						},
						StillExists: true,
					}), nil).
					Times(1)
				return articleServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.ArticlesInDTO {
					in, _ := dto.NewArticlesInDTO(dto.ArticlesInWithFirst(1), dto.ArticlesInWithAfter("Article0"))
					return in
				}(),
			},
			want: want{
				out: dto.NewArticlesOutDTO(
					[]dto.ArticleTag{
						dto.NewArticleTag(
							"Article1",
							"happy_path/next_paging",
							"## happy_path/next_paging",
							utils.MustURLParse("example.com/example.png"),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
							}),
					},
					dto.ArticlesOutDTOWithHasNext(true),
				),
			},
		},
		"unhappy_path/next_paging_returns_error": {
			articleServiceClient: func(ctrl *gomock.Controller) articleconnect.ArticleServiceClient {
				articleServiceClient := marticleconnect.NewMockArticleServiceClient(ctrl)
				articleServiceClient.EXPECT().
					GetNextArticles(gomock.Any(), gomock.Any()).
					Return(connect.NewResponse(&grpc.GetNextArticlesResponse{}), errTestArticles).
					Times(1)
				return articleServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.ArticlesInDTO {
					in, _ := dto.NewArticlesInDTO(dto.ArticlesInWithFirst(1), dto.ArticlesInWithAfter("Article0"))
					return in
				}(),
			},
			want: want{
				out: dto.ArticlesOutDTO{},
				err: errTestArticles,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			articleServiceClient := tt.articleServiceClient(ctrl)
			u := NewArticles(articleServiceClient)
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
		in  dto.ArticlesInDTO
	}
	type want struct {
		out dto.ArticlesOutDTO
		err error
	}
	type testCase struct {
		articleServiceClient func(ctrl *gomock.Controller) articleconnect.ArticleServiceClient
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
			articleServiceClient: func(ctrl *gomock.Controller) articleconnect.ArticleServiceClient {
				articleServiceClient := marticleconnect.NewMockArticleServiceClient(ctrl)
				articleServiceClient.EXPECT().
					GetPrevArticles(gomock.Any(), gomock.Any()).
					Return(connect.NewResponse(&grpc.GetPrevArticlesResponse{
						Articles: []*grpc.Article{
							{
								Id:           "Article1",
								Title:        "happy_path/prev_paging",
								Body:         "## happy_path/prev_paging",
								ThumbnailUrl: "example.com/example.png",
								CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
								UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
								Tags: []*grpc.Tag{
									{
										Id:   "Tag1",
										Name: "Tag1",
									},
								},
							},
						},
						StillExists: true,
					}), nil).
					Times(1)
				return articleServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.ArticlesInDTO {
					in, _ := dto.NewArticlesInDTO(dto.ArticlesInWithLast(2), dto.ArticlesInWithBefore("Article2"))
					return in
				}(),
			},
			want: want{
				out: dto.NewArticlesOutDTO(
					[]dto.ArticleTag{
						dto.NewArticleTag(
							"Article1",
							"happy_path/prev_paging",
							"## happy_path/prev_paging",
							utils.MustURLParse("example.com/example.png"),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
							}),
					},
					dto.ArticlesOutDTOWithHasPrev(true),
				),
			},
		},
		"unhappy_path/prev_paging_returns_error": {
			articleServiceClient: func(ctrl *gomock.Controller) articleconnect.ArticleServiceClient {
				articleServiceClient := marticleconnect.NewMockArticleServiceClient(ctrl)
				articleServiceClient.EXPECT().
					GetPrevArticles(gomock.Any(), gomock.Any()).
					Return(connect.NewResponse(&grpc.GetPrevArticlesResponse{}), errTestArticles).
					Times(1)
				return articleServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in: func() dto.ArticlesInDTO {
					in, _ := dto.NewArticlesInDTO(dto.ArticlesInWithLast(2), dto.ArticlesInWithBefore("Article2"))
					return in
				}(),
			},
			want: want{
				out: dto.ArticlesOutDTO{},
				err: errTestArticles,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			articleServiceClient := tt.articleServiceClient(ctrl)
			u := NewArticles(articleServiceClient)
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
		out dto.ArticlesOutDTO
		err error
	}
	type testCase struct {
		articleServiceClient func(ctrl *gomock.Controller) articleconnect.ArticleServiceClient
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
			articleServiceClient: func(ctrl *gomock.Controller) articleconnect.ArticleServiceClient {
				articleServiceClient := marticleconnect.NewMockArticleServiceClient(ctrl)
				articleServiceClient.EXPECT().
					GetAllArticles(gomock.Any(), gomock.Any()).
					Return(connect.NewResponse(&grpc.GetAllArticlesResponse{
						Articles: []*grpc.Article{
							{
								Id:           "Article1",
								Title:        "happy_path/all_articles",
								Body:         "## happy_path/all_articles",
								ThumbnailUrl: "example.com/example.png",
								CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
								UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
								Tags: []*grpc.Tag{
									{
										Id:   "Tag1",
										Name: "Tag1",
									},
								},
							},
						},
					}), nil).
					Times(1)
				return articleServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
			},
			want: want{
				out: dto.NewArticlesOutDTO(
					[]dto.ArticleTag{
						dto.NewArticleTag(
							"Article1",
							"happy_path/all_articles",
							"## happy_path/all_articles",
							utils.MustURLParse("example.com/example.png"),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							[]dto.Tag{
								dto.NewTag("Tag1", "Tag1"),
							}),
					},
				),
			},
		},
		"unhappy_path/all_articles_returns_error": {
			articleServiceClient: func(ctrl *gomock.Controller) articleconnect.ArticleServiceClient {
				articleServiceClient := marticleconnect.NewMockArticleServiceClient(ctrl)
				articleServiceClient.EXPECT().
					GetAllArticles(gomock.Any(), gomock.Any()).
					Return(connect.NewResponse(&grpc.GetAllArticlesResponse{}), errTestArticles).
					Times(1)
				return articleServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
			},
			want: want{
				out: dto.ArticlesOutDTO{},
				err: errTestArticles,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			articleServiceClient := tt.articleServiceClient(ctrl)
			u := NewArticles(articleServiceClient)
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
