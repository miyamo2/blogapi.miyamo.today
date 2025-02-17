package main

import (
	"blogapi.miyamo.today/article-service/internal/configs/di"
	"blogapi.miyamo.today/article-service/internal/infra/grpc"
	gwrapper "blogapi.miyamo.today/core/db/gorm"
	"blogapi.miyamo.today/core/util/tcp"
	"log/slog"
	"os"
	"os/signal"

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

	errChan := make(chan error, 1)
	go func() {
		if err := server.Serve(listener); err != nil {
			errChan <- err
			return
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	select {
	case err := <-errChan:
		slog.Error(err.Error())
	case <-quit:
		slog.Info("stopping gRPC server...")
		server.GracefulStop()
	}
}
