package provider

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/google/wire"
	"github.com/miyamo2/blogapi.miyamo.today/core/graphql/middleware"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/controller/graphql/resolver"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/controller/graphql/resolver/presenter/converter"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/controller/graphql/resolver/usecase"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/infra/fw/gqlgen"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func Usecases(
	article usecase.Article,
	articles usecase.Articles,
	tag usecase.Tag,
	tags usecase.Tags,
) *resolver.Usecases {
	return resolver.NewUsecases(
		resolver.WithArticlesUsecase(articles),
		resolver.WithArticleUsecase(article),
		resolver.WithTagUsecase(tag),
		resolver.WithTagsUsecase(tags))
}

func Converters(
	article converter.ArticleConverter,
	articles converter.ArticlesConverter,
	tag converter.TagConverter,
	tags converter.TagsConverter,
) *resolver.Converters {
	return resolver.NewConverters(
		resolver.WithArticleConverter(article),
		resolver.WithArticlesConverter(articles),
		resolver.WithTagConverter(tag),
		resolver.WithTagsConverter(tags))
}

func GqlgenConfig(rslvr *resolver.Resolver) *gqlgen.Config {
	return &gqlgen.Config{
		Resolvers: rslvr,
	}
}

func GqlgenExecutableSchema(config *gqlgen.Config) *graphql.ExecutableSchema {
	xschema := gqlgen.NewExecutableSchema(*config)
	return &xschema
}

func GqlgenServer(schema *graphql.ExecutableSchema, nr *newrelic.Application) *handler.Server {
	srv := handler.New(*schema)
	srv.AroundOperations(middleware.StartNewRelicTransaction(nr))
	srv.AroundOperations(middleware.SetBlogAPIContextToContext)
	srv.AroundRootFields(middleware.StartNewRelicSegment)
	srv.AroundOperations(middleware.SetLoggerToContext(nr))
	return srv
}

var GqlgenSet = wire.NewSet(
	Usecases,
	Converters,
	resolver.NewResolver,
	GqlgenConfig,
	GqlgenExecutableSchema,
	GqlgenServer,
)
