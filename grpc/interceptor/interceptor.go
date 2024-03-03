package interceptor

import (
	"context"
	"github.com/miyamo2/altnrslog"

	"github.com/miyamo2/blogapi-core/log"
	"github.com/newrelic/go-agent/v3/newrelic"

	blogapicontext "github.com/miyamo2/blogapi-core/context"
	"github.com/oklog/ulid/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func SetBlogAPIContextToContext(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	defer newrelic.FromContext(ctx).StartSegment("BlogAPICore: Set BlogAPIContext to Context").End()
	md, ok := metadata.FromIncomingContext(ctx)
	rid := func() string {
		if ok {
			vs := md.Get("request_id")
			if len(vs) > 0 {
				return vs[0]
			}
		}
		return ulid.Make().String()
	}()
	ctx = blogapicontext.StoreToContext(ctx, blogapicontext.New(rid, info.FullMethod, blogapicontext.RequestTypeGRPC, nil, req))
	res, err := handler(ctx, req)
	return res, err
}

func SetLoggerToContext(app *newrelic.Application) func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if app == nil {
		return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			return handler(ctx, req)
		}
	}
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		nrtx := newrelic.FromContext(ctx)
		defer nrtx.StartSegment("BlogAPICore: Set Transactional Logger to Context").End()
		lgr := log.New(log.WithAltNRSlogTransactionalHandler(app, nrtx))
		ctx, err := altnrslog.StoreToContext(ctx, lgr)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}
