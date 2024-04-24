package pb

import (
	"github.com/miyamo2/blogapi.miyamo.today/protogen/blogging_event/client/pb/internal"
	"google.golang.org/grpc"
)

type BloggingEventServiceClient = internal.BloggingEventServiceClient

func NewBloggingEventServiceClient(cc grpc.ClientConnInterface) BloggingEventServiceClient {
	return internal.NewBloggingEventServiceClient(cc)
}
