package resolver

import (
	"blogapi.miyamo.today/federator/internal/if-adapter/controller/graphql/resolver/presenter/converters"
	"blogapi.miyamo.today/federator/internal/if-adapter/controller/graphql/resolver/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	usecases   *Usecases
	converters *Converters
}

type Usecases struct {
	article                usecase.Article
	articles               usecase.Articles
	tag                    usecase.Tag
	tags                   usecase.Tags
	createArticle          usecase.CreateArticle
	updateArticleTitle     usecase.UpdateArticleTitle
	updateArticleBody      usecase.UpdateArticleBody
	updateArticleThumbnail usecase.UpdateArticleThumbnail
	attachTags             usecase.AttachTags
	detachTags             usecase.DetachTags
	uploadImage            usecase.UploadImage
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

// WithUpdateArticleBodyUsecase option for Usecases.
func WithUpdateArticleBodyUsecase(updateArticleBody usecase.UpdateArticleBody) UsecasesOption {
	return func(u *Usecases) {
		u.updateArticleBody = updateArticleBody
	}
}

// WithUpdateArticleThumbnailUsecase option for Usecases.
func WithUpdateArticleThumbnailUsecase(updateArticleThumbnail usecase.UpdateArticleThumbnail) UsecasesOption {
	return func(u *Usecases) {
		u.updateArticleThumbnail = updateArticleThumbnail
	}
}

// WithAttachTagsUsecase option for Usecases.
func WithAttachTagsUsecase(attachTags usecase.AttachTags) UsecasesOption {
	return func(u *Usecases) {
		u.attachTags = attachTags
	}
}

// WithDetachTagsUsecase option for Usecases.
func WithDetachTagsUsecase(detachTags usecase.DetachTags) UsecasesOption {
	return func(u *Usecases) {
		u.detachTags = detachTags
	}
}

// WithUploadImageUsecase option for Usecases.
func WithUploadImageUsecase(uploadImage usecase.UploadImage) UsecasesOption {
	return func(u *Usecases) {
		u.uploadImage = uploadImage
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
	article                converters.ArticleConverter
	articles               converters.ArticlesConverter
	tag                    converters.TagConverter
	tags                   converters.TagsConverter
	createArticle          converters.CreateArticleConverter
	updateArticleTitle     converters.UpdateArticleTitleConverter
	updateArticleBody      converters.UpdateArticleBodyConverter
	updateArticleThumbnail converters.UpdateArticleThumbnailConverter
	attachTags             converters.AttachTagsConverter
	detachTags             converters.DetachTagsConverter
	uploadImage            converters.UploadImageConverter
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

// WithUpdateArticleBodyConverter option for Converters.
func WithUpdateArticleBodyConverter(updateArticleBody converters.UpdateArticleBodyConverter) ConvertersOption {
	return func(c *Converters) {
		c.updateArticleBody = updateArticleBody
	}
}

// WithUpdateArticleThumbnailConverter option for Converters.
func WithUpdateArticleThumbnailConverter(updateArticleThumbnail converters.UpdateArticleThumbnailConverter) ConvertersOption {
	return func(c *Converters) {
		c.updateArticleThumbnail = updateArticleThumbnail
	}
}

// WithAttachTagsConverter option for Converters.
func WithAttachTagsConverter(attachTags converters.AttachTagsConverter) ConvertersOption {
	return func(c *Converters) {
		c.attachTags = attachTags
	}
}

// WithDetachTagsConverter option for Converters.
func WithDetachTagsConverter(detachTags converters.DetachTagsConverter) ConvertersOption {
	return func(c *Converters) {
		c.detachTags = detachTags
	}
}

// WithUploadImageConverter option for Converters.
func WithUploadImageConverter(uploadImage converters.UploadImageConverter) ConvertersOption {
	return func(c *Converters) {
		c.uploadImage = uploadImage
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
