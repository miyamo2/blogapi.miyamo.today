package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cockroachdb/errors"
	"golang.org/x/net/http2"

	"blogapi.miyamo.today/tag-service/internal/configs/di"
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

	e := di.GetEchoApp()

	errChan := make(chan error, 1)
	go func() {
		slog.Info("start gRPC server.", slog.String("port", port))
		if err := e.StartH2CServer(fmt.Sprintf(":%s", port), &http2.Server{}); err != nil {
			errChan <- err
			return
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	select {
	case err := <-errChan:
		if !errors.Is(err, http.ErrServerClosed) {
			slog.Error(err.Error())
		}
	case <-quit:
		slog.Info("stopping gRPC server...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		e.Shutdown(ctx)
	}
}
