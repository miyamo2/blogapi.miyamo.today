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
				t.Errorf("ToGetNextArticlesResponse() = %v, want %v", got, tt.want)
			}
			if !errors.Is(err, tt.want.err) {
				t.Errorf("ToGetNextArticlesResponse() error = %v, want %v", err, tt.want.err)
			}
		})
	}
}
