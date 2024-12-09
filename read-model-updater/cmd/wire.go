//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/app/usecase"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/app/usecase/command"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/app/usecase/externalapi"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/app/usecase/query"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/if-adapters/converter"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/if-adapters/lambda"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/infra/dynamo"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/infra/githubactions"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/infra/rdb"
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
	wire.Bind(new(command.ArticleService), new(*rdb.ArticleCommandService)),
	rdb.NewTagCommandService,
	wire.Bind(new(command.TagService), new(*rdb.TagCommandService)),
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
	wire.Bind(new(SyncHandler), new(*lambda.SyncHandler)),
)

var dependenciesSet = wire.NewSet(newDependencies)

func getDependecies() *dependencies {
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
