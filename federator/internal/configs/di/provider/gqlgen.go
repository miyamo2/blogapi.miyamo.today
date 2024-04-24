package provider

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/miyamo2/api.miyamo.today/core/graphql/middleware"
	"github.com/miyamo2/blogapi/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi/internal/if-adapter/controller/graphql/resolver"
	"github.com/miyamo2/blogapi/internal/if-adapter/controller/graphql/resolver/presenter/converter"
	"github.com/miyamo2/blogapi/internal/if-adapter/controller/graphql/resolver/usecase"
	"github.com/miyamo2/blogapi/internal/infra/fw/gqlgen"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/fx"
)

var Gqlgen = fx.Options(
	fx.Provide(func(
		article usecase.Article[dto.ArticleInDto, dto.Tag, dto.ArticleTag, dto.ArticleOutDto],
		articles usecase.Articles[dto.ArticlesInDto, dto.Tag, dto.ArticleTag, dto.ArticlesOutDto],
		tag usecase.Tag[dto.TagInDto, dto.Article, dto.TagArticle, dto.TagOutDto],
		tags usecase.Tags[dto.TagsInDto, dto.Article, dto.TagArticle, dto.TagsOutDto],
	) *resolver.Usecases {
		return resolver.NewUsecases(
			resolver.WithArticlesUsecase(articles),
			resolver.WithArticleUsecase(article),
			resolver.WithTagUsecase(tag),
			resolver.WithTagsUsecase(tags))
	}),
	fx.Provide(func(
		article converter.ArticleConverter[dto.Tag, dto.ArticleTag, dto.ArticleOutDto],
		articles converter.ArticlesConverter[dto.Tag, dto.ArticleTag, dto.ArticlesOutDto],
		tag converter.TagConverter[dto.Article, dto.TagArticle, dto.TagOutDto],
		tags converter.TagsConverter[dto.Article, dto.TagArticle, dto.TagsOutDto],
	) *resolver.Converters {
		return resolver.NewConverters(
			resolver.WithArticleConverter(article),
			resolver.WithArticlesConverter(articles),
			resolver.WithTagConverter(tag),
			resolver.WithTagsConverter(tags))
	}),
	fx.Provide(resolver.NewResolver),
	fx.Provide(func(rslvr *resolver.Resolver) gqlgen.Config {
		return gqlgen.Config{
			Resolvers: rslvr,
		}
	}),
	fx.Provide(gqlgen.NewExecutableSchema),
	fx.Provide(func(schema graphql.ExecutableSchema, nr *newrelic.Application) *handler.Server {
		srv := handler.NewDefaultServer(schema)
		srv.AroundOperations(middleware.StartNewRelicTransaction(nr))
		srv.AroundOperations(middleware.SetBlogAPIContextToContext)
		srv.AroundRootFields(middleware.StartNewRelicSegment)
		srv.AroundOperations(middleware.SetLoggerToContext(nr))
		return srv
	}),
)
