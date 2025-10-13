package di

import (
	"blogapi.miyamo.today/read-model-updater/internal/if-adapters/handler"
	"blogapi.miyamo.today/read-model-updater/internal/infra/streams"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Dependencies struct {
	AWSConfig    *aws.Config
	NewRelicApp  *newrelic.Application
	SyncHandler  *handler.SyncHandler
	StreamClient streams.Client
	StreamARN    StreamARN
}

type StreamARN *string

func newDependencies(
	awsConfig *aws.Config,
	newrelicApp *newrelic.Application,
	syncHandler *handler.SyncHandler,
	streamClient streams.Client,
	streamARN StreamARN,
) *Dependencies {
	return &Dependencies{
		AWSConfig:    awsConfig,
		NewRelicApp:  newrelicApp,
		SyncHandler:  syncHandler,
		StreamClient: streamClient,
		StreamARN:    streamARN,
	}
}
