package resolver

import (
	"context"
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/google/go-cmp/cmp"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/presenters/graphql/model"
	mconverter "github.com/miyamo2/blogapi.miyamo.today/federator/internal/mock/if-adapter/controller/graphql/resolver/presenter/converter"
	musecase "github.com/miyamo2/blogapi.miyamo.today/federator/internal/mock/if-adapter/controller/graphql/resolver/usecase"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/pkg/gqlscalar"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/utils"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_mutationResolver_CreateArticle(t *testing.T) {
	type args struct {
		ctx   context.Context
		input model.CreateArticleInput
	}
	type want struct {
		out *model.CreateArticlePayload
		err error
	}
	type usecaseResult struct {
		out dto.CreateArticleOutDTO
		err error
	}
	type converterResult struct {
		out *model.CreateArticlePayload
		err error
	}
	type testCase struct {
		sut                func(resolver *Resolver) *mutationResolver
		createArticleInDTO dto.CreateArticleInDTO
		setupMockUsecase   func(uc *musecase.MockCreateArticle, input dto.CreateArticleInDTO, usecaseResult usecaseResult)
		usecaseResult      usecaseResult
		setupMockConverter func(converter *mconverter.MockCreateArticleConverter, from dto.CreateArticleOutDTO, converterResult converterResult)
		converterResult    converterResult
		args               args
		want               want
		wantErr            bool
	}
	errFailedToUsecase := errors.New("failed to usecase")
	errFailedToConverter := errors.New("failed to converter")
	tests := map[string]testCase{
		"happy_path": {
			sut: func(resolver *Resolver) *mutationResolver {
				return &mutationResolver{resolver}
			},
			createArticleInDTO: dto.NewCreateArticleInDTO("Title1", "Content1", utils.MustURLParse("https://example.com/example.jpg"), []string{"Tag1", "Tag2"}, "Mutation1"),
			setupMockUsecase: func(uc *musecase.MockCreateArticle, input dto.CreateArticleInDTO, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), NewCreateArticleInputMatcher(input)).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewCreateArticleOutDTO(
					"Event1",
					"Article1",
					"Mutation1",
				),
				err: nil,
			},
			setupMockConverter: func(converter *mconverter.MockCreateArticleConverter, from dto.CreateArticleOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToCreateArticle(gomock.Any(), from).
					Return(converterResult.out, converterResult.err).
					Times(1)
			},
			converterResult: converterResult{
				out: &model.CreateArticlePayload{
					EventID:          "Event1",
					ArticleID:        "Article1",
					ClientMutationID: toPointerString("Mutation1"),
				},
			},
			args: args{
				ctx: context.Background(),
				input: model.CreateArticleInput{
					Title:            "Title1",
					Content:          "Content1",
					TagNames:         []string{"Tag1", "Tag2"},
					ThumbnailURL:     gqlscalar.URL(utils.MustURLParse("https://example.com/example.jpg")),
					ClientMutationID: toPointerString("Mutation1"),
				},
			},
			want: want{
				out: &model.CreateArticlePayload{
					EventID:          "Event1",
					ArticleID:        "Article1",
					ClientMutationID: toPointerString("Mutation1"),
				},
			},
		},
		"unhappy_path:usecase-returns-error": {
			sut: func(resolver *Resolver) *mutationResolver {
				return &mutationResolver{resolver}
			},
			createArticleInDTO: dto.NewCreateArticleInDTO("Title1", "Content1", utils.MustURLParse("https://example.com/example.jpg"), []string{"Tag1", "Tag2"}, "Mutation1"),
			setupMockUsecase: func(uc *musecase.MockCreateArticle, input dto.CreateArticleInDTO, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), NewCreateArticleInputMatcher(input)).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				err: errFailedToUsecase,
			},
			setupMockConverter: func(converter *mconverter.MockCreateArticleConverter, from dto.CreateArticleOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToCreateArticle(gomock.Any(), gomock.Any()).
					Times(0)
			},
			args: args{
				ctx: context.Background(),
				input: model.CreateArticleInput{
					Title:            "Title1",
					Content:          "Content1",
					TagNames:         []string{"Tag1", "Tag2"},
					ThumbnailURL:     gqlscalar.URL(utils.MustURLParse("https://example.com/example.jpg")),
					ClientMutationID: toPointerString("Mutation1"),
				},
			},
			want: want{
				err: errFailedToUsecase,
			},
		},
		"unhappy_path:converter-returns-error": {
			sut: func(resolver *Resolver) *mutationResolver {
				return &mutationResolver{resolver}
			},
			createArticleInDTO: dto.NewCreateArticleInDTO("Title1", "Content1", utils.MustURLParse("https://example.com/example.jpg"), []string{"Tag1", "Tag2"}, "Mutation1"),
			setupMockUsecase: func(uc *musecase.MockCreateArticle, input dto.CreateArticleInDTO, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), NewCreateArticleInputMatcher(input)).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewCreateArticleOutDTO(
					"Event1",
					"Article1",
					"Mutation1",
				),
				err: nil,
			},
			setupMockConverter: func(converter *mconverter.MockCreateArticleConverter, from dto.CreateArticleOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToCreateArticle(gomock.Any(), from).
					Return(converterResult.out, converterResult.err).
					Times(1)
			},
			converterResult: converterResult{
				err: errFailedToConverter,
			},
			args: args{
				ctx: context.Background(),
				input: model.CreateArticleInput{
					Title:            "Title1",
					Content:          "Content1",
					TagNames:         []string{"Tag1", "Tag2"},
					ThumbnailURL:     gqlscalar.URL(utils.MustURLParse("https://example.com/example.jpg")),
					ClientMutationID: toPointerString("Mutation1"),
				},
			},
			want: want{
				err: errFailedToConverter,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			uc := musecase.NewMockCreateArticle(ctrl)
			tt.setupMockUsecase(uc, tt.createArticleInDTO, tt.usecaseResult)

			converter := mconverter.NewMockCreateArticleConverter(ctrl)
			tt.setupMockConverter(converter, tt.usecaseResult.out, tt.converterResult)

			sut := tt.sut(NewResolver(NewUsecases(WithCreateArticleUsecase(uc)), NewConverters(WithCreateArticleConverter(converter))))
			got, err := sut.CreateArticle(tt.args.ctx, tt.args.input)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("CreateArticle() got = %v, want %v", err, tt.want.err)
				return
			}
			if diff := cmp.Diff(got, tt.want.out, cmpOpts...); diff != "" {
				t.Error(diff)
				return
			}
		})
	}
}

func toPointerString(s string) *string {
	return &s
}

func NewCreateArticleInputMatcher(expect dto.CreateArticleInDTO) gomock.Matcher {
	return &CreateArticleInDTOMatcher{
		expect: expect,
	}
}

type CreateArticleInDTOMatcher struct {
	gomock.Matcher
	expect dto.CreateArticleInDTO
}

func (m *CreateArticleInDTOMatcher) Matches(x interface{}) bool {
	switch x := x.(type) {
	case dto.CreateArticleInDTO:
		return m.expect.ClientMutationID() == x.ClientMutationID() &&
			m.expect.Body() == x.Body() &&
			cmp.Diff(x.TagNames(), m.expect.TagNames(), cmpOpts...) == "" &&
			cmp.Diff(x.ThumbnailURL(), m.expect.ThumbnailURL(), cmpOpts...) == ""
	}
	return false
}

func (m *CreateArticleInDTOMatcher) String() string {
	return fmt.Sprintf("is equal to %+v", m.expect)
}

func Test_mutationResolver_UpdateArticleTitle(t *testing.T) {
	type args struct {
		ctx   context.Context
		input model.UpdateArticleTitleInput
	}
	type want struct {
		out *model.UpdateArticleTitlePayload
		err error
	}
	type usecaseResult struct {
		out dto.UpdateArticleTitleOutDTO
		err error
	}
	type converterResult struct {
		out *model.UpdateArticleTitlePayload
		err error
	}
	type testCase struct {
		sut                func(resolver *Resolver) *mutationResolver
		updateArticleInDTO dto.UpdateArticleTitleInDTO
		setupMockUsecase   func(uc *musecase.MockUpdateArticleTitle, input dto.UpdateArticleTitleInDTO, usecaseResult usecaseResult)
		usecaseResult      usecaseResult
		setupMockConverter func(converter *mconverter.MockUpdateArticleTitleConverter, from dto.UpdateArticleTitleOutDTO, converterResult converterResult)
		converterResult    converterResult
		args               args
		want               want
		wantErr            bool
	}
	errFailedToUsecase := errors.New("failed to usecase")
	errFailedToConverter := errors.New("failed to converter")
	tests := map[string]testCase{
		"happy_path": {
			sut: func(resolver *Resolver) *mutationResolver {
				return &mutationResolver{resolver}
			},
			updateArticleInDTO: dto.NewUpdateArticleTitleInDTO("Article1", "Title1", "Mutation1"),
			setupMockUsecase: func(uc *musecase.MockUpdateArticleTitle, input dto.UpdateArticleTitleInDTO, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), NewUpdateArticleTitleInputMatcher(input)).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewUpdateArticleTitleOutDTO(
					"Event1",
					"Article1",
					"Mutation1",
				),
				err: nil,
			},
			setupMockConverter: func(converter *mconverter.MockUpdateArticleTitleConverter, from dto.UpdateArticleTitleOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToUpdateArticleTitle(gomock.Any(), from).
					Return(converterResult.out, converterResult.err).
					Times(1)
			},
			converterResult: converterResult{
				out: &model.UpdateArticleTitlePayload{
					EventID:          "Event1",
					ArticleID:        "Article1",
					ClientMutationID: toPointerString("Mutation1"),
				},
			},
			args: args{
				ctx: context.Background(),
				input: model.UpdateArticleTitleInput{
					ArticleID:        "Article1",
					Title:            "Title1",
					ClientMutationID: toPointerString("Mutation1"),
				},
			},
			want: want{
				out: &model.UpdateArticleTitlePayload{
					EventID:          "Event1",
					ArticleID:        "Article1",
					ClientMutationID: toPointerString("Mutation1"),
				},
			},
		},
		"unhappy_path:usecase-returns-error": {
			sut: func(resolver *Resolver) *mutationResolver {
				return &mutationResolver{resolver}
			},
			updateArticleInDTO: dto.NewUpdateArticleTitleInDTO("Article1", "Title1", "Mutation1"),
			setupMockUsecase: func(uc *musecase.MockUpdateArticleTitle, input dto.UpdateArticleTitleInDTO, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), NewUpdateArticleTitleInputMatcher(input)).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				err: errFailedToUsecase,
			},
			setupMockConverter: func(converter *mconverter.MockUpdateArticleTitleConverter, from dto.UpdateArticleTitleOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToUpdateArticleTitle(gomock.Any(), gomock.Any()).
					Times(0)
			},
			args: args{
				ctx: context.Background(),
				input: model.UpdateArticleTitleInput{
					ArticleID:        "Article1",
					Title:            "Title1",
					ClientMutationID: toPointerString("Mutation1"),
				},
			},
			want: want{
				err: errFailedToUsecase,
			},
		},
		"unhappy_path:converter-returns-error": {
			sut: func(resolver *Resolver) *mutationResolver {
				return &mutationResolver{resolver}
			},
			updateArticleInDTO: dto.NewUpdateArticleTitleInDTO("Article1", "Title1", "Mutation1"),
			setupMockUsecase: func(uc *musecase.MockUpdateArticleTitle, input dto.UpdateArticleTitleInDTO, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), NewUpdateArticleTitleInputMatcher(input)).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewUpdateArticleTitleOutDTO(
					"Event1",
					"Article1",
					"Mutation1",
				),
			},
			setupMockConverter: func(converter *mconverter.MockUpdateArticleTitleConverter, from dto.UpdateArticleTitleOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToUpdateArticleTitle(gomock.Any(), from).
					Return(converterResult.out, converterResult.err).
					Times(1)
			},
			converterResult: converterResult{
				err: errFailedToConverter,
			},
			args: args{
				ctx: context.Background(),
				input: model.UpdateArticleTitleInput{
					ArticleID:        "Article1",
					Title:            "Title1",
					ClientMutationID: toPointerString("Mutation1"),
				},
			},
			want: want{
				err: errFailedToConverter,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			uc := musecase.NewMockUpdateArticleTitle(ctrl)
			tt.setupMockUsecase(uc, tt.updateArticleInDTO, tt.usecaseResult)

			converter := mconverter.NewMockUpdateArticleTitleConverter(ctrl)
			tt.setupMockConverter(converter, tt.usecaseResult.out, tt.converterResult)

			sut := tt.sut(NewResolver(NewUsecases(WithUpdateArticleTitleUsecase(uc)), NewConverters(WithUpdateArticleTitleConverter(converter))))
			got, err := sut.UpdateArticleTitle(tt.args.ctx, tt.args.input)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("UpdateArticleTitle() got = %v, want %v", err, tt.want.err)
				return
			}
			if diff := cmp.Diff(got, tt.want.out, cmpOpts...); diff != "" {
				t.Error(diff)
				return
			}
		})
	}
}

func NewUpdateArticleTitleInputMatcher(expect dto.UpdateArticleTitleInDTO) gomock.Matcher {
	return &UpdateArticleTitleInDTOMatcher{
		expect: expect,
	}
}

type UpdateArticleTitleInDTOMatcher struct {
	gomock.Matcher
	expect dto.UpdateArticleTitleInDTO
}

func (m *UpdateArticleTitleInDTOMatcher) Matches(x interface{}) bool {
	switch x := x.(type) {
	case dto.UpdateArticleTitleInDTO:
		return cmp.Diff(x.ID(), m.expect.ID(), cmpOpts...) == "" &&
			cmp.Diff(x.Title(), m.expect.Title(), cmpOpts...) == "" &&
			cmp.Diff(x.ClientMutationID(), m.expect.ClientMutationID(), cmpOpts...) == ""
	}
	return false
}

func (m *UpdateArticleTitleInDTOMatcher) String() string {
	return fmt.Sprintf("is equal to %+v", m.expect)
}
