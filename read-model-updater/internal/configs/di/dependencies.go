package di

import (
	"blogapi.miyamo.today/read-model-updater/internal/if-adapters/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Dependencies struct {
	AWSConfig   *aws.Config
	NewRelicApp *newrelic.Application
	SyncHandler *lambda.SyncHandler
}

func newDependencies(
	awsConfig *aws.Config,
	newrelicApp *newrelic.Application,
	syncHandler *lambda.SyncHandler,
) *Dependencies {
	return &Dependencies{
		AWSConfig:   awsConfig,
		NewRelicApp: newrelicApp,
		SyncHandler: syncHandler,
	}
}
