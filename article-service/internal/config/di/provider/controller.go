package provider

import (
	impl "github.com/miyamo2/blogapi.miyamo.today/article-service/internal/if-adapter/controller/pb"
	"github.com/miyamo2/blogapi.miyamo.today/core/grpc/interceptor"
	"github.com/miyamo2/blogapi.miyamo.today/protogen/article/server/pb"
	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

// compatibility check
var _ pb.ArticleServiceServer = (*impl.ArticleServiceServer)(nil)

var Controller = fx.Options(
	fx.Provide(
		fx.Annotate(
			impl.NewArticleServiceServer,
			fx.As(new(pb.ArticleServiceServer)),
		),
	),
	fx.Provide(func(nr *newrelic.Application) *grpc.Server {
		return grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				nrgrpc.UnaryServerInterceptor(nr),
				interceptor.SetBlogAPIContextToContext,
				interceptor.SetLoggerToContext(nr)))
	}),
	fx.Invoke(func(aSvcSrv pb.ArticleServiceServer, srv *grpc.Server) {
		pb.RegisterArticleServiceServer(srv, aSvcSrv)
	}),
)
