//go:generate mockgen -source=$GOFILE -destination=../../../../../../mock/if-adapter/controller/graphql/resolver/presenter/converter/$GOFILE -package=$GOPACKAGE
package converters

import (
	"context"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"

	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/presenters/graphql/model"
)

// ArticleConverter is the converter for an article.
type ArticleConverter interface {
	// ToArticle converts an article.
	ToArticle(ctx context.Context, from dto.ArticleOutDTO) (*model.ArticleNode, bool)
}

// ArticlesConverter is the converter for articles.
type ArticlesConverter interface {
	// ToArticles converts articles.
	ToArticles(ctx context.Context, from dto.ArticlesOutDTO) (*model.ArticleConnection, bool)
}

// TagConverter is the converter for a tag.
type TagConverter interface {
	// ToTag converts a tag.
	ToTag(ctx context.Context, from dto.TagOutDTO) (*model.TagNode, error)
}

// TagsConverter is the converter for tags.
type TagsConverter interface {
	// ToTags converts tags.
	ToTags(ctx context.Context, from dto.TagsOutDTO) (*model.TagConnection, error)
}

// CreateArticleConverter is the converter for creating an article.
type CreateArticleConverter interface {
	// ToCreateArticle converts creating an article.
	ToCreateArticle(ctx context.Context, from dto.CreateArticleOutDTO) (*model.CreateArticlePayload, error)
}

// UpdateArticleTitleConverter is the converter for updating an article title.
type UpdateArticleTitleConverter interface {
	// ToUpdateArticleTitle converts updating an article title.
	ToUpdateArticleTitle(ctx context.Context, from dto.UpdateArticleTitleOutDTO) (*model.UpdateArticleTitlePayload, error)
}

// UpdateArticleBodyConverter is the converter for updating an article body.
type UpdateArticleBodyConverter interface {
	// ToUpdateArticleBody converts updating an article body.
	ToUpdateArticleBody(ctx context.Context, from dto.UpdateArticleBodyOutDTO) (*model.UpdateArticleBodyPayload, error)
}

// UpdateArticleThumbnailConverter is the converter for updating an article thumbnail.
type UpdateArticleThumbnailConverter interface {
	// ToUpdateArticleThumbnail converts updating an article thumbnail.
	ToUpdateArticleThumbnail(ctx context.Context, from dto.UpdateArticleThumbnailOutDTO) (*model.UpdateArticleThumbnailPayload, error)
}

// AttachTagsConverter is the converter for attaching tags.
type AttachTagsConverter interface {
	// ToAttachTags converts attaching tags.
	ToAttachTags(ctx context.Context, from dto.AttachTagsOutDTO) (*model.AttachTagsPayload, error)
}

// DetachTagsConverter is the converter for detaching tags.
type DetachTagsConverter interface {
	// ToDetachTags converts detaching tags.
	ToDetachTags(ctx context.Context, from dto.DetachTagsOutDTO) (*model.DetachTagsPayload, error)
}

// UploadImageConverter is the converter for uploading an image.
type UploadImageConverter interface {
	// ToUploadImage converts uploading an image.
	ToUploadImage(ctx context.Context, from dto.UploadImageOutDTO) (*model.UploadImagePayload, error)
}
