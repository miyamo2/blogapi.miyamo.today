package pb

import (
	"github.com/miyamo2/blogapi.miyamo.today/protogen/tag/server/pb/internal"
	"google.golang.org/grpc"
)

type TagServiceServer internal.TagServiceServer

type UnimplementedTagServiceServer internal.UnimplementedTagServiceServer

type UnsafeTagServiceServer internal.UnsafeTagServiceServer

func RegisterTagServiceServer(s grpc.ServiceRegistrar, srv TagServiceServer) {
	internal.RegisterTagServiceServer(s, srv)
}
