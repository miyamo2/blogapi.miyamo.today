package resolver

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/google/go-cmp/cmp"
	"github.com/miyamo2/blogapi/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi/internal/if-adapter/presenters/graphql/model"
	mconverter "github.com/miyamo2/blogapi/internal/mock/if-adapter/controller/graphql/resolver/presenter/converter"
	musecase "github.com/miyamo2/blogapi/internal/mock/if-adapter/controller/graphql/resolver/usecase"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

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
		out dto.ArticleOutDto
		err error
	}
	type converterResult struct {
		out *model.ArticleNode
		ok  bool
	}
	type testCase struct {
		sut                func(rslvr *Resolver) *queryResolver
		setupMockUsecase   func(uc *musecase.MockArticle[dto.ArticleInDto, dto.Tag, dto.ArticleTag, dto.ArticleOutDto], usecaseResult usecaseResult)
		usecaseResult      usecaseResult
		setupMockConverter func(cnvrtr *mconverter.MockArticleConverter[dto.Tag, dto.ArticleTag, dto.ArticleOutDto], from dto.ArticleOutDto, converterResult converterResult)
		converterResult    converterResult
		args               args
		want               want
		wantErr            bool
	}
	errFailedToUsecase := errors.New("failed to usecase")
	tests := map[string]testCase{
		"happy_path": {
			sut: func(rslvr *Resolver) *queryResolver {
				return &queryResolver{rslvr}
			},
			setupMockUsecase: func(uc *musecase.MockArticle[dto.ArticleInDto, dto.Tag, dto.ArticleTag, dto.ArticleOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewArticleOutDto(
					dto.NewArticleTag(
						"Article1",
						"Article1",
						"## Article1",
						"example.test",
						"2020-01-01T00:00:00Z",
						"2020-01-01T00:00:00Z",
						[]dto.Tag{
							dto.NewTag("Tag1", "Tag1"),
						})),
				err: nil,
			},
			setupMockConverter: func(cnvrtr *mconverter.MockArticleConverter[dto.Tag, dto.ArticleTag, dto.ArticleOutDto], from dto.ArticleOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
					ToArticle(gomock.Any(), gomock.Any()).
					Return(converterResult.out, converterResult.ok).
					Times(1)
			},
			converterResult: converterResult{
				out: &model.ArticleNode{
					ID:      "Article1",
					Title:   "Article1",
					Content: "## Article1",
					ThumbnailURL: func() *string {
						v := "example.test"
						return &v
					}(),
					CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				ok: true,
			},
			args: args{
				ctx: context.Background(),
				id:  "Article1",
			},
			want: want{
				out: &model.ArticleNode{
					ID:      "Article1",
					Title:   "Article1",
					Content: "## Article1",
					ThumbnailURL: func() *string {
						v := "example.test"
						return &v
					}(),
					CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				err: nil,
			},
		},
		"unhappy_path/usecase_returned_error": {
			sut: func(rslvr *Resolver) *queryResolver {
				return &queryResolver{rslvr}
			},
			setupMockUsecase: func(uc *musecase.MockArticle[dto.ArticleInDto, dto.Tag, dto.ArticleTag, dto.ArticleOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				err: errFailedToUsecase,
			},
			setupMockConverter: func(cnvrtr *mconverter.MockArticleConverter[dto.Tag, dto.ArticleTag, dto.ArticleOutDto], from dto.ArticleOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
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
			sut: func(rslvr *Resolver) *queryResolver {
				return &queryResolver{rslvr}
			},
			setupMockUsecase: func(uc *musecase.MockArticle[dto.ArticleInDto, dto.Tag, dto.ArticleTag, dto.ArticleOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewArticleOutDto(
					dto.NewArticleTag(
						"Article1",
						"Article1",
						"## Article1",
						"example.test",
						"2020-01-01T00:00:00Z",
						"2020-01-01T00:00:00Z",
						[]dto.Tag{},
					),
				),
				err: nil,
			},
			setupMockConverter: func(cnvrtr *mconverter.MockArticleConverter[dto.Tag, dto.ArticleTag, dto.ArticleOutDto], from dto.ArticleOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
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
			uc := musecase.NewMockArticle[dto.ArticleInDto, dto.Tag, dto.ArticleTag, dto.ArticleOutDto](ctrl)
			tt.setupMockUsecase(uc, tt.usecaseResult)
			cnvrtr := mconverter.NewMockArticleConverter[dto.Tag, dto.ArticleTag, dto.ArticleOutDto](ctrl)
			tt.setupMockConverter(cnvrtr, tt.usecaseResult.out, tt.converterResult)
			sut := tt.sut(NewResolver(NewUsecases(WithArticleUsecase(uc)), NewConverters(WithArticleConverter(cnvrtr))))
			got, err := sut.Article(tt.args.ctx, tt.args.id)
			if tt.wantErr {
				if !errors.Is(err, tt.want.err) {
					t.Errorf("Article() got = %v, want %v", err, tt.want.err)
					return
				}
			} else if err != nil {
				t.Errorf("Article() got = %v, want %v", err, tt.want.err)
				return
			}
			if diff := cmp.Diff(got, tt.want.out); diff != "" {
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
		out dto.ArticlesOutDto
		err error
	}
	type converterResult struct {
		out *model.ArticleConnection
		ok  bool
	}
	type testCase struct {
		sut                func(rslvr *Resolver) *queryResolver
		setupMockUsecase   func(uc *musecase.MockArticles[dto.ArticlesInDto, dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], usecaseResult usecaseResult)
		usecaseResult      usecaseResult
		setupMockConverter func(cnvrtr *mconverter.MockArticlesConverter[dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], from dto.ArticlesOutDto, converterResult converterResult)
		converterResult    converterResult
		args               args
		want               want
		wantErr            bool
	}
	errFailedToUsecase := errors.New("failed to usecase")
	tests := map[string]testCase{
		"happy_path/without_paging": {
			sut: func(rslvr *Resolver) *queryResolver {
				return &queryResolver{rslvr}
			},
			setupMockUsecase: func(uc *musecase.MockArticles[dto.ArticlesInDto, dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewArticlesOutDto([]dto.ArticleTag{
					dto.NewArticleTag(
						"Article1",
						"Article1",
						"## Article1",
						"example.test",
						"2020-01-01T00:00:00Z",
						"2020-01-01T00:00:00Z",
						[]dto.Tag{
							dto.NewTag("Tag1", "Tag1"),
						}),
				}),
				err: nil,
			},
			setupMockConverter: func(cnvrtr *mconverter.MockArticlesConverter[dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], from dto.ArticlesOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
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
								ID:      "Article1",
								Title:   "Article1",
								Content: "## Article1",
								ThumbnailURL: func() *string {
									v := "example.test"
									return &v
								}(),
								CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
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
								ID:      "Article1",
								Title:   "Article1",
								Content: "## Article1",
								ThumbnailURL: func() *string {
									v := "example.test"
									return &v
								}(),
								CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
								UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
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
			sut: func(rslvr *Resolver) *queryResolver {
				return &queryResolver{rslvr}
			},
			setupMockUsecase: func(uc *musecase.MockArticles[dto.ArticlesInDto, dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				err: errFailedToUsecase,
			},
			setupMockConverter: func(cnvrtr *mconverter.MockArticlesConverter[dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], from dto.ArticlesOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
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
			sut: func(rslvr *Resolver) *queryResolver {
				return &queryResolver{rslvr}
			},
			setupMockUsecase: func(uc *musecase.MockArticles[dto.ArticlesInDto, dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewArticlesOutDto([]dto.ArticleTag{}),
				err: nil,
			},
			setupMockConverter: func(cnvrtr *mconverter.MockArticlesConverter[dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], from dto.ArticlesOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
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
			sut: func(rslvr *Resolver) *queryResolver {
				return &queryResolver{rslvr}
			},
			setupMockUsecase: func(uc *musecase.MockArticles[dto.ArticlesInDto, dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			setupMockConverter: func(cnvrtr *mconverter.MockArticlesConverter[dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], from dto.ArticlesOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
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
			sut: func(rslvr *Resolver) *queryResolver {
				return &queryResolver{rslvr}
			},
			setupMockUsecase: func(uc *musecase.MockArticles[dto.ArticlesInDto, dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			setupMockConverter: func(cnvrtr *mconverter.MockArticlesConverter[dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], from dto.ArticlesOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
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
			sut: func(rslvr *Resolver) *queryResolver {
				return &queryResolver{rslvr}
			},
			setupMockUsecase: func(uc *musecase.MockArticles[dto.ArticlesInDto, dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			setupMockConverter: func(cnvrtr *mconverter.MockArticlesConverter[dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], from dto.ArticlesOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
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
			sut: func(rslvr *Resolver) *queryResolver {
				return &queryResolver{rslvr}
			},
			setupMockUsecase: func(uc *musecase.MockArticles[dto.ArticlesInDto, dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			setupMockConverter: func(cnvrtr *mconverter.MockArticlesConverter[dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], from dto.ArticlesOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
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
			sut: func(rslvr *Resolver) *queryResolver {
				return &queryResolver{rslvr}
			},
			args: args{
				ctx:   context.Background(),
				first: func() *int { v := 10; return &v }(),
				last:  func() *int { v := 10; return &v }(),
			},
			want: want{
				out: nil,
				err: dto.ErrInvalidateArticlesInDto,
			},
			setupMockUsecase: func(uc *musecase.MockArticles[dto.ArticlesInDto, dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Times(0)
			},
			setupMockConverter: func(cnvrtr *mconverter.MockArticlesConverter[dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], from dto.ArticlesOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
					ToArticles(gomock.Any(), gomock.Any()).
					Times(0)
			},
			wantErr: true,
		},
		"unhappy_path/with_first/with_before": {
			sut: func(rslvr *Resolver) *queryResolver {
				return &queryResolver{rslvr}
			},
			args: args{
				ctx:    context.Background(),
				first:  func() *int { v := 10; return &v }(),
				before: func() *string { v := "Article1"; return &v }(),
			},
			want: want{
				out: nil,
				err: dto.ErrInvalidateArticlesInDto,
			},
			setupMockUsecase: func(uc *musecase.MockArticles[dto.ArticlesInDto, dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Times(0)
			},
			setupMockConverter: func(cnvrtr *mconverter.MockArticlesConverter[dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], from dto.ArticlesOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
					ToArticles(gomock.Any(), gomock.Any()).
					Times(0)
			},
			wantErr: true,
		},
		"unhappy_path/with_last/with_after": {
			sut: func(rslvr *Resolver) *queryResolver {
				return &queryResolver{rslvr}
			},
			args: args{
				ctx:   context.Background(),
				last:  func() *int { v := 10; return &v }(),
				after: func() *string { v := "Article1"; return &v }(),
			},
			want: want{
				out: nil,
				err: dto.ErrInvalidateArticlesInDto,
			},
			setupMockUsecase: func(uc *musecase.MockArticles[dto.ArticlesInDto, dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Times(0)
			},
			setupMockConverter: func(cnvrtr *mconverter.MockArticlesConverter[dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], from dto.ArticlesOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
					ToArticles(gomock.Any(), gomock.Any()).
					Times(0)
			},
			wantErr: true,
		},
		"unhappy_path/with_before": {
			sut: func(rslvr *Resolver) *queryResolver {
				return &queryResolver{rslvr}
			},
			args: args{
				ctx:    context.Background(),
				before: func() *string { v := "Article1"; return &v }(),
			},
			want: want{
				out: nil,
				err: dto.ErrInvalidateArticlesInDto,
			},
			setupMockUsecase: func(uc *musecase.MockArticles[dto.ArticlesInDto, dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Times(0)
			},
			setupMockConverter: func(cnvrtr *mconverter.MockArticlesConverter[dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], from dto.ArticlesOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
					ToArticles(gomock.Any(), gomock.Any()).
					Times(0)
			},
			wantErr: true,
		},
		"unhappy_path/with_after": {
			sut: func(rslvr *Resolver) *queryResolver {
				return &queryResolver{rslvr}
			},
			args: args{
				ctx:   context.Background(),
				after: func() *string { v := "Article1"; return &v }(),
			},
			want: want{
				out: nil,
				err: dto.ErrInvalidateArticlesInDto,
			},
			setupMockUsecase: func(uc *musecase.MockArticles[dto.ArticlesInDto, dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Times(0)
			},
			setupMockConverter: func(cnvrtr *mconverter.MockArticlesConverter[dto.Tag, dto.ArticleTag, dto.ArticlesOutDto], from dto.ArticlesOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
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
			uc := musecase.NewMockArticles[dto.ArticlesInDto, dto.Tag, dto.ArticleTag, dto.ArticlesOutDto](ctrl)
			tt.setupMockUsecase(uc, tt.usecaseResult)
			cnvrtr := mconverter.NewMockArticlesConverter[dto.Tag, dto.ArticleTag, dto.ArticlesOutDto](ctrl)
			tt.setupMockConverter(cnvrtr, tt.usecaseResult.out, tt.converterResult)
			sut := tt.sut(NewResolver(NewUsecases(WithArticlesUsecase(uc)), NewConverters(WithArticlesConverter(cnvrtr))))
			got, err := sut.Articles(tt.args.ctx, tt.args.first, tt.args.last, tt.args.after, tt.args.before)
			if tt.wantErr {
				if !errors.Is(err, tt.want.err) {
					t.Errorf("Article() got = %v, want %v", err, tt.want.err)
					return
				}
			} else if err != nil {
				t.Errorf("Article() got = %v, want %v", err, tt.want.err)
				return
			}
			if diff := cmp.Diff(got, tt.want.out); diff != "" {
				t.Error(diff)
				return
			}
		})
	}
}
