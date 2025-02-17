//go:generate mockgen -source=$GOFILE -destination=../../../../../mock/if-adapter/controller/graphql/resolver/usecase/$GOFILE -package=$GOPACKAGE
package usecase

import (
	"blogapi.miyamo.today/federator/internal/app/usecase/dto"
	"context"
)

// CreateArticle is the usecase for creating an article.
type CreateArticle interface {
	// Execute creates an article.
	Execute(ctx context.Context, in dto.CreateArticleInDTO) (dto.CreateArticleOutDTO, error)
}

// UpdateArticleTitle is the usecase for updating an article title.
type UpdateArticleTitle interface {
	// Execute updates an article title.
	Execute(ctx context.Context, in dto.UpdateArticleTitleInDTO) (dto.UpdateArticleTitleOutDTO, error)
}

// UpdateArticleBody is the usecase for updating an article body.
type UpdateArticleBody interface {
	// Execute updates an article body.
	Execute(ctx context.Context, in dto.UpdateArticleBodyInDTO) (dto.UpdateArticleBodyOutDTO, error)
}

// UpdateArticleThumbnail is the usecase for updating an article thumbnail.
type UpdateArticleThumbnail interface {
	// Execute updates an article thumbnail.
	Execute(ctx context.Context, in dto.UpdateArticleThumbnailInDTO) (dto.UpdateArticleThumbnailOutDTO, error)
}

// AttachTags is a use-case for attaching tags to an article.
type AttachTags interface {
	// Execute attaches tags to an article.
	Execute(ctx context.Context, in dto.AttachTagsInDTO) (dto.AttachTagsOutDTO, error)
}

// DetachTags is a use-case for detaching tags from an article.
type DetachTags interface {
	// Execute detaches tags from an article.
	Execute(ctx context.Context, in dto.DetachTagsInDTO) (dto.DetachTagsOutDTO, error)
}

// UploadImage is a use-case for uploading an image.
type UploadImage interface {
	// Execute uploads an image.
	Execute(ctx context.Context, in dto.UploadImageInDTO) (dto.UploadImageOutDTO, error)
}
