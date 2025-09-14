package provider

import (
	"blogapi.miyamo.today/tag-service/internal/if-adapter/controller/pb"
	"blogapi.miyamo.today/tag-service/internal/if-adapter/controller/pb/presenter/convert"
	"blogapi.miyamo.today/tag-service/internal/if-adapter/controller/pb/usecase"
	"blogapi.miyamo.today/tag-service/internal/infra/grpc/grpcconnect"
	"github.com/google/wire"
)

// compatibility check
var _ grpcconnect.TagServiceHandler = (*pb.TagServiceServer)(nil)

func TagServiceServer(
	getByIdUsecase usecase.GetById,
	getByIdConverter convert.ToGetById,
	listAllUsecase usecase.ListAll,
	getAllConverter convert.ToGetAll,
	listAfterUsecase usecase.ListAfter,
	getNextConverter convert.ToGetNext,
	listBeforeUsecase usecase.ListBefore,
	getPrevConverter convert.ToGetPrev,
) *pb.TagServiceServer {
	return pb.NewTagServiceServer(
		pb.WithGetById(getByIdUsecase, getByIdConverter),
		pb.WithListAll(listAllUsecase, getAllConverter),
		pb.WithListAfter(listAfterUsecase, getNextConverter),
		pb.WithListBefore(listBeforeUsecase, getPrevConverter),
	)
}

var TagServiceSet = wire.NewSet(
	TagServiceServer,
	wire.Bind(new(grpcconnect.TagServiceHandler), new(*pb.TagServiceServer)),
)
