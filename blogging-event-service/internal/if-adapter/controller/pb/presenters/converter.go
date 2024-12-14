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
