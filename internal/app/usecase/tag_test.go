package usecase

import (
	"context"
	"github.com/cockroachdb/errors"
	blogapictx "github.com/miyamo2/blogapi-core/context"
	"github.com/miyamo2/blogapi/internal/app/usecase/dto"
	mpb "github.com/miyamo2/blogapi/internal/mock/github.com/miyamo2/blogproto-gen/tag/client/pb"
	"github.com/miyamo2/blogproto-gen/tag/client/pb"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
)

func TestTag_Execute(t *testing.T) {
	type args struct {
		ctx context.Context
		in  dto.TagInDto
	}
	type want struct {
		out dto.TagOutDto
		err error
	}
	type testCase struct {
		tagServiceClient func(ctrl *gomock.Controller) pb.TagServiceClient
		args             args
		want             want
		wantErr          bool
	}
	errTestTag := errors.New("test error")
	mockBlogAPIContext := func() context.Context {
		return blogapictx.StoreToContext(
			context.Background(),
			blogapictx.New(
				"1234567890",
				"0987654321",
				"Path",
				blogapictx.RequestTypeGRPC,
				nil,
				nil))
	}
	tests := map[string]testCase{
		"happy_path/single_article": {
			tagServiceClient: func(ctrl *gomock.Controller) pb.TagServiceClient {
				tSvcClt := mpb.NewMockTagServiceClient(ctrl)
				tSvcClt.EXPECT().
					GetTagById(gomock.Any(), gomock.Any()).
					Return(&pb.GetTagByIdResponse{
						Tag: &pb.Tag{
							Id:   "Tag1",
							Name: "Tag1",
							Articles: []*pb.Article{
								{
									Id:           "Article1",
									Title:        "Article1",
									ThumbnailUrl: "example.test",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
					}, nil).
					Times(1)
				return tSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.NewTagInDto("Tag1"),
			},
			want: want{
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
								"2020-01-01T00:00:00Z",
							),
						},
					),
				),
			},
		},
		"happy_path/multiple_article": {
			tagServiceClient: func(ctrl *gomock.Controller) pb.TagServiceClient {
				tSvcClt := mpb.NewMockTagServiceClient(ctrl)
				tSvcClt.EXPECT().
					GetTagById(gomock.Any(), gomock.Any()).
					Return(&pb.GetTagByIdResponse{
						Tag: &pb.Tag{
							Id:   "Tag1",
							Name: "Tag1",
							Articles: []*pb.Article{
								{
									Id:           "Article1",
									Title:        "Article1",
									ThumbnailUrl: "example.test",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
								{
									Id:           "Article2",
									Title:        "Article2",
									ThumbnailUrl: "example.test",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
					}, nil).
					Times(1)
				return tSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.NewTagInDto("Tag1"),
			},
			want: want{
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
								"2020-01-01T00:00:00Z",
							),
							dto.NewArticle(
								"Article2",
								"Article2",
								"",
								"example.test",
								"2020-01-01T00:00:00Z",
								"2020-01-01T00:00:00Z",
							),
						},
					),
				),
			},
		},
		"happy_path/no_article": {
			tagServiceClient: func(ctrl *gomock.Controller) pb.TagServiceClient {
				tSvcClt := mpb.NewMockTagServiceClient(ctrl)
				tSvcClt.EXPECT().
					GetTagById(gomock.Any(), gomock.Any()).
					Return(&pb.GetTagByIdResponse{
						Tag: &pb.Tag{
							Id:   "Tag1",
							Name: "Tag1",
						},
					}, nil).
					Times(1)
				return tSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.NewTagInDto("Tag1"),
			},
			want: want{
				out: dto.NewTagOutDto(
					dto.NewTagArticle(
						"Tag1",
						"Tag1",
						[]dto.Article{},
					),
				),
			},
		},
		"unhappy_path/grpc_returns_error": {
			tagServiceClient: func(ctrl *gomock.Controller) pb.TagServiceClient {
				tSvcClt := mpb.NewMockTagServiceClient(ctrl)
				tSvcClt.EXPECT().
					GetTagById(gomock.Any(), gomock.Any()).
					Return(&pb.GetTagByIdResponse{}, errTestTag).
					Times(1)
				return tSvcClt
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.NewTagInDto("Tag1"),
			},
			want: want{
				out: dto.TagOutDto{},
				err: errTestTag,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			tSvcClt := tt.tagServiceClient(ctrl)
			u := NewTag(tSvcClt)
			got, err := u.Execute(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Execute() expected error but got nil")
					return
				}
				if !errors.Is(err, tt.want.err) {
					t.Errorf("Execute() error = %v, want %v", err, tt.wantErr)
					return
				}
			} else if err != nil {
				t.Errorf("Execute() expected nil but got error = %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want.out) {
				t.Errorf("Execute() got = %v, want %v", got, tt.want.out)
			}
		})
	}
}
