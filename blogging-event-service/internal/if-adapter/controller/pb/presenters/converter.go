//go:generate mockgen -source=$GOFILE -destination=../../../../mock/if-adapter/controller/pb/presenter/$GOFILE -package=$GOPACKAGE
package presenters

import (
	"context"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/infra/grpc"
)

// ToCreateArticleResponse is a converter interface for converting from CreateArticle use-case's dto to pb response.
type ToCreateArticleResponse interface {
	// ToCreateArticleArticleResponse converts from CreateArticle use-case's dto to pb response.
	ToCreateArticleArticleResponse(ctx context.Context, from *dto.CreateArticleOutDto) (response *grpc.BloggingEventResponse, err error)
}

// ToUpdateArticleTitleResponse is a converter interface for converting from UpdateArticleTitle use-case's dto to pb response.
type ToUpdateArticleTitleResponse interface {
	// ToUpdateArticleTitleResponse converts from UpdateArticleTitle use-case's dto to pb response.
	ToUpdateArticleTitleResponse(ctx context.Context, from *dto.UpdateArticleTitleOutDto) (response *grpc.BloggingEventResponse, err error)
}

// ToUpdateArticleBodyResponse is a converter interface for converting from UpdateArticleBody use-case's dto to pb response.
type ToUpdateArticleBodyResponse interface {
	// ToUpdateArticleBodyResponse converts from UpdateArticleBody use-case's dto to pb response.
	ToUpdateArticleBodyResponse(ctx context.Context, from *dto.UpdateArticleBodyOutDto) (response *grpc.BloggingEventResponse, err error)
}

// ToUpdateArticleThumbnailResponse is a converter interface for converting from UpdateArticleThumbnail use-case's dto to pb response.
type ToUpdateArticleThumbnailResponse interface {
	// ToUpdateArticleThumbnailResponse converts from UpdateArticleThumbnail use-case's dto to pb response.
	ToUpdateArticleThumbnailResponse(ctx context.Context, from *dto.UpdateArticleThumbnailOutDto) (response *grpc.BloggingEventResponse, err error)
}

// ToAttachTagsResponse is a converter interface for converting from AttachTags use-case's dto to pb response.
type ToAttachTagsResponse interface {
	// ToAttachTagsResponse converts from AttachTags use-case's dto to pb response.
	ToAttachTagsResponse(ctx context.Context, from *dto.AttachTagsOutDto) (response *grpc.BloggingEventResponse, err error)
}

// ToDetachTagsResponse is a converter interface for converting from DetachTags use-case's dto to pb response.
type ToDetachTagsResponse interface {
	// ToDetachTagsResponse converts from DetachTags use-case's dto to pb response.
	ToDetachTagsResponse(ctx context.Context, from *dto.DetachTagsOutDto) (response *grpc.BloggingEventResponse, err error)
}
