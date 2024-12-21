//go:generate mockgen -source=$GOFILE -destination=../../../../mock/if-adapter/controller/pb/usecase/$GOFILE -package=$GOPACKAGE
package usecase

import (
	"context"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/app/usecase/dto"
)

// DetachTags is a use-case interface for detaching tags from an article.
type DetachTags interface {
	// Execute detaches tags from an article.
	Execute(ctx context.Context, in *dto.DetachTagsInDto) (*dto.DetachTagsOutDto, error)
}
