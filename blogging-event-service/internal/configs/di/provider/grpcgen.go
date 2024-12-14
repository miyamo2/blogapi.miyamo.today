package provider

import (
	"github.com/google/wire"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/if-adapter/controller/pb"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/infra/grpc"
)

// compatibility check
var _ grpc.BloggingEventServiceServer = (*pb.BloggingEventServiceServer)(nil)

var BloggingEventServiceServerSet = wire.NewSet(
	pb.NewBloggingEventServiceServer,
	wire.Bind(new(grpc.BloggingEventServiceServer), new(*pb.BloggingEventServiceServer)),
)
