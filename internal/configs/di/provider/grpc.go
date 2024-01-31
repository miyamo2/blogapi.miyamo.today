package provider

import (
	"log/slog"
	"os"

	apb "github.com/miyamo2/blogproto-gen/article/client/pb"
	tpb "github.com/miyamo2/blogproto-gen/tag/client/pb"
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
					grpc.WithTransportCredentials(insecure.NewCredentials()),
					grpc.WithBlock(),
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
					grpc.WithTransportCredentials(insecure.NewCredentials()),
					grpc.WithBlock(),
					grpc.WithUnaryInterceptor(nrgrpc.UnaryClientInterceptor),
				)
				if err != nil {
					slog.Info(err.Error())
				}
				slog.Info("grpc connection established")
				return tpb.NewTagServiceClient(conn)
			}, fx.As(new(tpb.TagServiceClient)))),
)
