//go:generate mockgen -source=$GOFILE -destination=../../../../mock/if-adapter/controller/pb/usecase/$GOFILE -package=$GOPACKAGE
package usecase

import (
	"blogapi.miyamo.today/blogging-event-service/internal/app/usecase/dto"
	"context"
)

// AttachTags is a use-case interface for attaching tags to an article.
type AttachTags interface {
	// Execute attaches tags to an article.
	Execute(ctx context.Context, in *dto.AttachTagsInDto) (*dto.AttachTagsOutDto, error)
}
