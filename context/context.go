package blogapictx

import (
	"context"
	"time"
)

type contextKey struct{}

type TraceIDKey struct{}

type RequestIDKey struct{}

type BlogAPIContext struct {
	Fingerprint Fingerprint
	Tracing     Tracing
	Incoming    Request
	Local       interface{}
	Outgoing    *Request
}

type Fingerprint struct {
	ID string
}

type Tracing struct {
	ID string
}

type RequestType string

const (
	RequestTypeGRPC    RequestType = "grpc"
	RequestTypeHTTP    RequestType = "http"
	RequestTypeGraphQL RequestType = "graphql"
)

type Request struct {
	Type       RequestType
	Service    string  // micro services service name
	GRPCMethod *string // gRPC method, REST path, GraphQL path
	Path       *string // REST path, GraphQL path
	Headers    map[string]string
	StartTime  time.Time
	DurationMS *float32
	Status     *string // HTTP status code
	Body       interface{}
}

// New returns a new BlogAPIContext.
func New(
	ctx context.Context,
	serviceName string,
	requestType RequestType,
	requestHeader map[string]string,
	requestBody interface{},
) BlogAPIContext {
	requestID := ctx.Value(RequestIDKey{}).(string)
	incoming := Request{
		Type:      requestType,
		Service:   serviceName,
		StartTime: time.Now(),
		Headers:   requestHeader,
		Body:      requestBody,
	}
	bctx := BlogAPIContext{
		Incoming: incoming,
		Fingerprint: Fingerprint{
			ID: requestID,
		},
	}
	traceID := ctx.Value(TraceIDKey{}).(string)
	bctx.Tracing = Tracing{
		ID: traceID,
	}
	return bctx
}

// StoreToContext stores the BlogAPIContext in context.Context.
func StoreToContext(ctx context.Context, bctx BlogAPIContext) context.Context {
	return context.WithValue(ctx, contextKey{}, bctx)
}

// FromContext returns the BlogAPIContext stored in context.Context.
func FromContext(ctx context.Context) *BlogAPIContext {
	bctx := ctx.Value(contextKey{})
	if bctx == nil {
		return nil
	}
	return bctx.(*BlogAPIContext)
}
