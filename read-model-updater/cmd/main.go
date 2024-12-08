package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/if-adapters/model"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/gorm"
	"log/slog"
)

// SyncHandler handles blogging event
type SyncHandler interface {
	Invoke(ctx context.Context, stream model.EventStream) error
}

type dependencies struct {
	awsConfig   *aws.Config
	newrelicApp *newrelic.Application
	gormDB      *gorm.DB
	synHandler  SyncHandler
}

func newDependencies(
	awsConfig *aws.Config,
	newrelicApp *newrelic.Application,
	gormDB *gorm.DB,
	syncHandler SyncHandler,
) *dependencies {
	return &dependencies{
		awsConfig:   awsConfig,
		newrelicApp: newrelicApp,
		gormDB:      gormDB,
		synHandler:  syncHandler,
	}
}

func main() {
	lambda.Start(func(event events.DynamoDBEvent) {
		slog.Info("receive", slog.Any("event", event))
	})
	//dep := getDependecies()
	//gw.Initialize(dep.gormDB)
	//nraws.AppendMiddlewares(&dep.awsConfig.APIOptions, nil)
	//nrlambda.Start(dep.synHandler, dep.newrelicApp)
}
