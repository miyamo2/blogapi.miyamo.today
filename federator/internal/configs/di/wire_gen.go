// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/configs/di/provider"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/controller/graphql/resolver"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/presenters/graphql/converters"
)

// Injectors from wire.go:

func GetDependencies() *Dependencies {
	articleServiceClient := provider.ArticleClient()
	article := usecase.NewArticle(articleServiceClient)
	articles := usecase.NewArticles(articleServiceClient)
	tagServiceClient := provider.TagClient()
	tag := usecase.NewTag(tagServiceClient)
	tags := usecase.NewTags(tagServiceClient)
	bloggingEventServiceClient := provider.BloggingEventClient()
	createArticle := usecase.NewCreateArticle(bloggingEventServiceClient)
	updateArticleTitle := usecase.NewUpdateArticleTitle(bloggingEventServiceClient)
	updateArticleBody := usecase.NewUpdateArticleBody(bloggingEventServiceClient)
	updateArticleThumbnail := usecase.NewUpdateArticleThumbnail(bloggingEventServiceClient)
	attachTags := usecase.NewAttachTags(bloggingEventServiceClient)
	detachTags := usecase.NewDetachTags(bloggingEventServiceClient)
	uploadImage := usecase.NewUploadImage(bloggingEventServiceClient)
	usecases := provider.Usecases(article, articles, tag, tags, createArticle, updateArticleTitle, updateArticleBody, updateArticleThumbnail, attachTags, detachTags, uploadImage)
	converter := converters.NewConverter()
	resolverConverters := provider.Converters(converter, converter, converter, converter, converter, converter, converter, converter, converter, converter, converter)
	resolverResolver := resolver.NewResolver(usecases, resolverConverters)
	config := provider.GqlgenConfig(resolverResolver)
	executableSchema := provider.GqlgenExecutableSchema(config)
	application := provider.NewRelic()
	server := provider.GqlgenServer(executableSchema, application)
	echo := provider.Echo(server, application)
	dependencies := NewDependencies(echo)
	return dependencies
}
