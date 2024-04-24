package middleware

import (
	"context"
	"fmt"
	"net/url"

	"github.com/99designs/gqlgen/graphql"
	"github.com/miyamo2/altnrslog"
	blogapicontext "github.com/miyamo2/api.miyamo.today/core/context"
	"github.com/miyamo2/api.miyamo.today/core/log"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/oklog/ulid/v2"
)

func SetBlogAPIContextToContext(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	seg := newrelic.FromContext(ctx).StartSegment("BlogAPICore: Set BlogAPIContext to Context")
	octx := graphql.GetOperationContext(ctx)
	headers := octx.Headers
	rid := func() string {
		v := headers.Get("x-request-id")
		if len(v) > 0 {
			return v
		}
		return ulid.Make().String()
	}()
	ctx = blogapicontext.StoreToContext(ctx, blogapicontext.New(rid, octx.OperationName, blogapicontext.RequestTypeGraphQL, headers, octx.Variables))
	seg.End()
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

func SetLoggerToContext(app *newrelic.Application) func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	if app == nil {
		return func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
			return next(ctx)
		}
	}
	return func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		nrtx := newrelic.FromContext(ctx)
		seg := nrtx.StartSegment("BlogAPICore: Set Transactional Logger to Context")
		lgr := log.New(log.WithAltNRSlogTransactionalHandler(app, nrtx))
		ctx, err := altnrslog.StoreToContext(ctx, lgr)
		if err != nil {
			er := graphql.ErrorResponse(ctx, err.Error())
			return func(ctx context.Context) *graphql.Response { return er }
		}
		seg.End()
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
