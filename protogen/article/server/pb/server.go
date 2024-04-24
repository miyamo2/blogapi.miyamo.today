package pb

import (
	"github.com/miyamo2/blogapi.miyamo.today/protogen/article/server/pb/internal"
	"google.golang.org/grpc"
)

type ArticleServiceServer = internal.ArticleServiceServer

type UnimplementedArticleServiceServer = internal.UnimplementedArticleServiceServer

type UnsafeArticleServiceServer = internal.UnsafeArticleServiceServer

func RegisterArticleServiceServer(s grpc.ServiceRegistrar, srv ArticleServiceServer) {
	internal.RegisterArticleServiceServer(s, srv)
}
