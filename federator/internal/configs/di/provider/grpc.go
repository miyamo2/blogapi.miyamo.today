package provider

import (
	"fmt"
	"log/slog"
	"os"

	"google.golang.org/grpc/balancer/roundrobin"

	apb "github.com/miyamo2/api.miyamo.today/protogen/article/client/pb"
	tpb "github.com/miyamo2/api.miyamo.today/protogen/tag/client/pb"
	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Grpc = fx.Options(
	fx.Provide(
		fx.Annotate(
			func(nr *newrelic.Application) apb.ArticleServiceClient {
				address := os.Getenv("ARTICLE_SERVICE_ADDRESS")
				conn, err := grpc.Dial(
					address,
					grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
					grpc.WithTransportCredentials(insecure.NewCredentials()),
					grpc.WithUnaryInterceptor(nrgrpc.UnaryClientInterceptor),
				)
				if err != nil {
					slog.Info(err.Error())
				}
				slog.Info("grpc connection established")
				return apb.NewArticleServiceClient(conn)
			}, fx.As(new(apb.ArticleServiceClient)))),
	fx.Provide(
		fx.Annotate(
			func(nr *newrelic.Application) tpb.TagServiceClient {
				address := os.Getenv("TAG_SERVICE_ADDRESS")
				conn, err := grpc.Dial(
					address,
					grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
					grpc.WithTransportCredentials(insecure.NewCredentials()),
					grpc.WithUnaryInterceptor(nrgrpc.UnaryClientInterceptor),
				)
				if err != nil {
					slog.Info(err.Error())
				}
				slog.Info("grpc connection established")
				return tpb.NewTagServiceClient(conn)
			}, fx.As(new(tpb.TagServiceClient)))),
)
