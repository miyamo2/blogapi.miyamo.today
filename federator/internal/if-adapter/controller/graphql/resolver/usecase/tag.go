//go:generate mockgen -source=tag.go -destination=../../../../../mock/if-adapter/controller/graphql/resolver/usecase/mock_tag.go -package=usecase
package usecase

import (
	"context"

	"github.com/miyamo2/api.miyamo.today/federator/internal/if-adapter/controller/graphql/resolver/usecase/dto"
)

// Tag is a use-case of getting a tag by id.
type Tag[I dto.TagInDto, A dto.Article, TA dto.TagArticle[A], O dto.TagOutDto[A, TA]] interface {
	// Execute gets a tag by id.
	Execute(ctx context.Context, in I) (O, error)
}

type Tags[I dto.TagsInDto, A dto.Article, TA dto.TagArticle[A], O dto.TagsOutDto[A, TA]] interface {
	// Execute gets tags.
	Execute(ctx context.Context, in I) (O, error)
}
