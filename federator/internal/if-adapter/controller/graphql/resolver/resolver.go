package resolver

import (
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/controller/graphql/resolver/presenter/converters"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/controller/graphql/resolver/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	usecases   *Usecases
	converters *Converters
}

type Usecases struct {
	article            usecase.Article
	articles           usecase.Articles
	tag                usecase.Tag
	tags               usecase.Tags
	createArticle      usecase.CreateArticle
	updateArticleTitle usecase.UpdateArticleTitle
}

type UsecasesOption func(*Usecases)

// WithArticleUsecase option for Usecases.
func WithArticleUsecase(article usecase.Article) UsecasesOption {
	return func(u *Usecases) {
		u.article = article
	}
}

// WithArticlesUsecase option for Usecases.
func WithArticlesUsecase(articles usecase.Articles) UsecasesOption {
	return func(u *Usecases) {
		u.articles = articles
	}
}

// WithTagUsecase option for Usecases.
func WithTagUsecase(tag usecase.Tag) UsecasesOption {
	return func(u *Usecases) {
		u.tag = tag
	}
}

// WithTagsUsecase option for Usecases.
func WithTagsUsecase(tags usecase.Tags) UsecasesOption {
	return func(u *Usecases) {
		u.tags = tags
	}
}

// WithCreateArticleUsecase option for Usecases.
func WithCreateArticleUsecase(createArticle usecase.CreateArticle) UsecasesOption {
	return func(u *Usecases) {
		u.createArticle = createArticle
	}
}

// WithUpdateArticleTitleUsecase option for Usecases.
func WithUpdateArticleTitleUsecase(updateArticleTitle usecase.UpdateArticleTitle) UsecasesOption {
	return func(u *Usecases) {
		u.updateArticleTitle = updateArticleTitle
	}
}

// NewUsecases constructor of Usecases.
func NewUsecases(options ...UsecasesOption) *Usecases {
	u := &Usecases{}
	for _, option := range options {
		option(u)
	}
	return u
}

type Converters struct {
	article            converters.ArticleConverter
	articles           converters.ArticlesConverter
	tag                converters.TagConverter
	tags               converters.TagsConverter
	createArticle      converters.CreateArticleConverter
	updateArticleTitle converters.UpdateArticleTitleConverter
}

type ConvertersOption func(*Converters)

// WithArticleConverter option for Converters.
func WithArticleConverter(article converters.ArticleConverter) ConvertersOption {
	return func(c *Converters) {
		c.article = article
	}
}

// WithArticlesConverter option for Converters.
func WithArticlesConverter(articles converters.ArticlesConverter) ConvertersOption {
	return func(c *Converters) {
		c.articles = articles
	}
}

// WithTagConverter option for Converters.
func WithTagConverter(tag converters.TagConverter) ConvertersOption {
	return func(c *Converters) {
		c.tag = tag
	}
}

// WithTagsConverter option for Converters.
func WithTagsConverter(tags converters.TagsConverter) ConvertersOption {
	return func(c *Converters) {
		c.tags = tags
	}
}

// WithCreateArticleConverter option for Converters.
func WithCreateArticleConverter(createArticle converters.CreateArticleConverter) ConvertersOption {
	return func(c *Converters) {
		c.createArticle = createArticle
	}
}

// WithUpdateArticleTitleConverter option for Converters.
func WithUpdateArticleTitleConverter(updateArticleTitle converters.UpdateArticleTitleConverter) ConvertersOption {
	return func(c *Converters) {
		c.updateArticleTitle = updateArticleTitle
	}
}

// NewConverters constructor of Converters.
func NewConverters(options ...ConvertersOption) *Converters {
	c := &Converters{}
	for _, option := range options {
		option(c)
	}
	return c
}

// New constructor of Resolver.
func NewResolver(usecases *Usecases, converters *Converters) *Resolver {
	return &Resolver{
		usecases:   usecases,
		converters: converters,
	}
}
