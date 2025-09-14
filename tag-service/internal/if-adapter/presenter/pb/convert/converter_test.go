package convert

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/protobuf/types/known/timestamppb"

	"blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/tag-service/internal/infra/grpc"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestGetByIdTag_ToResponse(t *testing.T) {
	type args struct {
		ctx  context.Context
		from func() *dto.GetByIdOutput
	}
	type want struct {
		result *connect.Response[grpc.GetTagByIdResponse]
		ok     bool
	}
	type testCase struct {
		args args
		want want
	}
	tests := map[string]testCase{
		"happy_path": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.GetByIdOutput {
					o := dto.NewTag(
						"tag1", "1",
						dto.NewArticle(
							"1",
							"happy_path",
							"1234567890",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						),
					)
					return &o
				},
			},
			want: want{
				result: connect.NewResponse(
					&grpc.GetTagByIdResponse{
						Tag: &grpc.Tag{
							Id:   "tag1",
							Name: "1",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path",
									ThumbnailUrl: "1234567890",
									CreatedAt: timestamppb.New(
										synchro.New[tz.UTC](
											2020,
											1,
											1,
											0,
											0,
											0,
											0,
										).StdTime(),
									),
									UpdatedAt: timestamppb.New(
										synchro.New[tz.UTC](
											2020,
											1,
											1,
											0,
											0,
											0,
											0,
										).StdTime(),
									),
								},
							},
						},
					},
				),
				ok: true,
			},
		},
	}
	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				c := NewGetByIdTag()
				got, ok := c.ToResponse(tt.args.ctx, tt.args.from())
				if tt.want.ok != ok {
					t.Errorf("ToResponse() ok = %v, want %v", ok, tt.want.ok)
				}
				if diff := cmp.Diff(
					got.Msg,
					tt.want.result.Msg,
					protocmp.Transform(),
					cmpopts.IgnoreUnexported(),
				); diff != "" {
					t.Errorf("ToResponse() = %v, want %v", got.Msg, tt.want.result.Msg)
				}
			},
		)
	}
}

func TestGetAllTags_ToResponse(t *testing.T) {
	type args struct {
		ctx  context.Context
		from func() *dto.ListAllOutput
	}
	type want struct {
		result *connect.Response[grpc.GetAllTagsResponse]
		ok     bool
	}
	type testCase struct {
		args args
		want want
	}
	tests := map[string]testCase{
		"happy_path": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.ListAllOutput {
					v := dto.NewListAllOutput(
						dto.NewTag(
							"tag1", "1",
							dto.NewArticle(
								"1",
								"happy_path",
								"1234567890",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							),
						),
						dto.NewTag(
							"tag2", "2",
							dto.NewArticle(
								"1",
								"happy_path",
								"1234567890",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							),
						),
					)
					return &v
				},
			},
			want: want{
				result: connect.NewResponse(
					&grpc.GetAllTagsResponse{
						Tags: []*grpc.Tag{
							{
								Id:   "tag1",
								Name: "1",
								Articles: []*grpc.Article{
									{
										Id:           "1",
										Title:        "happy_path",
										ThumbnailUrl: "1234567890",
										CreatedAt: timestamppb.New(
											synchro.New[tz.UTC](
												2020,
												1,
												1,
												0,
												0,
												0,
												0,
											).StdTime(),
										),
										UpdatedAt: timestamppb.New(
											synchro.New[tz.UTC](
												2020,
												1,
												1,
												0,
												0,
												0,
												0,
											).StdTime(),
										),
									},
								},
							},
							{
								Id:   "tag2",
								Name: "2",
								Articles: []*grpc.Article{
									{
										Id:           "1",
										Title:        "happy_path",
										ThumbnailUrl: "1234567890",
										CreatedAt: timestamppb.New(
											synchro.New[tz.UTC](
												2020,
												1,
												1,
												0,
												0,
												0,
												0,
											).StdTime(),
										),
										UpdatedAt: timestamppb.New(
											synchro.New[tz.UTC](
												2020,
												1,
												1,
												0,
												0,
												0,
												0,
											).StdTime(),
										),
									},
								},
							},
						},
					},
				),
				ok: true,
			},
		},
	}
	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				c := NewGetAllTags()
				got, ok := c.ToResponse(tt.args.ctx, tt.args.from())
				if tt.want.ok != ok {
					t.Errorf("ToResponse() ok = %v, want %v", ok, tt.want.ok)
				}
				if diff := cmp.Diff(
					got.Msg,
					tt.want.result.Msg,
					protocmp.Transform(),
					cmpopts.IgnoreUnexported(),
				); diff != "" {
					t.Errorf("ToResponse() = %v, want %v", got.Msg, tt.want.result.Msg)
				}
			},
		)
	}
}

func TestGetNextTags_ToResponse(t *testing.T) {
	type args struct {
		ctx  context.Context
		from func() *dto.ListAfterOutput
	}
	type want struct {
		result *connect.Response[grpc.GetNextTagResponse]
		ok     bool
	}
	type testCase struct {
		args args
		want want
	}

	tests := map[string]testCase{
		"happy_path/still_exists": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.ListAfterOutput {
					o := dto.NewListAfterOutput(
						true,
						dto.NewTag(
							"tag1", "1",
							dto.NewArticle(
								"1",
								"happy_path/still_exists",
								"1234567890",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							),
						),
						dto.NewTag(
							"tag2", "2",
							dto.NewArticle(
								"1",
								"happy_path/still_exists",
								"1234567890",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							),
						),
					)
					return &o
				},
			},
			want: want{
				result: connect.NewResponse(
					&grpc.GetNextTagResponse{
						Tags: []*grpc.Tag{
							{
								Id:   "tag1",
								Name: "1",
								Articles: []*grpc.Article{
									{
										Id:           "1",
										Title:        "happy_path/still_exists",
										ThumbnailUrl: "1234567890",
										CreatedAt: timestamppb.New(
											synchro.New[tz.UTC](
												2020,
												1,
												1,
												0,
												0,
												0,
												0,
											).StdTime(),
										),
										UpdatedAt: timestamppb.New(
											synchro.New[tz.UTC](
												2020,
												1,
												1,
												0,
												0,
												0,
												0,
											).StdTime(),
										),
									},
								},
							},
							{
								Id:   "tag2",
								Name: "2",
								Articles: []*grpc.Article{
									{
										Id:           "1",
										Title:        "happy_path/still_exists",
										ThumbnailUrl: "1234567890",
										CreatedAt: timestamppb.New(
											synchro.New[tz.UTC](
												2020,
												1,
												1,
												0,
												0,
												0,
												0,
											).StdTime(),
										),
										UpdatedAt: timestamppb.New(
											synchro.New[tz.UTC](
												2020,
												1,
												1,
												0,
												0,
												0,
												0,
											).StdTime(),
										),
									},
								},
							},
						},
						StillExists: true,
					},
				),
				ok: true,
			},
		},
		"happy_path/not_anymore": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.ListAfterOutput {
					o := dto.NewListAfterOutput(
						false,
						dto.NewTag(
							"tag1", "1",
							dto.NewArticle(
								"1",
								"happy_path/not_anymore",
								"1234567890",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							),
						),
						dto.NewTag(
							"tag2", "2",
							dto.NewArticle(
								"1",
								"happy_path/not_anymore",
								"1234567890",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							),
						),
					)
					return &o
				},
			},
			want: want{
				result: connect.NewResponse(
					&grpc.GetNextTagResponse{
						Tags: []*grpc.Tag{
							{
								Id:   "tag1",
								Name: "1",
								Articles: []*grpc.Article{
									{
										Id:           "1",
										Title:        "happy_path/not_anymore",
										ThumbnailUrl: "1234567890",
										CreatedAt: timestamppb.New(
											synchro.New[tz.UTC](
												2020,
												1,
												1,
												0,
												0,
												0,
												0,
											).StdTime(),
										),
										UpdatedAt: timestamppb.New(
											synchro.New[tz.UTC](
												2020,
												1,
												1,
												0,
												0,
												0,
												0,
											).StdTime(),
										),
									},
								},
							},
							{
								Id:   "tag2",
								Name: "2",
								Articles: []*grpc.Article{
									{
										Id:           "1",
										Title:        "happy_path/not_anymore",
										ThumbnailUrl: "1234567890",
										CreatedAt: timestamppb.New(
											synchro.New[tz.UTC](
												2020,
												1,
												1,
												0,
												0,
												0,
												0,
											).StdTime(),
										),
										UpdatedAt: timestamppb.New(
											synchro.New[tz.UTC](
												2020,
												1,
												1,
												0,
												0,
												0,
												0,
											).StdTime(),
										),
									},
								},
							},
						},
					},
				),
				ok: true,
			},
		},
	}
	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				c := NewGetNextTags()
				got, ok := c.ToResponse(tt.args.ctx, tt.args.from())
				if tt.want.ok != ok {
					t.Errorf("ToResponse() ok = %v, want %v", ok, tt.want.ok)
				}
				if diff := cmp.Diff(
					got.Msg,
					tt.want.result.Msg,
					protocmp.Transform(),
					cmpopts.IgnoreUnexported(),
				); diff != "" {
					t.Errorf("ToResponse() = %v, want %v", got.Msg, tt.want.result.Msg)
				}
			},
		)
	}
}

func TestGetPrevTags_ToResponse(t *testing.T) {
	type args struct {
		ctx  context.Context
		from func() *dto.ListBeforeOutput
	}
	type want struct {
		result *connect.Response[grpc.GetPrevTagResponse]
		ok     bool
	}
	type testCase struct {
		args args
		want want
	}

	tests := map[string]testCase{
		"happy_path/still_exists": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.ListBeforeOutput {
					o := dto.NewListBeforeOutput(
						true,
						dto.NewTag(
							"tag1", "1",
							dto.NewArticle(
								"1",
								"happy_path/still_exists",
								"1234567890",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							),
						),
						dto.NewTag(
							"tag2", "2",
							dto.NewArticle(
								"1",
								"happy_path/still_exists",
								"1234567890",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							),
						),
					)
					return &o
				},
			},
			want: want{
				result: connect.NewResponse(
					&grpc.GetPrevTagResponse{
						Tags: []*grpc.Tag{
							{
								Id:   "tag1",
								Name: "1",
								Articles: []*grpc.Article{
									{
										Id:           "1",
										Title:        "happy_path/still_exists",
										ThumbnailUrl: "1234567890",
										CreatedAt: timestamppb.New(
											synchro.New[tz.UTC](
												2020,
												1,
												1,
												0,
												0,
												0,
												0,
											).StdTime(),
										),
										UpdatedAt: timestamppb.New(
											synchro.New[tz.UTC](
												2020,
												1,
												1,
												0,
												0,
												0,
												0,
											).StdTime(),
										),
									},
								},
							},
							{
								Id:   "tag2",
								Name: "2",
								Articles: []*grpc.Article{
									{
										Id:           "1",
										Title:        "happy_path/still_exists",
										ThumbnailUrl: "1234567890",
										CreatedAt: timestamppb.New(
											synchro.New[tz.UTC](
												2020,
												1,
												1,
												0,
												0,
												0,
												0,
											).StdTime(),
										),
										UpdatedAt: timestamppb.New(
											synchro.New[tz.UTC](
												2020,
												1,
												1,
												0,
												0,
												0,
												0,
											).StdTime(),
										),
									},
								},
							},
						},
						StillExists: true,
					},
				),
				ok: true,
			},
		},
		"happy_path/not_anymore": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.ListBeforeOutput {
					o := dto.NewListBeforeOutput(
						false,
						dto.NewTag(
							"tag1", "1",
							dto.NewArticle(
								"1",
								"happy_path/not_anymore",
								"1234567890",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							),
						),
						dto.NewTag(
							"tag2", "2",
							dto.NewArticle(
								"1",
								"happy_path/not_anymore",
								"1234567890",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							),
						),
					)
					return &o
				},
			},
			want: want{
				result: connect.NewResponse(
					&grpc.GetPrevTagResponse{
						Tags: []*grpc.Tag{
							{
								Id:   "tag1",
								Name: "1",
								Articles: []*grpc.Article{
									{
										Id:           "1",
										Title:        "happy_path/not_anymore",
										ThumbnailUrl: "1234567890",
										CreatedAt: timestamppb.New(
											synchro.New[tz.UTC](
												2020,
												1,
												1,
												0,
												0,
												0,
												0,
											).StdTime(),
										),
										UpdatedAt: timestamppb.New(
											synchro.New[tz.UTC](
												2020,
												1,
												1,
												0,
												0,
												0,
												0,
											).StdTime(),
										),
									},
								},
							},
							{
								Id:   "tag2",
								Name: "2",
								Articles: []*grpc.Article{
									{
										Id:           "1",
										Title:        "happy_path/not_anymore",
										ThumbnailUrl: "1234567890",
										CreatedAt: timestamppb.New(
											synchro.New[tz.UTC](
												2020,
												1,
												1,
												0,
												0,
												0,
												0,
											).StdTime(),
										),
										UpdatedAt: timestamppb.New(
											synchro.New[tz.UTC](
												2020,
												1,
												1,
												0,
												0,
												0,
												0,
											).StdTime(),
										),
									},
								},
							},
						},
					},
				),
				ok: true,
			},
		},
	}
	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				c := NewGetPrevTags()
				got, ok := c.ToResponse(tt.args.ctx, tt.args.from())
				if tt.want.ok != ok {
					t.Errorf("ToResponse() ok = %v, want %v", ok, tt.want.ok)
				}
				if diff := cmp.Diff(
					got.Msg,
					tt.want.result.Msg,
					protocmp.Transform(),
					cmpopts.IgnoreUnexported(),
				); diff != "" {
					t.Errorf("ToResponse() = %v, want %v", got.Msg, tt.want.result.Msg)
				}
			},
		)
	}
}
