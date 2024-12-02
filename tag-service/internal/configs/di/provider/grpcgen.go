package provider

import (
	"github.com/google/wire"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/if-adapter/controller/pb"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/infra/grpc"
)

// compatibility check
var _ grpc.TagServiceServer = (*pb.TagServiceServer)(nil)

var TagServiceSet = wire.NewSet(
	pb.NewTagServiceServer,
	wire.Bind(new(grpc.TagServiceServer), new(*pb.TagServiceServer)),
)
