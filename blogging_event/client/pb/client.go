package pb

import (
	"github.com/miyamo2/blogproto-gen/blogging_event/client/pb/internal"
	"google.golang.org/grpc"
)

type BloggingEventServiceClient = internal.BloggingEventServiceClient

func NewBloggingEventServiceClient(cc grpc.ClientConnInterface) BloggingEventServiceClient {
	return internal.NewBloggingEventServiceClient(cc)
}
