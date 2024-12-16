package resolver

import (
	"context"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/pkg/gqlscalar"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/utils"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/google/go-cmp/cmp"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/presenters/graphql/model"
	mconverter "github.com/miyamo2/blogapi.miyamo.today/federator/internal/mock/if-adapter/controller/graphql/resolver/presenter/converter"
	musecase "github.com/miyamo2/blogapi.miyamo.today/federator/internal/mock/if-adapter/controller/graphql/resolver/usecase"
	"go.uber.org/mock/gomock"
)

var cmpOpts = []cmp.Option{
	cmp.AllowUnexported(gqlscalar.URL{}),
	cmp.AllowUnexported(gqlscalar.UTC{}),
}

func Test_queryResolver_Article(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	type want struct {
		out *model.ArticleNode
		err error
	}
	type usecaseResult struct {
		out dto.ArticleOutDTO
		err error
	}
	type converterResult struct {
		out *model.ArticleNode
		ok  bool
	}
	type testCase struct {
		sut                func(resolver *Resolver) *queryResolver
		setupMockUsecase   func(uc *musecase.MockArticle, usecaseResult usecaseResult)
		usecaseResult      usecaseResult
		setupMockConverter func(converter *mconverter.MockArticleConverter, from dto.ArticleOutDTO, converterResult converterResult)
		converterResult    converterResult
		args               args
		want               want
		wantErr            bool
	}
	errFailedToUsecase := errors.New("failed to usecase")
	tests := map[string]testCase{
		"happy_path": {
			sut: func(resolver *Resolver) *queryResolver {
				return &queryResolver{resolver}
			},
			setupMockUsecase: func(uc *musecase.MockArticle, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewArticleOutDTO(
					dto.NewArticleTag(
						"Article1",
						"Article1",
						"## Article1",
						utils.MustURLParse("example.com/example.png"),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						[]dto.Tag{
							dto.NewTag("Tag1", "Tag1"),
						})),
				err: nil,
			},
			setupMockConverter: func(converter *mconverter.MockArticleConverter, from dto.ArticleOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToArticle(gomock.Any(), gomock.Any()).
					Return(converterResult.out, converterResult.ok).
					Times(1)
			},
			converterResult: converterResult{
				out: &model.ArticleNode{
					ID:           "Article1",
					Title:        "Article1",
					Content:      "## Article1",
					ThumbnailURL: gqlscalar.URL(utils.MustURLParse("example.com/example.png")),
					CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
					UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
				},
				ok: true,
			},
			args: args{
				ctx: context.Background(),
				id:  "Article1",
			},
			want: want{
				out: &model.ArticleNode{
					ID:           "Article1",
					Title:        "Article1",
					Content:      "## Article1",
					ThumbnailURL: gqlscalar.URL(utils.MustURLParse("example.com/example.png")),
					CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
					UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
				},
				err: nil,
			},
		},
		"unhappy_path/usecase_returned_error": {
			sut: func(resolver *Resolver) *queryResolver {
				return &queryResolver{resolver}
			},
			setupMockUsecase: func(uc *musecase.MockArticle, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				err: errFailedToUsecase,
			},
			setupMockConverter: func(converter *mconverter.MockArticleConverter, from dto.ArticleOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToArticle(gomock.Any(), gomock.Any()).
					Times(0)
			},
			args: args{
				ctx: context.Background(),
				id:  "Article1",
			},
			want: want{
				out: nil,
				err: errFailedToUsecase,
			},
			wantErr: true,
		},
		"unhappy_path/converter_returned_error": {
			sut: func(resolver *Resolver) *queryResolver {
				return &queryResolver{resolver}
			},
			setupMockUsecase: func(uc *musecase.MockArticle, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewArticleOutDTO(
					dto.NewArticleTag(
						"Article1",
						"Article1",
						"## Article1",
						utils.MustURLParse("example.com/example.png"),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						[]dto.Tag{},
					),
				),
				err: nil,
			},
			setupMockConverter: func(converter *mconverter.MockArticleConverter, from dto.ArticleOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToArticle(gomock.Any(), gomock.Any()).
					Return(converterResult.out, converterResult.ok).
					Times(1)
			},
			converterResult: converterResult{
				out: nil,
				ok:  false,
			},
			args: args{
				ctx: context.Background(),
				id:  "Article1",
			},
			want: want{
				out: nil,
				err: ErrFailedToConvertToArticleNode,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			uc := musecase.NewMockArticle(ctrl)
			tt.setupMockUsecase(uc, tt.usecaseResult)
			converter := mconverter.NewMockArticleConverter(ctrl)
			tt.setupMockConverter(converter, tt.usecaseResult.out, tt.converterResult)
			sut := tt.sut(NewResolver(NewUsecases(WithArticleUsecase(uc)), NewConverters(WithArticleConverter(converter))))
			got, err := sut.Article(tt.args.ctx, tt.args.id)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("Article() got = %v, want %v", err, tt.want.err)
				return
			}
			if diff := cmp.Diff(got, tt.want.out, cmpOpts...); diff != "" {
				t.Error(diff)
				return
			}
		})
	}
}

func Test_queryResolver_Articles(t *testing.T) {
	type args struct {
		ctx    context.Context
		first  *int
		last   *int
		after  *string
		before *string
	}
	type want struct {
		out *model.ArticleConnection
		err error
	}
	type usecaseResult struct {
		out dto.ArticlesOutDTO
		err error
	}
	type converterResult struct {
		out *model.ArticleConnection
		ok  bool
	}
	type testCase struct {
		sut                func(resolver *Resolver) *queryResolver
		setupMockUsecase   func(uc *musecase.MockArticles, usecaseResult usecaseResult)
		usecaseResult      usecaseResult
		setupMockConverter func(converter *mconverter.MockArticlesConverter, from dto.ArticlesOutDTO, converterResult converterResult)
		converterResult    converterResult
		args               args
		want               want
		wantErr            bool
	}
	errFailedToUsecase := errors.New("failed to usecase")
	tests := map[string]testCase{
		"happy_path/without_paging": {
			sut: func(resolver *Resolver) *queryResolver {
				return &queryResolver{resolver}
			},
			setupMockUsecase: func(uc *musecase.MockArticles, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewArticlesOutDTO([]dto.ArticleTag{
					dto.NewArticleTag(
						"Article1",
						"Article1",
						"## Article1",
						utils.MustURLParse("example.com/example.png"),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						[]dto.Tag{
							dto.NewTag("Tag1", "Tag1"),
						}),
				}),
				err: nil,
			},
			setupMockConverter: func(converter *mconverter.MockArticlesConverter, from dto.ArticlesOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToArticles(gomock.Any(), gomock.Any()).
					Return(converterResult.out, converterResult.ok).
					Times(1)
			},
			converterResult: converterResult{
				out: &model.ArticleConnection{
					Edges: []*model.ArticleEdge{
						{
							Cursor: "Article1",
							Node: &model.ArticleNode{
								ID:           "Article1",
								Title:        "Article1",
								Content:      "## Article1",
								ThumbnailURL: gqlscalar.URL(utils.MustURLParse("example.com/example.png")),
								CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								Tags: &model.ArticleTagConnection{
									Edges: []*model.ArticleTagEdge{
										{
											Cursor: "Tag1",
											Node: &model.ArticleTagNode{
												ID:   "Tag1",
												Name: "Tag1",
											},
										},
									},
									PageInfo: &model.PageInfo{
										HasNextPage:     nil,
										HasPreviousPage: nil,
										StartCursor:     "Tag1",
										EndCursor:       "Tag1",
									},
									TotalCount: 1,
								},
							},
						},
					},
					PageInfo: &model.PageInfo{
						HasNextPage:     nil,
						HasPreviousPage: nil,
						StartCursor:     "Article1",
						EndCursor:       "Article1",
					},
					TotalCount: 1,
				},
				ok: true,
			},
			args: args{
				ctx: context.Background(),
			},
			want: want{
				out: &model.ArticleConnection{
					Edges: []*model.ArticleEdge{
						{
							Cursor: "Article1",
							Node: &model.ArticleNode{
								ID:           "Article1",
								Title:        "Article1",
								Content:      "## Article1",
								ThumbnailURL: gqlscalar.URL(utils.MustURLParse("example.com/example.png")),
								CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
								Tags: &model.ArticleTagConnection{
									Edges: []*model.ArticleTagEdge{
										{
											Cursor: "Tag1",
											Node: &model.ArticleTagNode{
												ID:   "Tag1",
												Name: "Tag1",
											},
										},
									},
									PageInfo: &model.PageInfo{
										StartCursor: "Tag1",
										EndCursor:   "Tag1",
									},
									TotalCount: 1,
								},
							},
						},
					},
					PageInfo: &model.PageInfo{
						StartCursor: "Article1",
						EndCursor:   "Article1",
					},
					TotalCount: 1,
				},
			},
		},
		"unhappy_path/usecase_returned_error": {
			sut: func(resolver *Resolver) *queryResolver {
				return &queryResolver{resolver}
			},
			setupMockUsecase: func(uc *musecase.MockArticles, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				err: errFailedToUsecase,
			},
			setupMockConverter: func(converter *mconverter.MockArticlesConverter, from dto.ArticlesOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToArticles(gomock.Any(), gomock.Any()).
					Times(0)
			},
			args: args{
				ctx: context.Background(),
			},
			want: want{
				out: nil,
				err: errFailedToUsecase,
			},
			wantErr: true,
		},
		"unhappy_path/converter_returned_error": {
			sut: func(resolver *Resolver) *queryResolver {
				return &queryResolver{resolver}
			},
			setupMockUsecase: func(uc *musecase.MockArticles, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewArticlesOutDTO([]dto.ArticleTag{}),
				err: nil,
			},
			setupMockConverter: func(converter *mconverter.MockArticlesConverter, from dto.ArticlesOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToArticles(gomock.Any(), gomock.Any()).
					Return(converterResult.out, converterResult.ok).
					Times(1)
			},
			converterResult: converterResult{
				out: nil,
				ok:  false,
			},
			args: args{
				ctx: context.Background(),
			},
			want: want{
				out: nil,
				err: ErrFailedToConvertToArticleConnection,
			},
			wantErr: true,
		},
		"happy_path/with_first": {
			sut: func(resolver *Resolver) *queryResolver {
				return &queryResolver{resolver}
			},
			setupMockUsecase: func(uc *musecase.MockArticles, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			setupMockConverter: func(converter *mconverter.MockArticlesConverter, from dto.ArticlesOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToArticles(gomock.Any(), gomock.Any()).
					Return(converterResult.out, converterResult.ok).
					Times(1)
			},
			converterResult: converterResult{
				out: nil,
				ok:  true,
			},
			args: args{
				ctx:   context.Background(),
				first: func() *int { v := 10; return &v }(),
			},
		},
		"happy_path/with_last": {
			sut: func(resolver *Resolver) *queryResolver {
				return &queryResolver{resolver}
			},
			setupMockUsecase: func(uc *musecase.MockArticles, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			setupMockConverter: func(converter *mconverter.MockArticlesConverter, from dto.ArticlesOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToArticles(gomock.Any(), gomock.Any()).
					Return(converterResult.out, converterResult.ok).
					Times(1)
			},
			converterResult: converterResult{
				out: nil,
				ok:  true,
			},
			args: args{
				ctx: context.Background(),
				last: func() *int {
					v := 10
					return &v
				}(),
			},
		},
		"happy_path/with_first/with_after": {
			sut: func(resolver *Resolver) *queryResolver {
				return &queryResolver{resolver}
			},
			setupMockUsecase: func(uc *musecase.MockArticles, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			setupMockConverter: func(converter *mconverter.MockArticlesConverter, from dto.ArticlesOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToArticles(gomock.Any(), gomock.Any()).
					Return(converterResult.out, converterResult.ok).
					Times(1)
			},
			converterResult: converterResult{
				out: nil,
				ok:  true,
			},
			args: args{
				ctx:   context.Background(),
				first: func() *int { v := 10; return &v }(),
				after: func() *string { v := "Article1"; return &v }(),
			},
		},
		"happy_path/with_last/with_before": {
			sut: func(resolver *Resolver) *queryResolver {
				return &queryResolver{resolver}
			},
			setupMockUsecase: func(uc *musecase.MockArticles, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			setupMockConverter: func(converter *mconverter.MockArticlesConverter, from dto.ArticlesOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToArticles(gomock.Any(), gomock.Any()).
					Return(converterResult.out, converterResult.ok).
					Times(1)
			},
			converterResult: converterResult{
				out: nil,
				ok:  true,
			},
			args: args{
				ctx:    context.Background(),
				last:   func() *int { v := 10; return &v }(),
				before: func() *string { v := "Article1"; return &v }(),
			},
		},
		"unhappy_path/with_first/with_last": {
			sut: func(resolver *Resolver) *queryResolver {
				return &queryResolver{resolver}
			},
			args: args{
				ctx:   context.Background(),
				first: func() *int { v := 10; return &v }(),
				last:  func() *int { v := 10; return &v }(),
			},
			want: want{
				out: nil,
				err: dto.ErrInvalidateArticlesInDTO,
			},
			setupMockUsecase: func(uc *musecase.MockArticles, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Times(0)
			},
			setupMockConverter: func(converter *mconverter.MockArticlesConverter, from dto.ArticlesOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToArticles(gomock.Any(), gomock.Any()).
					Times(0)
			},
			wantErr: true,
		},
		"unhappy_path/with_first/with_before": {
			sut: func(resolver *Resolver) *queryResolver {
				return &queryResolver{resolver}
			},
			args: args{
				ctx:    context.Background(),
				first:  func() *int { v := 10; return &v }(),
				before: func() *string { v := "Article1"; return &v }(),
			},
			want: want{
				out: nil,
				err: dto.ErrInvalidateArticlesInDTO,
			},
			setupMockUsecase: func(uc *musecase.MockArticles, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Times(0)
			},
			setupMockConverter: func(converter *mconverter.MockArticlesConverter, from dto.ArticlesOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToArticles(gomock.Any(), gomock.Any()).
					Times(0)
			},
			wantErr: true,
		},
		"unhappy_path/with_last/with_after": {
			sut: func(resolver *Resolver) *queryResolver {
				return &queryResolver{resolver}
			},
			args: args{
				ctx:   context.Background(),
				last:  func() *int { v := 10; return &v }(),
				after: func() *string { v := "Article1"; return &v }(),
			},
			want: want{
				out: nil,
				err: dto.ErrInvalidateArticlesInDTO,
			},
			setupMockUsecase: func(uc *musecase.MockArticles, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Times(0)
			},
			setupMockConverter: func(converter *mconverter.MockArticlesConverter, from dto.ArticlesOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToArticles(gomock.Any(), gomock.Any()).
					Times(0)
			},
			wantErr: true,
		},
		"unhappy_path/with_before": {
			sut: func(resolver *Resolver) *queryResolver {
				return &queryResolver{resolver}
			},
			args: args{
				ctx:    context.Background(),
				before: func() *string { v := "Article1"; return &v }(),
			},
			want: want{
				out: nil,
				err: dto.ErrInvalidateArticlesInDTO,
			},
			setupMockUsecase: func(uc *musecase.MockArticles, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Times(0)
			},
			setupMockConverter: func(converter *mconverter.MockArticlesConverter, from dto.ArticlesOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToArticles(gomock.Any(), gomock.Any()).
					Times(0)
			},
			wantErr: true,
		},
		"unhappy_path/with_after": {
			sut: func(resolver *Resolver) *queryResolver {
				return &queryResolver{resolver}
			},
			args: args{
				ctx:   context.Background(),
				after: func() *string { v := "Article1"; return &v }(),
			},
			want: want{
				out: nil,
				err: dto.ErrInvalidateArticlesInDTO,
			},
			setupMockUsecase: func(uc *musecase.MockArticles, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Times(0)
			},
			setupMockConverter: func(converter *mconverter.MockArticlesConverter, from dto.ArticlesOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToArticles(gomock.Any(), gomock.Any()).
					Times(0)
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			uc := musecase.NewMockArticles(ctrl)
			tt.setupMockUsecase(uc, tt.usecaseResult)
			converter := mconverter.NewMockArticlesConverter(ctrl)
			tt.setupMockConverter(converter, tt.usecaseResult.out, tt.converterResult)
			sut := tt.sut(NewResolver(NewUsecases(WithArticlesUsecase(uc)), NewConverters(WithArticlesConverter(converter))))
			got, err := sut.Articles(tt.args.ctx, tt.args.first, tt.args.last, tt.args.after, tt.args.before)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("Article() got = %v, want %v", err, tt.want.err)
				return
			}
			if diff := cmp.Diff(got, tt.want.out, cmpOpts...); diff != "" {
				t.Error(diff)
				return
			}
		})
	}
}
