package middleware

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	blogapicontext "github.com/miyamo2/blogapi-core/context"
	"github.com/oklog/ulid/v2"
	"strings"
)

func SetBlogAPIContextToContext(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	octx := graphql.GetOperationContext(ctx)
	headers := octx.Headers
	tid := func() string {
		v := headers.Get("trace_id")
		if len(v) > 0 {
			return v
		}
		// UUID: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
		// X-Ray Trace ID: 1-xxxxxxxx-xxxxxxxxxxxxxxxxxxxxxxxx
		suuid := strings.ReplaceAll(uuid.New().String(), "-", "")
		return fmt.Sprintf("1-%v-%v", suuid[0:8], suuid[8:])
	}()
	rid := func() string {
		v := headers.Get("request_id")
		if len(v) > 0 {
			return v
		}
		return ulid.Make().String()
	}()
	ctx = blogapicontext.StoreToContext(ctx, blogapicontext.New(tid, rid, octx.OperationName, blogapicontext.RequestTypeGraphQL, headers, octx.Variables))
	return next(ctx)
}
