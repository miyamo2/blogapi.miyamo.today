package usecase

import (
	blogapictx "blogapi.miyamo.today/core/context"
	"blogapi.miyamo.today/federator/internal/app/usecase/dto"
	grpc "blogapi.miyamo.today/federator/internal/infra/grpc/blogging_event"
	"blogapi.miyamo.today/federator/internal/infra/grpc/blogging_event/blogging_eventconnect"
	mbloggingeventconnect "blogapi.miyamo.today/federator/internal/mock/infra/grpc/blogging_event/blogging_eventconnect"
	"blogapi.miyamo.today/federator/internal/utils"
	"connectrpc.com/connect"
	"context"
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/testing/protocmp"
	"testing"
)

func TestCreateArticle_Execute(t *testing.T) {
	type args struct {
		ctx context.Context
		in  dto.CreateArticleInDTO
	}
	type want struct {
		out dto.CreateArticleOutDTO
		err error
	}
	type testCase struct {
		bloggingEventServiceClient func(t *testing.T, ctrl *gomock.Controller, req *connect.Request[grpc.CreateArticleRequest]) blogging_eventconnect.BloggingEventServiceClient
		args                       args
		expectedReq                *connect.Request[grpc.CreateArticleRequest]
		want                       want
	}
	errTestCreateArticle := errors.New("test error")
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
			bloggingEventServiceClient: func(t *testing.T, ctrl *gomock.Controller, req *connect.Request[grpc.CreateArticleRequest]) blogging_eventconnect.BloggingEventServiceClient {
				bloggingEventServiceClient := mbloggingeventconnect.NewMockBloggingEventServiceClient(ctrl)
				bloggingEventServiceClient.EXPECT().
					CreateArticle(gomock.Any(), NewCreateArticleRequestMatcher(t, req)).
					Return(connect.NewResponse(&grpc.BloggingEventResponse{EventId: "Event1", ArticleId: "Article1"}), nil).Times(1)
				return bloggingEventServiceClient
			},
			expectedReq: connect.NewRequest(&grpc.CreateArticleRequest{
				Title:        "Title1",
				Body:         "happy_path",
				ThumbnailUrl: "https://example.com/example.png",
				TagNames:     []string{"Tag1"},
			}),
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.NewCreateArticleInDTO("Title1", "happy_path", utils.MustURLParse("https://example.com/example.png"), []string{"Tag1"}, "Mutation1"),
			},
			want: want{
				out: dto.NewCreateArticleOutDTO("Event1", "Article1", "Mutation1"),
			},
		},
		"unhappy_path:grpc-return-error": {
			bloggingEventServiceClient: func(t *testing.T, ctrl *gomock.Controller, req *connect.Request[grpc.CreateArticleRequest]) blogging_eventconnect.BloggingEventServiceClient {
				bloggingEventServiceClient := mbloggingeventconnect.NewMockBloggingEventServiceClient(ctrl)
				bloggingEventServiceClient.EXPECT().
					CreateArticle(gomock.Any(), NewCreateArticleRequestMatcher(t, req)).
					Return(nil, errTestCreateArticle).Times(1)
				return bloggingEventServiceClient
			},
			expectedReq: connect.NewRequest(&grpc.CreateArticleRequest{
				Title:        "Title1",
				Body:         "happy_path",
				ThumbnailUrl: "https://example.com/example.png",
				TagNames:     []string{"Tag1"},
			}),
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.NewCreateArticleInDTO("Title1", "happy_path", utils.MustURLParse("https://example.com/example.png"), []string{"Tag1"}, "Mutation1"),
			},
			want: want{
				out: dto.CreateArticleOutDTO{},
				err: errTestCreateArticle,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			bloggingEventServiceClient := tt.bloggingEventServiceClient(t, ctrl, tt.expectedReq)
			u := NewCreateArticle(bloggingEventServiceClient)
			got, err := u.Execute(tt.args.ctx, tt.args.in)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.want.err)
			}
			if diff := cmp.Diff(tt.want.out, got, cmp.AllowUnexported(dto.CreateArticleOutDTO{})); diff != "" {
				t.Errorf("Execute() result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func NewCreateArticleRequestMatcher(t *testing.T, expect *connect.Request[grpc.CreateArticleRequest]) gomock.Matcher {
	return &CreateArticleRequestMatcher{
		expect: expect,
		t:      t,
	}
}

type CreateArticleRequestMatcher struct {
	gomock.Matcher
	expect *connect.Request[grpc.CreateArticleRequest]
	t      *testing.T
}

func (m *CreateArticleRequestMatcher) Matches(x interface{}) bool {
	switch x := x.(type) {
	case *connect.Request[grpc.CreateArticleRequest]:
		if x == nil {
			return m.expect == nil
		}
		diff := cmp.Diff(x.Msg, m.expect.Msg, protocmp.Transform())
		if diff != "" {
			m.t.Errorf("CreateArticleRequest mismatch (-want +got):\n%s", diff)
			return false
		}
		return true
	}
	return false
}
func (m *CreateArticleRequestMatcher) String() string {
	return fmt.Sprintf("is equal to %+v", m.expect)
}
