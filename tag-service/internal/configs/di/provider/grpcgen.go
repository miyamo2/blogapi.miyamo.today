package provider

import (
	"blogapi.miyamo.today/tag-service/internal/if-adapter/controller/pb"
	"blogapi.miyamo.today/tag-service/internal/infra/grpc/grpcconnect"
	"github.com/google/wire"
)

// compatibility check
var _ grpcconnect.TagServiceHandler = (*pb.TagServiceServer)(nil)

var TagServiceSet = wire.NewSet(
	pb.NewTagServiceServer,
	wire.Bind(new(grpcconnect.TagServiceHandler), new(*pb.TagServiceServer)),
)
