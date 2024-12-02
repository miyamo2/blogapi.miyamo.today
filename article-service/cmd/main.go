package main

import (
	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/configs/di"
	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/infra/grpc"
	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/infra/tcp"
	gwrapper "github.com/miyamo2/blogapi.miyamo.today/core/db/gorm"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc/health"
	hpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func init() {
	// only for local development
	_ = godotenv.Load(".env")
}

func main() {
	dependencies := di.GetDependencies()
	gormDialector := dependencies.GORMDialector
	gwrapper.InitializeDialector(gormDialector)

	server := dependencies.GRPCServer
	service := dependencies.ArticleServiceServer
	grpc.RegisterArticleServiceServer(server, service)

	listener := tcp.MustListen(tcp.WithPort(os.Getenv("PORT")))
	slog.Info("start gRPC server.", slog.String("address", listener.Addr().String()))
	reflection.Register(server)
	hSrv := health.NewServer()
	hpb.RegisterHealthServer(server, hSrv)
	hSrv.SetServingStatus(os.Getenv("SERVICE_NAME"), hpb.HealthCheckResponse_SERVING)
	go func() {
		if err := server.Serve(listener); err != nil {
			slog.Info(err.Error())
		}
	}()
}
