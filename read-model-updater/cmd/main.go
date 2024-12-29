package main

import (
	"blogapi.miyamo.today/core/log"
	"blogapi.miyamo.today/read-model-updater/internal/configs/di"
	nraws "github.com/newrelic/go-agent/v3/integrations/nrawssdk-v2"
	"github.com/newrelic/go-agent/v3/integrations/nrlambda"
	"log/slog"
)

func main() {
	dep := di.GetDependecies()
	if dep == nil {
		panic("failed to initialize Dependencies")
	}
	slog.SetDefault(log.New())
	nraws.AppendMiddlewares(&dep.AWSConfig.APIOptions, nil)
	nrlambda.Start(dep.SyncHandler.Invoke, dep.NewRelicApp)
}
