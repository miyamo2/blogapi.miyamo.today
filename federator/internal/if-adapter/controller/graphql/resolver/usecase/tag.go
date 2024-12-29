//go:generate mockgen -source=$GOFILE -destination=../../../../../mock/if-adapter/controller/graphql/resolver/usecase/$GOFILE -package=$GOPACKAGE
package usecase

import (
	"blogapi.miyamo.today/federator/internal/app/usecase/dto"
	"context"
)

// Tag is a use-case of getting a tag by id.
type Tag interface {
	// Execute gets a tag by id.
	Execute(ctx context.Context, in dto.TagInDTO) (dto.TagOutDTO, error)
}

type Tags interface {
	// Execute gets tags.
	Execute(ctx context.Context, in dto.TagsInDTO) (dto.TagsOutDTO, error)
}
