//go:generate mockgen -source=tag.go -destination=../../../../../mock/if-adapter/controller/graphql/resolver/usecase/mock_tag.go -package=usecase
package usecase

import (
	"context"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"
)

// Tag is a use-case of getting a tag by id.
type Tag interface {
	// Execute gets a tag by id.
	Execute(ctx context.Context, in dto.TagInDto) (dto.TagOutDto, error)
}

type Tags interface {
	// Execute gets tags.
	Execute(ctx context.Context, in dto.TagsInDto) (dto.TagsOutDto, error)
}
