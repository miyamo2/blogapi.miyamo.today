package main

import (
	"blogapi.miyamo.today/blogging-event-service/internal/configs/di"
	"blogapi.miyamo.today/blogging-event-service/internal/infra/grpc"
	gwrapper "blogapi.miyamo.today/core/db/gorm"
	"blogapi.miyamo.today/core/util/tcp"
	nraws "github.com/newrelic/go-agent/v3/integrations/nrawssdk-v2"
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

	nraws.AppendMiddlewares(&dependencies.AWSConfig.APIOptions, nil)

	server := dependencies.GRPCServer
	service := dependencies.BloggingEventServiceServer
	grpc.RegisterBloggingEventServiceServer(server, service)

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
