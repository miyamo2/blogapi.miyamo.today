package provider

import (
	"github.com/google/wire"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/if-adapter/controller/pb"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/infra/grpc/grpcconnect"
)

// compatibility check
var _ grpcconnect.TagServiceHandler = (*pb.TagServiceServer)(nil)

var TagServiceSet = wire.NewSet(
	pb.NewTagServiceServer,
	wire.Bind(new(grpcconnect.TagServiceHandler), new(*pb.TagServiceServer)),
)
