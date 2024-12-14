package provider

import (
	"github.com/google/wire"
	impl "github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/app/usecase"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/app/usecase/command"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/if-adapter/controller/pb/usecase"
)

// compatibility check
var (
	_ usecase.CreateArticle = (*impl.CreateArticle)(nil)
)

func CreateArticleUsecase(bloggingEventCommand command.BloggingEventService) *impl.CreateArticle {
	return impl.NewCreateArticle(bloggingEventCommand)
}

var UsecaseSet = wire.NewSet(
	CreateArticleUsecase,
	wire.Bind(new(usecase.CreateArticle), new(*impl.CreateArticle)),
)
