//go:generate mockgen -source=$GOFILE -destination=../../../../mock/if-adapter/controller/pb/usecase/$GOFILE -package=$GOPACKAGE
package usecase

import (
	"context"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/app/usecase/dto"
)

type UpdateArticleThumbnail interface {
	Execute(ctx context.Context, in *dto.UpdateArticleThumbnailInDto) (*dto.UpdateArticleThumbnailOutDto, error)
}
