package main

import (
	"blogapi.miyamo.today/blogging-event-service/internal/configs/di"
	gwrapper "blogapi.miyamo.today/core/db/gorm"
	"context"
	"fmt"
	"github.com/cockroachdb/errors"
	nraws "github.com/newrelic/go-agent/v3/integrations/nrawssdk-v2"
	"golang.org/x/net/http2"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

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

	dependencies := di.GetDependencies()
	gormDialector := dependencies.GORMDialector
	gwrapper.InitializeDialector(gormDialector)
	nraws.AppendMiddlewares(&dependencies.AWSConfig.APIOptions, nil)

	e := dependencies.Echo

	errChan := make(chan error, 1)
	go func() {
		slog.Info("start graphql server.", slog.String("port", port))
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
		slog.Info("stopping graphql server...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		e.Shutdown(ctx)
	}
}
