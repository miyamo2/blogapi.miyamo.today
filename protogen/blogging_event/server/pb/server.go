package pb

import (
	"github.com/miyamo2/blogapi.miyamo.today/protogen/blogging_event/server/pb/internal"
	"google.golang.org/grpc"
)

type BloggingEventServiceServer = internal.BloggingEventServiceServer

type UnimplementedBloggingEventServiceServer = internal.UnimplementedBloggingEventServiceServer

type UnsafeBloggingEventServiceServer = internal.UnsafeBloggingEventServiceServer

func RegisterBloggingEventServiceServer(s grpc.ServiceRegistrar, srv BloggingEventServiceServer) {
	internal.RegisterBloggingEventServiceServer(s, srv)
}
