package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/miyamo2/blogapi/internal/configs/di"
	"go.uber.org/fx"

	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func init() {
	// only for local development
	_ = godotenv.Load(".env")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	fx.New(
		di.Container,
		fx.Invoke(
			func(lc fx.Lifecycle, e *echo.Echo) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						go func() {
							slog.Info(e.Start(fmt.Sprintf(":%s", port)).Error())
						}()
						return nil
					},
					OnStop: func(ctx context.Context) error {
						err := e.Shutdown(ctx)
						if err != nil {
							panic(err)
						}
						slog.InfoContext(ctx, "Stopped Echo server")
						return nil
					},
				})
			}),
	).Run()
}
