package pb

import (
	"connectrpc.com/connect"
	"context"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"

	"blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/tag-service/internal/infra/grpc"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestConverter_ToGetByIdTagResponse(t *testing.T) {
	type args struct {
		ctx  context.Context
		from func() *dto.GetByIdOutDto
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
				from: func() *dto.GetByIdOutDto {
					o := dto.NewGetByIdOutDto(
						"tag1", "1",
						[]dto.Article{
							dto.NewArticle(
								"1",
								"happy_path",
								"1234567890",
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							),
						},
					)
					return &o
				},
			},
			want: want{
				result: connect.NewResponse(&grpc.GetTagByIdResponse{
					Tag: &grpc.Tag{
						Id:   "tag1",
						Name: "1",
						Articles: []*grpc.Article{
							{
								Id:           "1",
								Title:        "happy_path",
								ThumbnailUrl: "1234567890",
								CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
								UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
							},
						},
					},
				}),
				ok: true,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := NewConverter()
			got, ok := c.ToGetByIdTagResponse(tt.args.ctx, tt.args.from())
			if tt.want.ok != ok {
				t.Errorf("ToGetByIdTagResponse() ok = %v, want %v", ok, tt.want.ok)
			}
			if diff := cmp.Diff(got.Msg, tt.want.result.Msg, protocmp.Transform(), cmpopts.IgnoreUnexported()); diff != "" {
				t.Errorf("ToGetByIdTagResponse() = %v, want %v", got.Msg, tt.want.result.Msg)
			}
		})
	}
}

func TestConverter_ToGetAllTagsResponse(t *testing.T) {
	type args struct {
		ctx  context.Context
		from func() *dto.GetAllOutDto
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
				from: func() *dto.GetAllOutDto {
					o := dto.NewGetAllOutDto()
					o = o.WithTagDto(
						dto.NewTag(
							"tag1", "1",
							[]dto.Article{
								dto.NewArticle(
									"1",
									"happy_path",
									"1234567890",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								),
							}),
					)
					o = o.WithTagDto(
						dto.NewTag(
							"tag2", "2",
							[]dto.Article{
								dto.NewArticle(
									"1",
									"happy_path",
									"1234567890",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								),
							}),
					)
					return &o
				},
			},
			want: want{
				result: connect.NewResponse(&grpc.GetAllTagsResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "tag1",
							Name: "1",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path",
									ThumbnailUrl: "1234567890",
									CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
									UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
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
									CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
									UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
								},
							},
						},
					},
				}),
				ok: true,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := NewConverter()
			got, ok := c.ToGetAllTagsResponse(tt.args.ctx, tt.args.from())
			if tt.want.ok != ok {
				t.Errorf("ToGetByIdTagResponse() ok = %v, want %v", ok, tt.want.ok)
			}
			if diff := cmp.Diff(got.Msg, tt.want.result.Msg, protocmp.Transform(), cmpopts.IgnoreUnexported()); diff != "" {
				t.Errorf("ToGetByIdTagResponse() = %v, want %v", got.Msg, tt.want.result.Msg)
			}
		})
	}
}

func TestConverter_ToGetNextTagsResponse(t *testing.T) {
	type args struct {
		ctx  context.Context
		from func() *dto.GetNextOutDto
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
				from: func() *dto.GetNextOutDto {
					o := dto.NewGetNextOutDto(true)
					o = o.WithTagDto(
						dto.NewTag(
							"tag1", "1",
							[]dto.Article{
								dto.NewArticle(
									"1",
									"happy_path/still_exists",
									"1234567890",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								),
							}),
					)
					o = o.WithTagDto(
						dto.NewTag(
							"tag2", "2",
							[]dto.Article{
								dto.NewArticle(
									"1",
									"happy_path/still_exists",
									"1234567890",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								),
							}),
					)
					return &o
				},
			},
			want: want{
				result: connect.NewResponse(&grpc.GetNextTagResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "tag1",
							Name: "1",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/still_exists",
									ThumbnailUrl: "1234567890",
									CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
									UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
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
									CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
									UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
								},
							},
						},
					},
					StillExists: true,
				}),
				ok: true,
			},
		},
		"happy_path/not_anymore": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.GetNextOutDto {
					o := dto.NewGetNextOutDto(false)
					o = o.WithTagDto(
						dto.NewTag(
							"tag1", "1",
							[]dto.Article{
								dto.NewArticle(
									"1",
									"happy_path/not_anymore",
									"1234567890",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								),
							}),
					)
					o = o.WithTagDto(
						dto.NewTag(
							"tag2", "2",
							[]dto.Article{
								dto.NewArticle(
									"1",
									"happy_path/not_anymore",
									"1234567890",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								),
							}),
					)
					return &o
				},
			},
			want: want{
				result: connect.NewResponse(&grpc.GetNextTagResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "tag1",
							Name: "1",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/not_anymore",
									ThumbnailUrl: "1234567890",
									CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
									UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
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
									CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
									UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
								},
							},
						},
					},
				}),
				ok: true,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := NewConverter()
			got, ok := c.ToGetNextTagsResponse(tt.args.ctx, tt.args.from())
			if tt.want.ok != ok {
				t.Errorf("ToGetNextTagsResponse() ok = %v, want %v", ok, tt.want.ok)
			}
			if diff := cmp.Diff(got.Msg, tt.want.result.Msg, protocmp.Transform(), cmpopts.IgnoreUnexported()); diff != "" {
				t.Errorf("ToGetNextTagsResponse() = %v, want %v", got.Msg, tt.want.result.Msg)
			}
		})
	}
}

func TestConverter_ToGetPrevTagsResponse(t *testing.T) {
	type args struct {
		ctx  context.Context
		from func() *dto.GetPrevOutDto
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
				from: func() *dto.GetPrevOutDto {
					o := dto.NewGetPrevOutDto(true)
					o = o.WithTagDto(
						dto.NewTag(
							"tag1", "1",
							[]dto.Article{
								dto.NewArticle(
									"1",
									"happy_path/still_exists",
									"1234567890",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								),
							}),
					)
					o = o.WithTagDto(
						dto.NewTag(
							"tag2", "2",
							[]dto.Article{
								dto.NewArticle(
									"1",
									"happy_path/still_exists",
									"1234567890",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								),
							}),
					)
					return &o
				},
			},
			want: want{
				result: connect.NewResponse(&grpc.GetPrevTagResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "tag1",
							Name: "1",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/still_exists",
									ThumbnailUrl: "1234567890",
									CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
									UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
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
									CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
									UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
								},
							},
						},
					},
					StillExists: true,
				}),
				ok: true,
			},
		},
		"happy_path/not_anymore": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.GetPrevOutDto {
					o := dto.NewGetPrevOutDto(false)
					o = o.WithTagDto(
						dto.NewTag(
							"tag1", "1",
							[]dto.Article{
								dto.NewArticle(
									"1",
									"happy_path/not_anymore",
									"1234567890",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								),
							}),
					)
					o = o.WithTagDto(
						dto.NewTag(
							"tag2", "2",
							[]dto.Article{
								dto.NewArticle(
									"1",
									"happy_path/not_anymore",
									"1234567890",
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
									synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
								),
							}),
					)
					return &o
				},
			},
			want: want{
				result: connect.NewResponse(&grpc.GetPrevTagResponse{
					Tags: []*grpc.Tag{
						{
							Id:   "tag1",
							Name: "1",
							Articles: []*grpc.Article{
								{
									Id:           "1",
									Title:        "happy_path/not_anymore",
									ThumbnailUrl: "1234567890",
									CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
									UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
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
									CreatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
									UpdatedAt:    timestamppb.New(synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0).StdTime()),
								},
							},
						},
					},
				}),
				ok: true,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := NewConverter()
			got, ok := c.ToGetPrevTagsResponse(tt.args.ctx, tt.args.from())
			if tt.want.ok != ok {
				t.Errorf("ToGetPrevTagsResponse() ok = %v, want %v", ok, tt.want.ok)
			}
			if diff := cmp.Diff(got.Msg, tt.want.result.Msg, protocmp.Transform(), cmpopts.IgnoreUnexported()); diff != "" {
				t.Errorf("ToGetPrevTagsResponse() = %v, want %v", got.Msg, tt.want.result.Msg)
			}
		})
	}
}
