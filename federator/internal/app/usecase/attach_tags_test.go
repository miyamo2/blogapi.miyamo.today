package usecase

import (
	blogapictx "blogapi.miyamo.today/core/context"
	"blogapi.miyamo.today/federator/internal/app/usecase/dto"
	grpc "blogapi.miyamo.today/federator/internal/infra/grpc/blogging_event"
	"blogapi.miyamo.today/federator/internal/infra/grpc/blogging_event/blogging_eventconnect"
	mbloggingeventconnect "blogapi.miyamo.today/federator/internal/mock/infra/grpc/blogging_event/blogging_eventconnect"
	"connectrpc.com/connect"
	"context"
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/testing/protocmp"
	"testing"
)

func TestAttachTags_Execute(t *testing.T) {
	type args struct {
		ctx context.Context
		in  dto.AttachTagsInDTO
	}
	type want struct {
		out dto.AttachTagsOutDTO
		err error
	}
	type testCase struct {
		bloggingEventServiceClient func(t *testing.T, ctrl *gomock.Controller, req *connect.Request[grpc.AttachTagsRequest]) blogging_eventconnect.BloggingEventServiceClient
		args                       args
		expectedReq                *connect.Request[grpc.AttachTagsRequest]
		want                       want
	}
	errTestAttachTags := errors.New("test error")
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
			bloggingEventServiceClient: func(t *testing.T, ctrl *gomock.Controller, req *connect.Request[grpc.AttachTagsRequest]) blogging_eventconnect.BloggingEventServiceClient {
				bloggingEventServiceClient := mbloggingeventconnect.NewMockBloggingEventServiceClient(ctrl)
				bloggingEventServiceClient.EXPECT().
					AttachTags(gomock.Any(), NewAttachTagsRequestMatcher(t, req)).
					Return(connect.NewResponse(&grpc.BloggingEventResponse{EventId: "Event1", ArticleId: "Article1"}), nil).Times(1)
				return bloggingEventServiceClient
			},
			expectedReq: connect.NewRequest(&grpc.AttachTagsRequest{
				Id:       "Article1",
				TagNames: []string{"Tag1"},
			}),
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.NewAttachTagsInDTO("Article1", []string{"Tag1"}, "ClientMutationID1"),
			},
			want: want{
				out: dto.NewAttachTagsOutDTO("Event1", "Article1", "ClientMutationID1"),
			},
		},
		"unhappy_path:grpc-return-error": {
			bloggingEventServiceClient: func(t *testing.T, ctrl *gomock.Controller, req *connect.Request[grpc.AttachTagsRequest]) blogging_eventconnect.BloggingEventServiceClient {
				bloggingEventServiceClient := mbloggingeventconnect.NewMockBloggingEventServiceClient(ctrl)
				bloggingEventServiceClient.EXPECT().
					AttachTags(gomock.Any(), NewAttachTagsRequestMatcher(t, req)).
					Return(nil, errTestAttachTags).Times(1)
				return bloggingEventServiceClient
			},
			expectedReq: connect.NewRequest(&grpc.AttachTagsRequest{
				Id:       "Article1",
				TagNames: []string{"Tag1"},
			}),
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.NewAttachTagsInDTO("Article1", []string{"Tag1"}, "ClientMutationID1"),
			},
			want: want{
				err: errTestAttachTags,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			uc := NewAttachTags(tt.bloggingEventServiceClient(t, ctrl, tt.expectedReq))
			out, err := uc.Execute(tt.args.ctx, tt.args.in)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("Execute() err = %v, want %v", err, tt.want.err)
			}
			if diff := cmp.Diff(tt.want.out, out, cmp.AllowUnexported((dto.AttachTagsOutDTO{}))); diff != "" {
				t.Errorf("Execute() out = %v, want %v", out, tt.want.out)
			}
		})
	}
}

func NewAttachTagsRequestMatcher(t *testing.T, expect *connect.Request[grpc.AttachTagsRequest]) gomock.Matcher {
	return &AttachTagsRequestMatcher{
		expect: expect,
		t:      t,
	}
}

type AttachTagsRequestMatcher struct {
	gomock.Matcher
	expect *connect.Request[grpc.AttachTagsRequest]
	t      *testing.T
}

func (m *AttachTagsRequestMatcher) Matches(x interface{}) bool {
	switch x := x.(type) {
	case *connect.Request[grpc.AttachTagsRequest]:
		if x == nil {
			return m.expect == nil
		}
		diff := cmp.Diff(x.Msg, m.expect.Msg, protocmp.Transform())
		if diff != "" {
			m.t.Errorf("AttachTagsRequest mismatch (-want +got):\n%s", diff)
			return false
		}
		return true
	}
	return false
}

func (m *AttachTagsRequestMatcher) String() string {
	return fmt.Sprintf("is equal to %+v", m.expect)
}
