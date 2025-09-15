package convert

import (
	"context"
	"testing"

	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"google.golang.org/protobuf/types/known/timestamppb"

	"blogapi.miyamo.today/article-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/article-service/internal/infra/grpc"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestListAfter_ToResponse(t *testing.T) {
	type args struct {
		ctx  context.Context
		from func() *dto.ListAfterOutput
	}
	type want struct {
		result *grpc.GetNextArticlesResponse
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
				from: func() *dto.ListAfterOutput {
					o := dto.NewListAfterOutput(
						false,
						dto.NewArticle(
							"1",
							"happy_path/multiple/still_exists1",
							"## happy_path/multiple/still_exists1",
							"1234567890",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							dto.NewTag("tag1", "1"),
							dto.NewTag("tag2", "2"),
						),
						dto.NewArticle(
							"2",
							"happy_path/multiple/still_exists2",
							"## happy_path/multiple_/still_exists2",
							"1234567890",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							dto.NewTag("tag1", "1"),
							dto.NewTag("tag2", "2"),
						),
					)
					return &o
				},
			},
			want: want{
				result: &grpc.GetNextArticlesResponse{
					Articles: []*grpc.Article{
						{
							Id:           "1",
							Title:        "happy_path/multiple/still_exists1",
							Body:         "## happy_path/multiple/still_exists1",
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
							Title:        "happy_path/multiple/still_exists2",
							Body:         "## happy_path/multiple_/still_exists2",
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
					StillExists: false,
				},
				ok: true,
			},
		},
		"happy_path/multiple/still_exists": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.ListAfterOutput {
					o := dto.NewListAfterOutput(
						true,
						dto.NewArticle(
							"1",
							"happy_path/multiple/still_exists1",
							"## happy_path/multiple/still_exists1",
							"1234567890",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							dto.NewTag("tag1", "1"),
							dto.NewTag("tag2", "2"),
						),
						dto.NewArticle(
							"2",
							"happy_path/multiple/still_exists2",
							"## happy_path/multiple_/still_exists2",
							"1234567890",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							dto.NewTag("tag1", "1"),
							dto.NewTag("tag2", "2"),
						),
					)
					return &o
				},
			},
			want: want{
				result: &grpc.GetNextArticlesResponse{
					Articles: []*grpc.Article{
						{
							Id:           "1",
							Title:        "happy_path/multiple/still_exists1",
							Body:         "## happy_path/multiple/still_exists1",
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
							Title:        "happy_path/multiple/still_exists2",
							Body:         "## happy_path/multiple_/still_exists2",
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
				},
				ok: true,
			},
		},
		"happy_path/single": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.ListAfterOutput {
					o := dto.NewListAfterOutput(
						false,
						dto.NewArticle(
							"1",
							"happy_path/multiple/still_exists1",
							"## happy_path/multiple/still_exists1",
							"1234567890",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							dto.NewTag("tag1", "1"),
							dto.NewTag("tag2", "2"),
						),
					)
					return &o
				},
			},
			want: want{
				result: &grpc.GetNextArticlesResponse{
					Articles: []*grpc.Article{
						{
							Id:           "1",
							Title:        "happy_path/multiple/still_exists1",
							Body:         "## happy_path/multiple/still_exists1",
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
					StillExists: false,
				},
				ok: true,
			},
		},
		"happy_path/single/still_exists": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.ListAfterOutput {
					o := dto.NewListAfterOutput(
						true,
						dto.NewArticle(
							"1",
							"happy_path/multiple/still_exists1",
							"## happy_path/multiple/still_exists1",
							"1234567890",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							dto.NewTag("tag1", "1"),
							dto.NewTag("tag2", "2"),
						),
					)
					return &o
				},
			},
			want: want{
				result: &grpc.GetNextArticlesResponse{
					Articles: []*grpc.Article{
						{
							Id:           "1",
							Title:        "happy_path/multiple/still_exists1",
							Body:         "## happy_path/multiple/still_exists1",
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
				},
				ok: true,
			},
		},
	}
	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				c := NewListAfter()
				got, ok := c.ToResponse(tt.args.ctx, tt.args.from())
				if tt.want.ok != ok {
					t.Errorf("ToResponse() ok = %v, want %v", ok, tt.want.ok)
				}
				if diff := cmp.Diff(got, tt.want.result, protocmp.Transform()); diff != "" {
					t.Errorf("ToResponse() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestListAll_ToResponse(t *testing.T) {
	type args struct {
		ctx  context.Context
		from func() *dto.ListAllOutput
	}
	type want struct {
		result *grpc.GetAllArticlesResponse
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
				from: func() *dto.ListAllOutput {
					o := dto.NewListAllOutput(
						dto.NewArticle(
							"1",
							"happy_path/multiple/still_exists1",
							"## happy_path/multiple/still_exists1",
							"1234567890",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							dto.NewTag("tag1", "1"),
							dto.NewTag("tag2", "2"),
						),
						dto.NewArticle(
							"2",
							"happy_path/multiple/still_exists2",
							"## happy_path/multiple_/still_exists2",
							"1234567890",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							dto.NewTag("tag1", "1"),
							dto.NewTag("tag2", "2"),
						),
					)
					return &o
				},
			},
			want: want{
				result: &grpc.GetAllArticlesResponse{
					Articles: []*grpc.Article{
						{
							Id:           "1",
							Title:        "happy_path/multiple/still_exists1",
							Body:         "## happy_path/multiple/still_exists1",
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
							Title:        "happy_path/multiple/still_exists2",
							Body:         "## happy_path/multiple_/still_exists2",
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
				},
				ok: true,
			},
		},
		"happy_path/multiple/still_exists": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.ListAllOutput {
					o := dto.NewListAllOutput(
						dto.NewArticle(
							"1",
							"happy_path/multiple/still_exists1",
							"## happy_path/multiple/still_exists1",
							"1234567890",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							dto.NewTag("tag1", "1"),
							dto.NewTag("tag2", "2"),
						),
						dto.NewArticle(
							"2",
							"happy_path/multiple/still_exists2",
							"## happy_path/multiple_/still_exists2",
							"1234567890",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							dto.NewTag("tag1", "1"),
							dto.NewTag("tag2", "2"),
						),
					)
					return &o
				},
			},
			want: want{
				result: &grpc.GetAllArticlesResponse{
					Articles: []*grpc.Article{
						{
							Id:           "1",
							Title:        "happy_path/multiple/still_exists1",
							Body:         "## happy_path/multiple/still_exists1",
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
							Title:        "happy_path/multiple/still_exists2",
							Body:         "## happy_path/multiple_/still_exists2",
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
				},
				ok: true,
			},
		},
		"happy_path/single": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.ListAllOutput {
					o := dto.NewListAllOutput(
						dto.NewArticle(
							"1",
							"happy_path/multiple/still_exists1",
							"## happy_path/multiple/still_exists1",
							"1234567890",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							dto.NewTag("tag1", "1"),
							dto.NewTag("tag2", "2"),
						),
					)
					return &o
				},
			},
			want: want{
				result: &grpc.GetAllArticlesResponse{
					Articles: []*grpc.Article{
						{
							Id:           "1",
							Title:        "happy_path/multiple/still_exists1",
							Body:         "## happy_path/multiple/still_exists1",
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
				},
				ok: true,
			},
		},
		"happy_path/single/still_exists": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.ListAllOutput {
					o := dto.NewListAllOutput(
						dto.NewArticle(
							"1",
							"happy_path/multiple/still_exists1",
							"## happy_path/multiple/still_exists1",
							"1234567890",
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
							dto.NewTag("tag1", "1"),
							dto.NewTag("tag2", "2"),
						),
					)
					return &o
				},
			},
			want: want{
				result: &grpc.GetAllArticlesResponse{
					Articles: []*grpc.Article{
						{
							Id:           "1",
							Title:        "happy_path/multiple/still_exists1",
							Body:         "## happy_path/multiple/still_exists1",
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
				},
				ok: true,
			},
		},
	}
	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				c := NewListAll()
				got, ok := c.ToResponse(tt.args.ctx, tt.args.from())
				if tt.want.ok != ok {
					t.Errorf("ToResponse() ok = %v, want %v", ok, tt.want.ok)
				}
				if diff := cmp.Diff(got, tt.want.result, protocmp.Transform()); diff != "" {
					t.Errorf("ToResponse() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestGetByID_ToResponse(t *testing.T) {
	type args struct {
		ctx  context.Context
		from func() *dto.GetByIDOutput
	}
	type want struct {
		result *grpc.GetArticleByIdResponse
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
				from: func() *dto.GetByIDOutput {
					o := dto.NewGetByIDOutput(
						"1",
						"happy_path/multiple/still_exists1",
						"## happy_path/multiple/still_exists1",
						"1234567890",
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						synchro.New[tz.UTC](2020, 1, 1, 0, 0, 0, 0),
						dto.NewTag("tag1", "1"),
						dto.NewTag("tag2", "2"),
					)
					return &o
				},
			},
			want: want{
				result: &grpc.GetArticleByIdResponse{
					Article: &grpc.Article{
						Id:           "1",
						Title:        "happy_path/multiple/still_exists1",
						Body:         "## happy_path/multiple/still_exists1",
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
				ok: true,
			},
		},
	}
	for name, tt := range tests {
		t.Run(
			name, func(t *testing.T) {
				c := NewGetByID()
				got, ok := c.ToResponse(tt.args.ctx, tt.args.from())
				if tt.want.ok != ok {
					t.Errorf("ToResponse() ok = %v, want %v", ok, tt.want.ok)
				}
				if diff := cmp.Diff(got, tt.want.result, protocmp.Transform()); diff != "" {
					t.Errorf("ToResponse() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
