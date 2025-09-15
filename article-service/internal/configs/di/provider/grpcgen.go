package provider

import (
	"blogapi.miyamo.today/article-service/internal/if-adapter/controller/pb"
	"blogapi.miyamo.today/article-service/internal/if-adapter/controller/pb/presenter/convert"
	"blogapi.miyamo.today/article-service/internal/if-adapter/controller/pb/usecase"
	"blogapi.miyamo.today/article-service/internal/infra/grpc/grpcconnect"
	"github.com/google/wire"
)

// compatibility check
var _ grpcconnect.ArticleServiceHandler = (*pb.ArticleServiceServer)(nil)

var ArticleServiceSet = wire.NewSet(
	ArticleServiceServer,
	wire.Bind(new(grpcconnect.ArticleServiceHandler), new(*pb.ArticleServiceServer)),
)

func ArticleServiceServer(
	getByIDUsecase usecase.GetByID,
	listAllUsecase usecase.ListAll,
	listAfterUsecase usecase.ListAfter,
	listBeforeUsecase usecase.ListBefore,
	getByIDConverter convert.GetByID,
	listAllConverter convert.ListAll,
	listAfterConverter convert.ListAfter,
	listBeforeConverter convert.ListBefore,
) *pb.ArticleServiceServer {
	return pb.NewArticleServiceServer(
		pb.WithGetByID(getByIDUsecase, getByIDConverter),
		pb.WithListAll(listAllUsecase, listAllConverter),
		pb.WithListAfter(listAfterUsecase, listAfterConverter),
		pb.WithListBefore(listBeforeUsecase, listBeforeConverter),
	)
}
