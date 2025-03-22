package pb

import (
	"blogapi.miyamo.today/article-service/internal/infra/grpc"
	"connectrpc.com/connect"
	"context"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"

	"blogapi.miyamo.today/article-service/internal/app/usecase/dto"
	mpresenter "blogapi.miyamo.today/article-service/internal/mock/if-adapter/controller/pb/presenter"
	musecase "blogapi.miyamo.today/article-service/internal/mock/if-adapter/controller/pb/usecase"
	"github.com/cockroachdb/errors"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestArticleServiceServer_GetArticleById(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *connect.Request[grpc.GetArticleByIdRequest]
	}
	type want struct {
		response *connect.Response[grpc.GetArticleByIdResponse]
		err      error
	}
	type testCase struct {
		outDto         dto.GetByIdOutDto
		setupUsecase   func(out dto.GetByIdOutDto, u *musecase.MockGetById)
		setupConverter func(from dto.GetByIdOutDto, res *grpc.GetArticleByIdResponse, conv *mpresenter.MockToGetByIdConverter)
		args           args
		want           want
	}
	errGetArticleById := errors.New("error get article by id")
	tests := map[string]testCase{
		"happy_path/article_has_tag": {
			outDto: dto.NewGetByIdOutDto(
				"1",
				"happy_path/article_has_tag",
				"## happy_path/article_has_tag",
				"1234567890",
				synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
				synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
				[]dto.Tag{
					dto.NewTag("1", "happy_path")}),
			setupUsecase: func(out dto.GetByIdOutDto, u *musecase.MockGetById) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetByIdInDto("1")).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetByIdOutDto, res *grpc.GetArticleByIdResponse, conv *mpresenter.MockToGetByIdConverter) {
				conv.EXPECT().
					ToGetByIdArticlesResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: connect.NewRequest(&grpc.GetArticleByIdRequest{
					Id: "1",
				}),
			},
			want: want{
				response: connect.NewResponse(&grpc.GetArticleByIdResponse{
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
				}),
				err: nil,
			},
		},
		"unhappy_path/usecase_returns_error": {
			setupUsecase: func(out dto.GetByIdOutDto, u *musecase.MockGetById) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetByIdInDto("1")).
					Return(nil, errGetArticleById).Times(1)
			},
			setupConverter: func(from dto.GetByIdOutDto, res *grpc.GetArticleByIdResponse, conv *mpresenter.MockToGetByIdConverter) {
				conv.EXPECT().
					ToGetByIdArticlesResponse(gomock.Any(), gomock.Any()).
					Times(0)
			},
			args: args{
				ctx: context.Background(),
				in: connect.NewRequest(&grpc.GetArticleByIdRequest{
					Id: "1",
				}),
			},
			want: want{
				response: nil,
				err:      errGetArticleById,
			},
		},
		"unhappy_path/failed_to_convert": {
			setupUsecase: func(out dto.GetByIdOutDto, u *musecase.MockGetById) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetByIdInDto("1")).
					Return(&out, nil).Times(1)
			},
			setupConverter: func(from dto.GetByIdOutDto, res *grpc.GetArticleByIdResponse, conv *mpresenter.MockToGetByIdConverter) {
				conv.EXPECT().
					ToGetByIdArticlesResponse(gomock.Any(), &from).
					Return(nil, false).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: connect.NewRequest(&grpc.GetArticleByIdRequest{
					Id: "1",
				}),
			},
			want: want{
				response: nil,
				err:      ErrConversionToGetArticleByIdFailed,
			},
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
			var message *grpc.GetArticleByIdResponse
			if response != nil {
				message = response.Msg
			}
			tt.setupConverter(out, message, conv)
			s := NewArticleServiceServer(u, nil, nil, nil, conv, nil, nil, nil)
			got, err := s.GetArticleById(tt.args.ctx, tt.args.in)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("GetArticleById() error = %v, want %v", err, tt.want.err)
				return
			}
			if diff := cmp.Diff(got, tt.want.response, []cmp.Option{protocmp.Transform(), cmpopts.IgnoreUnexported(connect.Response[grpc.GetArticleByIdResponse]{})}...); diff != "" {
				t.Errorf("GetArticleById() got = %v, want %v", got, tt.want.response)
			}
		})
	}
}

func TestArticleServiceServer_GetAllArticles(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	type want struct {
		response *connect.Response[grpc.GetAllArticlesResponse]
		err      error
	}
	type testCase struct {
		outDto         dto.GetAllOutDto
		setupUsecase   func(out dto.GetAllOutDto, u *musecase.MockGetAll)
		setupConverter func(from dto.GetAllOutDto, res *grpc.GetAllArticlesResponse, conv *mpresenter.MockToGetAllConverter)
		args           args
		want           want
	}
	errGetAllArticles := errors.New("error get all articles")
	tests := map[string]testCase{
		"happy_path/multiple_article": {
			outDto: dto.NewGetAllOutDto(
				[]dto.Article{
					dto.NewArticle(
						"1",
						"happy_path/multiple_article1",
						"## happy_path/multiple_article1",
						"1234567890",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						[]dto.Tag{
							dto.NewTag("tag1", "1"),
							dto.NewTag("tag2", "2"),
						},
					),
					dto.NewArticle(
						"2",
						"happy_path/multiple_article2",
						"## happy_path/multiple_article2",
						"1234567890",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						[]dto.Tag{
							dto.NewTag("tag1", "1"),
							dto.NewTag("tag2", "2"),
						},
					),
				}),
			setupUsecase: func(out dto.GetAllOutDto, u *musecase.MockGetAll) {
				u.EXPECT().
					Execute(gomock.Any()).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetAllOutDto, res *grpc.GetAllArticlesResponse, conv *mpresenter.MockToGetAllConverter) {
				conv.EXPECT().
					ToGetAllArticlesResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
			},
			want: want{
				response: connect.NewResponse(&grpc.GetAllArticlesResponse{
					Articles: []*grpc.Article{
						{
							Id:           "1",
							Title:        "happy_path/multiple_article1",
							Body:         "## happy_path/multiple_article1",
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
						{
							Id:           "2",
							Title:        "happy_path/multiple_article2",
							Body:         "## happy_path/multiple_article2",
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
				}),
				err: nil,
			},
		},
		"happy_path/single_article": {
			outDto: dto.NewGetAllOutDto(
				[]dto.Article{
					dto.NewArticle(
						"1",
						"happy_path/single_article",
						"## happy_path/single_article",
						"1234567890",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						[]dto.Tag{
							dto.NewTag("tag1", "1"),
							dto.NewTag("tag2", "2"),
						},
					),
				}),
			setupUsecase: func(out dto.GetAllOutDto, u *musecase.MockGetAll) {
				u.EXPECT().
					Execute(gomock.Any()).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetAllOutDto, res *grpc.GetAllArticlesResponse, conv *mpresenter.MockToGetAllConverter) {
				conv.EXPECT().
					ToGetAllArticlesResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
			},
			want: want{
				response: connect.NewResponse(&grpc.GetAllArticlesResponse{
					Articles: []*grpc.Article{
						{
							Id:           "1",
							Title:        "happy_path/single_article",
							Body:         "## happy_path/single_article",
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
				}),
				err: nil,
			},
		},
		"happy_path/zero_article": {
			outDto: dto.NewGetAllOutDto([]dto.Article{}),
			setupUsecase: func(out dto.GetAllOutDto, u *musecase.MockGetAll) {
				u.EXPECT().
					Execute(gomock.Any()).
					Return(&out, nil).Times(1)
			},
			setupConverter: func(from dto.GetAllOutDto, res *grpc.GetAllArticlesResponse, conv *mpresenter.MockToGetAllConverter) {
				conv.EXPECT().
					ToGetAllArticlesResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
			},
			want: want{
				response: connect.NewResponse(&grpc.GetAllArticlesResponse{
					Articles: []*grpc.Article{},
				}),
				err: nil,
			},
		},
		"unhappy_path/usecase_returns_error": {
			outDto: dto.GetAllOutDto{},
			setupUsecase: func(out dto.GetAllOutDto, u *musecase.MockGetAll) {
				u.EXPECT().
					Execute(gomock.Any()).
					Return(nil, errGetAllArticles).Times(1)
			},
			setupConverter: func(from dto.GetAllOutDto, res *grpc.GetAllArticlesResponse, conv *mpresenter.MockToGetAllConverter) {
				conv.EXPECT().
					ToGetAllArticlesResponse(gomock.Any(), gomock.Any()).
					Times(0)
			},
			args: args{
				ctx: context.Background(),
			},
			want: want{
				response: nil,
				err:      errGetAllArticles,
			},
		},
		"unhappy_path/failed_to_convert": {
			outDto: dto.GetAllOutDto{},
			setupUsecase: func(out dto.GetAllOutDto, u *musecase.MockGetAll) {
				u.EXPECT().
					Execute(gomock.Any()).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetAllOutDto, res *grpc.GetAllArticlesResponse, conv *mpresenter.MockToGetAllConverter) {
				conv.EXPECT().
					ToGetAllArticlesResponse(gomock.Any(), &from).
					Return(nil, false).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
			},
			want: want{
				response: nil,
				err:      ErrConversionToGetAllArticlesFailed,
			},
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

			var message *grpc.GetAllArticlesResponse
			if response != nil {
				message = response.Msg
			}
			tt.setupConverter(out, message, conv)
			s := NewArticleServiceServer(nil, u, nil, nil, nil, conv, nil, nil)

			got, err := s.GetAllArticles(tt.args.ctx, connect.NewRequest(&emptypb.Empty{}))
			if !errors.Is(err, tt.want.err) {
				t.Errorf("GetAllArticles() error = %v, want %v", err, tt.want.err)
				return
			}
			if diff := cmp.Diff(got, tt.want.response, []cmp.Option{protocmp.Transform(), cmpopts.IgnoreUnexported(connect.Response[grpc.GetAllArticlesResponse]{})}...); diff != "" {
				t.Errorf("GetAllArticles() got = %v, want %v", got, tt.want.response)
			}
		})
	}
}

func TestArticleServiceServer_GetNextArticles(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *connect.Request[grpc.GetNextArticlesRequest]
	}
	type want struct {
		response *connect.Response[grpc.GetNextArticlesResponse]
		err      error
	}
	type testCase struct {
		outDto         dto.GetNextOutDto
		setupUsecase   func(out dto.GetNextOutDto, u *musecase.MockGetNext)
		setupConverter func(from dto.GetNextOutDto, res *grpc.GetNextArticlesResponse, conv *mpresenter.MockToGetNextConverter)
		args           args
		want           want
	}
	cursor := "0"
	pCursor := &cursor
	errGetAllArticles := errors.New("error get all articles")
	tests := map[string]testCase{
		"happy_path/multiple_article": {
			outDto: dto.NewGetNextOutDto(
				[]dto.Article{
					dto.NewArticle(
						"1",
						"happy_path/multiple_article1",
						"## happy_path/multiple_article1",
						"1234567890",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						[]dto.Tag{
							dto.NewTag("tag1", "1"),
							dto.NewTag("tag2", "2"),
						},
					),
					dto.NewArticle(
						"2",
						"happy_path/multiple_article2",
						"## happy_path/multiple_article2",
						"1234567890",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						[]dto.Tag{
							dto.NewTag("tag1", "1"),
							dto.NewTag("tag2", "2"),
						},
					),
				}, true),
			setupUsecase: func(out dto.GetNextOutDto, u *musecase.MockGetNext) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetNextInDto(2, pCursor)).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetNextOutDto, res *grpc.GetNextArticlesResponse, conv *mpresenter.MockToGetNextConverter) {
				conv.EXPECT().
					ToGetNextArticlesResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: connect.NewRequest(&grpc.GetNextArticlesRequest{
					First: 2,
					After: &cursor,
				}),
			},
			want: want{
				response: connect.NewResponse(&grpc.GetNextArticlesResponse{
					Articles: []*grpc.Article{
						{
							Id:           "1",
							Title:        "happy_path/multiple_article1",
							Body:         "## happy_path/multiple_article1",
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
						{
							Id:           "2",
							Title:        "happy_path/multiple_article2",
							Body:         "## happy_path/multiple_article2",
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
					StillExists: true,
				}),
				err: nil,
			},
		},
		"happy_path/single_article": {
			outDto: dto.NewGetNextOutDto(
				[]dto.Article{
					dto.NewArticle(
						"1",
						"happy_path/single_article",
						"## happy_path/single_article",
						"1234567890",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						[]dto.Tag{
							dto.NewTag("tag1", "1"),
							dto.NewTag("tag2", "2"),
						},
					),
				}, true),
			setupUsecase: func(out dto.GetNextOutDto, u *musecase.MockGetNext) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetNextInDto(1, pCursor)).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetNextOutDto, res *grpc.GetNextArticlesResponse, conv *mpresenter.MockToGetNextConverter) {
				conv.EXPECT().
					ToGetNextArticlesResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: connect.NewRequest(&grpc.GetNextArticlesRequest{
					First: 1,
					After: pCursor,
				}),
			},
			want: want{
				response: connect.NewResponse(&grpc.GetNextArticlesResponse{
					Articles: []*grpc.Article{
						{
							Id:           "1",
							Title:        "happy_path/single_article",
							Body:         "## happy_path/single_article",
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
					StillExists: true,
				}),
				err: nil,
			},
		},
		"happy_path/zero_article": {
			outDto: dto.NewGetNextOutDto([]dto.Article{}, false),
			setupUsecase: func(out dto.GetNextOutDto, u *musecase.MockGetNext) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetNextInDto(2, pCursor)).
					Return(&out, nil).Times(1)
			},
			setupConverter: func(from dto.GetNextOutDto, res *grpc.GetNextArticlesResponse, conv *mpresenter.MockToGetNextConverter) {
				conv.EXPECT().
					ToGetNextArticlesResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: connect.NewRequest(&grpc.GetNextArticlesRequest{
					First: 2,
					After: pCursor,
				}),
			},
			want: want{
				response: connect.NewResponse(&grpc.GetNextArticlesResponse{
					Articles:    []*grpc.Article{},
					StillExists: false,
				}),
				err: nil,
			},
		},
		"unhappy_path/usecase_returns_error": {
			outDto: dto.GetNextOutDto{},
			setupUsecase: func(out dto.GetNextOutDto, u *musecase.MockGetNext) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetNextInDto(2, pCursor)).
					Return(nil, errGetAllArticles).Times(1)
			},
			setupConverter: func(from dto.GetNextOutDto, res *grpc.GetNextArticlesResponse, conv *mpresenter.MockToGetNextConverter) {
				conv.EXPECT().
					ToGetNextArticlesResponse(gomock.Any(), gomock.Any()).
					Times(0)
			},
			args: args{
				ctx: context.Background(),
				in: connect.NewRequest(&grpc.GetNextArticlesRequest{
					First: 2,
					After: pCursor,
				}),
			},
			want: want{
				response: nil,
				err:      errGetAllArticles,
			},
		},
		"unhappy_path/failed_to_convert": {
			outDto: dto.GetNextOutDto{},
			setupUsecase: func(out dto.GetNextOutDto, u *musecase.MockGetNext) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetNextInDto(1, pCursor)).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetNextOutDto, res *grpc.GetNextArticlesResponse, conv *mpresenter.MockToGetNextConverter) {
				conv.EXPECT().
					ToGetNextArticlesResponse(gomock.Any(), &from).
					Return(nil, false).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: connect.NewRequest(&grpc.GetNextArticlesRequest{
					First: 1,
					After: pCursor,
				}),
			},
			want: want{
				response: nil,
				err:      ErrConversionToGetNextArticlesFailed,
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

			var message *grpc.GetNextArticlesResponse
			if response != nil {
				message = response.Msg
			}
			tt.setupConverter(out, message, conv)
			s := NewArticleServiceServer(nil, nil, u, nil, nil, nil, conv, nil)

			got, err := s.GetNextArticles(tt.args.ctx, tt.args.in)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("GetNextArticles() error = %v, want %v", err, tt.want.err)
				return
			}
			if diff := cmp.Diff(got, tt.want.response, []cmp.Option{protocmp.Transform(), cmpopts.IgnoreUnexported(connect.Response[grpc.GetNextArticlesResponse]{})}...); diff != "" {
				t.Errorf("GetNextArticles() got = %v, want %v", got, tt.want.response)
			}
		})
	}
}

func TestArticleServiceServer_GetPrevArticles(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *connect.Request[grpc.GetPrevArticlesRequest]
	}
	type want struct {
		response *connect.Response[grpc.GetPrevArticlesResponse]
		err      error
	}
	type testCase struct {
		outDto         dto.GetPrevOutDto
		setupUsecase   func(out dto.GetPrevOutDto, u *musecase.MockGetPrev)
		setupConverter func(from dto.GetPrevOutDto, res *grpc.GetPrevArticlesResponse, conv *mpresenter.MockToGetPrevConverter)
		args           args
		want           want
	}
	cursor := "0"
	pCursor := &cursor
	errGetAllArticles := errors.New("error get all articles")
	tests := map[string]testCase{
		"happy_path/multiple_article": {
			outDto: dto.NewGetPrevOutDto(
				[]dto.Article{
					dto.NewArticle(
						"1",
						"happy_path/multiple_article1",
						"## happy_path/multiple_article1",
						"1234567890",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						[]dto.Tag{
							dto.NewTag("tag1", "1"),
							dto.NewTag("tag2", "2"),
						},
					),
					dto.NewArticle(
						"2",
						"happy_path/multiple_article2",
						"## happy_path/multiple_article2",
						"1234567890",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						[]dto.Tag{
							dto.NewTag("tag1", "1"),
							dto.NewTag("tag2", "2"),
						},
					),
				}, true),
			setupUsecase: func(out dto.GetPrevOutDto, u *musecase.MockGetPrev) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetPrevInDto(2, pCursor)).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetPrevOutDto, res *grpc.GetPrevArticlesResponse, conv *mpresenter.MockToGetPrevConverter) {
				conv.EXPECT().
					ToGetPrevArticlesResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: connect.NewRequest(&grpc.GetPrevArticlesRequest{
					Last:   2,
					Before: &cursor,
				}),
			},
			want: want{
				response: connect.NewResponse(&grpc.GetPrevArticlesResponse{
					Articles: []*grpc.Article{
						{
							Id:           "1",
							Title:        "happy_path/multiple_article1",
							Body:         "## happy_path/multiple_article1",
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
						{
							Id:           "2",
							Title:        "happy_path/multiple_article2",
							Body:         "## happy_path/multiple_article2",
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
					StillExists: true,
				}),
				err: nil,
			},
		},
		"happy_path/single_article": {
			outDto: dto.NewGetPrevOutDto(
				[]dto.Article{
					dto.NewArticle(
						"1",
						"happy_path/single_article",
						"## happy_path/single_article",
						"1234567890",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						[]dto.Tag{
							dto.NewTag("tag1", "1"),
							dto.NewTag("tag2", "2"),
						},
					),
				}, true),
			setupUsecase: func(out dto.GetPrevOutDto, u *musecase.MockGetPrev) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetPrevInDto(1, pCursor)).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetPrevOutDto, res *grpc.GetPrevArticlesResponse, conv *mpresenter.MockToGetPrevConverter) {
				conv.EXPECT().
					ToGetPrevArticlesResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: connect.NewRequest(&grpc.GetPrevArticlesRequest{
					Last:   1,
					Before: pCursor,
				}),
			},
			want: want{
				response: connect.NewResponse(&grpc.GetPrevArticlesResponse{
					Articles: []*grpc.Article{
						{
							Id:           "1",
							Title:        "happy_path/single_article",
							Body:         "## happy_path/single_article",
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
					StillExists: true,
				}),
				err: nil,
			},
		},
		"happy_path/zero_article": {
			outDto: dto.NewGetPrevOutDto([]dto.Article{}, false),
			setupUsecase: func(out dto.GetPrevOutDto, u *musecase.MockGetPrev) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetPrevInDto(2, pCursor)).
					Return(&out, nil).Times(1)
			},
			setupConverter: func(from dto.GetPrevOutDto, res *grpc.GetPrevArticlesResponse, conv *mpresenter.MockToGetPrevConverter) {
				conv.EXPECT().
					ToGetPrevArticlesResponse(gomock.Any(), &from).
					Return(res, true).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: connect.NewRequest(&grpc.GetPrevArticlesRequest{
					Last:   2,
					Before: pCursor,
				}),
			},
			want: want{
				response: connect.NewResponse(&grpc.GetPrevArticlesResponse{
					Articles:    []*grpc.Article{},
					StillExists: false,
				}),
				err: nil,
			},
		},
		"unhappy_path/usecase_returns_error": {
			outDto: dto.GetPrevOutDto{},
			setupUsecase: func(out dto.GetPrevOutDto, u *musecase.MockGetPrev) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetPrevInDto(2, pCursor)).
					Return(nil, errGetAllArticles).Times(1)
			},
			setupConverter: func(from dto.GetPrevOutDto, res *grpc.GetPrevArticlesResponse, conv *mpresenter.MockToGetPrevConverter) {
				conv.EXPECT().
					ToGetPrevArticlesResponse(gomock.Any(), gomock.Any()).
					Times(0)
			},
			args: args{
				ctx: context.Background(),
				in: connect.NewRequest(&grpc.GetPrevArticlesRequest{
					Last:   2,
					Before: pCursor,
				}),
			},
			want: want{
				response: nil,
				err:      errGetAllArticles,
			},
		},
		"unhappy_path/failed_to_convert": {
			outDto: dto.GetPrevOutDto{},
			setupUsecase: func(out dto.GetPrevOutDto, u *musecase.MockGetPrev) {
				u.EXPECT().
					Execute(gomock.Any(), dto.NewGetPrevInDto(1, pCursor)).
					Return(&out, nil).
					Times(1)
			},
			setupConverter: func(from dto.GetPrevOutDto, res *grpc.GetPrevArticlesResponse, conv *mpresenter.MockToGetPrevConverter) {
				conv.EXPECT().
					ToGetPrevArticlesResponse(gomock.Any(), &from).
					Return(nil, false).
					Times(1)
			},
			args: args{
				ctx: context.Background(),
				in: connect.NewRequest(&grpc.GetPrevArticlesRequest{
					Last:   1,
					Before: pCursor,
				}),
			},
			want: want{
				response: nil,
				err:      ErrConversionToGetPrevArticlesFailed,
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

			var message *grpc.GetPrevArticlesResponse
			if response != nil {
				message = response.Msg
			}
			tt.setupConverter(out, message, conv)
			s := NewArticleServiceServer(nil, nil, nil, u, nil, nil, nil, conv)

			got, err := s.GetPrevArticles(tt.args.ctx, tt.args.in)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("GetPrevArticles() error = %v, want %v", err, tt.want.err)
				return
			}
			if diff := cmp.Diff(got, tt.want.response, []cmp.Option{protocmp.Transform(), cmpopts.IgnoreUnexported(connect.Response[grpc.GetPrevArticlesResponse]{})}...); diff != "" {
				t.Errorf("GetPrevArticles() got = %v, want %v", got, tt.want.response)
			}
		})
	}
}
