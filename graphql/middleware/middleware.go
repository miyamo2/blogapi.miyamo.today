package middleware

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	blogapicontext "github.com/miyamo2/blogapi-core/context"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/oklog/ulid/v2"
	"net/url"
)

func SetBlogAPIContextToContext(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	octx := graphql.GetOperationContext(ctx)
	headers := octx.Headers
	tid := func() string {
		v := headers.Get("trace_id")
		if len(v) > 0 {
			return v
		}
		ntx := newrelic.FromContext(ctx)
		return ntx.GetLinkingMetadata().TraceID
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

func StartNewRelicTransaction(app *newrelic.Application) func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	if app == nil {
		return func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
			return next(ctx)
		}
	}
	return func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		oc := graphql.GetOperationContext(ctx)
		nrtx := newrelic.FromContext(ctx)
		if nrtx == nil {
			nrtx.SetWebRequest(newrelic.WebRequest{
				Header:    oc.Headers,
				URL:       &url.URL{Path: oc.Operation.Name},
				Method:    "POST",
				Transport: newrelic.TransportHTTP})
			nrtx = app.StartTransaction(fmt.Sprintf("POST/ query@GraphQL:%v", oc.Operation.Name))
			defer nrtx.End()
		}
		nrtx.SetName(fmt.Sprintf("%v@GraphQL:%v", nrtx.Name(), oc.Operation.Name))
		ctx = newrelic.NewContext(ctx, nrtx)
		res := next(ctx)
		return res
	}
}

func StartNewRelicSegment(ctx context.Context, next graphql.RootResolver) graphql.Marshaler {
	rslvrName := graphql.GetRootFieldContext(ctx).Object
	txn := newrelic.FromContext(ctx)
	sgm := txn.StartSegment(rslvrName)
	defer sgm.End()
	m := next(ctx)
	return m
}
