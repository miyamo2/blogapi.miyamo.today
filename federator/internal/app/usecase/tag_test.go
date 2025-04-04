package usecase

import (
	grpc "blogapi.miyamo.today/federator/internal/infra/grpc/tag"
	"blogapi.miyamo.today/federator/internal/infra/grpc/tag/tagconnect"
	mtagconnect "blogapi.miyamo.today/federator/internal/mock/infra/grpc/tag/tagconnect"
	"blogapi.miyamo.today/federator/internal/utils"
	"connectrpc.com/connect"
	"context"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"google.golang.org/protobuf/types/known/timestamppb"
	"reflect"
	"testing"

	blogapictx "blogapi.miyamo.today/core/context"
	"blogapi.miyamo.today/federator/internal/app/usecase/dto"
	"github.com/cockroachdb/errors"
	"go.uber.org/mock/gomock"
)

func TestTag_Execute(t *testing.T) {
	type args struct {
		ctx context.Context
		in  dto.TagInDTO
	}
	type want struct {
		out dto.TagOutDTO
		err error
	}
	type testCase struct {
		tagServiceClient func(ctrl *gomock.Controller) tagconnect.TagServiceClient
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
				blogapictx.RequestTypeGRPC,
				nil,
				nil))
	}
	tests := map[string]testCase{
		"happy_path/single_article": {
			tagServiceClient: func(ctrl *gomock.Controller) tagconnect.TagServiceClient {
				tagServiceClient := mtagconnect.NewMockTagServiceClient(ctrl)
				tagServiceClient.EXPECT().
					GetTagById(gomock.Any(), gomock.Any()).
					Return(connect.NewResponse(&grpc.GetTagByIdResponse{
						Tag: &grpc.Tag{
							Id:   "Tag1",
							Name: "Tag1",
							Articles: []*grpc.Article{
								{
									Id:           "Article1",
									Title:        "Article1",
									ThumbnailUrl: "example.com/example.png",
									CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
									UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
								},
							},
						},
					}), nil).
					Times(1)
				return tagServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.NewTagInDTO("Tag1"),
			},
			want: want{
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
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							),
						},
					),
				),
			},
		},
		"happy_path/multiple_article": {
			tagServiceClient: func(ctrl *gomock.Controller) tagconnect.TagServiceClient {
				tagServiceClient := mtagconnect.NewMockTagServiceClient(ctrl)
				tagServiceClient.EXPECT().
					GetTagById(gomock.Any(), gomock.Any()).
					Return(connect.NewResponse(&grpc.GetTagByIdResponse{
						Tag: &grpc.Tag{
							Id:   "Tag1",
							Name: "Tag1",
							Articles: []*grpc.Article{
								{
									Id:           "Article1",
									Title:        "Article1",
									ThumbnailUrl: "example.com/example.png",
									CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
									UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
								},
								{
									Id:           "Article2",
									Title:        "Article2",
									ThumbnailUrl: "example.com/example.png",
									CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
									UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
								},
							},
						},
					}), nil).
					Times(1)
				return tagServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.NewTagInDTO("Tag1"),
			},
			want: want{
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
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							),
							dto.NewArticle(
								"Article2",
								"Article2",
								"",
								utils.MustURLParse("example.com/example.png"),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							),
						},
					),
				),
			},
		},
		"happy_path/no_article": {
			tagServiceClient: func(ctrl *gomock.Controller) tagconnect.TagServiceClient {
				tagServiceClient := mtagconnect.NewMockTagServiceClient(ctrl)
				tagServiceClient.EXPECT().
					GetTagById(gomock.Any(), gomock.Any()).
					Return(connect.NewResponse(&grpc.GetTagByIdResponse{
						Tag: &grpc.Tag{
							Id:   "Tag1",
							Name: "Tag1",
						},
					}), nil).
					Times(1)
				return tagServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.NewTagInDTO("Tag1"),
			},
			want: want{
				out: dto.NewTagOutDTO(
					dto.NewTagArticle(
						"Tag1",
						"Tag1",
						[]dto.Article{},
					),
				),
			},
		},
		"unhappy_path/grpc_returns_error": {
			tagServiceClient: func(ctrl *gomock.Controller) tagconnect.TagServiceClient {
				tagServiceClient := mtagconnect.NewMockTagServiceClient(ctrl)
				tagServiceClient.EXPECT().
					GetTagById(gomock.Any(), gomock.Any()).
					Return(connect.NewResponse(&grpc.GetTagByIdResponse{}), errTestTag).
					Times(1)
				return tagServiceClient
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.NewTagInDTO("Tag1"),
			},
			want: want{
				out: dto.TagOutDTO{},
				err: errTestTag,
			},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			tagServiceClient := tt.tagServiceClient(ctrl)
			u := NewTag(tagServiceClient)
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
