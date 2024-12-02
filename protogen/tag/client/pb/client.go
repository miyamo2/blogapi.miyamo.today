package pb

import (
	"github.com/miyamo2/blogapi.miyamo.today/protogen/tag/client/pb/internal"
	"google.golang.org/grpc"
)

type TagServiceClient internal.TagServiceClient

func NewTagServiceClient(cc grpc.ClientConnInterface) TagServiceClient {
	return internal.NewTagServiceClient(cc)
}
