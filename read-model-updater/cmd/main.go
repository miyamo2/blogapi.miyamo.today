package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	gw "github.com/miyamo2/blogapi.miyamo.today/core/db/gorm"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/if-adapters/model"
	nraws "github.com/newrelic/go-agent/v3/integrations/nrawssdk-v2"
	"github.com/newrelic/go-agent/v3/integrations/nrlambda"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/gorm"
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
	dep := getDependecies()
	if dep == nil {
		panic("failed to initialize dependencies")
	}
	gw.Initialize(dep.gormDB)
	nraws.AppendMiddlewares(&dep.awsConfig.APIOptions, nil)
	nrlambda.Start(dep.synHandler, dep.newrelicApp)
}
