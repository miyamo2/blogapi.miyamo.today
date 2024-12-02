package provider

import (
	"github.com/google/wire"
	"github.com/miyamo2/blogapi.miyamo.today/core/grpc/interceptor"
	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"github.com/newrelic/go-agent/v3/newrelic"
	"google.golang.org/grpc"
)

func GRPCServer(nr *newrelic.Application) *grpc.Server {
	return grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			nrgrpc.UnaryServerInterceptor(nr),
			interceptor.SetBlogAPIContextToContext,
			interceptor.SetLoggerToContext(nr)))
}

var GRPCServerSet = wire.NewSet(GRPCServer)
