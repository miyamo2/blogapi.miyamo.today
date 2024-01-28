package pb

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/miyamo2/blogapi-tag-service/internal/app/usecase/dto"
	"github.com/miyamo2/blogproto-gen/tag/server/pb"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestConverter_ToGetByIdTagResponse(t *testing.T) {
	type args struct {
		ctx  context.Context
		from func() *dto.GetByIdOutDto
	}
	type want struct {
		result *pb.GetTagByIdResponse
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
								"2020-01-01T00:00:00Z",
								"2020-01-01T00:00:00Z",
							),
						},
					)
					return &o
				},
			},
			want: want{
				result: &pb.GetTagByIdResponse{
					Tag: &pb.Tag{
						Id:   "tag1",
						Name: "1",
						Articles: []*pb.Article{
							{
								Id:           "1",
								Title:        "happy_path",
								ThumbnailUrl: "1234567890",
								CreatedAt:    "2020-01-01T00:00:00Z",
								UpdatedAt:    "2020-01-01T00:00:00Z",
							},
						},
					},
				},
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
			if diff := cmp.Diff(got, tt.want.result, protocmp.Transform()); diff != "" {
				t.Errorf("ToGetByIdTagResponse() = %v, want %v", got, tt.want.result)
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
		result *pb.GetAllTagsResponse
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
									"2020-01-01T00:00:00Z",
									"2020-01-01T00:00:00Z",
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
									"2020-01-01T00:00:00Z",
									"2020-01-01T00:00:00Z",
								),
							}),
					)
					return &o
				},
			},
			want: want{
				result: &pb.GetAllTagsResponse{
					Tags: []*pb.Tag{
						{
							Id:   "tag1",
							Name: "1",
							Articles: []*pb.Article{
								{
									Id:           "1",
									Title:        "happy_path",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
						{
							Id:   "tag2",
							Name: "2",
							Articles: []*pb.Article{
								{
									Id:           "1",
									Title:        "happy_path",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
					},
				},
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
			if diff := cmp.Diff(got, tt.want.result, protocmp.Transform()); diff != "" {
				t.Errorf("ToGetByIdTagResponse() = %v, want %v", got, tt.want.result)
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
		result *pb.GetNextTagResponse
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
									"2020-01-01T00:00:00Z",
									"2020-01-01T00:00:00Z",
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
									"2020-01-01T00:00:00Z",
									"2020-01-01T00:00:00Z",
								),
							}),
					)
					return &o
				},
			},
			want: want{
				result: &pb.GetNextTagResponse{
					Tags: []*pb.Tag{
						{
							Id:   "tag1",
							Name: "1",
							Articles: []*pb.Article{
								{
									Id:           "1",
									Title:        "happy_path/still_exists",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
						{
							Id:   "tag2",
							Name: "2",
							Articles: []*pb.Article{
								{
									Id:           "1",
									Title:        "happy_path/still_exists",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
					},
					StillExists: true,
				},
				ok: true,
			},
		},
		"happy_path/not_anymore": {
			args: args{
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
									"2020-01-01T00:00:00Z",
									"2020-01-01T00:00:00Z",
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
									"2020-01-01T00:00:00Z",
									"2020-01-01T00:00:00Z",
								),
							}),
					)
					return &o
				},
			},
			want: want{
				result: &pb.GetNextTagResponse{
					Tags: []*pb.Tag{
						{
							Id:   "tag1",
							Name: "1",
							Articles: []*pb.Article{
								{
									Id:           "1",
									Title:        "happy_path/not_anymore",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
						{
							Id:   "tag2",
							Name: "2",
							Articles: []*pb.Article{
								{
									Id:           "1",
									Title:        "happy_path/not_anymore",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
					},
				},
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
			if diff := cmp.Diff(got, tt.want.result, protocmp.Transform()); diff != "" {
				t.Errorf("ToGetNextTagsResponse() = %v, want %v", got, tt.want.result)
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
		result *pb.GetPrevTagResponse
		ok     bool
	}
	type testCase struct {
		args args
		want want
	}

	tests := map[string]testCase{
		"happy_path/still_exists": {
			args: args{
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
									"2020-01-01T00:00:00Z",
									"2020-01-01T00:00:00Z",
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
									"2020-01-01T00:00:00Z",
									"2020-01-01T00:00:00Z",
								),
							}),
					)
					return &o
				},
			},
			want: want{
				result: &pb.GetPrevTagResponse{
					Tags: []*pb.Tag{
						{
							Id:   "tag1",
							Name: "1",
							Articles: []*pb.Article{
								{
									Id:           "1",
									Title:        "happy_path/still_exists",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
						{
							Id:   "tag2",
							Name: "2",
							Articles: []*pb.Article{
								{
									Id:           "1",
									Title:        "happy_path/still_exists",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
					},
					StillExists: true,
				},
				ok: true,
			},
		},
		"happy_path/not_anymore": {
			args: args{
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
									"2020-01-01T00:00:00Z",
									"2020-01-01T00:00:00Z",
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
									"2020-01-01T00:00:00Z",
									"2020-01-01T00:00:00Z",
								),
							}),
					)
					return &o
				},
			},
			want: want{
				result: &pb.GetPrevTagResponse{
					Tags: []*pb.Tag{
						{
							Id:   "tag1",
							Name: "1",
							Articles: []*pb.Article{
								{
									Id:           "1",
									Title:        "happy_path/not_anymore",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
						{
							Id:   "tag2",
							Name: "2",
							Articles: []*pb.Article{
								{
									Id:           "1",
									Title:        "happy_path/not_anymore",
									ThumbnailUrl: "1234567890",
									CreatedAt:    "2020-01-01T00:00:00Z",
									UpdatedAt:    "2020-01-01T00:00:00Z",
								},
							},
						},
					},
				},
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
			if diff := cmp.Diff(got, tt.want.result, protocmp.Transform()); diff != "" {
				t.Errorf("ToGetPrevTagsResponse() = %v, want %v", got, tt.want.result)
			}
		})
	}
}
