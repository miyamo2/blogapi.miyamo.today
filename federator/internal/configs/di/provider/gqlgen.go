package provider

import (
	"blogapi.miyamo.today/core/graphql/middleware"
	"blogapi.miyamo.today/federator/internal/if-adapter/controller/graphql/resolver"
	"blogapi.miyamo.today/federator/internal/if-adapter/controller/graphql/resolver/presenter/converters"
	"blogapi.miyamo.today/federator/internal/if-adapter/controller/graphql/resolver/usecase"
	"blogapi.miyamo.today/federator/internal/infra/fw/gqlgen"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/google/wire"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/vektah/gqlparser/v2/ast"
	"time"
)

func Usecases(
	article usecase.Article,
	articles usecase.Articles,
	tag usecase.Tag,
	tags usecase.Tags,
	createArticle usecase.CreateArticle,
	updateArticleTiele usecase.UpdateArticleTitle,
	updateArticleBody usecase.UpdateArticleBody,
	updateArticleThumbnail usecase.UpdateArticleThumbnail,
	attachTags usecase.AttachTags,
	detachTags usecase.DetachTags,
	uploadImage usecase.UploadImage,
) *resolver.Usecases {
	return resolver.NewUsecases(
		resolver.WithArticlesUsecase(articles),
		resolver.WithArticleUsecase(article),
		resolver.WithTagUsecase(tag),
		resolver.WithTagsUsecase(tags),
		resolver.WithCreateArticleUsecase(createArticle),
		resolver.WithUpdateArticleTitleUsecase(updateArticleTiele),
		resolver.WithUpdateArticleBodyUsecase(updateArticleBody),
		resolver.WithUpdateArticleThumbnailUsecase(updateArticleThumbnail),
		resolver.WithAttachTagsUsecase(attachTags),
		resolver.WithDetachTagsUsecase(detachTags),
		resolver.WithUploadImageUsecase(uploadImage))
}

func Converters(
	article converters.ArticleConverter,
	articles converters.ArticlesConverter,
	tag converters.TagConverter,
	tags converters.TagsConverter,
	createArticle converters.CreateArticleConverter,
	updateArticleTitle converters.UpdateArticleTitleConverter,
	updateArticleBodyConverter converters.UpdateArticleBodyConverter,
	updateArticleThumbnail converters.UpdateArticleThumbnailConverter,
	attachTags converters.AttachTagsConverter,
	detachTags converters.DetachTagsConverter,
	uploadImage converters.UploadImageConverter,
) *resolver.Converters {
	return resolver.NewConverters(
		resolver.WithArticleConverter(article),
		resolver.WithArticlesConverter(articles),
		resolver.WithTagConverter(tag),
		resolver.WithTagsConverter(tags),
		resolver.WithCreateArticleConverter(createArticle),
		resolver.WithUpdateArticleTitleConverter(updateArticleTitle),
		resolver.WithUpdateArticleBodyConverter(updateArticleBodyConverter),
		resolver.WithUpdateArticleThumbnailConverter(updateArticleThumbnail),
		resolver.WithAttachTagsConverter(attachTags),
		resolver.WithDetachTagsConverter(detachTags),
		resolver.WithUploadImageConverter(uploadImage))
}

func GqlgenConfig(resolver *resolver.Resolver) *gqlgen.Config {
	return &gqlgen.Config{
		Resolvers: resolver,
	}
}

func GqlgenExecutableSchema(config *gqlgen.Config) *graphql.ExecutableSchema {
	xschema := gqlgen.NewExecutableSchema(*config)
	return &xschema
}

func GqlgenServer(schema *graphql.ExecutableSchema, nr *newrelic.Application) *handler.Server {
	srv := handler.New(*schema)

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{
		AllowedMethods: []string{"OPTIONS", "GET", "POST"},
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{
		MaxUploadSize: 3 << 30,
		MaxMemory:     3 << 30,
	})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	srv.AroundOperations(middleware.StartNewRelicTransaction(nr))
	srv.AroundOperations(middleware.SetBlogAPIContextToContext)
	srv.AroundRootFields(middleware.StartNewRelicSegment)
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
