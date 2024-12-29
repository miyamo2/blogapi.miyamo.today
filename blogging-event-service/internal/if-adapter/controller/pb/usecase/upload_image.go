//go:generate mockgen -source=$GOFILE -destination=../../../../mock/if-adapter/controller/pb/usecase/$GOFILE -package=$GOPACKAGE
package usecase

import (
	"blogapi.miyamo.today/blogging-event-service/internal/app/usecase/dto"
	"context"
)

// UploadImage is an interface for uploading image
type UploadImage interface {
	// Execute uploads an image
	Execute(ctx context.Context, in *dto.UploadImageInDto) (*dto.UploadImageOutDto, error)
}
