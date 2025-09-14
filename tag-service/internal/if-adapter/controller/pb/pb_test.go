package pb

import (
	"testing"

	"blogapi.miyamo.today/tag-service/internal/if-adapter/controller/pb/presenter/convert"
	"blogapi.miyamo.today/tag-service/internal/if-adapter/controller/pb/usecase"
	"connectrpc.com/connect"
	"github.com/Code-Hex/synchro/tz"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/tag-service/internal/infra/grpc"
	"github.com/Code-Hex/synchro"
	"github.com/cockroachdb/errors"
	. "github.com/ovechkin-dm/mockio/v2/mock"
)

type TagServiceServerTestSuite struct {
	suite.Suite
}

func Test_TagServiceServerTestSuite(t *testing.T) {
	suite.Run(t, new(TagServiceServerTestSuite))
}

func (s *TagServiceServerTestSuite) TestTagServiceServer_GetTagById() {
	s.Run(
		"happy_path", func() {
			ctrl := NewMockController(s.T())
			uc := Mock[usecase.GetById](ctrl)

			getByIdOutput := dto.NewTag(
				"1",
				"tag1",
				dto.NewArticle(
					"1",
					"happy_path",
					"1234567890",
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
				),
			)

			WhenDouble(uc.Execute(AnyContext(), Equal(dto.NewGetByIdInput("1")))).
				ThenReturn(&getByIdOutput, nil)

			res := connect.NewResponse(
				&grpc.GetTagByIdResponse{
					Tag: &grpc.Tag{
						Id:   "1",
						Name: "happy_path",
						Articles: []*grpc.Article{
							{
								Id:           "1",
								Title:        "happy_path",
								ThumbnailUrl: "1234567890",
								CreatedAt: timestamppb.New(
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).
										StdTime(),
								),
								UpdatedAt: timestamppb.New(
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).
										StdTime(),
								),
							},
						},
					},
				},
			)

			conv := Mock[convert.ToGetById](ctrl)
			WhenDouble(conv.ToResponse(AnyContext(), Equal(&getByIdOutput))).
				ThenReturn(res, true)

			sut := NewTagServiceServer(WithGetById(uc, conv))
			got, err := sut.GetTagById(
				s.T().Context(),
				connect.NewRequest(
					&grpc.GetTagByIdRequest{
						Id: "1",
					},
				),
			)
			s.Require().NoError(err)
			s.Require().NotNil(got)
			s.Require().Equal(res.Msg, got.Msg)
		},
	)
	s.Run(
		"unhappy_path/usecase_returns_error", func() {
			errGetTagById := errors.New("error get article by id")

			ctrl := NewMockController(s.T())
			uc := Mock[usecase.GetById](ctrl)

			WhenDouble(uc.Execute(AnyContext(), Equal(dto.NewGetByIdInput("1")))).
				ThenReturn(nil, errGetTagById)

			sut := NewTagServiceServer(WithGetById(uc, nil))
			got, err := sut.GetTagById(
				s.T().Context(),
				connect.NewRequest(
					&grpc.GetTagByIdRequest{
						Id: "1",
					},
				),
			)
			s.Require().Error(err)
			s.Require().ErrorIs(err, errGetTagById)
			s.Require().Nil(got)
		},
	)
	s.Run(
		"unhappy_path/failed_to_convert", func() {
			ctrl := NewMockController(s.T())
			uc := Mock[usecase.GetById](ctrl)

			getByIdOutput := dto.NewTag(
				"1",
				"tag1",
				dto.NewArticle(
					"1",
					"unhappy_path/failed_to_convert",
					"1234567890",
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
				),
			)

			WhenDouble(uc.Execute(AnyContext(), Equal(dto.NewGetByIdInput("1")))).
				ThenReturn(&getByIdOutput, nil)

			conv := Mock[convert.ToGetById](ctrl)
			WhenDouble(conv.ToResponse(AnyContext(), Equal(&getByIdOutput))).
				ThenReturn(nil, false)

			sut := NewTagServiceServer(WithGetById(uc, conv))
			got, err := sut.GetTagById(
				s.T().Context(),
				connect.NewRequest(
					&grpc.GetTagByIdRequest{
						Id: "1",
					},
				),
			)
			s.Require().Error(err)
			s.Require().ErrorIs(err, ErrConversionToGetTagByIdFailed)
			s.Require().Nil(got)
		},
	)
}

func (s *TagServiceServerTestSuite) TestTagServiceServer_GetAllTags() {
	s.Run(
		"happy_path", func() {
			ctrl := NewMockController(s.T())
			uc := Mock[usecase.ListAll](ctrl)

			listAllOutput := dto.NewListAllOutput(
				dto.NewTag(
					"1",
					"tag1",
					dto.NewArticle(
						"1",
						"happy_path",
						"1234567890",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					),
				),
			)
			WhenDouble(uc.Execute(AnyContext())).
				ThenReturn(&listAllOutput, nil)

			res := connect.NewResponse(
				&grpc.GetAllTagsResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "tag1",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path",
									ThumbnailUrl: "1234567890",
									CreatedAt: timestamppb.New(
										synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).
											StdTime(),
									),
									UpdatedAt: timestamppb.New(
										synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).
											StdTime(),
									),
								},
							},
						},
					},
				},
			)

			conv := Mock[convert.ToGetAll](ctrl)
			WhenDouble(conv.ToResponse(AnyContext(), Equal(&listAllOutput))).
				ThenReturn(res, true)

			sut := NewTagServiceServer(WithListAll(uc, conv))
			got, err := sut.GetAllTags(
				s.T().Context(),
				connect.NewRequest(&emptypb.Empty{}),
			)
			s.Require().NoError(err)
			s.Require().NotNil(got)
			s.Require().Equal(res.Msg, got.Msg)
		},
	)
	s.Run(
		"unhappy_path/usecase_returns_error", func() {
			errGetAllTag := errors.New("error get all tags")

			ctrl := NewMockController(s.T())
			uc := Mock[usecase.ListAll](ctrl)

			WhenDouble(uc.Execute(AnyContext())).
				ThenReturn(nil, errGetAllTag)

			sut := NewTagServiceServer(WithListAll(uc, nil))
			got, err := sut.GetAllTags(
				s.T().Context(),
				connect.NewRequest(&emptypb.Empty{}),
			)
			s.Require().Error(err)
			s.Require().Nil(got)
			s.Require().ErrorIs(err, errGetAllTag)
		},
	)
	s.Run(
		"unhappy_path/failed_to_convert", func() {
			ctrl := NewMockController(s.T())
			uc := Mock[usecase.ListAll](ctrl)

			listAllOutput := dto.NewListAllOutput(
				dto.NewTag(
					"1",
					"tag1",
					dto.NewArticle(
						"1",
						"unhappy_path/failed_to_convert",
						"1234567890",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					),
				),
			)
			WhenDouble(uc.Execute(AnyContext())).
				ThenReturn(&listAllOutput, nil)

			conv := Mock[convert.ToGetAll](ctrl)
			WhenDouble(conv.ToResponse(AnyContext(), Equal(&listAllOutput))).
				ThenReturn(nil, false)

			sut := NewTagServiceServer(WithListAll(uc, conv))
			got, err := sut.GetAllTags(
				s.T().Context(),
				connect.NewRequest(&emptypb.Empty{}),
			)
			s.Require().Error(err)
			s.Require().Nil(got)
			s.Require().ErrorIs(err, ErrConversionToGetAllTagsFailed)
		},
	)
}

func (s *TagServiceServerTestSuite) TestTagServiceServer_GetNextTags() {
	s.Run(
		"happy_path/without-cursor", func() {
			ctrl := NewMockController(s.T())
			uc := Mock[usecase.ListAfter](ctrl)

			listAfterOutput := dto.NewListAfterOutput(
				true,
				dto.NewTag(
					"1",
					"tag1",
					dto.NewArticle(
						"1",
						"happy_path/without-cursor",
						"1234567890",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					),
				),
			)
			WhenDouble(
				uc.Execute(
					AnyContext(),
					Equal(dto.NewListAfterInput(1)),
				),
			).
				ThenReturn(&listAfterOutput, nil)

			res := connect.NewResponse(
				&grpc.GetNextTagResponse{
					StillExists: true,
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "tag1",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/without-cursor",
									ThumbnailUrl: "1234567890",
									CreatedAt: timestamppb.New(
										synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).
											StdTime(),
									),
									UpdatedAt: timestamppb.New(
										synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).
											StdTime(),
									),
								},
							},
						},
					},
				},
			)

			conv := Mock[convert.ToGetNext](ctrl)
			WhenDouble(conv.ToResponse(AnyContext(), Equal(&listAfterOutput))).
				ThenReturn(res, true)

			sut := NewTagServiceServer(WithListAfter(uc, conv))
			got, err := sut.GetNextTags(
				s.T().Context(),
				connect.NewRequest(
					&grpc.GetNextTagsRequest{
						First: 1,
					},
				),
			)
			s.Require().NoError(err)
			s.Require().NotNil(got)
			s.Require().Equal(res.Msg, got.Msg)
		},
	)
	s.Run(
		"happy_path/with-cursor", func() {
			ctrl := NewMockController(s.T())
			uc := Mock[usecase.ListAfter](ctrl)

			listAfterOutput := dto.NewListAfterOutput(
				true,
				dto.NewTag(
					"1",
					"tag1",
					dto.NewArticle(
						"1",
						"happy_path/with-cursor",
						"1234567890",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					),
				),
			)
			WhenDouble(
				uc.Execute(
					AnyContext(),
					Equal(dto.NewListAfterInput(1, dto.ListAfterInputWithCursor("0"))),
				),
			).
				ThenReturn(&listAfterOutput, nil)

			res := connect.NewResponse(
				&grpc.GetNextTagResponse{
					StillExists: true,
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "tag1",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/with-cursor",
									ThumbnailUrl: "1234567890",
									CreatedAt: timestamppb.New(
										synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).
											StdTime(),
									),
									UpdatedAt: timestamppb.New(
										synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).
											StdTime(),
									),
								},
							},
						},
					},
				},
			)

			conv := Mock[convert.ToGetNext](ctrl)
			WhenDouble(conv.ToResponse(AnyContext(), Equal(&listAfterOutput))).
				ThenReturn(res, true)

			sut := NewTagServiceServer(WithListAfter(uc, conv))
			got, err := sut.GetNextTags(
				s.T().Context(),
				connect.NewRequest(
					&grpc.GetNextTagsRequest{
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
			s.Require().Equal(res.Msg, got.Msg)
		},
	)
	s.Run(
		"unhappy_path/usecase_returns_error", func() {
			errGetNextTag := errors.New("error get next tags")

			ctrl := NewMockController(s.T())
			uc := Mock[usecase.ListAfter](ctrl)

			WhenDouble(
				uc.Execute(
					AnyContext(),
					Equal(dto.NewListAfterInput(1)),
				),
			).
				ThenReturn(nil, errGetNextTag)

			sut := NewTagServiceServer(WithListAfter(uc, nil))
			got, err := sut.GetNextTags(
				s.T().Context(),
				connect.NewRequest(
					&grpc.GetNextTagsRequest{
						First: 1,
					},
				),
			)
			s.Require().Error(err)
			s.Require().Nil(got)
			s.Require().ErrorIs(err, errGetNextTag)
		},
	)
	s.Run(
		"unhappy_path/failed_to_convert", func() {
			ctrl := NewMockController(s.T())
			uc := Mock[usecase.ListAfter](ctrl)

			listAfterOutput := dto.NewListAfterOutput(
				true,
				dto.NewTag(
					"1",
					"tag1",
					dto.NewArticle(
						"1",
						"unhappy_path/failed_to_convert",
						"1234567890",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					),
				),
			)
			WhenDouble(
				uc.Execute(
					AnyContext(),
					Equal(dto.NewListAfterInput(1)),
				),
			).
				ThenReturn(&listAfterOutput, nil)

			conv := Mock[convert.ToGetNext](ctrl)
			WhenDouble(conv.ToResponse(AnyContext(), Equal(&listAfterOutput))).
				ThenReturn(nil, false)

			sut := NewTagServiceServer(WithListAfter(uc, conv))
			got, err := sut.GetNextTags(
				s.T().Context(),
				connect.NewRequest(
					&grpc.GetNextTagsRequest{
						First: 1,
					},
				),
			)
			s.Require().Error(err)
			s.Require().Nil(got)
			s.Require().ErrorIs(err, ErrConversionToGetNextTagsFailed)
		},
	)
}

func (s *TagServiceServerTestSuite) TestTagServiceServer_GetPrevTags() {
	s.Run(
		"happy_path/without-cursor", func() {
			ctrl := NewMockController(s.T())
			uc := Mock[usecase.ListBefore](ctrl)

			listBeforeOutput := dto.NewListBeforeOutput(
				true,
				dto.NewTag(
					"1",
					"tag1",
					dto.NewArticle(
						"1",
						"happy_path/without-cursor",
						"1234567890",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					),
				),
			)
			WhenDouble(
				uc.Execute(
					AnyContext(),
					Equal(dto.NewListBeforeInput(1)),
				),
			).
				ThenReturn(&listBeforeOutput, nil)

			res := connect.NewResponse(
				&grpc.GetPrevTagResponse{
					StillExists: true,
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "tag1",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/without-cursor",
									ThumbnailUrl: "1234567890",
									CreatedAt: timestamppb.New(
										synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).
											StdTime(),
									),
									UpdatedAt: timestamppb.New(
										synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).
											StdTime(),
									),
								},
							},
						},
					},
				},
			)

			conv := Mock[convert.ToGetPrev](ctrl)
			WhenDouble(conv.ToResponse(AnyContext(), Equal(&listBeforeOutput))).
				ThenReturn(res, true)

			sut := NewTagServiceServer(WithListBefore(uc, conv))
			got, err := sut.GetPrevTags(
				s.T().Context(),
				connect.NewRequest(
					&grpc.GetPrevTagsRequest{
						Last: 1,
					},
				),
			)
			s.Require().NoError(err)
			s.Require().NotNil(got)
			s.Require().Equal(res.Msg, got.Msg)
		},
	)
	s.Run(
		"happy_path/with-cursor", func() {
			ctrl := NewMockController(s.T())
			uc := Mock[usecase.ListBefore](ctrl)

			listBeforeOutput := dto.NewListBeforeOutput(
				true,
				dto.NewTag(
					"1",
					"tag1",
					dto.NewArticle(
						"1",
						"happy_path/without-cursor",
						"1234567890",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					),
				),
			)
			WhenDouble(
				uc.Execute(
					AnyContext(),
					Equal(dto.NewListBeforeInput(1, dto.ListBeforeInputWithCursor("2"))),
				),
			).
				ThenReturn(&listBeforeOutput, nil)

			res := connect.NewResponse(
				&grpc.GetPrevTagResponse{
					StillExists: true,
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "tag1",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/without-cursor",
									ThumbnailUrl: "1234567890",
									CreatedAt: timestamppb.New(
										synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).
											StdTime(),
									),
									UpdatedAt: timestamppb.New(
										synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).
											StdTime(),
									),
								},
							},
						},
					},
				},
			)

			conv := Mock[convert.ToGetPrev](ctrl)
			WhenDouble(conv.ToResponse(AnyContext(), Equal(&listBeforeOutput))).
				ThenReturn(res, true)

			sut := NewTagServiceServer(WithListBefore(uc, conv))
			got, err := sut.GetPrevTags(
				s.T().Context(),
				connect.NewRequest(
					&grpc.GetPrevTagsRequest{
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
			s.Require().Equal(res.Msg, got.Msg)
		},
	)
	s.Run(
		"unhappy_path/usecase_returns_error", func() {
			errGetPrevTag := errors.New("error get prev tags")

			ctrl := NewMockController(s.T())
			uc := Mock[usecase.ListBefore](ctrl)

			WhenDouble(
				uc.Execute(
					AnyContext(),
					Equal(dto.NewListBeforeInput(1)),
				),
			).
				ThenReturn(nil, errGetPrevTag)

			sut := NewTagServiceServer(WithListBefore(uc, nil))
			got, err := sut.GetPrevTags(
				s.T().Context(),
				connect.NewRequest(
					&grpc.GetPrevTagsRequest{
						Last: 1,
					},
				),
			)
			s.Require().Error(err)
			s.Require().Nil(got)
			s.Require().ErrorIs(err, errGetPrevTag)
		},
	)
	s.Run(
		"unhappy_path/failed_to_convert", func() {
			ctrl := NewMockController(s.T())
			uc := Mock[usecase.ListBefore](ctrl)

			listBeforeOutput := dto.NewListBeforeOutput(
				true,
				dto.NewTag(
					"1",
					"tag1",
					dto.NewArticle(
						"1",
						"unhappy_path/failed_to_convert",
						"1234567890",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
					),
				),
			)
			WhenDouble(
				uc.Execute(
					AnyContext(),
					Equal(dto.NewListBeforeInput(1)),
				),
			).
				ThenReturn(&listBeforeOutput, nil)

			conv := Mock[convert.ToGetPrev](ctrl)
			WhenDouble(conv.ToResponse(AnyContext(), Equal(&listBeforeOutput))).
				ThenReturn(nil, false)

			sut := NewTagServiceServer(WithListBefore(uc, conv))
			got, err := sut.GetPrevTags(
				s.T().Context(),
				connect.NewRequest(
					&grpc.GetPrevTagsRequest{
						Last: 1,
					},
				),
			)
			s.Require().Error(err)
			s.Require().Nil(got)
			s.Require().ErrorIs(err, ErrConversionToGetPrevTagsFailed)
		},
	)
}
