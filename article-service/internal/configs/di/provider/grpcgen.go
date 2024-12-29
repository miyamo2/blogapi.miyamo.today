package provider

import (
	"blogapi.miyamo.today/article-service/internal/if-adapter/controller/pb"
	"blogapi.miyamo.today/article-service/internal/infra/grpc"
	"github.com/google/wire"
)

// compatibility check
var _ grpc.ArticleServiceServer = (*pb.ArticleServiceServer)(nil)

var ArticleServiceSet = wire.NewSet(
	pb.NewArticleServiceServer,
	wire.Bind(new(grpc.ArticleServiceServer), new(*pb.ArticleServiceServer)),
)
