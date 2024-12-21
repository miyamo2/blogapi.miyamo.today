package pb

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/infra/grpc"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestConverter_ToCreateArticleArticleResponse(t *testing.T) {
	type args struct {
		ctx  context.Context
		from func() *dto.CreateArticleOutDto
	}
	type want struct {
		result *grpc.BloggingEventResponse
		err    error
	}
	type testCase struct {
		args args
		want want
	}
	tests := map[string]testCase{
		"happy_path/single": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.CreateArticleOutDto {
					o := dto.NewCreateArticleOutDto("abc", "def")
					return &o
				},
			},
			want: want{
				result: &grpc.BloggingEventResponse{EventId: "abc", ArticleId: "def"},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := NewConverter()
			got, err := c.ToCreateArticleArticleResponse(tt.args.ctx, tt.args.from())
			if diff := cmp.Diff(got, tt.want.result, protocmp.Transform()); diff != "" {
				t.Errorf("ToCreateArticleArticleResponse() = %v, want %v", got, tt.want)
			}
			if !errors.Is(err, tt.want.err) {
				t.Errorf("ToCreateArticleArticleResponse() error = %v, want %v", err, tt.want.err)
			}
		})
	}
}

func TestConverter_ToUpdateArticleTitleResponse(t *testing.T) {
	type args struct {
		ctx  context.Context
		from func() *dto.UpdateArticleTitleOutDto
	}
	type want struct {
		result *grpc.BloggingEventResponse
		err    error
	}
	type testCase struct {
		args args
		want want
	}
	tests := map[string]testCase{
		"happy_path/single": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.UpdateArticleTitleOutDto {
					o := dto.NewUpdateArticleTitleOutDto("abc", "def")
					return &o
				},
			},
			want: want{
				result: &grpc.BloggingEventResponse{EventId: "abc", ArticleId: "def"},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := NewConverter()
			got, err := c.ToUpdateArticleTitleResponse(tt.args.ctx, tt.args.from())
			if diff := cmp.Diff(got, tt.want.result, protocmp.Transform()); diff != "" {
				t.Errorf("ToUpdateArticleTitleResponse() = %v, want %v", got, tt.want)
			}
			if !errors.Is(err, tt.want.err) {
				t.Errorf("ToUpdateArticleTitleResponse() error = %v, want %v", err, tt.want.err)
			}
		})
	}
}

func TestConverter_ToUpdateArticleBodyResponse(t *testing.T) {
	type args struct {
		ctx  context.Context
		from func() *dto.UpdateArticleBodyOutDto
	}
	type want struct {
		result *grpc.BloggingEventResponse
		err    error
	}
	type testCase struct {
		args args
		want want
	}
	tests := map[string]testCase{
		"happy_path/single": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.UpdateArticleBodyOutDto {
					o := dto.NewUpdateArticleBodyOutDto("abc", "def")
					return &o
				},
			},
			want: want{
				result: &grpc.BloggingEventResponse{EventId: "abc", ArticleId: "def"},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := NewConverter()
			got, err := c.ToUpdateArticleBodyResponse(tt.args.ctx, tt.args.from())
			if diff := cmp.Diff(got, tt.want.result, protocmp.Transform()); diff != "" {
				t.Errorf("ToUpdateArticleBodyResponse() = %v, want %v", got, tt.want)
			}
			if !errors.Is(err, tt.want.err) {
				t.Errorf("ToUpdateArticleBodyResponse() error = %v, want %v", err, tt.want.err)
			}
		})
	}
}

func TestConverter_ToUpdateArticleThumbnailResponse(t *testing.T) {
	type args struct {
		ctx  context.Context
		from func() *dto.UpdateArticleThumbnailOutDto
	}
	type want struct {
		result *grpc.BloggingEventResponse
		err    error
	}
	type testCase struct {
		args args
		want want
	}
	tests := map[string]testCase{
		"happy_path/single": {
			args: args{
				ctx: context.Background(),
				from: func() *dto.UpdateArticleThumbnailOutDto {
					o := dto.NewUpdateArticleThumbnailOutDto("abc", "def")
					return &o
				},
			},
			want: want{
				result: &grpc.BloggingEventResponse{EventId: "abc", ArticleId: "def"},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := NewConverter()
			got, err := c.ToUpdateArticleThumbnailResponse(tt.args.ctx, tt.args.from())
			if diff := cmp.Diff(got, tt.want.result, protocmp.Transform()); diff != "" {
				t.Errorf("ToUpdateArticleThumbnailResponse() = %v, want %v", got, tt.want)
			}
			if !errors.Is(err, tt.want.err) {
				t.Errorf("ToUpdateArticleThumbnailResponse() error = %v, want %v", err, tt.want.err)
			}
		})
	}
}
