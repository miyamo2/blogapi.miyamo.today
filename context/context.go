package blogapictx

import (
	"context"
	"os"
	"time"
)

type contextKey struct{}

type BlogAPIContext struct {
	TraceID  string
	SpanID   string
	Incoming Request
	Outgoing *Request
}

type RequestType string

const (
	RequestTypeGRPC    RequestType = "grpc"
	RequestTypeHTTP    RequestType = "http"
	RequestTypeGraphQL RequestType = "graphql"
)

type Request struct {
	Type       RequestType
	Service    string // micro services service name
	Path       string // gRPC method, REST path, GraphQL path
	Headers    map[string][]string
	StartTime  time.Time
	DurationMS *float32
	Status     *string
	Body       interface{}
}

// New returns a new BlogAPIContext.
func New(
	traceID string,
	spanID string,
	path string,
	requestType RequestType,
	requestHeader map[string][]string,
	requestBody interface{},
) BlogAPIContext {
	incoming := Request{
		Type:      requestType,
		Service:   os.Getenv("SERVICE_NAME"),
		Path:      path,
		StartTime: time.Now(),
		Headers:   requestHeader,
		Body:      requestBody,
	}
	bctx := BlogAPIContext{
		TraceID:  traceID,
		SpanID:   spanID,
		Incoming: incoming,
	}
	return bctx
}

// StoreToContext stores the BlogAPIContext in context.Context.
func StoreToContext(ctx context.Context, bctx BlogAPIContext) context.Context {
	return context.WithValue(ctx, contextKey{}, bctx)
}

// FromContext returns the BlogAPIContext stored in context.Context.
func FromContext(ctx context.Context) *BlogAPIContext {
	bctx, ok := ctx.Value(contextKey{}).(BlogAPIContext)
	if !ok {
		return nil
	}

	return &bctx
}
