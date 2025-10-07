//go:build wireinject

package di

import (
	"blogapi.miyamo.today/read-model-updater/internal/app/usecase"
	"blogapi.miyamo.today/read-model-updater/internal/app/usecase/command"
	"blogapi.miyamo.today/read-model-updater/internal/app/usecase/externalapi"
	"blogapi.miyamo.today/read-model-updater/internal/app/usecase/query"
	"blogapi.miyamo.today/read-model-updater/internal/if-adapters/converter"
	"blogapi.miyamo.today/read-model-updater/internal/if-adapters/lambda"
	"blogapi.miyamo.today/read-model-updater/internal/infra/dynamo"
	"blogapi.miyamo.today/read-model-updater/internal/infra/githubactions"
	"blogapi.miyamo.today/read-model-updater/internal/infra/rdb"
	"github.com/google/wire"
)

var awsConfigSet = wire.NewSet(provideAWSConfig)

var newRelicSet = wire.NewSet(provideNewRelicApp)

var rdbSet = wire.NewSet(provideRDBGORM)

var dynamodbSet = wire.NewSet(provideDynamoDBGORM)

var queryServiceSet = wire.NewSet(
	dynamo.NewBloggingEventQueryService,
	wire.Bind(new(query.BloggingEventService), new(*dynamo.BloggingEventQueryService)),
)

var commandServiceSet = wire.NewSet(
	rdb.NewArticleCommandService,
	wire.Bind(new(command.Article), new(*rdb.ArticleCommandService)),
	rdb.NewTagCommandService,
	wire.Bind(new(command.Tag), new(*rdb.TagCommandService)),
)

var externalAPISet = wire.NewSet(
	provideBlogPublisher,
	wire.Bind(new(externalapi.BlogPublisher), new(*githubactions.BlogPublisher)),
)

var usecaseSet = wire.NewSet(
	provideSynUsecaseSet,
	wire.Bind(new(lambda.SyncUsecase), new(*usecase.Sync)),
)

var converterSet = wire.NewSet(
	converter.NewConverter,
	wire.Bind(new(lambda.ToSyncUsecaseInDtoConverter), new(*converter.Converter)),
)

var handlerSet = wire.NewSet(
	lambda.NewSyncHandler,
)

var dependenciesSet = wire.NewSet(newDependencies)

func GetDependecies() *Dependencies {
	wire.Build(
		awsConfigSet,
		dynamodbSet,
		rdbSet,
		newRelicSet,
		queryServiceSet,
		commandServiceSet,
		externalAPISet,
		usecaseSet,
		converterSet,
		handlerSet,
		dependenciesSet,
	)
	return nil
}
