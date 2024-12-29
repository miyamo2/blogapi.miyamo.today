package resolver

import (
	"blogapi.miyamo.today/federator/internal/app/usecase/dto"
	"blogapi.miyamo.today/federator/internal/if-adapter/presenters/graphql/model"
	mconverter "blogapi.miyamo.today/federator/internal/mock/if-adapter/controller/graphql/resolver/presenter/converter"
	musecase "blogapi.miyamo.today/federator/internal/mock/if-adapter/controller/graphql/resolver/usecase"
	"blogapi.miyamo.today/federator/internal/pkg/gqlscalar"
	"blogapi.miyamo.today/federator/internal/utils"
	"context"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/cockroachdb/errors"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"
	"testing"
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
		out dto.TagOutDTO
		err error
	}
	type converterResult struct {
		out *model.TagNode
		err error
	}
	type testCase struct {
		sut                func(resolver *Resolver) *queryResolver
		setupMockUsecase   func(uc *musecase.MockTag, usecaseResult usecaseResult)
		usecaseResult      usecaseResult
		setupMockConverter func(converter *mconverter.MockTagConverter, from dto.TagOutDTO, converterResult converterResult)
		converterResult    converterResult
		args               args
		want               want
		wantErr            bool
	}
	errAtConverter := errors.New("error at converter")
	errAtUseCase := errors.New("error at usecase")
	tests := map[string]testCase{
		"happy_path": {
			sut: func(resolver *Resolver) *queryResolver {
				return &queryResolver{resolver}
			},
			setupMockUsecase: func(uc *musecase.MockTag, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewTagOutDTO(
					dto.NewTagArticle(
						"Tag1",
						"Tag1",
						[]dto.Article{
							dto.NewArticle(
								"Article1",
								"Article1",
								"",
								utils.MustURLParse("example.com/example.png"),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
						})),
				err: nil,
			},
			setupMockConverter: func(converter *mconverter.MockTagConverter, from dto.TagOutDTO, converterResult converterResult) {
				converter.EXPECT().
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
			sut: func(resolver *Resolver) *queryResolver {
				return &queryResolver{resolver}
			},
			setupMockUsecase: func(uc *musecase.MockTag, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.TagOutDTO{},
				err: errAtUseCase,
			},
			setupMockConverter: func(converter *mconverter.MockTagConverter, from dto.TagOutDTO, converterResult converterResult) {
				converter.EXPECT().
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
			sut: func(resolver *Resolver) *queryResolver {
				return &queryResolver{resolver}
			},
			setupMockUsecase: func(uc *musecase.MockTag, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewTagOutDTO(
					dto.NewTagArticle(
						"Tag1",
						"Tag1",
						[]dto.Article{
							dto.NewArticle(
								"Article1",
								"Article1",
								"",
								utils.MustURLParse("example.com/example.png"),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
						})),
				err: nil,
			},
			setupMockConverter: func(converter *mconverter.MockTagConverter, from dto.TagOutDTO, converterResult converterResult) {
				converter.EXPECT().
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
			uc := musecase.NewMockTag(ctrl)
			tt.setupMockUsecase(uc, tt.usecaseResult)
			converter := mconverter.NewMockTagConverter(ctrl)
			tt.setupMockConverter(converter, tt.usecaseResult.out, tt.converterResult)
			sut := tt.sut(NewResolver(NewUsecases(WithTagUsecase(uc)), NewConverters(WithTagConverter(converter))))
			got, err := sut.Tag(tt.args.ctx, tt.args.id)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("Tag() got = %v, want %v", err, tt.want.err)
				return
			}
			if diff := cmp.Diff(got, tt.want.out, cmpOpts...); diff != "" {
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
		out dto.TagsOutDTO
		err error
	}
	type converterResult struct {
		out *model.TagConnection
		err error
	}
	type testCase struct {
		sut                func(resolver *Resolver) *queryResolver
		setupMockUsecase   func(uc *musecase.MockTags, usecaseResult usecaseResult)
		usecaseResult      usecaseResult
		setupMockConverter func(converter *mconverter.MockTagsConverter, from dto.TagsOutDTO, converterResult converterResult)
		converterResult    converterResult
		args               args
		want               want
		wantErr            bool
	}
	errFailedToUsecase := errors.New("failed to usecase")
	errFailedToConvert := errors.New("failed to convert")
	tests := map[string]testCase{
		"happy_path/without_paging": {
			sut: func(resolver *Resolver) *queryResolver {
				return &queryResolver{resolver}
			},
			setupMockUsecase: func(uc *musecase.MockTags, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewTagsOutDTO([]dto.TagArticle{
					dto.NewTagArticle(
						"Tag1",
						"Tag1",
						[]dto.Article{
							dto.NewArticle(
								"Article1",
								"Article1",
								"",
								utils.MustURLParse("example.com/example.png"),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
						}),
				}),
				err: nil,
			},
			setupMockConverter: func(converter *mconverter.MockTagsConverter, from dto.TagsOutDTO, converterResult converterResult) {
				converter.EXPECT().
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
												ThumbnailURL: gqlscalar.URL(utils.MustURLParse("example.com/example.png")),
												CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
												UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
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
												ThumbnailURL: gqlscalar.URL(utils.MustURLParse("example.com/example.png")),
												CreatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
												UpdatedAt:    gqlscalar.UTC(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0)),
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
			sut: func(resolver *Resolver) *queryResolver {
				return &queryResolver{resolver}
			},
			setupMockUsecase: func(uc *musecase.MockTags, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				err: errFailedToUsecase,
			},
			setupMockConverter: func(converter *mconverter.MockTagsConverter, from dto.TagsOutDTO, converterResult converterResult) {
				converter.EXPECT().
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
			sut: func(resolver *Resolver) *queryResolver {
				return &queryResolver{resolver}
			},
			setupMockUsecase: func(uc *musecase.MockTags, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewTagsOutDTO([]dto.TagArticle{}),
				err: nil,
			},
			setupMockConverter: func(converter *mconverter.MockTagsConverter, from dto.TagsOutDTO, converterResult converterResult) {
				converter.EXPECT().
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
			sut: func(resolver *Resolver) *queryResolver {
				return &queryResolver{resolver}
			},
			setupMockUsecase: func(uc *musecase.MockTags, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			setupMockConverter: func(converter *mconverter.MockTagsConverter, from dto.TagsOutDTO, converterResult converterResult) {
				converter.EXPECT().
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
			sut: func(resolver *Resolver) *queryResolver {
				return &queryResolver{resolver}
			},
			setupMockUsecase: func(uc *musecase.MockTags, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			setupMockConverter: func(converter *mconverter.MockTagsConverter, from dto.TagsOutDTO, converterResult converterResult) {
				converter.EXPECT().
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
			sut: func(resolver *Resolver) *queryResolver {
				return &queryResolver{resolver}
			},
			setupMockUsecase: func(uc *musecase.MockTags, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			setupMockConverter: func(converter *mconverter.MockTagsConverter, from dto.TagsOutDTO, converterResult converterResult) {
				converter.EXPECT().
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
			sut: func(resolver *Resolver) *queryResolver {
				return &queryResolver{resolver}
			},
			setupMockUsecase: func(uc *musecase.MockTags, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			setupMockConverter: func(converter *mconverter.MockTagsConverter, from dto.TagsOutDTO, converterResult converterResult) {
				converter.EXPECT().
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
				err: dto.ErrInvalidateTagsInDTO,
			},
			setupMockUsecase: func(uc *musecase.MockTags, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Times(0)
			},
			setupMockConverter: func(converter *mconverter.MockTagsConverter, from dto.TagsOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToTags(gomock.Any(), gomock.Any()).
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
				before: func() *string { v := "Tag1"; return &v }(),
			},
			want: want{
				out: nil,
				err: dto.ErrInvalidateTagsInDTO,
			},
			setupMockUsecase: func(uc *musecase.MockTags, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Times(0)
			},
			setupMockConverter: func(converter *mconverter.MockTagsConverter, from dto.TagsOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToTags(gomock.Any(), gomock.Any()).
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
				after: func() *string { v := "Tag1"; return &v }(),
			},
			want: want{
				out: nil,
				err: dto.ErrInvalidateTagsInDTO,
			},
			setupMockUsecase: func(uc *musecase.MockTags, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Times(0)
			},
			setupMockConverter: func(converter *mconverter.MockTagsConverter, from dto.TagsOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToTags(gomock.Any(), gomock.Any()).
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
				before: func() *string { v := "Tag1"; return &v }(),
			},
			want: want{
				out: nil,
				err: dto.ErrInvalidateTagsInDTO,
			},
			setupMockUsecase: func(uc *musecase.MockTags, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Times(0)
			},
			setupMockConverter: func(converter *mconverter.MockTagsConverter, from dto.TagsOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToTags(gomock.Any(), gomock.Any()).
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
				after: func() *string { v := "Tag1"; return &v }(),
			},
			want: want{
				out: nil,
				err: dto.ErrInvalidateTagsInDTO,
			},
			setupMockUsecase: func(uc *musecase.MockTags, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), gomock.Any()).
					Times(0)
			},
			setupMockConverter: func(converter *mconverter.MockTagsConverter, from dto.TagsOutDTO, converterResult converterResult) {
				converter.EXPECT().
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
			uc := musecase.NewMockTags(ctrl)
			tt.setupMockUsecase(uc, tt.usecaseResult)
			converter := mconverter.NewMockTagsConverter(ctrl)
			tt.setupMockConverter(converter, tt.usecaseResult.out, tt.converterResult)
			sut := tt.sut(NewResolver(NewUsecases(WithTagsUsecase(uc)), NewConverters(WithTagsConverter(converter))))
			got, err := sut.Tags(tt.args.ctx, tt.args.first, tt.args.last, tt.args.after, tt.args.before)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("Tags() got = %v, want %v", err, tt.want.err)
				return
			}
			if diff := cmp.Diff(got, tt.want.out, cmpOpts...); diff != "" {
				t.Error(diff)
				return
			}
		})
	}
}
