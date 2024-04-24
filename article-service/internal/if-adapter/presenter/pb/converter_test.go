package pb

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/miyamo2/api.miyamo.today/article-service/internal/app/usecase/dto"
	"github.com/miyamo2/api.miyamo.today/protogen/article/server/pb"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestConverter_ToGetNextArticlesResponse(t *testing.T) {
	type args struct {
		ctx  context.Context
		from func() *dto.GetNextOutDto
	}
	type want struct {
		result *pb.GetNextArticlesResponse
		ok     bool
	}
	type testCase struct {
		args args
		want want
	}
	tests := map[string]testCase{
		"happy_path/multiple": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.GetNextOutDto {
					o := dto.NewGetNextOutDto(
						[]dto.Article{
							dto.NewArticle(
								"1",
								"happy_path/multiple/still_exists1",
								"## happy_path/multiple/still_exists1",
								"1234567890",
								"2020-01-01T00:00:00Z",
								"2020-01-01T00:00:00Z",
								[]dto.Tag{
									dto.NewTag("tag1", "1"),
									dto.NewTag("tag2", "2"),
								},
							),
							dto.NewArticle(
								"2",
								"happy_path/multiple/still_exists2",
								"## happy_path/multiple_/still_exists2",
								"1234567890",
								"2020-01-01T00:00:00Z",
								"2020-01-01T00:00:00Z",
								[]dto.Tag{
									dto.NewTag("tag1", "1"),
									dto.NewTag("tag2", "2"),
								},
							),
						}, false)
					return &o
				},
			},
			want: want{
				result: &pb.GetNextArticlesResponse{
					Articles: []*pb.Article{
						{
							Id:           "1",
							Title:        "happy_path/multiple/still_exists1",
							Body:         "## happy_path/multiple/still_exists1",
							ThumbnailUrl: "1234567890",
							CreatedAt:    "2020-01-01T00:00:00Z",
							UpdatedAt:    "2020-01-01T00:00:00Z",
							Tags: []*pb.Tag{
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
							Title:        "happy_path/multiple/still_exists2",
							Body:         "## happy_path/multiple_/still_exists2",
							ThumbnailUrl: "1234567890",
							CreatedAt:    "2020-01-01T00:00:00Z",
							UpdatedAt:    "2020-01-01T00:00:00Z",
							Tags: []*pb.Tag{
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
					StillExists: false,
				},
				ok: true,
			},
		},
		"happy_path/multiple/still_exists": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.GetNextOutDto {
					o := dto.NewGetNextOutDto(
						[]dto.Article{
							dto.NewArticle(
								"1",
								"happy_path/multiple/still_exists1",
								"## happy_path/multiple/still_exists1",
								"1234567890",
								"2020-01-01T00:00:00Z",
								"2020-01-01T00:00:00Z",
								[]dto.Tag{
									dto.NewTag("tag1", "1"),
									dto.NewTag("tag2", "2"),
								},
							),
							dto.NewArticle(
								"2",
								"happy_path/multiple/still_exists2",
								"## happy_path/multiple_/still_exists2",
								"1234567890",
								"2020-01-01T00:00:00Z",
								"2020-01-01T00:00:00Z",
								[]dto.Tag{
									dto.NewTag("tag1", "1"),
									dto.NewTag("tag2", "2"),
								},
							),
						}, true)
					return &o
				},
			},
			want: want{
				result: &pb.GetNextArticlesResponse{
					Articles: []*pb.Article{
						{
							Id:           "1",
							Title:        "happy_path/multiple/still_exists1",
							Body:         "## happy_path/multiple/still_exists1",
							ThumbnailUrl: "1234567890",
							CreatedAt:    "2020-01-01T00:00:00Z",
							UpdatedAt:    "2020-01-01T00:00:00Z",
							Tags: []*pb.Tag{
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
							Title:        "happy_path/multiple/still_exists2",
							Body:         "## happy_path/multiple_/still_exists2",
							ThumbnailUrl: "1234567890",
							CreatedAt:    "2020-01-01T00:00:00Z",
							UpdatedAt:    "2020-01-01T00:00:00Z",
							Tags: []*pb.Tag{
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
				},
				ok: true,
			},
		},
		"happy_path/single": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.GetNextOutDto {
					o := dto.NewGetNextOutDto(
						[]dto.Article{
							dto.NewArticle(
								"1",
								"happy_path/multiple/still_exists1",
								"## happy_path/multiple/still_exists1",
								"1234567890",
								"2020-01-01T00:00:00Z",
								"2020-01-01T00:00:00Z",
								[]dto.Tag{
									dto.NewTag("tag1", "1"),
									dto.NewTag("tag2", "2"),
								},
							),
						}, false)
					return &o
				},
			},
			want: want{
				result: &pb.GetNextArticlesResponse{
					Articles: []*pb.Article{
						{
							Id:           "1",
							Title:        "happy_path/multiple/still_exists1",
							Body:         "## happy_path/multiple/still_exists1",
							ThumbnailUrl: "1234567890",
							CreatedAt:    "2020-01-01T00:00:00Z",
							UpdatedAt:    "2020-01-01T00:00:00Z",
							Tags: []*pb.Tag{
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
					StillExists: false,
				},
				ok: true,
			},
		},
		"happy_path/single/still_exists": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.GetNextOutDto {
					o := dto.NewGetNextOutDto(
						[]dto.Article{
							dto.NewArticle(
								"1",
								"happy_path/multiple/still_exists1",
								"## happy_path/multiple/still_exists1",
								"1234567890",
								"2020-01-01T00:00:00Z",
								"2020-01-01T00:00:00Z",
								[]dto.Tag{
									dto.NewTag("tag1", "1"),
									dto.NewTag("tag2", "2"),
								},
							),
						}, true)
					return &o
				},
			},
			want: want{
				result: &pb.GetNextArticlesResponse{
					Articles: []*pb.Article{
						{
							Id:           "1",
							Title:        "happy_path/multiple/still_exists1",
							Body:         "## happy_path/multiple/still_exists1",
							ThumbnailUrl: "1234567890",
							CreatedAt:    "2020-01-01T00:00:00Z",
							UpdatedAt:    "2020-01-01T00:00:00Z",
							Tags: []*pb.Tag{
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
				},
				ok: true,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := NewConverter()
			got, ok := c.ToGetNextArticlesResponse(tt.args.ctx, tt.args.from())
			if tt.want.ok != ok {
				t.Errorf("ToGetNextArticlesResponse() ok = %v, want %v", ok, tt.want.ok)
			}
			if diff := cmp.Diff(got, tt.want.result, protocmp.Transform()); diff != "" {
				t.Errorf("ToGetNextArticlesResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConverter_ToGetAllArticlesResponse(t *testing.T) {
	type args struct {
		ctx  context.Context
		from func() *dto.GetAllOutDto
	}
	type want struct {
		result *pb.GetAllArticlesResponse
		ok     bool
	}
	type testCase struct {
		args args
		want want
	}
	tests := map[string]testCase{
		"happy_path/multiple": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.GetAllOutDto {
					o := dto.NewGetAllOutDto(
						[]dto.Article{
							dto.NewArticle(
								"1",
								"happy_path/multiple/still_exists1",
								"## happy_path/multiple/still_exists1",
								"1234567890",
								"2020-01-01T00:00:00Z",
								"2020-01-01T00:00:00Z",
								[]dto.Tag{
									dto.NewTag("tag1", "1"),
									dto.NewTag("tag2", "2"),
								},
							),
							dto.NewArticle(
								"2",
								"happy_path/multiple/still_exists2",
								"## happy_path/multiple_/still_exists2",
								"1234567890",
								"2020-01-01T00:00:00Z",
								"2020-01-01T00:00:00Z",
								[]dto.Tag{
									dto.NewTag("tag1", "1"),
									dto.NewTag("tag2", "2"),
								},
							),
						})
					return &o
				},
			},
			want: want{
				result: &pb.GetAllArticlesResponse{
					Articles: []*pb.Article{
						{
							Id:           "1",
							Title:        "happy_path/multiple/still_exists1",
							Body:         "## happy_path/multiple/still_exists1",
							ThumbnailUrl: "1234567890",
							CreatedAt:    "2020-01-01T00:00:00Z",
							UpdatedAt:    "2020-01-01T00:00:00Z",
							Tags: []*pb.Tag{
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
							Title:        "happy_path/multiple/still_exists2",
							Body:         "## happy_path/multiple_/still_exists2",
							ThumbnailUrl: "1234567890",
							CreatedAt:    "2020-01-01T00:00:00Z",
							UpdatedAt:    "2020-01-01T00:00:00Z",
							Tags: []*pb.Tag{
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
				},
				ok: true,
			},
		},
		"happy_path/multiple/still_exists": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.GetAllOutDto {
					o := dto.NewGetAllOutDto(
						[]dto.Article{
							dto.NewArticle(
								"1",
								"happy_path/multiple/still_exists1",
								"## happy_path/multiple/still_exists1",
								"1234567890",
								"2020-01-01T00:00:00Z",
								"2020-01-01T00:00:00Z",
								[]dto.Tag{
									dto.NewTag("tag1", "1"),
									dto.NewTag("tag2", "2"),
								},
							),
							dto.NewArticle(
								"2",
								"happy_path/multiple/still_exists2",
								"## happy_path/multiple_/still_exists2",
								"1234567890",
								"2020-01-01T00:00:00Z",
								"2020-01-01T00:00:00Z",
								[]dto.Tag{
									dto.NewTag("tag1", "1"),
									dto.NewTag("tag2", "2"),
								},
							),
						})
					return &o
				},
			},
			want: want{
				result: &pb.GetAllArticlesResponse{
					Articles: []*pb.Article{
						{
							Id:           "1",
							Title:        "happy_path/multiple/still_exists1",
							Body:         "## happy_path/multiple/still_exists1",
							ThumbnailUrl: "1234567890",
							CreatedAt:    "2020-01-01T00:00:00Z",
							UpdatedAt:    "2020-01-01T00:00:00Z",
							Tags: []*pb.Tag{
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
							Title:        "happy_path/multiple/still_exists2",
							Body:         "## happy_path/multiple_/still_exists2",
							ThumbnailUrl: "1234567890",
							CreatedAt:    "2020-01-01T00:00:00Z",
							UpdatedAt:    "2020-01-01T00:00:00Z",
							Tags: []*pb.Tag{
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
				},
				ok: true,
			},
		},
		"happy_path/single": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.GetAllOutDto {
					o := dto.NewGetAllOutDto(
						[]dto.Article{
							dto.NewArticle(
								"1",
								"happy_path/multiple/still_exists1",
								"## happy_path/multiple/still_exists1",
								"1234567890",
								"2020-01-01T00:00:00Z",
								"2020-01-01T00:00:00Z",
								[]dto.Tag{
									dto.NewTag("tag1", "1"),
									dto.NewTag("tag2", "2"),
								},
							),
						})
					return &o
				},
			},
			want: want{
				result: &pb.GetAllArticlesResponse{
					Articles: []*pb.Article{
						{
							Id:           "1",
							Title:        "happy_path/multiple/still_exists1",
							Body:         "## happy_path/multiple/still_exists1",
							ThumbnailUrl: "1234567890",
							CreatedAt:    "2020-01-01T00:00:00Z",
							UpdatedAt:    "2020-01-01T00:00:00Z",
							Tags: []*pb.Tag{
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
				},
				ok: true,
			},
		},
		"happy_path/single/still_exists": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.GetAllOutDto {
					o := dto.NewGetAllOutDto(
						[]dto.Article{
							dto.NewArticle(
								"1",
								"happy_path/multiple/still_exists1",
								"## happy_path/multiple/still_exists1",
								"1234567890",
								"2020-01-01T00:00:00Z",
								"2020-01-01T00:00:00Z",
								[]dto.Tag{
									dto.NewTag("tag1", "1"),
									dto.NewTag("tag2", "2"),
								},
							),
						})
					return &o
				},
			},
			want: want{
				result: &pb.GetAllArticlesResponse{
					Articles: []*pb.Article{
						{
							Id:           "1",
							Title:        "happy_path/multiple/still_exists1",
							Body:         "## happy_path/multiple/still_exists1",
							ThumbnailUrl: "1234567890",
							CreatedAt:    "2020-01-01T00:00:00Z",
							UpdatedAt:    "2020-01-01T00:00:00Z",
							Tags: []*pb.Tag{
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
				},
				ok: true,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := NewConverter()
			got, ok := c.ToGetAllArticlesResponse(tt.args.ctx, tt.args.from())
			if tt.want.ok != ok {
				t.Errorf("ToGetAllArticlesResponse() ok = %v, want %v", ok, tt.want.ok)
			}
			if diff := cmp.Diff(got, tt.want.result, protocmp.Transform()); diff != "" {
				t.Errorf("ToGetAllArticlesResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConverter_ToGetByIdArticlesResponse(t *testing.T) {
	type args struct {
		ctx  context.Context
		from func() *dto.GetByIdOutDto
	}
	type want struct {
		result *pb.GetArticleByIdResponse
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
						"1",
						"happy_path/multiple/still_exists1",
						"## happy_path/multiple/still_exists1",
						"1234567890",
						"2020-01-01T00:00:00Z",
						"2020-01-01T00:00:00Z",
						[]dto.Tag{
							dto.NewTag("tag1", "1"),
							dto.NewTag("tag2", "2"),
						},
					)
					return &o
				},
			},
			want: want{
				result: &pb.GetArticleByIdResponse{
					Article: &pb.Article{
						Id:           "1",
						Title:        "happy_path/multiple/still_exists1",
						Body:         "## happy_path/multiple/still_exists1",
						ThumbnailUrl: "1234567890",
						CreatedAt:    "2020-01-01T00:00:00Z",
						UpdatedAt:    "2020-01-01T00:00:00Z",
						Tags: []*pb.Tag{
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
				ok: true,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := NewConverter()
			got, ok := c.ToGetByIdArticlesResponse(tt.args.ctx, tt.args.from())
			if tt.want.ok != ok {
				t.Errorf("ToGetByIdArticlesResponse() ok = %v, want %v", ok, tt.want.ok)
			}
			if diff := cmp.Diff(got, tt.want.result, protocmp.Transform()); diff != "" {
				t.Errorf("ToGetByIdArticlesResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
