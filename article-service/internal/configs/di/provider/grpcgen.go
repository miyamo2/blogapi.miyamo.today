package provider

import (
	"github.com/google/wire"
	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/if-adapter/controller/pb"
	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/infra/grpc"
)

// compatibility check
var _ grpc.ArticleServiceServer = (*pb.ArticleServiceServer)(nil)

var ArticleServiceSet = wire.NewSet(
	pb.NewArticleServiceServer,
	wire.Bind(new(grpc.ArticleServiceServer), new(*pb.ArticleServiceServer)),
)
