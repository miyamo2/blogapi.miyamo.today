package resolver

import (
	"blogapi.miyamo.today/federator/internal/app/usecase/dto"
	"blogapi.miyamo.today/federator/internal/if-adapter/presenters/graphql/model"
	mconverter "blogapi.miyamo.today/federator/internal/mock/if-adapter/controller/graphql/resolver/presenter/converter"
	musecase "blogapi.miyamo.today/federator/internal/mock/if-adapter/controller/graphql/resolver/usecase"
	"blogapi.miyamo.today/federator/internal/pkg/gqlscalar"
	"blogapi.miyamo.today/federator/internal/utils"
	"bytes"
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/cockroachdb/errors"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"
	"io"
	"reflect"
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

func Test_mutationResolver_UpdateArticleBody(t *testing.T) {
	type args struct {
		ctx   context.Context
		input model.UpdateArticleBodyInput
	}
	type want struct {
		out *model.UpdateArticleBodyPayload
		err error
	}
	type usecaseResult struct {
		out dto.UpdateArticleBodyOutDTO
		err error
	}
	type converterResult struct {
		out *model.UpdateArticleBodyPayload
		err error
	}
	type testCase struct {
		sut                func(resolver *Resolver) *mutationResolver
		updateArticleInDTO dto.UpdateArticleBodyInDTO
		setupMockUsecase   func(uc *musecase.MockUpdateArticleBody, input dto.UpdateArticleBodyInDTO, usecaseResult usecaseResult)
		usecaseResult      usecaseResult
		setupMockConverter func(converter *mconverter.MockUpdateArticleBodyConverter, from dto.UpdateArticleBodyOutDTO, converterResult converterResult)
		converterResult    converterResult
		args               args
		want               want
	}
	errFailedToUsecase := errors.New("failed to usecase")
	errFailedToConverter := errors.New("failed to converter")
	tests := map[string]testCase{
		"happy_path": {
			sut: func(resolver *Resolver) *mutationResolver {
				return &mutationResolver{resolver}
			},
			updateArticleInDTO: dto.NewUpdateArticleBodyInDTO("Article1", "Content1", "Mutation1"),
			setupMockUsecase: func(uc *musecase.MockUpdateArticleBody, input dto.UpdateArticleBodyInDTO, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), NewUpdateArticleBodyInputMatcher(input)).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewUpdateArticleBodyOutDTO(
					"Event1",
					"Article1",
					"Mutation1",
				),
			},
			setupMockConverter: func(converter *mconverter.MockUpdateArticleBodyConverter, from dto.UpdateArticleBodyOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToUpdateArticleBody(gomock.Any(), from).
					Return(converterResult.out, converterResult.err).
					Times(1)
			},
			converterResult: converterResult{
				out: &model.UpdateArticleBodyPayload{
					EventID:          "Event1",
					ArticleID:        "Article1",
					ClientMutationID: toPointerString("Mutation1"),
				},
			},
			args: args{
				ctx: context.Background(),
				input: model.UpdateArticleBodyInput{
					ArticleID:        "Article1",
					Content:          "Content1",
					ClientMutationID: toPointerString("Mutation1"),
				},
			},
			want: want{
				out: &model.UpdateArticleBodyPayload{
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
			updateArticleInDTO: dto.NewUpdateArticleBodyInDTO("Article1", "Content1", "Mutation1"),
			setupMockUsecase: func(uc *musecase.MockUpdateArticleBody, input dto.UpdateArticleBodyInDTO, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), NewUpdateArticleBodyInputMatcher(input)).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				err: errFailedToUsecase,
			},
			setupMockConverter: func(converter *mconverter.MockUpdateArticleBodyConverter, from dto.UpdateArticleBodyOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToUpdateArticleBody(gomock.Any(), gomock.Any()).
					Times(0)
			},
			args: args{
				ctx: context.Background(),
				input: model.UpdateArticleBodyInput{
					ArticleID:        "Article1",
					Content:          "Content1",
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
			updateArticleInDTO: dto.NewUpdateArticleBodyInDTO("Article1", "Content1", "Mutation1"),
			setupMockUsecase: func(uc *musecase.MockUpdateArticleBody, input dto.UpdateArticleBodyInDTO, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), NewUpdateArticleBodyInputMatcher(input)).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewUpdateArticleBodyOutDTO(
					"Event1",
					"Article1",
					"Mutation1",
				),
			},
			setupMockConverter: func(converter *mconverter.MockUpdateArticleBodyConverter, from dto.UpdateArticleBodyOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToUpdateArticleBody(gomock.Any(), from).
					Return(converterResult.out, converterResult.err).
					Times(1)
			},
			converterResult: converterResult{
				err: errFailedToConverter,
			},
			args: args{
				ctx: context.Background(),
				input: model.UpdateArticleBodyInput{
					ArticleID:        "Article1",
					Content:          "Content1",
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

			uc := musecase.NewMockUpdateArticleBody(ctrl)
			tt.setupMockUsecase(uc, tt.updateArticleInDTO, tt.usecaseResult)

			converter := mconverter.NewMockUpdateArticleBodyConverter(ctrl)
			tt.setupMockConverter(converter, tt.usecaseResult.out, tt.converterResult)

			sut := tt.sut(NewResolver(NewUsecases(WithUpdateArticleBodyUsecase(uc)), NewConverters(WithUpdateArticleBodyConverter(converter))))
			got, err := sut.UpdateArticleBody(tt.args.ctx, tt.args.input)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("UpdateArticleBody() got = %v, want %v", err, tt.want.err)
				return
			}
			if diff := cmp.Diff(got, tt.want.out, cmpOpts...); diff != "" {
				t.Error(diff)
				return
			}
		})
	}
}

type UpdateArticleBodyInputMatcher struct {
	gomock.Matcher
	expect dto.UpdateArticleBodyInDTO
}

func NewUpdateArticleBodyInputMatcher(expect dto.UpdateArticleBodyInDTO) gomock.Matcher {
	return &UpdateArticleBodyInputMatcher{
		expect: expect,
	}
}

func (m *UpdateArticleBodyInputMatcher) Matches(x interface{}) bool {
	switch x := x.(type) {
	case dto.UpdateArticleBodyInDTO:
		return cmp.Diff(x.ID(), m.expect.ID(), cmpOpts...) == "" &&
			cmp.Diff(x.Content(), m.expect.Content(), cmpOpts...) == "" &&
			cmp.Diff(x.ClientMutationID(), m.expect.ClientMutationID(), cmpOpts...) == ""
	}
	return false
}

func (m *UpdateArticleBodyInputMatcher) String() string {
	return fmt.Sprintf("is equal to %+v", m.expect)
}

func Test_mutationResolver_UpdateArticleThumbnail(t *testing.T) {
	type args struct {
		ctx   context.Context
		input model.UpdateArticleThumbnailInput
	}
	type want struct {
		out *model.UpdateArticleThumbnailPayload
		err error
	}
	type usecaseResult struct {
		out dto.UpdateArticleThumbnailOutDTO
		err error
	}
	type converterResult struct {
		out *model.UpdateArticleThumbnailPayload
		err error
	}
	type testCase struct {
		sut                func(resolver *Resolver) *mutationResolver
		updateArticleInDTO dto.UpdateArticleThumbnailInDTO
		setupMockUsecase   func(uc *musecase.MockUpdateArticleThumbnail, input dto.UpdateArticleThumbnailInDTO, usecaseResult usecaseResult)
		usecaseResult      usecaseResult
		setupMockConverter func(converter *mconverter.MockUpdateArticleThumbnailConverter, from dto.UpdateArticleThumbnailOutDTO, converterResult converterResult)
		converterResult    converterResult
		args               args
		want               want
	}
	errFailedToUsecase := errors.New("failed to usecase")
	errFailedToConverter := errors.New("failed to converter")
	tests := map[string]testCase{
		"happy_path": {
			sut: func(resolver *Resolver) *mutationResolver {
				return &mutationResolver{resolver}
			},
			updateArticleInDTO: dto.NewUpdateArticleThumbnailInDTO("Article1", utils.MustURLParse("https://example.com/example.png"), "Mutation1"),
			setupMockUsecase: func(uc *musecase.MockUpdateArticleThumbnail, input dto.UpdateArticleThumbnailInDTO, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), NewUpdateArticleThumbnailInputMatcher(input)).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewUpdateArticleThumbnailOutDTO("Event1", "Article1", "Mutation1"),
			},
			setupMockConverter: func(converter *mconverter.MockUpdateArticleThumbnailConverter, from dto.UpdateArticleThumbnailOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToUpdateArticleThumbnail(gomock.Any(), from).
					Return(converterResult.out, converterResult.err).
					Times(1)
			},
			converterResult: converterResult{
				out: &model.UpdateArticleThumbnailPayload{
					EventID:          "Event1",
					ArticleID:        "Article1",
					ClientMutationID: toPointerString("Mutation1"),
				},
			},
			args: args{
				ctx: context.Background(),
				input: model.UpdateArticleThumbnailInput{
					ArticleID:        "Article1",
					ThumbnailURL:     gqlscalar.URL(utils.MustURLParse("https://example.com/example.png")),
					ClientMutationID: toPointerString("Mutation1"),
				},
			},
			want: want{
				out: &model.UpdateArticleThumbnailPayload{
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
			updateArticleInDTO: dto.NewUpdateArticleThumbnailInDTO("Article1", utils.MustURLParse("https://example.com/example.png"), "Mutation1"),
			setupMockUsecase: func(uc *musecase.MockUpdateArticleThumbnail, input dto.UpdateArticleThumbnailInDTO, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), NewUpdateArticleThumbnailInputMatcher(input)).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				err: errFailedToUsecase,
			},
			setupMockConverter: func(converter *mconverter.MockUpdateArticleThumbnailConverter, from dto.UpdateArticleThumbnailOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToUpdateArticleThumbnail(gomock.Any(), gomock.Any()).
					Times(0)
			},
			args: args{
				ctx: context.Background(),
				input: model.UpdateArticleThumbnailInput{
					ArticleID:        "Article1",
					ThumbnailURL:     gqlscalar.URL(utils.MustURLParse("https://example.com/example.png")),
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
			updateArticleInDTO: dto.NewUpdateArticleThumbnailInDTO("Article1", utils.MustURLParse("https://example.com/example.png"), "Mutation1"),
			setupMockUsecase: func(uc *musecase.MockUpdateArticleThumbnail, input dto.UpdateArticleThumbnailInDTO, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), NewUpdateArticleThumbnailInputMatcher(input)).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.UpdateArticleThumbnailOutDTO{},
				err: nil,
			},
			setupMockConverter: func(converter *mconverter.MockUpdateArticleThumbnailConverter, from dto.UpdateArticleThumbnailOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToUpdateArticleThumbnail(gomock.Any(), from).
					Return(converterResult.out, converterResult.err).
					Times(1)
			},
			converterResult: converterResult{
				err: errFailedToConverter,
			},
			args: args{
				ctx: context.Background(),
				input: model.UpdateArticleThumbnailInput{
					ArticleID:        "Article1",
					ThumbnailURL:     gqlscalar.URL(utils.MustURLParse("https://example.com/example.png")),
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

			uc := musecase.NewMockUpdateArticleThumbnail(ctrl)
			tt.setupMockUsecase(uc, tt.updateArticleInDTO, tt.usecaseResult)

			converter := mconverter.NewMockUpdateArticleThumbnailConverter(ctrl)
			tt.setupMockConverter(converter, tt.usecaseResult.out, tt.converterResult)

			sut := tt.sut(NewResolver(NewUsecases(WithUpdateArticleThumbnailUsecase(uc)), NewConverters(WithUpdateArticleThumbnailConverter(converter))))
			got, err := sut.UpdateArticleThumbnail(tt.args.ctx, tt.args.input)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("UpdateArticleThumbnail() got = %v, want %v", err, tt.want.err)
				return
			}
			if diff := cmp.Diff(got, tt.want.out, cmpOpts...); diff != "" {
				t.Error(diff)
				return
			}
		})
	}
}

type UpdateArticleThumbnailInputMatcher struct {
	gomock.Matcher
	expect dto.UpdateArticleThumbnailInDTO
}

func NewUpdateArticleThumbnailInputMatcher(expect dto.UpdateArticleThumbnailInDTO) gomock.Matcher {
	return &UpdateArticleThumbnailInputMatcher{
		expect: expect,
	}
}

func (m *UpdateArticleThumbnailInputMatcher) Matches(x interface{}) bool {
	switch x := x.(type) {
	case dto.UpdateArticleThumbnailInDTO:
		return cmp.Diff(x.ID(), m.expect.ID(), cmpOpts...) == "" &&
			cmp.Diff(x.Thumbnail(), m.expect.Thumbnail(), cmpOpts...) == "" &&
			cmp.Diff(x.ClientMutationID(), m.expect.ClientMutationID(), cmpOpts...) == ""
	}
	return false
}

func (m *UpdateArticleThumbnailInputMatcher) String() string {
	return fmt.Sprintf("is equal to %+v", m.expect)
}

func Test_mutationResolver_AttachTags(t *testing.T) {
	type args struct {
		ctx   context.Context
		input model.AttachTagsInput
	}
	type want struct {
		out *model.AttachTagsPayload
		err error
	}
	type usecaseResult struct {
		out dto.AttachTagsOutDTO
		err error
	}
	type converterResult struct {
		out *model.AttachTagsPayload
		err error
	}
	type testCase struct {
		sut                func(resolver *Resolver) *mutationResolver
		updateArticleInDTO dto.AttachTagsInDTO
		setupMockUsecase   func(uc *musecase.MockAttachTags, input dto.AttachTagsInDTO, usecaseResult usecaseResult)
		usecaseResult      usecaseResult
		setupMockConverter func(converter *mconverter.MockAttachTagsConverter, from dto.AttachTagsOutDTO, converterResult converterResult)
		converterResult    converterResult
		args               args
		want               want
	}
	errFailedToUsecase := errors.New("failed to usecase")
	errFailedToConverter := errors.New("failed to converter")
	tests := map[string]testCase{
		"happy_path": {
			sut: func(resolver *Resolver) *mutationResolver {
				return &mutationResolver{resolver}
			},
			updateArticleInDTO: dto.NewAttachTagsInDTO("Article1", []string{"Tag1", "Tag2"}, "Mutation1"),
			setupMockUsecase: func(uc *musecase.MockAttachTags, input dto.AttachTagsInDTO, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), NewAttachTagsInputMatcher(input)).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewAttachTagsOutDTO("Event1", "Article1", "Mutation1"),
			},
			setupMockConverter: func(converter *mconverter.MockAttachTagsConverter, from dto.AttachTagsOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToAttachTags(gomock.Any(), from).
					Return(converterResult.out, converterResult.err).
					Times(1)
			},
			converterResult: converterResult{
				out: &model.AttachTagsPayload{
					EventID:          "Event1",
					ArticleID:        "Article1",
					ClientMutationID: toPointerString("Mutation1"),
				},
			},
			args: args{
				ctx: context.Background(),
				input: model.AttachTagsInput{
					ArticleID:        "Article1",
					TagNames:         []string{"Tag1", "Tag2"},
					ClientMutationID: toPointerString("Mutation1"),
				},
			},
			want: want{
				out: &model.AttachTagsPayload{
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
			updateArticleInDTO: dto.NewAttachTagsInDTO("Article1", []string{"Tag1", "Tag2"}, "Mutation1"),
			setupMockUsecase: func(uc *musecase.MockAttachTags, input dto.AttachTagsInDTO, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), NewAttachTagsInputMatcher(input)).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				err: errFailedToUsecase,
			},
			setupMockConverter: func(converter *mconverter.MockAttachTagsConverter, from dto.AttachTagsOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToAttachTags(gomock.Any(), gomock.Any()).
					Times(0)
			},
			args: args{
				ctx: context.Background(),
				input: model.AttachTagsInput{
					ArticleID:        "Article1",
					TagNames:         []string{"Tag1", "Tag2"},
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
			updateArticleInDTO: dto.NewAttachTagsInDTO("Article1", []string{"Tag1", "Tag2"}, "Mutation1"),
			setupMockUsecase: func(uc *musecase.MockAttachTags, input dto.AttachTagsInDTO, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), NewAttachTagsInputMatcher(input)).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.AttachTagsOutDTO{},
				err: nil,
			},
			setupMockConverter: func(converter *mconverter.MockAttachTagsConverter, from dto.AttachTagsOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToAttachTags(gomock.Any(), from).
					Return(converterResult.out, converterResult.err).
					Times(1)
			},
			converterResult: converterResult{
				err: errFailedToConverter,
			},
			args: args{
				ctx: context.Background(),
				input: model.AttachTagsInput{
					ArticleID:        "Article1",
					TagNames:         []string{"Tag1", "Tag2"},
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

			uc := musecase.NewMockAttachTags(ctrl)
			tt.setupMockUsecase(uc, tt.updateArticleInDTO, tt.usecaseResult)

			converter := mconverter.NewMockAttachTagsConverter(ctrl)
			tt.setupMockConverter(converter, tt.usecaseResult.out, tt.converterResult)

			sut := tt.sut(NewResolver(NewUsecases(WithAttachTagsUsecase(uc)), NewConverters(WithAttachTagsConverter(converter))))
			got, err := sut.AttachTags(tt.args.ctx, tt.args.input)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("AttachTags() got = %v, want %v", err, tt.want.err)
				return
			}
			if diff := cmp.Diff(got, tt.want.out, cmpOpts...); diff != "" {
				t.Error(diff)
				return
			}
		})
	}
}

type AttachTagsInputMatcher struct {
	gomock.Matcher
	expect dto.AttachTagsInDTO
}

func NewAttachTagsInputMatcher(expect dto.AttachTagsInDTO) gomock.Matcher {
	return &AttachTagsInputMatcher{
		expect: expect,
	}
}

func (m *AttachTagsInputMatcher) Matches(x interface{}) bool {
	switch x := x.(type) {
	case dto.AttachTagsInDTO:
		return cmp.Diff(x.ID(), m.expect.ID(), cmpOpts...) == "" &&
			cmp.Diff(x.TagNames(), m.expect.TagNames(), cmpOpts...) == "" &&
			cmp.Diff(x.ClientMutationID(), m.expect.ClientMutationID(), cmpOpts...) == ""
	}
	return false
}

func (m *AttachTagsInputMatcher) String() string {
	return fmt.Sprintf("is equal to %+v", m.expect)
}

func Test_mutationResolver_DetachTags(t *testing.T) {
	type args struct {
		ctx   context.Context
		input model.DetachTagsInput
	}
	type want struct {
		out *model.DetachTagsPayload
		err error
	}
	type usecaseResult struct {
		out dto.DetachTagsOutDTO
		err error
	}
	type converterResult struct {
		out *model.DetachTagsPayload
		err error
	}
	type testCase struct {
		sut                func(resolver *Resolver) *mutationResolver
		updateArticleInDTO dto.DetachTagsInDTO
		setupMockUsecase   func(uc *musecase.MockDetachTags, input dto.DetachTagsInDTO, usecaseResult usecaseResult)
		usecaseResult      usecaseResult
		setupMockConverter func(converter *mconverter.MockDetachTagsConverter, from dto.DetachTagsOutDTO, converterResult converterResult)
		converterResult    converterResult
		args               args
		want               want
	}
	errFailedToUsecase := errors.New("failed to usecase")
	errFailedToConverter := errors.New("failed to converter")
	tests := map[string]testCase{
		"happy_path": {
			sut: func(resolver *Resolver) *mutationResolver {
				return &mutationResolver{resolver}
			},
			updateArticleInDTO: dto.NewDetachTagsInDTO("Article1", []string{"Tag1", "Tag2"}, "Mutation1"),
			setupMockUsecase: func(uc *musecase.MockDetachTags, input dto.DetachTagsInDTO, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), NewDetachTagsInputMatcher(input)).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewDetachTagsOutDTO("Event1", "Article1", "Mutation1"),
			},
			setupMockConverter: func(converter *mconverter.MockDetachTagsConverter, from dto.DetachTagsOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToDetachTags(gomock.Any(), from).
					Return(converterResult.out, converterResult.err).
					Times(1)
			},
			converterResult: converterResult{
				out: &model.DetachTagsPayload{
					EventID:          "Event1",
					ArticleID:        "Article1",
					ClientMutationID: toPointerString("Mutation1"),
				},
			},
			args: args{
				ctx: context.Background(),
				input: model.DetachTagsInput{
					ArticleID:        "Article1",
					TagNames:         []string{"Tag1", "Tag2"},
					ClientMutationID: toPointerString("Mutation1"),
				},
			},
			want: want{
				out: &model.DetachTagsPayload{
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
			updateArticleInDTO: dto.NewDetachTagsInDTO("Article1", []string{"Tag1", "Tag2"}, "Mutation1"),
			setupMockUsecase: func(uc *musecase.MockDetachTags, input dto.DetachTagsInDTO, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), NewDetachTagsInputMatcher(input)).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				err: errFailedToUsecase,
			},
			setupMockConverter: func(converter *mconverter.MockDetachTagsConverter, from dto.DetachTagsOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToDetachTags(gomock.Any(), gomock.Any()).
					Times(0)
			},
			args: args{
				ctx: context.Background(),
				input: model.DetachTagsInput{
					ArticleID:        "Article1",
					TagNames:         []string{"Tag1", "Tag2"},
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
			updateArticleInDTO: dto.NewDetachTagsInDTO("Article1", []string{"Tag1", "Tag2"}, "Mutation1"),
			setupMockUsecase: func(uc *musecase.MockDetachTags, input dto.DetachTagsInDTO, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), NewDetachTagsInputMatcher(input)).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.DetachTagsOutDTO{},
				err: nil,
			},
			setupMockConverter: func(converter *mconverter.MockDetachTagsConverter, from dto.DetachTagsOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToDetachTags(gomock.Any(), from).
					Return(converterResult.out, converterResult.err).
					Times(1)
			},
			converterResult: converterResult{
				err: errFailedToConverter,
			},
			args: args{
				ctx: context.Background(),
				input: model.DetachTagsInput{
					ArticleID:        "Article1",
					TagNames:         []string{"Tag1", "Tag2"},
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

			uc := musecase.NewMockDetachTags(ctrl)
			tt.setupMockUsecase(uc, tt.updateArticleInDTO, tt.usecaseResult)

			converter := mconverter.NewMockDetachTagsConverter(ctrl)
			tt.setupMockConverter(converter, tt.usecaseResult.out, tt.converterResult)

			sut := tt.sut(NewResolver(NewUsecases(WithDetachTagsUsecase(uc)), NewConverters(WithDetachTagsConverter(converter))))
			got, err := sut.DetachTags(tt.args.ctx, tt.args.input)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("DetachTags() got = %v, want %v", err, tt.want.err)
				return
			}
			if diff := cmp.Diff(got, tt.want.out, cmpOpts...); diff != "" {
				t.Error(diff)
				return
			}
		})
	}
}

type DetachTagsInputMatcher struct {
	gomock.Matcher
	expect dto.DetachTagsInDTO
}

func NewDetachTagsInputMatcher(expect dto.DetachTagsInDTO) gomock.Matcher {
	return &DetachTagsInputMatcher{
		expect: expect,
	}
}

func (m *DetachTagsInputMatcher) Matches(x interface{}) bool {
	switch x := x.(type) {
	case dto.DetachTagsInDTO:
		return cmp.Diff(x.ID(), m.expect.ID(), cmpOpts...) == "" &&
			cmp.Diff(x.TagNames(), m.expect.TagNames(), cmpOpts...) == "" &&
			cmp.Diff(x.ClientMutationID(), m.expect.ClientMutationID(), cmpOpts...) == ""
	}
	return false
}

func (m *DetachTagsInputMatcher) String() string {
	return fmt.Sprintf("is equal to %+v", m.expect)
}

func Test_mutationResolver_UploadImage(t *testing.T) {
	type args struct {
		ctx   context.Context
		input model.UploadImageInput
	}
	type want struct {
		out *model.UploadImagePayload
		err error
	}
	type usecaseResult struct {
		out dto.UploadImageOutDTO
		err error
	}
	type converterResult struct {
		out *model.UploadImagePayload
		err error
	}
	type testCase struct {
		sut                func(resolver *Resolver) *mutationResolver
		updateArticleInDTO dto.UploadImageInDTO
		setupMockUsecase   func(uc *musecase.MockUploadImage, input dto.UploadImageInDTO, usecaseResult usecaseResult)
		usecaseResult      usecaseResult
		setupMockConverter func(converter *mconverter.MockUploadImageConverter, from dto.UploadImageOutDTO, converterResult converterResult)
		converterResult    converterResult
		args               args
		want               want
	}

	errFailedToUsecase := errors.New("failed to usecase")
	errFailedToConverter := errors.New("failed to converter")
	tests := map[string]testCase{
		"happy_path": {
			sut: func(resolver *Resolver) *mutationResolver {
				return &mutationResolver{resolver}
			},
			updateArticleInDTO: dto.NewUploadImageInDTO(bytes.NewReader([]byte("abc")), "example.png", "image/png", "Mutation1"),
			setupMockUsecase: func(uc *musecase.MockUploadImage, input dto.UploadImageInDTO, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), NewUploadImageInputMatcher(input)).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.NewUploadImageOutDTO(utils.MustURLParse("https://example.com/example.png"), "Mutation1"),
			},
			setupMockConverter: func(converter *mconverter.MockUploadImageConverter, from dto.UploadImageOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToUploadImage(gomock.Any(), from).
					Return(converterResult.out, converterResult.err).
					Times(1)
			},
			converterResult: converterResult{
				out: &model.UploadImagePayload{
					ImageURL:         gqlscalar.URL(utils.MustURLParse("https://example.com/example.png")),
					ClientMutationID: toPointerString("Mutation1"),
				},
			},
			args: args{
				ctx: context.Background(),
				input: model.UploadImageInput{
					Image: graphql.Upload{
						Filename: "example.png",
						File:     bytes.NewReader([]byte("abc")),
					},
					ClientMutationID: toPointerString("Mutation1"),
				},
			},
			want: want{
				out: &model.UploadImagePayload{
					ImageURL:         gqlscalar.URL(utils.MustURLParse("https://example.com/example.png")),
					ClientMutationID: toPointerString("Mutation1"),
				},
			},
		},
		"unhappy_path:usecase-returns-error": {
			sut: func(resolver *Resolver) *mutationResolver {
				return &mutationResolver{resolver}
			},
			updateArticleInDTO: dto.NewUploadImageInDTO(bytes.NewReader([]byte("abc")), "example.png", "image/png", "Mutation1"),
			setupMockUsecase: func(uc *musecase.MockUploadImage, input dto.UploadImageInDTO, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), NewUploadImageInputMatcher(input)).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				err: errFailedToUsecase,
			},
			setupMockConverter: func(converter *mconverter.MockUploadImageConverter, from dto.UploadImageOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToUploadImage(gomock.Any(), gomock.Any()).
					Times(0)
			},
			args: args{
				ctx: context.Background(),
				input: model.UploadImageInput{
					Image: graphql.Upload{
						Filename: "example.png",
						File:     bytes.NewReader([]byte("abc")),
					},
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
			updateArticleInDTO: dto.NewUploadImageInDTO(bytes.NewReader([]byte("abc")), "example.png", "image/png", "Mutation1"),
			setupMockUsecase: func(uc *musecase.MockUploadImage, input dto.UploadImageInDTO, usecaseResult usecaseResult) {
				uc.EXPECT().
					Execute(gomock.Any(), NewUploadImageInputMatcher(input)).
					Return(usecaseResult.out, usecaseResult.err).
					Times(1)
			},
			usecaseResult: usecaseResult{
				out: dto.UploadImageOutDTO{},
				err: nil,
			},
			setupMockConverter: func(converter *mconverter.MockUploadImageConverter, from dto.UploadImageOutDTO, converterResult converterResult) {
				converter.EXPECT().
					ToUploadImage(gomock.Any(), from).
					Return(converterResult.out, converterResult.err).
					Times(1)
			},
			converterResult: converterResult{
				err: errFailedToConverter,
			},
			args: args{
				ctx: context.Background(),
				input: model.UploadImageInput{
					Image: graphql.Upload{
						Filename: "example.png",
						File:     bytes.NewReader([]byte("abc")),
					},
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

			uc := musecase.NewMockUploadImage(ctrl)
			tt.setupMockUsecase(uc, tt.updateArticleInDTO, tt.usecaseResult)

			converter := mconverter.NewMockUploadImageConverter(ctrl)
			tt.setupMockConverter(converter, tt.usecaseResult.out, tt.converterResult)

			sut := tt.sut(NewResolver(NewUsecases(WithUploadImageUsecase(uc)), NewConverters(WithUploadImageConverter(converter))))
			got, err := sut.UploadImage(tt.args.ctx, tt.args.input)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("UploadImage() got = %v, want %v", err, tt.want.err)
				return
			}
			if diff := cmp.Diff(got, tt.want.out, cmpOpts...); diff != "" {
				t.Error(diff)
				return
			}
		})
	}
}

type UploadImageInputMatcher struct {
	gomock.Matcher
	expect dto.UploadImageInDTO
}

func NewUploadImageInputMatcher(expect dto.UploadImageInDTO) gomock.Matcher {
	return &UploadImageInputMatcher{
		expect: expect,
	}
}

func (m *UploadImageInputMatcher) Matches(x interface{}) bool {
	switch x := x.(type) {
	case dto.UploadImageInDTO:
		if x.Filename() != m.expect.Filename() {
			return false
		}
		if x.ClientMutationID() != m.expect.ClientMutationID() {
			return false
		}
		expectBody, err := io.ReadAll(m.expect.Data())
		if err != nil {
			return false
		}
		xBody, err := io.ReadAll(x.Data())
		if err != nil {
			return false
		}
		return reflect.DeepEqual(xBody, expectBody)
	}
	return false
}

func (m *UploadImageInputMatcher) String() string {
	return fmt.Sprintf("is equal to %+v", m.expect)
}
