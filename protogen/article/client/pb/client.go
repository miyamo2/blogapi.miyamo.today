package pb

import (
	"github.com/miyamo2/api.miyamo.today/protogen/article/client/pb/internal"
	"google.golang.org/grpc"
)

type ArticleServiceClient = internal.ArticleServiceClient

func NewArticleServiceClient(cc grpc.ClientConnInterface) ArticleServiceClient {
	return internal.NewArticleServiceClient(cc)
}
