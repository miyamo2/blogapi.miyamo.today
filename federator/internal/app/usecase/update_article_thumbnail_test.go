package usecase

import (
	"context"
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/google/go-cmp/cmp"
	blogapictx "github.com/miyamo2/blogapi.miyamo.today/core/context"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"
	grpc "github.com/miyamo2/blogapi.miyamo.today/federator/internal/infra/grpc/bloggingevent"
	mgrpc "github.com/miyamo2/blogapi.miyamo.today/federator/internal/mock/infra/grpc/bloggingevent"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/utils"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/testing/protocmp"
	"testing"
)

func TestUpdateArticleThumbnail_Execute(t *testing.T) {
	type args struct {
		ctx context.Context
		in  dto.UpdateArticleThumbnailInDTO
	}
	type want struct {
		out dto.UpdateArticleThumbnailOutDTO
		err error
	}
	type testCase struct {
		bloggingEventServiceClient func(t *testing.T, ctrl *gomock.Controller, req *grpc.UpdateArticleThumbnailRequest) grpc.BloggingEventServiceClient
		args                       args
		expectedReq                *grpc.UpdateArticleThumbnailRequest
		want                       want
	}
	errTestUpdateArticleThumbnail := errors.New("test error")
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
		"happy_path": {
			bloggingEventServiceClient: func(t *testing.T, ctrl *gomock.Controller, req *grpc.UpdateArticleThumbnailRequest) grpc.BloggingEventServiceClient {
				bloggingEventServiceClient := mgrpc.NewMockBloggingEventServiceClient(ctrl)
				bloggingEventServiceClient.EXPECT().
					UpdateArticleThumbnail(gomock.Any(), NewUpdateArticleThumbnailRequestMatcher(t, req)).
					Return(&grpc.BloggingEventResponse{EventId: "Event1", ArticleId: "Article1"}, nil).Times(1)
				return bloggingEventServiceClient
			},
			expectedReq: &grpc.UpdateArticleThumbnailRequest{
				Id:           "Article1",
				ThumbnailUrl: "https://example.com/example.png",
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.NewUpdateArticleThumbnailInDTO("Article1", utils.MustURLParse("https://example.com/example.png"), "Mutation1"),
			},
			want: want{
				out: dto.NewUpdateArticleThumbnailOutDTO("Event1", "Article1", "Mutation1"),
			},
		},
		"unhappy_path:grpc-return-error": {
			bloggingEventServiceClient: func(t *testing.T, ctrl *gomock.Controller, req *grpc.UpdateArticleThumbnailRequest) grpc.BloggingEventServiceClient {
				bloggingEventServiceClient := mgrpc.NewMockBloggingEventServiceClient(ctrl)
				bloggingEventServiceClient.EXPECT().
					UpdateArticleThumbnail(gomock.Any(), NewUpdateArticleThumbnailRequestMatcher(t, req)).
					Return(nil, errTestUpdateArticleThumbnail).Times(1)
				return bloggingEventServiceClient
			},
			expectedReq: &grpc.UpdateArticleThumbnailRequest{
				Id:           "Article1",
				ThumbnailUrl: "https://example.com/example.png",
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.NewUpdateArticleThumbnailInDTO("Article1", utils.MustURLParse("https://example.com/example.png"), "Mutation1"),
			},
			want: want{
				out: dto.UpdateArticleThumbnailOutDTO{},
				err: errTestUpdateArticleThumbnail,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			bloggingEventServiceClient := tt.bloggingEventServiceClient(t, ctrl, tt.expectedReq)
			u := NewUpdateArticleThumbnail(bloggingEventServiceClient)
			got, err := u.Execute(tt.args.ctx, tt.args.in)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.want.err)
			}
			if diff := cmp.Diff(tt.want.out, got, cmp.AllowUnexported(dto.UpdateArticleThumbnailOutDTO{})); diff != "" {
				t.Errorf("Execute() result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func NewUpdateArticleThumbnailRequestMatcher(t *testing.T, expect *grpc.UpdateArticleThumbnailRequest) gomock.Matcher {
	return &UpdateArticleThumbnailRequestMatcher{
		expect: expect,
		t:      t,
	}
}

type UpdateArticleThumbnailRequestMatcher struct {
	gomock.Matcher
	expect *grpc.UpdateArticleThumbnailRequest
	t      *testing.T
}

func (m *UpdateArticleThumbnailRequestMatcher) Matches(x interface{}) bool {
	switch x := x.(type) {
	case *grpc.UpdateArticleThumbnailRequest:
		if x == nil {
			return m.expect == nil
		}
		diff := cmp.Diff(x, m.expect, protocmp.Transform())
		if diff != "" {
			m.t.Errorf("UpdateArticleThumbnailRequest mismatch (-want +got):\n%s", diff)
			return false
		}
		return true
	}
	return false
}
func (m *UpdateArticleThumbnailRequestMatcher) String() string {
	return fmt.Sprintf("is equal to %+v", m.expect)
}
