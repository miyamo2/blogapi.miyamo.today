package usecase

import (
	blogapictx "blogapi.miyamo.today/core/context"
	"blogapi.miyamo.today/federator/internal/app/usecase/dto"
	grpc "blogapi.miyamo.today/federator/internal/infra/grpc/bloggingevent"
	mgrpc "blogapi.miyamo.today/federator/internal/mock/infra/grpc/bloggingevent"
	"context"
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/testing/protocmp"
	"testing"
)

func TestUpdateArticleBody_Execute(t *testing.T) {
	type args struct {
		ctx context.Context
		in  dto.UpdateArticleBodyInDTO
	}
	type want struct {
		out dto.UpdateArticleBodyOutDTO
		err error
	}
	type testCase struct {
		bloggingEventServiceClient func(t *testing.T, ctrl *gomock.Controller, req *grpc.UpdateArticleBodyRequest) grpc.BloggingEventServiceClient
		args                       args
		expectedReq                *grpc.UpdateArticleBodyRequest
		want                       want
	}
	errTestUpdateArticleBody := errors.New("test error")
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
			bloggingEventServiceClient: func(t *testing.T, ctrl *gomock.Controller, req *grpc.UpdateArticleBodyRequest) grpc.BloggingEventServiceClient {
				bloggingEventServiceClient := mgrpc.NewMockBloggingEventServiceClient(ctrl)
				bloggingEventServiceClient.EXPECT().
					UpdateArticleBody(gomock.Any(), NewUpdateArticleBodyRequestMatcher(t, req)).
					Return(&grpc.BloggingEventResponse{EventId: "Event1", ArticleId: "Article1"}, nil).Times(1)
				return bloggingEventServiceClient
			},
			expectedReq: &grpc.UpdateArticleBodyRequest{
				Id:   "Article1",
				Body: "happy_path",
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.NewUpdateArticleBodyInDTO("Article1", "happy_path", "Mutation1"),
			},
			want: want{
				out: dto.NewUpdateArticleBodyOutDTO("Event1", "Article1", "Mutation1"),
			},
		},
		"unhappy_path:grpc-return-error": {
			bloggingEventServiceClient: func(t *testing.T, ctrl *gomock.Controller, req *grpc.UpdateArticleBodyRequest) grpc.BloggingEventServiceClient {
				bloggingEventServiceClient := mgrpc.NewMockBloggingEventServiceClient(ctrl)
				bloggingEventServiceClient.EXPECT().
					UpdateArticleBody(gomock.Any(), NewUpdateArticleBodyRequestMatcher(t, req)).
					Return(nil, errTestUpdateArticleBody).Times(1)
				return bloggingEventServiceClient
			},
			expectedReq: &grpc.UpdateArticleBodyRequest{
				Id:   "Article1",
				Body: "happy_path",
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.NewUpdateArticleBodyInDTO("Article1", "happy_path", "Mutation1"),
			},
			want: want{
				out: dto.UpdateArticleBodyOutDTO{},
				err: errTestUpdateArticleBody,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			bloggingEventServiceClient := tt.bloggingEventServiceClient(t, ctrl, tt.expectedReq)
			u := NewUpdateArticleBody(bloggingEventServiceClient)
			got, err := u.Execute(tt.args.ctx, tt.args.in)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.want.err)
			}
			if diff := cmp.Diff(tt.want.out, got, cmp.AllowUnexported(dto.UpdateArticleBodyOutDTO{})); diff != "" {
				t.Errorf("Execute() result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func NewUpdateArticleBodyRequestMatcher(t *testing.T, expect *grpc.UpdateArticleBodyRequest) gomock.Matcher {
	return &UpdateArticleBodyRequestMatcher{
		expect: expect,
		t:      t,
	}
}

type UpdateArticleBodyRequestMatcher struct {
	gomock.Matcher
	expect *grpc.UpdateArticleBodyRequest
	t      *testing.T
}

func (m *UpdateArticleBodyRequestMatcher) Matches(x interface{}) bool {
	switch x := x.(type) {
	case *grpc.UpdateArticleBodyRequest:
		if x == nil {
			return m.expect == nil
		}
		diff := cmp.Diff(x, m.expect, protocmp.Transform())
		if diff != "" {
			m.t.Errorf("UpdateArticleBodyRequest mismatch (-want +got):\n%s", diff)
			return false
		}
		return true
	}
	return false
}
func (m *UpdateArticleBodyRequestMatcher) String() string {
	return fmt.Sprintf("is equal to %+v", m.expect)
}
