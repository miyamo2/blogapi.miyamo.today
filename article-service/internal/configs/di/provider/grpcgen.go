package provider

import (
	"blogapi.miyamo.today/article-service/internal/if-adapter/controller/pb"
	"blogapi.miyamo.today/article-service/internal/infra/grpc/grpcconnect"
	"github.com/google/wire"
)

// compatibility check
var _ grpcconnect.ArticleServiceHandler = (*pb.ArticleServiceServer)(nil)

var ArticleServiceSet = wire.NewSet(
	pb.NewArticleServiceServer,
	wire.Bind(new(grpcconnect.ArticleServiceHandler), new(*pb.ArticleServiceServer)),
)
