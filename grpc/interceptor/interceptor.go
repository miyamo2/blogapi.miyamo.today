package interceptor

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	blogapicontext "github.com/miyamo2/blogapi-core/context"
	"github.com/oklog/ulid/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

func SetBlogAPIContextToContext(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	tid := func() string {
		if ok {
			vs := md.Get("trace_id")
			if len(vs) > 0 {
				return vs[0]
			}
		}
		// UUID: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
		// X-Ray Trace ID: 1-xxxxxxxx-xxxxxxxxxxxxxxxxxxxxxxxx
		suuid := strings.ReplaceAll(uuid.New().String(), "-", "")
		return fmt.Sprintf("1-%v-%v", suuid[0:8], suuid[8:])
	}()
	rid := func() string {
		if ok {
			vs := md.Get("request_id")
			if len(vs) > 0 {
				return vs[0]
			}
		}
		return ulid.Make().String()
	}()
	ctx = blogapicontext.StoreToContext(ctx, blogapicontext.New(tid, rid, info.FullMethod, blogapicontext.RequestTypeGRPC, nil, req))
	res, err := handler(ctx, req)
	return res, err
}
