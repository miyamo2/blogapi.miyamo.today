//go:build wireinject

package di

import (
	"blogapi.miyamo.today/read-model-updater/internal/app/usecase"
	"blogapi.miyamo.today/read-model-updater/internal/app/usecase/command"
	"blogapi.miyamo.today/read-model-updater/internal/app/usecase/externalapi"
	"blogapi.miyamo.today/read-model-updater/internal/app/usecase/query"
	"blogapi.miyamo.today/read-model-updater/internal/if-adapters/converter"
	"blogapi.miyamo.today/read-model-updater/internal/if-adapters/handler"
	"blogapi.miyamo.today/read-model-updater/internal/infra/dynamo"
	"blogapi.miyamo.today/read-model-updater/internal/infra/githubactions"
	"blogapi.miyamo.today/read-model-updater/internal/infra/queue"
	"blogapi.miyamo.today/read-model-updater/internal/infra/rdb/sqlc/article"
	"blogapi.miyamo.today/read-model-updater/internal/infra/rdb/sqlc/tag"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/google/wire"
)

var awsConfigSet = wire.NewSet(provideAWSConfig)

var newRelicSet = wire.NewSet(provideNewRelicApp)

var rdbSet = wire.NewSet(provideArticleDBPool, provideTagDBPool)

var dynamodbSet = wire.NewSet(provideDynamoDB)

var queryServiceSet = wire.NewSet(
	dynamo.NewBloggingEventQueryService,
	wire.Bind(new(query.BloggingEventService), new(*dynamo.BloggingEventQueryService)),
)

var commandSet = wire.NewSet(
	provideArticleQuery,
	wire.Bind(new(command.Article), new(*article.Queries)),
	provideTagQuery,
	wire.Bind(new(command.Tag), new(*tag.Queries)),
)

var txSet = wire.NewSet(
	command.NewArticleTx,
	command.NewTagTx,
)

var externalAPISet = wire.NewSet(
	provideBlogPublisher,
	wire.Bind(new(externalapi.BlogPublisher), new(*githubactions.BlogPublisher)),
)

var usecaseSet = wire.NewSet(
	provideSynUsecaseSet,
	wire.Bind(new(handler.SyncUsecase), new(*usecase.Sync)),
)

var converterSet = wire.NewSet(
	converter.NewConverter,
	wire.Bind(new(handler.ToSyncUsecaseInDtoConverter), new(*converter.Converter)),
)

var handlerSet = wire.NewSet(
	handler.NewSyncHandler,
)

var queueSet = wire.NewSet(
	provideSQSClient,
	wire.Bind(new(queue.Client), new(*sqs.Client)),
)

var queueURLSet = wire.NewSet(provideQueueURL)

var dependenciesSet = wire.NewSet(newDependencies)

func GetDependecies() *Dependencies {
	wire.Build(
		awsConfigSet,
		dynamodbSet,
		queueSet,
		queueURLSet,
		rdbSet,
		commandSet,
		txSet,
		newRelicSet,
		queryServiceSet,
		externalAPISet,
		usecaseSet,
		converterSet,
		handlerSet,
		dependenciesSet,
	)
	return nil
}
