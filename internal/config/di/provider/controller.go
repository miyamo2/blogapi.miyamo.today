package provider

import (
	impl "github.com/miyamo2/blogapi-article-service/internal/if-adapter/controller/pb"
	"github.com/miyamo2/blogapi-core/grpc/interceptor"
	"github.com/miyamo2/blogproto-gen/article/server/pb"
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
	fx.Provide(func() *grpc.Server {
		return grpc.NewServer(grpc.UnaryInterceptor(interceptor.SetBlogAPIContextToContext))
	}),
	fx.Invoke(func(aSvcSrv pb.ArticleServiceServer, srv *grpc.Server) {
		pb.RegisterArticleServiceServer(srv, aSvcSrv)
	}),
)
