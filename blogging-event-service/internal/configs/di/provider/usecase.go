package provider

import (
	"github.com/google/wire"
	impl "github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/app/usecase"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/app/usecase/command"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/if-adapter/controller/pb/usecase"
)

// compatibility check
var (
	_ usecase.CreateArticle      = (*impl.CreateArticle)(nil)
	_ usecase.UpdateArticleTitle = (*impl.UpdateArticleTitle)(nil)
	_ usecase.UpdateArticleBody  = (*impl.UpdateArticleBody)(nil)
)

func CreateArticleUsecase(bloggingEventCommand command.BloggingEventService) *impl.CreateArticle {
	return impl.NewCreateArticle(bloggingEventCommand)
}

func UpdateArticleTitleUsecase(bloggingEventCommand command.BloggingEventService) *impl.UpdateArticleTitle {
	return impl.NewUpdateArticleTitle(bloggingEventCommand)
}

func UpdateArticleBodyUsecase(bloggingEventCommand command.BloggingEventService) *impl.UpdateArticleBody {
	return impl.NewUpdateArticleBody(bloggingEventCommand)
}

var UsecaseSet = wire.NewSet(
	CreateArticleUsecase,
	wire.Bind(new(usecase.CreateArticle), new(*impl.CreateArticle)),
	UpdateArticleTitleUsecase,
	wire.Bind(new(usecase.UpdateArticleTitle), new(*impl.UpdateArticleTitle)),
	UpdateArticleBodyUsecase,
	wire.Bind(new(usecase.UpdateArticleBody), new(*impl.UpdateArticleBody)),
)
