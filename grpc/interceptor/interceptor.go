package interceptor

import (
	"context"
	blogapicontext "github.com/miyamo2/blogapi-core/context"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/oklog/ulid/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func SetBlogAPIContextToContext(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	tid := func() string {
		if ok {
			vs := md.Get("trace_id")
			if len(vs) > 0 {
				return vs[0]
			}
		}
		ntx := newrelic.FromContext(ctx)
		return ntx.GetLinkingMetadata().TraceID
	}()
	rid := func() string {
		if ok {
			vs := md.Get("request_id")
			if len(vs) > 0 {
				return vs[0]
			}
		}
		return ulid.Make().String()
	}()
	ctx = blogapicontext.StoreToContext(ctx, blogapicontext.New(tid, rid, info.FullMethod, blogapicontext.RequestTypeGRPC, nil, req))
	res, err := handler(ctx, req)
	return res, err
}
