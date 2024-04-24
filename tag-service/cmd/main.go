package main

import (
	"context"
	"log/slog"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/config/di"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	hpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func init() {
	// only for local development
	_ = godotenv.Load(".env")
}

func main() {
	fx.New(
		di.Container,
		fx.Invoke(
			func(lc fx.Lifecycle, listener net.Listener, srv *grpc.Server) {
				lc.Append(fx.Hook{
					OnStart: func(context.Context) error {
						slog.Info("start gRPC server.", slog.String("address", listener.Addr().String()))
						reflection.Register(srv)
						hSrv := health.NewServer()
						hpb.RegisterHealthServer(srv, hSrv)
						hSrv.SetServingStatus(os.Getenv("SERVICE_NAME"), hpb.HealthCheckResponse_SERVING)
						go func() {
							if err := srv.Serve(listener); err != nil {
								slog.Info(err.Error())
							}
						}()
						return nil
					},
					OnStop: func(ctx context.Context) error {
						slog.Info("stopping gRPC server...")
						srv.GracefulStop()
						return nil
					},
				})
			}),
	).Run()
}
