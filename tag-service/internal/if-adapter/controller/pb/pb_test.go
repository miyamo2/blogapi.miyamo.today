package pb

import (
	"context"
	"testing"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/cockroachdb/errors"
	"github.com/google/go-cmp/cmp"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/infra/grpc"
	mpresenter "github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/mock/if-adapter/controller/pb/presenter"
	musecase "github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/mock/if-adapter/controller/pb/usecase"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestTagServiceServer_GetTagById(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *grpc.GetTagByIdRequest
	}
	type want struct {
		response *grpc.GetTagByIdResponse
		err      error
	}
	type testCase struct {
		outDto         dto.GetByIdOutDto
		setupUsecase   func(out dto.GetByIdOutDto, u *musecase.MockGetById)
		setupConverter func(from dto.GetByIdOutDto, res *grpc.GetTagByIdResponse, conv *mpresenter.MockToGetByIdConverter)
		args           args
		want           want
		wantErr        bool
	}
	errGetTagById := errors.New("error get article by id")
	tests := map[string]testCase{
		"happy_path/tag_has_article": {
			outDto: dto.NewGetByIdOutDto(
				"1",
				"happy_path/tag_has_article",
				[]dto.Article{
					dto.NewArticle(
						"1",
						"happy_path/article_has_tag",
						"1234567890",
						"2020-01-01T00:00:00Z",
						"2020-01-01T00:00:00Z"),
				}),
			setupUsecase: func(out dto.GetByIdOutDto, u *musecase.MockGetById) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetByIdInDto("1")).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetByIdOutDto, res *grpc.GetTagByIdResponse, conv *mpresenter.MockToGetByIdConverter) {
				conv.EXPECT().
					ToGetByIdTagResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.GetTagByIdRequest{
					Id: "1",
				},
			},
			want: want{
				response: &grpc.GetTagByIdResponse{
					Tag: &grpc.Tag{
						Id:   "1",
						Name: "happy_path/tag_has_article",
						Articles: []*grpc.Article{
							{
								Id:           "1",
								Title:        "happy_path/article_has_tag",
								ThumbnailUrl: "1234567890",
								CreatedAt:    "2020-01-01T00:00:00Z",
								UpdatedAt:    "2020-01-01T00:00:00Z",
							},
						},
					},
				},
				err: nil,
			},
		},
		"unhappy_path/usecase_returns_error": {
			setupUsecase: func(out dto.GetByIdOutDto, u *musecase.MockGetById) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetByIdInDto("1")).
					Return(nil, errGetTagById).Times(1)
			},
			setupConverter: func(from dto.GetByIdOutDto, res *grpc.GetTagByIdResponse, conv *mpresenter.MockToGetByIdConverter) {
				conv.EXPECT().
					ToGetByIdTagResponse(gomock.Any(), gomock.Any()).
					Times(0)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.GetTagByIdRequest{
					Id: "1",
				},
			},
			want: want{
				response: nil,
				err:      errGetTagById,
			},
			wantErr: true,
		},
		"unhappy_path/failed_to_convert": {
			setupUsecase: func(out dto.GetByIdOutDto, u *musecase.MockGetById) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetByIdInDto("1")).
					Return(&out, nil).Times(1)
			},
			setupConverter: func(from dto.GetByIdOutDto, res *grpc.GetTagByIdResponse, conv *mpresenter.MockToGetByIdConverter) {
				conv.EXPECT().
					ToGetByIdTagResponse(gomock.Any(), &from).
					Return(nil, false).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.GetTagByIdRequest{
					Id: "1",
				},
			},
			want: want{
				response: nil,
				err:      ErrConversionToGetTagByIdFailed,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			out := tt.outDto
			u := musecase.NewMockGetById(ctrl)
			tt.setupUsecase(out, u)
			response := tt.want.response
			conv := mpresenter.NewMockToGetByIdConverter(ctrl)
			tt.setupConverter(out, response, conv)
			s := NewTagServiceServer(u, conv, nil, nil, nil, nil, nil, nil)
			got, err := s.GetTagById(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				if err == nil {
					t.Errorf("GetTagById() expected to return an error, but it was nil. want: %+v", err)
					return
				}
				if !errors.Is(err, tt.want.err) {
					t.Errorf("GetTagById() error = %v, want %v", err, tt.want.err)
					return
				}
				return
			}
			if diff := cmp.Diff(got, tt.want.response, protocmp.Transform()); diff != "" {
				t.Errorf("GetTagById() got = %v, want %v", got, tt.want.response)
			}
		})
	}
}

func TestTagServiceServer_GetAllTags(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *emptypb.Empty
	}
	type want struct {
		response *grpc.GetAllTagsResponse
		err      error
	}
	type testCase struct {
		outDto         dto.GetAllOutDto
		setupUsecase   func(out dto.GetAllOutDto, u *musecase.MockGetAll)
		setupConverter func(from dto.GetAllOutDto, res *grpc.GetAllTagsResponse, conv *mpresenter.MockToGetAllConverter)
		args           args
		want           want
		wantErr        bool
	}
	errGetAllTag := errors.New("error get all tags")
	tests := map[string]testCase{
		"happy_path/single_tag/single_article": {
			outDto: func() dto.GetAllOutDto {
				o := dto.NewGetAllOutDto()
				o = o.WithTagDto(dto.NewTag(
					"1",
					"happy_path/single_tag/single_article",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/single_tag/single_article",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z",
						),
					}))
				return o
			}(),
			setupUsecase: func(out dto.GetAllOutDto, u *musecase.MockGetAll) {
				u.EXPECT().
					Execute(gomock.Any()).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetAllOutDto, res *grpc.GetAllTagsResponse, conv *mpresenter.MockToGetAllConverter) {
				conv.EXPECT().
					ToGetAllTagsResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in:  &emptypb.Empty{},
			},
			want: want{
				response: &grpc.GetAllTagsResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "happy_path/single_tag/single_article",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/single_tag/single_article",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
					},
				},
				err: nil,
			},
		},
		"happy_path/single_tag/multiple_article": {
			outDto: func() dto.GetAllOutDto {
				o := dto.NewGetAllOutDto()
				o = o.WithTagDto(dto.NewTag(
					"1",
					"happy_path/single_tag/multiple_article",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/single_tag/multiple_article1",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z",
						),
						dto.NewArticle(
							"2",
							"happy_path/single_tag/multiple_article2",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z",
						),
					}))
				return o
			}(),
			setupUsecase: func(out dto.GetAllOutDto, u *musecase.MockGetAll) {
				u.EXPECT().
					Execute(gomock.Any()).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetAllOutDto, res *grpc.GetAllTagsResponse, conv *mpresenter.MockToGetAllConverter) {
				conv.EXPECT().
					ToGetAllTagsResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in:  &emptypb.Empty{},
			},
			want: want{
				response: &grpc.GetAllTagsResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "happy_path/single_tag/multiple_article",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/single_tag/multiple_article1",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
								{
									Id:           "2",
									Title:        "happy_path/single_tag/multiple_article2",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
					},
				},
				err: nil,
			},
		},
		"happy_path/multiple_tag/single_article": {
			outDto: func() dto.GetAllOutDto {
				o := dto.NewGetAllOutDto()
				o = o.WithTagDto(dto.NewTag(
					"1",
					"happy_path/multiple_tag/single_article1",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/multiple_tag/single_article",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z",
						),
					}))
				o = o.WithTagDto(dto.NewTag(
					"2",
					"happy_path/multiple_tag/single_article2",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/multiple_tag/single_article",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z",
						),
					}))
				return o
			}(),
			setupUsecase: func(out dto.GetAllOutDto, u *musecase.MockGetAll) {
				u.EXPECT().
					Execute(gomock.Any()).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetAllOutDto, res *grpc.GetAllTagsResponse, conv *mpresenter.MockToGetAllConverter) {
				conv.EXPECT().
					ToGetAllTagsResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in:  &emptypb.Empty{},
			},
			want: want{
				response: &grpc.GetAllTagsResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "happy_path/multiple_tag/single_article1",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/multiple_tag/single_article",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
						{
							Id:   "2",
							Name: "happy_path/multiple_tag/single_article2",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/multiple_tag/single_article",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
					},
				},
				err: nil,
			},
		},
		"happy_path/multiple_tag/multiple_article": {
			outDto: func() dto.GetAllOutDto {
				o := dto.NewGetAllOutDto()
				o = o.WithTagDto(dto.NewTag(
					"1",
					"happy_path/multiple_tag/multiple_article1",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/multiple_tag/multiple_article1",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z",
						),
						dto.NewArticle(
							"2",
							"happy_path/multiple_tag/multiple_article2",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z",
						),
					}))
				o = o.WithTagDto(dto.NewTag(
					"2",
					"happy_path/multiple_tag/multiple_article2",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/multiple_tag/multiple_article1",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z",
						),
						dto.NewArticle(
							"2",
							"happy_path/multiple_tag/multiple_article2",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z",
						),
					}))
				return o
			}(),
			setupUsecase: func(out dto.GetAllOutDto, u *musecase.MockGetAll) {
				u.EXPECT().
					Execute(gomock.Any()).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetAllOutDto, res *grpc.GetAllTagsResponse, conv *mpresenter.MockToGetAllConverter) {
				conv.EXPECT().
					ToGetAllTagsResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in:  &emptypb.Empty{},
			},
			want: want{
				response: &grpc.GetAllTagsResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "happy_path/multiple_tag/multiple_article1",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/multiple_tag/multiple_article1",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
								{
									Id:           "2",
									Title:        "happy_path/multiple_tag/multiple_article2",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
						{
							Id:   "2",
							Name: "happy_path/multiple_tag/multiple_article2",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/multiple_tag/multiple_article1",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
								{
									Id:           "2",
									Title:        "happy_path/multiple_tag/multiple_article2",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
					},
				},
				err: nil,
			},
		},
		"unhappy_path/usecase_returns_error": {
			setupUsecase: func(out dto.GetAllOutDto, u *musecase.MockGetAll) {
				u.EXPECT().
					Execute(gomock.Any()).
					Return(nil, errGetAllTag).Times(1)
			},
			setupConverter: func(from dto.GetAllOutDto, res *grpc.GetAllTagsResponse, conv *mpresenter.MockToGetAllConverter) {
				conv.EXPECT().
					ToGetAllTagsResponse(gomock.Any(), &from).
					Times(0)
			},
			args: args{
				ctx: context.Background(),
				in:  &emptypb.Empty{},
			},
			want: want{
				response: nil,
				err:      errGetAllTag,
			},
			wantErr: true,
		},
		"unhappy_path/failed_to_convert": {
			setupUsecase: func(out dto.GetAllOutDto, u *musecase.MockGetAll) {
				u.EXPECT().
					Execute(gomock.Any()).
					Return(&out, nil).Times(1)
			},
			setupConverter: func(from dto.GetAllOutDto, res *grpc.GetAllTagsResponse, conv *mpresenter.MockToGetAllConverter) {
				conv.EXPECT().
					ToGetAllTagsResponse(gomock.Any(), &from).
					Return(nil, false).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in:  &emptypb.Empty{},
			},
			want: want{
				response: nil,
				err:      ErrConversionToGetAllTagsFailed,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			out := tt.outDto
			u := musecase.NewMockGetAll(ctrl)
			tt.setupUsecase(out, u)
			response := tt.want.response
			conv := mpresenter.NewMockToGetAllConverter(ctrl)
			tt.setupConverter(out, response, conv)
			s := NewTagServiceServer(nil, nil, u, conv, nil, nil, nil, nil)
			got, err := s.GetAllTags(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				if err == nil {
					t.Errorf("GetAllTags() expected to return an error, but it was nil. want: %+v", err)
					return
				}
				if !errors.Is(err, tt.want.err) {
					t.Errorf("GetAllTags() error = %v, want %v", err, tt.want.err)
					return
				}
				return
			}
			if diff := cmp.Diff(got, tt.want.response, protocmp.Transform()); diff != "" {
				t.Errorf("GetAllTags() got = %v, want %v", got, tt.want.response)
			}
		})
	}
}

func TestTagServiceServer_GetNextTags(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *grpc.GetNextTagsRequest
	}
	type want struct {
		response *grpc.GetNextTagResponse
		err      error
	}
	type testCase struct {
		outDto         dto.GetNextOutDto
		setupUsecase   func(out dto.GetNextOutDto, u *musecase.MockGetNext)
		setupConverter func(from dto.GetNextOutDto, res *grpc.GetNextTagResponse, conv *mpresenter.MockToGetNextConverter)
		args           args
		want           want
		wantErr        bool
	}
	errGetNextTag := errors.New("error get next tags")
	tests := map[string]testCase{
		"happy_path/single_tag/single_article/has_next": {
			outDto: func() dto.GetNextOutDto {
				o := dto.NewGetNextOutDto(true)
				o = o.WithTagDto(dto.NewTag(
					"1",
					"happy_path/single_tag/single_article/has_next",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/single_tag/single_article/has_next",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z",
						),
					}))
				return o
			}(),
			setupUsecase: func(out dto.GetNextOutDto, u *musecase.MockGetNext) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetNextInDto(1, nil)).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetNextOutDto, res *grpc.GetNextTagResponse, conv *mpresenter.MockToGetNextConverter) {
				conv.EXPECT().
					ToGetNextTagsResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.GetNextTagsRequest{
					First: 1,
				},
			},
			want: want{
				response: &grpc.GetNextTagResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "happy_path/single_tag/single_article/has_next",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/single_tag/single_article/has_next",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
					},
					StillExists: true,
				},
			},
		},
		"happy_path/single_tag/single_article/not_anymore": {
			outDto: func() dto.GetNextOutDto {
				o := dto.NewGetNextOutDto(false)
				o = o.WithTagDto(dto.NewTag(
					"1",
					"happy_path/single_tag/single_article/not_anymore",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/single_tag/single_article/not_anymore",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z",
						),
					}))
				return o
			}(),
			setupUsecase: func(out dto.GetNextOutDto, u *musecase.MockGetNext) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetNextInDto(1, nil)).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetNextOutDto, res *grpc.GetNextTagResponse, conv *mpresenter.MockToGetNextConverter) {
				conv.EXPECT().
					ToGetNextTagsResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.GetNextTagsRequest{
					First: 1,
				},
			},
			want: want{
				response: &grpc.GetNextTagResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "happy_path/single_tag/single_article/not_anymore",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/single_tag/single_article/not_anymore",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
					},
					StillExists: false,
				},
			},
		},
		"happy_path/single_tag/multiple_article/has_next": {
			outDto: func() dto.GetNextOutDto {
				o := dto.NewGetNextOutDto(true)
				o = o.WithTagDto(dto.NewTag(
					"1",
					"happy_path/single_tag/multiple_article/has_next",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/single_tag/multiple_article/has_next1",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
						dto.NewArticle(
							"2",
							"happy_path/single_tag/multiple_article/has_next2",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
					}))
				return o
			}(),
			setupUsecase: func(out dto.GetNextOutDto, u *musecase.MockGetNext) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetNextInDto(1, nil)).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetNextOutDto, res *grpc.GetNextTagResponse, conv *mpresenter.MockToGetNextConverter) {
				conv.EXPECT().
					ToGetNextTagsResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.GetNextTagsRequest{
					First: 1,
				},
			},
			want: want{
				response: &grpc.GetNextTagResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "happy_path/single_tag/multiple_article/has_next",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/single_tag/multiple_article/has_next1",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
								{
									Id:           "2",
									Title:        "happy_path/single_tag/multiple_article/has_next2",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
					},
					StillExists: true,
				},
			},
		},
		"happy_path/single_tag/multiple_article/not_anymore": {
			outDto: func() dto.GetNextOutDto {
				o := dto.NewGetNextOutDto(false)
				o = o.WithTagDto(dto.NewTag(
					"1",
					"happy_path/single_tag/multiple_article/not_anymore",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/single_tag/multiple_article/not_anymore1",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
						dto.NewArticle(
							"2",
							"happy_path/single_tag/multiple_article/not_anymore2",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
					}))
				return o
			}(),
			setupUsecase: func(out dto.GetNextOutDto, u *musecase.MockGetNext) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetNextInDto(1, nil)).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetNextOutDto, res *grpc.GetNextTagResponse, conv *mpresenter.MockToGetNextConverter) {
				conv.EXPECT().
					ToGetNextTagsResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.GetNextTagsRequest{
					First: 1,
				},
			},
			want: want{
				response: &grpc.GetNextTagResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "happy_path/single_tag/multiple_article/not_anymore",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/single_tag/multiple_article/not_anymore1",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
								{
									Id:           "2",
									Title:        "happy_path/single_tag/multiple_article/not_anymore2",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
					},
				},
			},
		},
		"happy_path/multiple_tag/single_article/has_next": {
			outDto: func() dto.GetNextOutDto {
				o := dto.NewGetNextOutDto(true)
				o = o.WithTagDto(dto.NewTag(
					"1",
					"happy_path/multiple_tag/single_article/has_next1",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/multiple_tag/single_article/has_next",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
					}))
				o = o.WithTagDto(dto.NewTag(
					"2",
					"happy_path/multiple_tag/single_article/has_next2",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/multiple_tag/single_article/has_next",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
					}))
				return o
			}(),
			setupUsecase: func(out dto.GetNextOutDto, u *musecase.MockGetNext) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetNextInDto(2, nil)).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetNextOutDto, res *grpc.GetNextTagResponse, conv *mpresenter.MockToGetNextConverter) {
				conv.EXPECT().
					ToGetNextTagsResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.GetNextTagsRequest{
					First: 2,
				},
			},
			want: want{
				response: &grpc.GetNextTagResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "happy_path/multiple_tag/single_article/has_next1",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/multiple_tag/single_article/has_next",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
						{
							Id:   "2",
							Name: "happy_path/multiple_tag/single_article/has_next2",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/multiple_tag/single_article/has_next",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
					},
					StillExists: true,
				},
			},
		},
		"happy_path/multiple_tag/single_article/not_anymore": {
			outDto: func() dto.GetNextOutDto {
				o := dto.NewGetNextOutDto(false)
				o = o.WithTagDto(dto.NewTag(
					"1",
					"happy_path/multiple_tag/single_article/not_anymore1",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/multiple_tag/single_article/not_anymore",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
					}))
				o = o.WithTagDto(dto.NewTag(
					"2",
					"happy_path/multiple_tag/single_article/not_anymore2",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/multiple_tag/single_article/not_anymore",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
					}))
				return o
			}(),
			setupUsecase: func(out dto.GetNextOutDto, u *musecase.MockGetNext) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetNextInDto(2, nil)).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetNextOutDto, res *grpc.GetNextTagResponse, conv *mpresenter.MockToGetNextConverter) {
				conv.EXPECT().
					ToGetNextTagsResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.GetNextTagsRequest{
					First: 2,
				},
			},
			want: want{
				response: &grpc.GetNextTagResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "happy_path/multiple_tag/single_article/not_anymore1",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/multiple_tag/single_article/not_anymore",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
						{
							Id:   "2",
							Name: "happy_path/multiple_tag/single_article/not_anymore2",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/multiple_tag/single_article/not_anymore",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
					},
				},
			},
		},
		"happy_path/multiple_tag/multiple_article/has_next": {
			outDto: func() dto.GetNextOutDto {
				o := dto.NewGetNextOutDto(true)
				o = o.WithTagDto(dto.NewTag(
					"1",
					"happy_path/multiple_tag/multiple_article/has_next1",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/multiple_tag/multiple_article/has_next1",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
						dto.NewArticle(
							"2",
							"happy_path/multiple_tag/multiple_article/has_next2",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
					}))
				o = o.WithTagDto(dto.NewTag(
					"2",
					"happy_path/multiple_tag/multiple_article/has_next2",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/multiple_tag/multiple_article/has_next1",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
						dto.NewArticle(
							"2",
							"happy_path/multiple_tag/multiple_article/has_next2",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
					}))
				return o
			}(),
			setupUsecase: func(out dto.GetNextOutDto, u *musecase.MockGetNext) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetNextInDto(2, nil)).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetNextOutDto, res *grpc.GetNextTagResponse, conv *mpresenter.MockToGetNextConverter) {
				conv.EXPECT().
					ToGetNextTagsResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.GetNextTagsRequest{
					First: 2,
				},
			},
			want: want{
				response: &grpc.GetNextTagResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "happy_path/multiple_tag/multiple_article/has_next1",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/multiple_tag/multiple_article/has_next1",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
								{
									Id:           "2",
									Title:        "happy_path/multiple_tag/multiple_article/has_next2",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
						{
							Id:   "2",
							Name: "happy_path/multiple_tag/multiple_article/has_next2",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/multiple_tag/multiple_article/has_next1",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
								{
									Id:           "2",
									Title:        "happy_path/multiple_tag/multiple_article/has_next2",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
					},
					StillExists: true,
				},
			},
		},
		"happy_path/multiple_tag/multiple_article/not_anymore": {
			outDto: func() dto.GetNextOutDto {
				o := dto.NewGetNextOutDto(false)
				o = o.WithTagDto(dto.NewTag(
					"1",
					"happy_path/multiple_tag/multiple_article/not_anymore1",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/multiple_tag/multiple_article/not_anymore1",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
						dto.NewArticle(
							"2",
							"happy_path/multiple_tag/multiple_article/not_anymore2",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
					}))
				o = o.WithTagDto(dto.NewTag(
					"2",
					"happy_path/multiple_tag/multiple_article/not_anymore2",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/multiple_tag/multiple_article/not_anymore1",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
						dto.NewArticle(
							"2",
							"happy_path/multiple_tag/multiple_article/not_anymore2",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
					}))
				return o
			}(),
			setupUsecase: func(out dto.GetNextOutDto, u *musecase.MockGetNext) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetNextInDto(2, nil)).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetNextOutDto, res *grpc.GetNextTagResponse, conv *mpresenter.MockToGetNextConverter) {
				conv.EXPECT().
					ToGetNextTagsResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.GetNextTagsRequest{
					First: 2,
				},
			},
			want: want{
				response: &grpc.GetNextTagResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "happy_path/multiple_tag/multiple_article/not_anymore1",
							Articles: []*grpc.Article{
								{
									Id:    "1",
									Title: "happy_path/multiple_tag/multiple_article/not_anymore1",
								},
							},
						},
					},
				},
			},
		},
		"unhappy_path/usecase_returns_error": {
			setupUsecase: func(out dto.GetNextOutDto, u *musecase.MockGetNext) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetNextInDto(1, nil)).
					Return(nil, errGetNextTag).
					Times(1)
			},
			setupConverter: func(from dto.GetNextOutDto, res *grpc.GetNextTagResponse, conv *mpresenter.MockToGetNextConverter) {
				conv.EXPECT().
					ToGetNextTagsResponse(gomock.Any(), gomock.Any()).
					Times(0)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.GetNextTagsRequest{
					First: 1,
				},
			},
			want: want{
				response: nil,
				err:      errGetNextTag,
			},
		},
		"unhappy_path/failed_to_convert": {
			outDto: dto.NewGetNextOutDto(true),
			setupUsecase: func(out dto.GetNextOutDto, u *musecase.MockGetNext) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetNextInDto(1, nil)).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetNextOutDto, res *grpc.GetNextTagResponse, conv *mpresenter.MockToGetNextConverter) {
				conv.EXPECT().
					ToGetNextTagsResponse(gomock.Any(), &from).
					Return(nil, false).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.GetNextTagsRequest{
					First: 1,
				},
			},
			want: want{
				response: nil,
				err:      ErrConversionToGetNextTagsFailed,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			out := tt.outDto
			u := musecase.NewMockGetNext(ctrl)
			tt.setupUsecase(out, u)
			response := tt.want.response
			conv := mpresenter.NewMockToGetNextConverter(ctrl)
			tt.setupConverter(out, response, conv)
			s := NewTagServiceServer(nil, nil, nil, nil, u, conv, nil, nil)
			got, err := s.GetNextTags(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				if err == nil {
					t.Errorf("GetNextTags() expected to return an error, but it was nil. want: %+v", err)
					return
				}
				if !errors.Is(err, tt.want.err) {
					t.Errorf("GetNextTags() error = %v, want %v", err, tt.want.err)
					return
				}
				return
			}
			if diff := cmp.Diff(got, tt.want.response, protocmp.Transform()); diff != "" {
				t.Errorf("GetNextTags() got = %v, want %v", got, tt.want.response)
			}
		})
	}
}

func TestTagServiceServer_GetPrevTags(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *grpc.GetPrevTagsRequest
	}
	type want struct {
		response *grpc.GetPrevTagResponse
		err      error
	}
	type testCase struct {
		outDto         dto.GetPrevOutDto
		setupUsecase   func(out dto.GetPrevOutDto, u *musecase.MockGetPrev)
		setupConverter func(from dto.GetPrevOutDto, res *grpc.GetPrevTagResponse, conv *mpresenter.MockToGetPrevConverter)
		args           args
		want           want
		wantErr        bool
	}
	errGetPrevTag := errors.New("error get prev tags")
	tests := map[string]testCase{
		"happy_path/single_tag/single_article/has_prev": {
			outDto: func() dto.GetPrevOutDto {
				o := dto.NewGetPrevOutDto(true)
				o = o.WithTagDto(dto.NewTag(
					"1",
					"happy_path/single_tag/single_article/has_prev",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/single_tag/single_article/has_prev",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z",
						),
					}))
				return o
			}(),
			setupUsecase: func(out dto.GetPrevOutDto, u *musecase.MockGetPrev) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetPrevInDto(1, nil)).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetPrevOutDto, res *grpc.GetPrevTagResponse, conv *mpresenter.MockToGetPrevConverter) {
				conv.EXPECT().
					ToGetPrevTagsResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.GetPrevTagsRequest{
					Last: 1,
				},
			},
			want: want{
				response: &grpc.GetPrevTagResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "happy_path/single_tag/single_article/has_prev",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/single_tag/single_article/has_prev",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
					},
					StillExists: true,
				},
			},
		},
		"happy_path/single_tag/single_article/not_anymore": {
			outDto: func() dto.GetPrevOutDto {
				o := dto.NewGetPrevOutDto(false)
				o = o.WithTagDto(dto.NewTag(
					"1",
					"happy_path/single_tag/single_article/not_anymore",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/single_tag/single_article/not_anymore",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z",
						),
					}))
				return o
			}(),
			setupUsecase: func(out dto.GetPrevOutDto, u *musecase.MockGetPrev) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetPrevInDto(1, nil)).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetPrevOutDto, res *grpc.GetPrevTagResponse, conv *mpresenter.MockToGetPrevConverter) {
				conv.EXPECT().
					ToGetPrevTagsResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.GetPrevTagsRequest{
					Last: 1,
				},
			},
			want: want{
				response: &grpc.GetPrevTagResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "happy_path/single_tag/single_article/not_anymore",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/single_tag/single_article/not_anymore",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
					},
					StillExists: false,
				},
			},
		},
		"happy_path/single_tag/multiple_article/has_prev": {
			outDto: func() dto.GetPrevOutDto {
				o := dto.NewGetPrevOutDto(true)
				o = o.WithTagDto(dto.NewTag(
					"1",
					"happy_path/single_tag/multiple_article/has_prev",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/single_tag/multiple_article/has_prev1",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
						dto.NewArticle(
							"2",
							"happy_path/single_tag/multiple_article/has_prev2",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
					}))
				return o
			}(),
			setupUsecase: func(out dto.GetPrevOutDto, u *musecase.MockGetPrev) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetPrevInDto(1, nil)).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetPrevOutDto, res *grpc.GetPrevTagResponse, conv *mpresenter.MockToGetPrevConverter) {
				conv.EXPECT().
					ToGetPrevTagsResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.GetPrevTagsRequest{
					Last: 1,
				},
			},
			want: want{
				response: &grpc.GetPrevTagResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "happy_path/single_tag/multiple_article/has_prev",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/single_tag/multiple_article/has_prev1",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
								{
									Id:           "2",
									Title:        "happy_path/single_tag/multiple_article/has_prev2",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
					},
					StillExists: true,
				},
			},
		},
		"happy_path/single_tag/multiple_article/not_anymore": {
			outDto: func() dto.GetPrevOutDto {
				o := dto.NewGetPrevOutDto(false)
				o = o.WithTagDto(dto.NewTag(
					"1",
					"happy_path/single_tag/multiple_article/not_anymore",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/single_tag/multiple_article/not_anymore1",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
						dto.NewArticle(
							"2",
							"happy_path/single_tag/multiple_article/not_anymore2",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
					}))
				return o
			}(),
			setupUsecase: func(out dto.GetPrevOutDto, u *musecase.MockGetPrev) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetPrevInDto(1, nil)).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetPrevOutDto, res *grpc.GetPrevTagResponse, conv *mpresenter.MockToGetPrevConverter) {
				conv.EXPECT().
					ToGetPrevTagsResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.GetPrevTagsRequest{
					Last: 1,
				},
			},
			want: want{
				response: &grpc.GetPrevTagResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "happy_path/single_tag/multiple_article/not_anymore",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/single_tag/multiple_article/not_anymore1",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
								{
									Id:           "2",
									Title:        "happy_path/single_tag/multiple_article/not_anymore2",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
					},
				},
			},
		},
		"happy_path/multiple_tag/single_article/has_prev": {
			outDto: func() dto.GetPrevOutDto {
				o := dto.NewGetPrevOutDto(true)
				o = o.WithTagDto(dto.NewTag(
					"1",
					"happy_path/multiple_tag/single_article/has_prev1",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/multiple_tag/single_article/has_prev",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
					}))
				o = o.WithTagDto(dto.NewTag(
					"2",
					"happy_path/multiple_tag/single_article/has_prev2",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/multiple_tag/single_article/has_prev",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
					}))
				return o
			}(),
			setupUsecase: func(out dto.GetPrevOutDto, u *musecase.MockGetPrev) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetPrevInDto(2, nil)).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetPrevOutDto, res *grpc.GetPrevTagResponse, conv *mpresenter.MockToGetPrevConverter) {
				conv.EXPECT().
					ToGetPrevTagsResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.GetPrevTagsRequest{
					Last: 2,
				},
			},
			want: want{
				response: &grpc.GetPrevTagResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "happy_path/multiple_tag/single_article/has_prev1",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/multiple_tag/single_article/has_prev",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
						{
							Id:   "2",
							Name: "happy_path/multiple_tag/single_article/has_prev2",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/multiple_tag/single_article/has_prev",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
					},
					StillExists: true,
				},
			},
		},
		"happy_path/multiple_tag/single_article/not_anymore": {
			outDto: func() dto.GetPrevOutDto {
				o := dto.NewGetPrevOutDto(false)
				o = o.WithTagDto(dto.NewTag(
					"1",
					"happy_path/multiple_tag/single_article/not_anymore1",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/multiple_tag/single_article/not_anymore",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
					}))
				o = o.WithTagDto(dto.NewTag(
					"2",
					"happy_path/multiple_tag/single_article/not_anymore2",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/multiple_tag/single_article/not_anymore",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
					}))
				return o
			}(),
			setupUsecase: func(out dto.GetPrevOutDto, u *musecase.MockGetPrev) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetPrevInDto(2, nil)).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetPrevOutDto, res *grpc.GetPrevTagResponse, conv *mpresenter.MockToGetPrevConverter) {
				conv.EXPECT().
					ToGetPrevTagsResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.GetPrevTagsRequest{
					Last: 2,
				},
			},
			want: want{
				response: &grpc.GetPrevTagResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "happy_path/multiple_tag/single_article/not_anymore1",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/multiple_tag/single_article/not_anymore",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
						{
							Id:   "2",
							Name: "happy_path/multiple_tag/single_article/not_anymore2",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/multiple_tag/single_article/not_anymore",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
					},
				},
			},
		},
		"happy_path/multiple_tag/multiple_article/has_prev": {
			outDto: func() dto.GetPrevOutDto {
				o := dto.NewGetPrevOutDto(true)
				o = o.WithTagDto(dto.NewTag(
					"1",
					"happy_path/multiple_tag/multiple_article/has_prev1",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/multiple_tag/multiple_article/has_prev1",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
						dto.NewArticle(
							"2",
							"happy_path/multiple_tag/multiple_article/has_prev2",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
					}))
				o = o.WithTagDto(dto.NewTag(
					"2",
					"happy_path/multiple_tag/multiple_article/has_prev2",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/multiple_tag/multiple_article/has_prev1",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
						dto.NewArticle(
							"2",
							"happy_path/multiple_tag/multiple_article/has_prev2",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
					}))
				return o
			}(),
			setupUsecase: func(out dto.GetPrevOutDto, u *musecase.MockGetPrev) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetPrevInDto(2, nil)).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetPrevOutDto, res *grpc.GetPrevTagResponse, conv *mpresenter.MockToGetPrevConverter) {
				conv.EXPECT().
					ToGetPrevTagsResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.GetPrevTagsRequest{
					Last: 2,
				},
			},
			want: want{
				response: &grpc.GetPrevTagResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "happy_path/multiple_tag/multiple_article/has_prev1",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/multiple_tag/multiple_article/has_prev1",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
								{
									Id:           "2",
									Title:        "happy_path/multiple_tag/multiple_article/has_prev2",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
						{
							Id:   "2",
							Name: "happy_path/multiple_tag/multiple_article/has_prev2",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/multiple_tag/multiple_article/has_prev1",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
								{
									Id:           "2",
									Title:        "happy_path/multiple_tag/multiple_article/has_prev2",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
					},
					StillExists: true,
				},
			},
		},
		"happy_path/multiple_tag/multiple_article/not_anymore": {
			outDto: func() dto.GetPrevOutDto {
				o := dto.NewGetPrevOutDto(false)
				o = o.WithTagDto(dto.NewTag(
					"1",
					"happy_path/multiple_tag/multiple_article/not_anymore1",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/multiple_tag/multiple_article/not_anymore1",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
						dto.NewArticle(
							"2",
							"happy_path/multiple_tag/multiple_article/not_anymore2",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
					}))
				o = o.WithTagDto(dto.NewTag(
					"2",
					"happy_path/multiple_tag/multiple_article/not_anymore2",
					[]dto.Article{
						dto.NewArticle(
							"1",
							"happy_path/multiple_tag/multiple_article/not_anymore1",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
						dto.NewArticle(
							"2",
							"happy_path/multiple_tag/multiple_article/not_anymore2",
							"1234567890",
							"2020-01-01T00:00:00Z",
							"2020-01-01T00:00:00Z"),
					}))
				return o
			}(),
			setupUsecase: func(out dto.GetPrevOutDto, u *musecase.MockGetPrev) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetPrevInDto(2, nil)).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetPrevOutDto, res *grpc.GetPrevTagResponse, conv *mpresenter.MockToGetPrevConverter) {
				conv.EXPECT().
					ToGetPrevTagsResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.GetPrevTagsRequest{
					Last: 2,
				},
			},
			want: want{
				response: &grpc.GetPrevTagResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "1",
							Name: "happy_path/multiple_tag/multiple_article/not_anymore1",
							Articles: []*grpc.Article{
								{
									Id:    "1",
									Title: "happy_path/multiple_tag/multiple_article/not_anymore1",
								},
							},
						},
					},
				},
			},
		},
		"unhappy_path/usecase_returns_error": {
			setupUsecase: func(out dto.GetPrevOutDto, u *musecase.MockGetPrev) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetPrevInDto(1, nil)).
					Return(nil, errGetPrevTag).
					Times(1)
			},
			setupConverter: func(from dto.GetPrevOutDto, res *grpc.GetPrevTagResponse, conv *mpresenter.MockToGetPrevConverter) {
				conv.EXPECT().
					ToGetPrevTagsResponse(gomock.Any(), gomock.Any()).
					Times(0)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.GetPrevTagsRequest{
					Last: 1,
				},
			},
			want: want{
				response: nil,
				err:      errGetPrevTag,
			},
		},
		"unhappy_path/failed_to_convert": {
			outDto: dto.NewGetPrevOutDto(true),
			setupUsecase: func(out dto.GetPrevOutDto, u *musecase.MockGetPrev) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetPrevInDto(1, nil)).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetPrevOutDto, res *grpc.GetPrevTagResponse, conv *mpresenter.MockToGetPrevConverter) {
				conv.EXPECT().
					ToGetPrevTagsResponse(gomock.Any(), &from).
					Return(nil, false).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: &grpc.GetPrevTagsRequest{
					Last: 1,
				},
			},
			want: want{
				response: nil,
				err:      ErrConversionToGetPrevTagsFailed,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			out := tt.outDto
			u := musecase.NewMockGetPrev(ctrl)
			tt.setupUsecase(out, u)
			response := tt.want.response
			conv := mpresenter.NewMockToGetPrevConverter(ctrl)
			tt.setupConverter(out, response, conv)
			s := NewTagServiceServer(nil, nil, nil, nil, nil, nil, u, conv)
			got, err := s.GetPrevTags(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				if err == nil {
					t.Errorf("GetPrevTags() expected to return an error, but it was nil. want: %+v", err)
					return
				}
				if !errors.Is(err, tt.want.err) {
					t.Errorf("GetPrevTags() error = %v, want %v", err, tt.want.err)
					return
				}
				return
			}
			if diff := cmp.Diff(got, tt.want.response, protocmp.Transform()); diff != "" {
				t.Errorf("GetPrevTags() got = %v, want %v", got, tt.want.response)
			}
		})
	}
}
