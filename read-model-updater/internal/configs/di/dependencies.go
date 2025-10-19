package di

import (
	"blogapi.miyamo.today/read-model-updater/internal/if-adapters/handler"
	"blogapi.miyamo.today/read-model-updater/internal/infra/queue"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Dependencies struct {
	AWSConfig   *aws.Config
	NewRelicApp *newrelic.Application
	SyncHandler *handler.SyncHandler
	QueueClient queue.Client
	QueueURL    QueueURL
}

type QueueURL *string

func newDependencies(
	awsConfig *aws.Config,
	newrelicApp *newrelic.Application,
	syncHandler *handler.SyncHandler,
	queueClient queue.Client,
	queueURL QueueURL,
) *Dependencies {
	return &Dependencies{
		AWSConfig:   awsConfig,
		NewRelicApp: newrelicApp,
		SyncHandler: syncHandler,
		QueueClient: queueClient,
		QueueURL:    queueURL,
	}
}
