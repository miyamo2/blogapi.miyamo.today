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

func TestDetachTags_Execute(t *testing.T) {
	type args struct {
		ctx context.Context
		in  dto.DetachTagsInDTO
	}
	type want struct {
		out dto.DetachTagsOutDTO
		err error
	}
	type testCase struct {
		bloggingEventServiceClient func(t *testing.T, ctrl *gomock.Controller, req *grpc.DetachTagsRequest) grpc.BloggingEventServiceClient
		args                       args
		expectedReq                *grpc.DetachTagsRequest
		want                       want
	}
	errTestDetachTags := errors.New("test error")
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
			bloggingEventServiceClient: func(t *testing.T, ctrl *gomock.Controller, req *grpc.DetachTagsRequest) grpc.BloggingEventServiceClient {
				bloggingEventServiceClient := mgrpc.NewMockBloggingEventServiceClient(ctrl)
				bloggingEventServiceClient.EXPECT().
					DetachTags(gomock.Any(), NewDetachTagsRequestMatcher(t, req)).
					Return(&grpc.BloggingEventResponse{EventId: "Event1", ArticleId: "Article1"}, nil).Times(1)
				return bloggingEventServiceClient
			},
			expectedReq: &grpc.DetachTagsRequest{
				Id:       "Article1",
				TagNames: []string{"Tag1"},
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.NewDetachTagsInDTO("Article1", []string{"Tag1"}, "ClientMutationID1"),
			},
			want: want{
				out: dto.NewDetachTagsOutDTO("Event1", "Article1", "ClientMutationID1"),
			},
		},
		"unhappy_path:grpc-return-error": {
			bloggingEventServiceClient: func(t *testing.T, ctrl *gomock.Controller, req *grpc.DetachTagsRequest) grpc.BloggingEventServiceClient {
				bloggingEventServiceClient := mgrpc.NewMockBloggingEventServiceClient(ctrl)
				bloggingEventServiceClient.EXPECT().
					DetachTags(gomock.Any(), NewDetachTagsRequestMatcher(t, req)).
					Return(nil, errTestDetachTags).Times(1)
				return bloggingEventServiceClient
			},
			expectedReq: &grpc.DetachTagsRequest{
				Id:       "Article1",
				TagNames: []string{"Tag1"},
			},
			args: args{
				ctx: mockBlogAPIContext(),
				in:  dto.NewDetachTagsInDTO("Article1", []string{"Tag1"}, "ClientMutationID1"),
			},
			want: want{
				err: errTestDetachTags,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			uc := NewDetachTags(tt.bloggingEventServiceClient(t, ctrl, tt.expectedReq))
			out, err := uc.Execute(tt.args.ctx, tt.args.in)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("Execute() err = %v, want %v", err, tt.want.err)
			}
			if diff := cmp.Diff(tt.want.out, out, cmp.AllowUnexported((dto.DetachTagsOutDTO{}))); diff != "" {
				t.Errorf("Execute() out = %v, want %v", out, tt.want.out)
			}
		})
	}
}

func NewDetachTagsRequestMatcher(t *testing.T, expect *grpc.DetachTagsRequest) gomock.Matcher {
	return &DetachTagsRequestMatcher{
		expect: expect,
		t:      t,
	}
}

type DetachTagsRequestMatcher struct {
	gomock.Matcher
	expect *grpc.DetachTagsRequest
	t      *testing.T
}

func (m *DetachTagsRequestMatcher) Matches(x interface{}) bool {
	switch x := x.(type) {
	case *grpc.DetachTagsRequest:
		if x == nil {
			return m.expect == nil
		}
		diff := cmp.Diff(x, m.expect, protocmp.Transform())
		if diff != "" {
			m.t.Errorf("DetachTagsRequest mismatch (-want +got):\n%s", diff)
			return false
		}
		return true
	}
	return false
}

func (m *DetachTagsRequestMatcher) String() string {
	return fmt.Sprintf("is equal to %+v", m.expect)
}
