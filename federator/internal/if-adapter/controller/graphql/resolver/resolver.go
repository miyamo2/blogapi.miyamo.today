package resolver

import (
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/controller/graphql/resolver/presenter/converter"
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
	article  usecase.Article
	articles usecase.Articles
	tag      usecase.Tag
	tags     usecase.Tags
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

// NewUsecases constructor of Usecases.
func NewUsecases(options ...UsecasesOption) *Usecases {
	u := &Usecases{}
	for _, option := range options {
		option(u)
	}
	return u
}

type ConvertersOption func(*Converters)

// WithArticleConverter option for Converters.
func WithArticleConverter(article converter.ArticleConverter) ConvertersOption {
	return func(c *Converters) {
		c.article = article
	}
}

// WithArticlesConverter option for Converters.
func WithArticlesConverter(articles converter.ArticlesConverter) ConvertersOption {
	return func(c *Converters) {
		c.articles = articles
	}
}

// WithTagConverter option for Converters.
func WithTagConverter(tag converter.TagConverter) ConvertersOption {
	return func(c *Converters) {
		c.tag = tag
	}
}

// WithTagsConverter option for Converters.
func WithTagsConverter(tags converter.TagsConverter) ConvertersOption {
	return func(c *Converters) {
		c.tags = tags
	}
}

type Converters struct {
	article  converter.ArticleConverter
	articles converter.ArticlesConverter
	tag      converter.TagConverter
	tags     converter.TagsConverter
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
