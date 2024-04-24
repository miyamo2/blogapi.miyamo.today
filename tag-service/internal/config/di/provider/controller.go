package provider

import (
	"github.com/miyamo2/blogapi-core/grpc/interceptor"
	impl "github.com/miyamo2/blogapi-tag-service/internal/if-adapter/controller/pb"
	"github.com/miyamo2/blogproto-gen/tag/server/pb"
	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

// compatibility check
var _ pb.TagServiceServer = (*impl.TagServiceServer)(nil)

var Controller = fx.Options(
	fx.Provide(
		fx.Annotate(
			impl.NewTagServiceServer,
			fx.As(new(pb.TagServiceServer)),
		),
	),
	fx.Provide(func(nr *newrelic.Application) *grpc.Server {
		return grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				nrgrpc.UnaryServerInterceptor(nr),
				interceptor.SetBlogAPIContextToContext,
				interceptor.SetLoggerToContext(nr)))
	}),
	fx.Invoke(func(aSvcSrv pb.TagServiceServer, srv *grpc.Server) {
		pb.RegisterTagServiceServer(srv, aSvcSrv)
	}),
)
