package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"blogapi.miyamo.today/federator/internal/configs/di"
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

	dep := di.GetDependencies()
	e := dep.Echo

	errChan := make(chan error, 1)
	go func() {
		slog.Info("start graphql server.", slog.String("port", port))
		if err := e.Start(fmt.Sprintf(":%s", port)); err != nil {
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
		slog.Info("stopping graphql server...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		e.Shutdown(ctx)
	}
}
