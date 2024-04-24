package resolver

import (
	"context"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/google/go-cmp/cmp"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/presenters/graphql/model"
	mconverter "github.com/miyamo2/blogapi.miyamo.today/federator/internal/mock/if-adapter/controller/graphql/resolver/presenter/converter"
	musecase "github.com/miyamo2/blogapi.miyamo.today/federator/internal/mock/if-adapter/controller/graphql/resolver/usecase"
	"go.uber.org/mock/gomock"
)

func Test_queryResolver_Tag(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	type want struct {
		out *model.TagNode
		err error
	}
	type usecaseResult struct {
		out dto.TagOutDto
		err error
	}
	type converterResult struct {
		out *model.TagNode
		err error
	}
	type testCase struct {
		sut                func(rslvr *Resolver) *queryResolver
		setupMockUsecase   func(uc *musecase.MockTag[dto.TagInDto, dto.Article, dto.TagArticle, dto.TagOutDto], usecaseResult usecaseResult)
		usecaseResult      usecaseResult
		setupMockConverter func(cnvrtr *mconverter.MockTagConverter[dto.Article, dto.TagArticle, dto.TagOutDto], from dto.TagOutDto, converterResult converterResult)
		converterResult    converterResult
		args               args
		want               want
		wantErr            bool
	}
	errAtConverter := errors.New("error at converter")
	errAtUseCase := errors.New("error at usecase")
	tests := map[string]testCase{
		"happy_path": {
			sut: func(rslvr *Resolver) *queryResolver {
				return &queryResolver{rslvr}
			},
			setupMockUsecase: func(uc *musecase.MockTag[dto.TagInDto, dto.Article, dto.TagArticle, dto.TagOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewTagOutDto(
					dto.NewTagArticle(
						"Tag1",
						"Tag1",
						[]dto.Article{
							dto.NewArticle(
								"Article1",
								"Article1",
								"",
								"example.test",
								"2020-01-01T00:00:00Z",
								"2020-01-01T00:00:00Z"),
						})),
				err: nil,
			},
			setupMockConverter: func(cnvrtr *mconverter.MockTagConverter[dto.Article, dto.TagArticle, dto.TagOutDto], from dto.TagOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
					ToTag(gomock.Any(), gomock.Any()).
					Return(converterResult.out, converterResult.err).
					Times(1)
			},
			converterResult: converterResult{
				out: &model.TagNode{
					ID:   "Tag1",
					Name: "Tag1",
				},
			},
			args: args{
				ctx: context.Background(),
				id:  "Tag1",
			},
			want: want{
				out: &model.TagNode{
					ID:   "Tag1",
					Name: "Tag1",
				},
			},
		},
		"error_at_usecase": {
			sut: func(rslvr *Resolver) *queryResolver {
				return &queryResolver{rslvr}
			},
			setupMockUsecase: func(uc *musecase.MockTag[dto.TagInDto, dto.Article, dto.TagArticle, dto.TagOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.TagOutDto{},
				err: errAtUseCase,
			},
			setupMockConverter: func(cnvrtr *mconverter.MockTagConverter[dto.Article, dto.TagArticle, dto.TagOutDto], from dto.TagOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
					ToTag(gomock.Any(), gomock.Any()).
					Times(0)
			},
			args: args{
				ctx: context.Background(),
				id:  "Tag1",
			},
			want: want{
				out: nil,
				err: errAtUseCase,
			},
			wantErr: true,
		},
		"error_at_converter": {
			sut: func(rslvr *Resolver) *queryResolver {
				return &queryResolver{rslvr}
			},
			setupMockUsecase: func(uc *musecase.MockTag[dto.TagInDto, dto.Article, dto.TagArticle, dto.TagOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewTagOutDto(
					dto.NewTagArticle(
						"Tag1",
						"Tag1",
						[]dto.Article{
							dto.NewArticle(
								"Article1",
								"Article1",
								"",
								"example.test",
								"2020-01-01T00:00:00Z",
								"2020-01-01T00:00:00Z"),
						})),
				err: nil,
			},
			setupMockConverter: func(cnvrtr *mconverter.MockTagConverter[dto.Article, dto.TagArticle, dto.TagOutDto], from dto.TagOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
					ToTag(gomock.Any(), gomock.Any()).
					Return(converterResult.out, converterResult.err).
					Times(1)
			},
			converterResult: converterResult{
				out: nil,
				err: errAtConverter,
			},
			args: args{
				ctx: context.Background(),
				id:  "Tag1",
			},
			want: want{
				out: nil,
				err: errAtConverter,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			uc := musecase.NewMockTag[dto.TagInDto, dto.Article, dto.TagArticle, dto.TagOutDto](ctrl)
			tt.setupMockUsecase(uc, tt.usecaseResult)
			cnvrtr := mconverter.NewMockTagConverter[dto.Article, dto.TagArticle, dto.TagOutDto](ctrl)
			tt.setupMockConverter(cnvrtr, tt.usecaseResult.out, tt.converterResult)
			sut := tt.sut(NewResolver(NewUsecases(WithTagUsecase(uc)), NewConverters(WithTagConverter(cnvrtr))))
			got, err := sut.Tag(tt.args.ctx, tt.args.id)
			if tt.wantErr {
				if !errors.Is(err, tt.want.err) {
					t.Errorf("Tag() got = %v, want %v", err, tt.want.err)
					return
				}
			} else if err != nil {
				t.Errorf("Tag() got = %v, want %v", err, tt.want.err)
				return
			}
			if diff := cmp.Diff(got, tt.want.out); diff != "" {
				t.Error(diff)
				return
			}
		})
	}
}

func Test_queryResolver_Tags(t *testing.T) {
	type args struct {
		ctx    context.Context
		first  *int
		last   *int
		after  *string
		before *string
	}
	type want struct {
		out *model.TagConnection
		err error
	}
	type usecaseResult struct {
		out dto.TagsOutDto
		err error
	}
	type converterResult struct {
		out *model.TagConnection
		err error
	}
	type testCase struct {
		sut                func(rslvr *Resolver) *queryResolver
		setupMockUsecase   func(uc *musecase.MockTags[dto.TagsInDto, dto.Article, dto.TagArticle, dto.TagsOutDto], usecaseResult usecaseResult)
		usecaseResult      usecaseResult
		setupMockConverter func(cnvrtr *mconverter.MockTagsConverter[dto.Article, dto.TagArticle, dto.TagsOutDto], from dto.TagsOutDto, converterResult converterResult)
		converterResult    converterResult
		args               args
		want               want
		wantErr            bool
	}
	errFailedToUsecase := errors.New("failed to usecase")
	errFailedToConvert := errors.New("failed to convert")
	tests := map[string]testCase{
		"happy_path/without_paging": {
			sut: func(rslvr *Resolver) *queryResolver {
				return &queryResolver{rslvr}
			},
			setupMockUsecase: func(uc *musecase.MockTags[dto.TagsInDto, dto.Article, dto.TagArticle, dto.TagsOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewTagsOutDto([]dto.TagArticle{
					dto.NewTagArticle(
						"Tag1",
						"Tag1",
						[]dto.Article{
							dto.NewArticle(
								"Article1",
								"Article1",
								"",
								"example.test",
								"2020-01-01T00:00:00Z",
								"2020-01-01T00:00:00Z"),
						}),
				}),
				err: nil,
			},
			setupMockConverter: func(cnvrtr *mconverter.MockTagsConverter[dto.Article, dto.TagArticle, dto.TagsOutDto], from dto.TagsOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
					ToTags(gomock.Any(), gomock.Any()).
					Return(converterResult.out, converterResult.err).
					Times(1)
			},
			converterResult: converterResult{
				out: &model.TagConnection{
					Edges: []*model.TagEdge{
						{
							Cursor: "Tag1",
							Node: &model.TagNode{
								ID:   "Tag1",
								Name: "Tag1",
								Articles: &model.TagArticleConnection{
									Edges: []*model.TagArticleEdge{
										{
											Cursor: "Article1",
											Node: &model.TagArticleNode{
												ID:           "Article1",
												Title:        "Article1",
												ThumbnailURL: "example.test",
												CreatedAt:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
												UpdatedAt:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
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
					},
					PageInfo: &model.PageInfo{
						StartCursor: "Tag1",
						EndCursor:   "Tag1",
					},
					TotalCount: 1,
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want: want{
				out: &model.TagConnection{
					Edges: []*model.TagEdge{
						{
							Cursor: "Tag1",
							Node: &model.TagNode{
								ID:   "Tag1",
								Name: "Tag1",
								Articles: &model.TagArticleConnection{
									Edges: []*model.TagArticleEdge{
										{
											Cursor: "Article1",
											Node: &model.TagArticleNode{
												ID:           "Article1",
												Title:        "Article1",
												ThumbnailURL: "example.test",
												CreatedAt:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
												UpdatedAt:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
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
					},
					PageInfo: &model.PageInfo{
						StartCursor: "Tag1",
						EndCursor:   "Tag1",
					},
					TotalCount: 1,
				},
			},
		},
		"unhappy_path/usecase_returned_error": {
			sut: func(rslvr *Resolver) *queryResolver {
				return &queryResolver{rslvr}
			},
			setupMockUsecase: func(uc *musecase.MockTags[dto.TagsInDto, dto.Article, dto.TagArticle, dto.TagsOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				err: errFailedToUsecase,
			},
			setupMockConverter: func(cnvrtr *mconverter.MockTagsConverter[dto.Article, dto.TagArticle, dto.TagsOutDto], from dto.TagsOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
					ToTags(gomock.Any(), gomock.Any()).
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
			setupMockUsecase: func(uc *musecase.MockTags[dto.TagsInDto, dto.Article, dto.TagArticle, dto.TagsOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewTagsOutDto([]dto.TagArticle{}),
				err: nil,
			},
			setupMockConverter: func(cnvrtr *mconverter.MockTagsConverter[dto.Article, dto.TagArticle, dto.TagsOutDto], from dto.TagsOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
					ToTags(gomock.Any(), gomock.Any()).
					Return(converterResult.out, converterResult.err).
					Times(1)
			},
			converterResult: converterResult{
				out: nil,
				err: errFailedToConvert,
			},
			args: args{
				ctx: context.Background(),
			},
			want: want{
				out: nil,
				err: errFailedToConvert,
			},
			wantErr: true,
		},
		"happy_path/with_first": {
			sut: func(rslvr *Resolver) *queryResolver {
				return &queryResolver{rslvr}
			},
			setupMockUsecase: func(uc *musecase.MockTags[dto.TagsInDto, dto.Article, dto.TagArticle, dto.TagsOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			setupMockConverter: func(cnvrtr *mconverter.MockTagsConverter[dto.Article, dto.TagArticle, dto.TagsOutDto], from dto.TagsOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
					ToTags(gomock.Any(), gomock.Any()).
					Return(converterResult.out, converterResult.err).
					Times(1)
			},
			converterResult: converterResult{
				out: nil,
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
			setupMockUsecase: func(uc *musecase.MockTags[dto.TagsInDto, dto.Article, dto.TagArticle, dto.TagsOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			setupMockConverter: func(cnvrtr *mconverter.MockTagsConverter[dto.Article, dto.TagArticle, dto.TagsOutDto], from dto.TagsOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
					ToTags(gomock.Any(), gomock.Any()).
					Return(converterResult.out, converterResult.err).
					Times(1)
			},
			converterResult: converterResult{
				out: nil,
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
			setupMockUsecase: func(uc *musecase.MockTags[dto.TagsInDto, dto.Article, dto.TagArticle, dto.TagsOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			setupMockConverter: func(cnvrtr *mconverter.MockTagsConverter[dto.Article, dto.TagArticle, dto.TagsOutDto], from dto.TagsOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
					ToTags(gomock.Any(), gomock.Any()).
					Return(converterResult.out, converterResult.err).
					Times(1)
			},
			converterResult: converterResult{
				out: nil,
			},
			args: args{
				ctx:   context.Background(),
				first: func() *int { v := 10; return &v }(),
				after: func() *string { v := "Tag1"; return &v }(),
			},
		},
		"happy_path/with_last/with_before": {
			sut: func(rslvr *Resolver) *queryResolver {
				return &queryResolver{rslvr}
			},
			setupMockUsecase: func(uc *musecase.MockTags[dto.TagsInDto, dto.Article, dto.TagArticle, dto.TagsOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			setupMockConverter: func(cnvrtr *mconverter.MockTagsConverter[dto.Article, dto.TagArticle, dto.TagsOutDto], from dto.TagsOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
					ToTags(gomock.Any(), gomock.Any()).
					Return(converterResult.out, converterResult.err).
					Times(1)
			},
			converterResult: converterResult{
				out: nil,
			},
			args: args{
				ctx:    context.Background(),
				last:   func() *int { v := 10; return &v }(),
				before: func() *string { v := "Tag1"; return &v }(),
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
				err: dto.ErrInvalidateTagsInDto,
			},
			setupMockUsecase: func(uc *musecase.MockTags[dto.TagsInDto, dto.Article, dto.TagArticle, dto.TagsOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Times(0)
			},
			setupMockConverter: func(cnvrtr *mconverter.MockTagsConverter[dto.Article, dto.TagArticle, dto.TagsOutDto], from dto.TagsOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
					ToTags(gomock.Any(), gomock.Any()).
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
				before: func() *string { v := "Tag1"; return &v }(),
			},
			want: want{
				out: nil,
				err: dto.ErrInvalidateTagsInDto,
			},
			setupMockUsecase: func(uc *musecase.MockTags[dto.TagsInDto, dto.Article, dto.TagArticle, dto.TagsOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Times(0)
			},
			setupMockConverter: func(cnvrtr *mconverter.MockTagsConverter[dto.Article, dto.TagArticle, dto.TagsOutDto], from dto.TagsOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
					ToTags(gomock.Any(), gomock.Any()).
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
				after: func() *string { v := "Tag1"; return &v }(),
			},
			want: want{
				out: nil,
				err: dto.ErrInvalidateTagsInDto,
			},
			setupMockUsecase: func(uc *musecase.MockTags[dto.TagsInDto, dto.Article, dto.TagArticle, dto.TagsOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Times(0)
			},
			setupMockConverter: func(cnvrtr *mconverter.MockTagsConverter[dto.Article, dto.TagArticle, dto.TagsOutDto], from dto.TagsOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
					ToTags(gomock.Any(), gomock.Any()).
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
				before: func() *string { v := "Tag1"; return &v }(),
			},
			want: want{
				out: nil,
				err: dto.ErrInvalidateTagsInDto,
			},
			setupMockUsecase: func(uc *musecase.MockTags[dto.TagsInDto, dto.Article, dto.TagArticle, dto.TagsOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Times(0)
			},
			setupMockConverter: func(cnvrtr *mconverter.MockTagsConverter[dto.Article, dto.TagArticle, dto.TagsOutDto], from dto.TagsOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
					ToTags(gomock.Any(), gomock.Any()).
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
				after: func() *string { v := "Tag1"; return &v }(),
			},
			want: want{
				out: nil,
				err: dto.ErrInvalidateTagsInDto,
			},
			setupMockUsecase: func(uc *musecase.MockTags[dto.TagsInDto, dto.Article, dto.TagArticle, dto.TagsOutDto], usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Times(0)
			},
			setupMockConverter: func(cnvrtr *mconverter.MockTagsConverter[dto.Article, dto.TagArticle, dto.TagsOutDto], from dto.TagsOutDto, converterResult converterResult) {
				cnvrtr.EXPECT().
					ToTags(gomock.Any(), gomock.Any()).
					Times(0)
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			uc := musecase.NewMockTags[dto.TagsInDto, dto.Article, dto.TagArticle, dto.TagsOutDto](ctrl)
			tt.setupMockUsecase(uc, tt.usecaseResult)
			cnvrtr := mconverter.NewMockTagsConverter[dto.Article, dto.TagArticle, dto.TagsOutDto](ctrl)
			tt.setupMockConverter(cnvrtr, tt.usecaseResult.out, tt.converterResult)
			sut := tt.sut(NewResolver(NewUsecases(WithTagsUsecase(uc)), NewConverters(WithTagsConverter(cnvrtr))))
			got, err := sut.Tags(tt.args.ctx, tt.args.first, tt.args.last, tt.args.after, tt.args.before)
			if tt.wantErr {
				if !errors.Is(err, tt.want.err) {
					t.Errorf("Tags() got = %v, want %v", err, tt.want.err)
					return
				}
			} else if err != nil {
				t.Errorf("Tags() got = %v, want %v", err, tt.want.err)
				return
			}
			if diff := cmp.Diff(got, tt.want.out); diff != "" {
				t.Error(diff)
				return
			}
		})
	}
}
