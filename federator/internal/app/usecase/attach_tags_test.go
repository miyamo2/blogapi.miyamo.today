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
		bloggingEventServiceClient func(t *testing.T, ctrl *gomock.Controller, req *grpc.AttachTagsRequest) grpc.BloggingEventServiceClient
		args                       args
		expectedReq                *grpc.AttachTagsRequest
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
			bloggingEventServiceClient: func(t *testing.T, ctrl *gomock.Controller, req *grpc.AttachTagsRequest) grpc.BloggingEventServiceClient {
				bloggingEventServiceClient := mgrpc.NewMockBloggingEventServiceClient(ctrl)
				bloggingEventServiceClient.EXPECT().
					AttachTags(gomock.Any(), NewAttachTagsRequestMatcher(t, req)).
					Return(&grpc.BloggingEventResponse{EventId: "Event1", ArticleId: "Article1"}, nil).Times(1)
				return bloggingEventServiceClient
			},
			expectedReq: &grpc.AttachTagsRequest{
				Id:       "Article1",
				TagNames: []string{"Tag1"},
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.NewAttachTagsInDTO("Article1", []string{"Tag1"}, "ClientMutationID1"),
			},
			want: want{
				out: dto.NewAttachTagsOutDTO("Event1", "Article1", "ClientMutationID1"),
			},
		},
		"unhappy_path:grpc-return-error": {
			bloggingEventServiceClient: func(t *testing.T, ctrl *gomock.Controller, req *grpc.AttachTagsRequest) grpc.BloggingEventServiceClient {
				bloggingEventServiceClient := mgrpc.NewMockBloggingEventServiceClient(ctrl)
				bloggingEventServiceClient.EXPECT().
					AttachTags(gomock.Any(), NewAttachTagsRequestMatcher(t, req)).
					Return(nil, errTestAttachTags).Times(1)
				return bloggingEventServiceClient
			},
			expectedReq: &grpc.AttachTagsRequest{
				Id:       "Article1",
				TagNames: []string{"Tag1"},
			},
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

func NewAttachTagsRequestMatcher(t *testing.T, expect *grpc.AttachTagsRequest) gomock.Matcher {
	return &AttachTagsRequestMatcher{
		expect: expect,
		t:      t,
	}
}

type AttachTagsRequestMatcher struct {
	gomock.Matcher
	expect *grpc.AttachTagsRequest
	t      *testing.T
}

func (m *AttachTagsRequestMatcher) Matches(x interface{}) bool {
	switch x := x.(type) {
	case *grpc.AttachTagsRequest:
		if x == nil {
			return m.expect == nil
		}
		diff := cmp.Diff(x, m.expect, protocmp.Transform())
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
