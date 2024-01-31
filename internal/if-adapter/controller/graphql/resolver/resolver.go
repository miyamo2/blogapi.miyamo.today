package resolver

import (
	"github.com/miyamo2/blogapi/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi/internal/if-adapter/controller/graphql/resolver/presenter/converter"
	"github.com/miyamo2/blogapi/internal/if-adapter/controller/graphql/resolver/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	usecases   *Usecases
	converters *Converters
}

type Usecases struct {
	article  usecase.Article[dto.ArticleInDto, dto.Tag, dto.ArticleTag, dto.ArticleOutDto]
	articles usecase.Articles[dto.ArticlesInDto, dto.Tag, dto.ArticleTag, dto.ArticlesOutDto]
	tag      usecase.Tag[dto.TagInDto, dto.Article, dto.TagArticle, dto.TagOutDto]
	tags     usecase.Tags[dto.TagsInDto, dto.Article, dto.TagArticle, dto.TagsOutDto]
}

type UsecasesOption func(*Usecases)

// WithArticleUsecase option for Usecases.
func WithArticleUsecase(article usecase.Article[dto.ArticleInDto, dto.Tag, dto.ArticleTag, dto.ArticleOutDto]) UsecasesOption {
	return func(u *Usecases) {
		u.article = article
	}
}

// WithArticlesUsecase option for Usecases.
func WithArticlesUsecase(articles usecase.Articles[dto.ArticlesInDto, dto.Tag, dto.ArticleTag, dto.ArticlesOutDto]) UsecasesOption {
	return func(u *Usecases) {
		u.articles = articles
	}
}

// WithTagUsecase option for Usecases.
func WithTagUsecase(tag usecase.Tag[dto.TagInDto, dto.Article, dto.TagArticle, dto.TagOutDto]) UsecasesOption {
	return func(u *Usecases) {
		u.tag = tag
	}
}

// WithTagsUsecase option for Usecases.
func WithTagsUsecase(tags usecase.Tags[dto.TagsInDto, dto.Article, dto.TagArticle, dto.TagsOutDto]) UsecasesOption {
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
func WithArticleConverter(article converter.ArticleConverter[dto.Tag, dto.ArticleTag, dto.ArticleOutDto]) ConvertersOption {
	return func(c *Converters) {
		c.article = article
	}
}

// WithArticlesConverter option for Converters.
func WithArticlesConverter(articles converter.ArticlesConverter[dto.Tag, dto.ArticleTag, dto.ArticlesOutDto]) ConvertersOption {
	return func(c *Converters) {
		c.articles = articles
	}
}

// WithTagConverter option for Converters.
func WithTagConverter(tag converter.TagConverter[dto.Article, dto.TagArticle, dto.TagOutDto]) ConvertersOption {
	return func(c *Converters) {
		c.tag = tag
	}
}

// WithTagsConverter option for Converters.
func WithTagsConverter(tags converter.TagsConverter[dto.Article, dto.TagArticle, dto.TagsOutDto]) ConvertersOption {
	return func(c *Converters) {
		c.tags = tags
	}
}

type Converters struct {
	article  converter.ArticleConverter[dto.Tag, dto.ArticleTag, dto.ArticleOutDto]
	articles converter.ArticlesConverter[dto.Tag, dto.ArticleTag, dto.ArticlesOutDto]
	tag      converter.TagConverter[dto.Article, dto.TagArticle, dto.TagOutDto]
	tags     converter.TagsConverter[dto.Article, dto.TagArticle, dto.TagsOutDto]
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
