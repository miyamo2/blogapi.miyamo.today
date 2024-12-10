package main

import (
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/configs/di"
	nraws "github.com/newrelic/go-agent/v3/integrations/nrawssdk-v2"
	"github.com/newrelic/go-agent/v3/integrations/nrlambda"
)

func main() {
	dep := di.GetDependecies()
	if dep == nil {
		panic("failed to initialize Dependencies")
	}
	nraws.AppendMiddlewares(&dep.AWSConfig.APIOptions, nil)
	nrlambda.Start(dep.SyncHandler.Invoke, dep.NewRelicApp)
}
