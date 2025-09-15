package pb

import (
	"testing"

	"blogapi.miyamo.today/article-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/article-service/internal/if-adapter/controller/pb/presenter/convert"
	"blogapi.miyamo.today/article-service/internal/if-adapter/controller/pb/usecase"
	"blogapi.miyamo.today/article-service/internal/infra/grpc"
	"connectrpc.com/connect"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/cockroachdb/errors"
	. "github.com/ovechkin-dm/mockio/v2/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ArticleServiceServerTestSuite struct {
	suite.Suite
}

func Test_ArticleServiceServerTestSuite(t *testing.T) {
	suite.Run(t, new(ArticleServiceServerTestSuite))
}

func (s *ArticleServiceServerTestSuite) TestArticleServiceServer_GetArticleById() {
	s.Run(
		"happy_path", func() {
			ctrl := NewMockController(s.T())
			uc := Mock[usecase.GetByID](ctrl)

			getByIdOutput := dto.NewArticle(
				"1",
				"happy_path/article_has_tag",
				"## happy_path/article_has_tag",
				"1234567890",
				synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
				synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
				dto.NewTag("1", "happy_path"),
			)

			WhenDouble(uc.Execute(AnyContext(), Equal(dto.NewGetByIDInput("1")))).
				ThenReturn(&getByIdOutput, nil)

			res := &grpc.GetArticleByIdResponse{
				Article: &grpc.Article{
					Id:           "1",
					Title:        "happy_path/article_has_tag",
					Body:         "## happy_path/article_has_tag",
					ThumbnailUrl: "1234567890",
					CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
					UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "happy_path",
						},
					},
				},
			}

			conv := Mock[convert.GetByID](ctrl)
			WhenDouble(conv.ToResponse(AnyContext(), Equal(&getByIdOutput))).
				ThenReturn(res, true)

			sut := NewArticleServiceServer(WithGetByID(uc, conv))
			got, err := sut.GetArticleById(
				s.T().Context(),
				connect.NewRequest(
					&grpc.GetArticleByIdRequest{
						Id: "1",
					},
				),
			)
			s.Require().NoError(err)
			s.Require().NotNil(got)
			s.Require().Equal(res, got.Msg)
		},
	)
	s.Run(
		"unhappy_path/usecase_returns_error", func() {
			errGetArticleById := errors.New("error get article by id")

			ctrl := NewMockController(s.T())
			uc := Mock[usecase.GetByID](ctrl)

			WhenDouble(uc.Execute(AnyContext(), Equal(dto.NewGetByIDInput("1")))).
				ThenReturn(nil, errGetArticleById)

			sut := NewArticleServiceServer(WithGetByID(uc, nil))
			got, err := sut.GetArticleById(
				s.T().Context(),
				connect.NewRequest(
					&grpc.GetArticleByIdRequest{
						Id: "1",
					},
				),
			)
			s.Require().Error(err)
			s.Require().ErrorIs(err, errGetArticleById)
			s.Require().Nil(got)
		},
	)
	s.Run(
		"unhappy_path/failed_to_convert", func() {
			ctrl := NewMockController(s.T())
			uc := Mock[usecase.GetByID](ctrl)

			getByIdOutput := dto.NewArticle(
				"1",
				"happy_path/article_has_tag",
				"## happy_path/article_has_tag",
				"1234567890",
				synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
				synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
				dto.NewTag("1", "happy_path"),
			)

			WhenDouble(uc.Execute(AnyContext(), Equal(dto.NewGetByIDInput("1")))).
				ThenReturn(&getByIdOutput, nil)

			conv := Mock[convert.GetByID](ctrl)
			WhenDouble(conv.ToResponse(AnyContext(), Equal(&getByIdOutput))).
				ThenReturn(nil, false)

			sut := NewArticleServiceServer(WithGetByID(uc, conv))
			got, err := sut.GetArticleById(
				s.T().Context(),
				connect.NewRequest(
					&grpc.GetArticleByIdRequest{
						Id: "1",
					},
				),
			)
			s.Require().Error(err)
			s.Require().ErrorIs(err, ErrConversionToGetByIDFailed)
			s.Require().Nil(got)
		},
	)
}

func (s *ArticleServiceServerTestSuite) TestArticleServiceServer_GetAllArticles() {
	s.Run(
		"happy_path", func() {
			ctrl := NewMockController(s.T())
			uc := Mock[usecase.ListAll](ctrl)

			listAllOutput := dto.NewListAllOutput(
				dto.NewArticle(
					"1",
					"happy_path1",
					"## happy_path1",
					"1234567890",
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					dto.NewTag("tag1", "1"),
					dto.NewTag("tag2", "2"),
				),
			)
			WhenDouble(uc.Execute(AnyContext())).
				ThenReturn(&listAllOutput, nil)

			res := &grpc.GetAllArticlesResponse{
				Articles: []*grpc.Article{
					{
						Id:           "1",
						Title:        "happy_path1",
						Body:         "## happy_path1",
						ThumbnailUrl: "1234567890",
						CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
						UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
						Tags: []*grpc.Tag{
							{
								Id:   "tag1",
								Name: "1",
							},
							{
								Id:   "tag2",
								Name: "2",
							},
						},
					},
				},
			}

			conv := Mock[convert.ListAll](ctrl)
			WhenDouble(conv.ToResponse(AnyContext(), Equal(&listAllOutput))).
				ThenReturn(res, true)

			sut := NewArticleServiceServer(WithListAll(uc, conv))
			got, err := sut.GetAllArticles(
				s.T().Context(),
				connect.NewRequest(&emptypb.Empty{}),
			)
			s.Require().NoError(err)
			s.Require().NotNil(got)
			s.Require().Equal(res, got.Msg)
		},
	)
	s.Run(
		"unhappy_path/usecase_returns_error", func() {
			errGetAllArticle := errors.New("error get all Articles")

			ctrl := NewMockController(s.T())
			uc := Mock[usecase.ListAll](ctrl)

			WhenDouble(uc.Execute(AnyContext())).
				ThenReturn(nil, errGetAllArticle)

			sut := NewArticleServiceServer(WithListAll(uc, nil))
			got, err := sut.GetAllArticles(
				s.T().Context(),
				connect.NewRequest(&emptypb.Empty{}),
			)
			s.Require().Error(err)
			s.Require().Nil(got)
			s.Require().ErrorIs(err, errGetAllArticle)
		},
	)
	s.Run(
		"unhappy_path/failed_to_convert", func() {
			ctrl := NewMockController(s.T())
			uc := Mock[usecase.ListAll](ctrl)

			listAllOutput := dto.NewListAllOutput(
				dto.NewArticle(
					"1",
					"happy_path1",
					"## happy_path1",
					"1234567890",
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					dto.NewTag("tag1", "1"),
					dto.NewTag("tag2", "2"),
				),
			)
			WhenDouble(uc.Execute(AnyContext())).
				ThenReturn(&listAllOutput, nil)

			conv := Mock[convert.ListAll](ctrl)
			WhenDouble(conv.ToResponse(AnyContext(), Equal(&listAllOutput))).
				ThenReturn(nil, false)

			sut := NewArticleServiceServer(WithListAll(uc, conv))
			got, err := sut.GetAllArticles(
				s.T().Context(),
				connect.NewRequest(&emptypb.Empty{}),
			)
			s.Require().Error(err)
			s.Require().Nil(got)
			s.Require().ErrorIs(err, ErrConversionToListAllFailed)
		},
	)
}

func (s *ArticleServiceServerTestSuite) TestArticleServiceServer_GetNextArticles() {
	s.Run(
		"happy_path/without-cursor", func() {
			ctrl := NewMockController(s.T())
			uc := Mock[usecase.ListAfter](ctrl)

			listAfterOutput := dto.NewListAfterOutput(
				true,
				dto.NewArticle(
					"1",
					"happy_path1",
					"## happy_path1",
					"1234567890",
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					dto.NewTag("tag1", "1"),
					dto.NewTag("tag2", "2"),
				),
			)
			WhenDouble(
				uc.Execute(
					AnyContext(),
					Equal(dto.NewListAfterInput(1)),
				),
			).
				ThenReturn(&listAfterOutput, nil)

			res := &grpc.GetNextArticlesResponse{
				StillExists: true,
				Articles: []*grpc.Article{
					{
						Id:           "1",
						Title:        "happy_path1",
						Body:         "## happy_path1",
						ThumbnailUrl: "1234567890",
						CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
						UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
						Tags: []*grpc.Tag{
							{
								Id:   "tag1",
								Name: "1",
							},
							{
								Id:   "tag2",
								Name: "2",
							},
						},
					},
				},
			}

			conv := Mock[convert.ListAfter](ctrl)
			WhenDouble(conv.ToResponse(AnyContext(), Equal(&listAfterOutput))).
				ThenReturn(res, true)

			sut := NewArticleServiceServer(WithListAfter(uc, conv))
			got, err := sut.GetNextArticles(
				s.T().Context(),
				connect.NewRequest(
					&grpc.GetNextArticlesRequest{
						First: 1,
					},
				),
			)
			s.Require().NoError(err)
			s.Require().NotNil(got)
			s.Require().Equal(res, got.Msg)
		},
	)
	s.Run(
		"happy_path/with-cursor", func() {
			ctrl := NewMockController(s.T())
			uc := Mock[usecase.ListAfter](ctrl)

			listAfterOutput := dto.NewListAfterOutput(
				true,
				dto.NewArticle(
					"1",
					"happy_path1",
					"## happy_path1",
					"1234567890",
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					dto.NewTag("tag1", "1"),
					dto.NewTag("tag2", "2"),
				),
			)
			WhenDouble(
				uc.Execute(
					AnyContext(),
					Equal(dto.NewListAfterInput(1, dto.ListAfterInputWithCursor("0"))),
				),
			).
				ThenReturn(&listAfterOutput, nil)

			res := &grpc.GetNextArticlesResponse{
				StillExists: true,
				Articles: []*grpc.Article{
					{
						Id:           "1",
						Title:        "happy_path1",
						Body:         "## happy_path1",
						ThumbnailUrl: "1234567890",
						CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
						UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
						Tags: []*grpc.Tag{
							{
								Id:   "tag1",
								Name: "1",
							},
							{
								Id:   "tag2",
								Name: "2",
							},
						},
					},
				},
			}

			conv := Mock[convert.ListAfter](ctrl)
			WhenDouble(conv.ToResponse(AnyContext(), Equal(&listAfterOutput))).
				ThenReturn(res, true)

			sut := NewArticleServiceServer(WithListAfter(uc, conv))
			got, err := sut.GetNextArticles(
				s.T().Context(),
				connect.NewRequest(
					&grpc.GetNextArticlesRequest{
						First: 1,
						After: func() *string {
							v := "0"
							return &v
						}(),
					},
				),
			)
			s.Require().NoError(err)
			s.Require().NotNil(got)
			s.Require().Equal(res, got.Msg)
		},
	)
	s.Run(
		"unhappy_path/usecase_returns_error", func() {
			errGetNextArticle := errors.New("error get next Articles")

			ctrl := NewMockController(s.T())
			uc := Mock[usecase.ListAfter](ctrl)

			WhenDouble(
				uc.Execute(
					AnyContext(),
					Equal(dto.NewListAfterInput(1)),
				),
			).
				ThenReturn(nil, errGetNextArticle)

			sut := NewArticleServiceServer(WithListAfter(uc, nil))
			got, err := sut.GetNextArticles(
				s.T().Context(),
				connect.NewRequest(
					&grpc.GetNextArticlesRequest{
						First: 1,
					},
				),
			)
			s.Require().Error(err)
			s.Require().Nil(got)
			s.Require().ErrorIs(err, errGetNextArticle)
		},
	)
	s.Run(
		"unhappy_path/failed_to_convert", func() {
			ctrl := NewMockController(s.T())
			uc := Mock[usecase.ListAfter](ctrl)

			listAfterOutput := dto.NewListAfterOutput(
				true,
				dto.NewArticle(
					"1",
					"happy_path1",
					"## happy_path1",
					"1234567890",
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					dto.NewTag("tag1", "1"),
					dto.NewTag("tag2", "2"),
				),
			)
			WhenDouble(
				uc.Execute(
					AnyContext(),
					Equal(dto.NewListAfterInput(1)),
				),
			).
				ThenReturn(&listAfterOutput, nil)

			conv := Mock[convert.ListAfter](ctrl)
			WhenDouble(conv.ToResponse(AnyContext(), Equal(&listAfterOutput))).
				ThenReturn(nil, false)

			sut := NewArticleServiceServer(WithListAfter(uc, conv))
			got, err := sut.GetNextArticles(
				s.T().Context(),
				connect.NewRequest(
					&grpc.GetNextArticlesRequest{
						First: 1,
					},
				),
			)
			s.Require().Error(err)
			s.Require().Nil(got)
			s.Require().ErrorIs(err, ErrConversionToListNextFailed)
		},
	)
}

func (s *ArticleServiceServerTestSuite) TestArticleServiceServer_GetPrevArticles() {
	s.Run(
		"happy_path/without-cursor", func() {
			ctrl := NewMockController(s.T())
			uc := Mock[usecase.ListBefore](ctrl)

			listBeforeOutput := dto.NewListBeforeOutput(
				true,
				dto.NewArticle(
					"1",
					"happy_path1",
					"## happy_path1",
					"1234567890",
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					dto.NewTag("tag1", "1"),
					dto.NewTag("tag2", "2"),
				),
			)
			WhenDouble(
				uc.Execute(
					AnyContext(),
					Equal(dto.NewListBeforeInput(1)),
				),
			).
				ThenReturn(&listBeforeOutput, nil)

			res := &grpc.GetPrevArticlesResponse{
				StillExists: true,
				Articles: []*grpc.Article{
					{
						Id:           "1",
						Title:        "happy_path1",
						Body:         "## happy_path1",
						ThumbnailUrl: "1234567890",
						CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
						UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
						Tags: []*grpc.Tag{
							{
								Id:   "tag1",
								Name: "1",
							},
							{
								Id:   "tag2",
								Name: "2",
							},
						},
					},
				},
			}

			conv := Mock[convert.ListBefore](ctrl)
			WhenDouble(conv.ToResponse(AnyContext(), Equal(&listBeforeOutput))).
				ThenReturn(res, true)

			sut := NewArticleServiceServer(WithListBefore(uc, conv))
			got, err := sut.GetPrevArticles(
				s.T().Context(),
				connect.NewRequest(
					&grpc.GetPrevArticlesRequest{
						Last: 1,
					},
				),
			)
			s.Require().NoError(err)
			s.Require().NotNil(got)
			s.Require().Equal(res, got.Msg)
		},
	)
	s.Run(
		"happy_path/with-cursor", func() {
			ctrl := NewMockController(s.T())
			uc := Mock[usecase.ListBefore](ctrl)

			listBeforeOutput := dto.NewListBeforeOutput(
				true,
				dto.NewArticle(
					"1",
					"happy_path1",
					"## happy_path1",
					"1234567890",
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					dto.NewTag("tag1", "1"),
					dto.NewTag("tag2", "2"),
				),
			)
			WhenDouble(
				uc.Execute(
					AnyContext(),
					Equal(dto.NewListBeforeInput(1, dto.ListBeforeInputWithCursor("2"))),
				),
			).
				ThenReturn(&listBeforeOutput, nil)

			res := &grpc.GetPrevArticlesResponse{
				StillExists: true,
				Articles: []*grpc.Article{
					{
						Id:           "1",
						Title:        "happy_path1",
						Body:         "## happy_path1",
						ThumbnailUrl: "1234567890",
						CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
						UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
						Tags: []*grpc.Tag{
							{
								Id:   "tag1",
								Name: "1",
							},
							{
								Id:   "tag2",
								Name: "2",
							},
						},
					},
				},
			}

			conv := Mock[convert.ListBefore](ctrl)
			WhenDouble(conv.ToResponse(AnyContext(), Equal(&listBeforeOutput))).
				ThenReturn(res, true)

			sut := NewArticleServiceServer(WithListBefore(uc, conv))
			got, err := sut.GetPrevArticles(
				s.T().Context(),
				connect.NewRequest(
					&grpc.GetPrevArticlesRequest{
						Last: 1,
						Before: func() *string {
							v := "2"
							return &v
						}(),
					},
				),
			)
			s.Require().NoError(err)
			s.Require().NotNil(got)
			s.Require().Equal(res, got.Msg)
		},
	)
	s.Run(
		"unhappy_path/usecase_returns_error", func() {
			errGetPrevArticle := errors.New("error get prev Articles")

			ctrl := NewMockController(s.T())
			uc := Mock[usecase.ListBefore](ctrl)

			WhenDouble(
				uc.Execute(
					AnyContext(),
					Equal(dto.NewListBeforeInput(1)),
				),
			).
				ThenReturn(nil, errGetPrevArticle)

			sut := NewArticleServiceServer(WithListBefore(uc, nil))
			got, err := sut.GetPrevArticles(
				s.T().Context(),
				connect.NewRequest(
					&grpc.GetPrevArticlesRequest{
						Last: 1,
					},
				),
			)
			s.Require().Error(err)
			s.Require().Nil(got)
			s.Require().ErrorIs(err, errGetPrevArticle)
		},
	)
	s.Run(
		"unhappy_path/failed_to_convert", func() {
			ctrl := NewMockController(s.T())
			uc := Mock[usecase.ListBefore](ctrl)

			listBeforeOutput := dto.NewListBeforeOutput(
				true,
				dto.NewArticle(
					"1",
					"happy_path1",
					"## happy_path1",
					"1234567890",
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					dto.NewTag("tag1", "1"),
					dto.NewTag("tag2", "2"),
				),
			)
			WhenDouble(
				uc.Execute(
					AnyContext(),
					Equal(dto.NewListBeforeInput(1)),
				),
			).
				ThenReturn(&listBeforeOutput, nil)

			conv := Mock[convert.ListBefore](ctrl)
			WhenDouble(conv.ToResponse(AnyContext(), Equal(&listBeforeOutput))).
				ThenReturn(nil, false)

			sut := NewArticleServiceServer(WithListBefore(uc, conv))
			got, err := sut.GetPrevArticles(
				s.T().Context(),
				connect.NewRequest(
					&grpc.GetPrevArticlesRequest{
						Last: 1,
					},
				),
			)
			s.Require().Error(err)
			s.Require().Nil(got)
			s.Require().ErrorIs(err, ErrConversionToListPrevFailed)
		},
	)
}
