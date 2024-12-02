package provider

import (
	"fmt"
	"github.com/google/wire"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/infra/grpc/article"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/infra/grpc/tag"
	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"os"
)

func ArticleClient() article.ArticleServiceClient {
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
	return article.NewArticleServiceClient(conn)
}

func TagClient() tag.TagServiceClient {
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
	return tag.NewTagServiceClient(conn)
}

var GRPCClientSet = wire.NewSet(ArticleClient, TagClient)
