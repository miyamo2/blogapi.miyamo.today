package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	nraws "github.com/newrelic/go-agent/v3/integrations/nrawssdk-v2"
	"github.com/newrelic/go-agent/v3/integrations/nrlambda"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// SyncHandler handles blogging event
type SyncHandler interface {
	Invoke(ctx context.Context, stream events.DynamoDBEvent) error
}

type dependencies struct {
	awsConfig   *aws.Config
	newrelicApp *newrelic.Application
	syncHandler SyncHandler
}

func newDependencies(
	awsConfig *aws.Config,
	newrelicApp *newrelic.Application,
	syncHandler SyncHandler,
) *dependencies {
	return &dependencies{
		awsConfig:   awsConfig,
		newrelicApp: newrelicApp,
		syncHandler: syncHandler,
	}
}

func main() {
	dep := getDependecies()
	if dep == nil {
		panic("failed to initialize dependencies")
	}
	nraws.AppendMiddlewares(&dep.awsConfig.APIOptions, nil)
	nrlambda.Start(dep.syncHandler.Invoke, dep.newrelicApp)
}
