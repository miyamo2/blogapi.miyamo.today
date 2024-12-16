package provider

import (
	"github.com/google/wire"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/if-adapter/controller/pb/presenters"
	impl "github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/if-adapter/presenter/pb"
)

// compatibility check
var (
	_ presenters.ToCreateArticleResponse      = (*impl.Converter)(nil)
	_ presenters.ToUpdateArticleTitleResponse = (*impl.Converter)(nil)
)

var PresenterSet = wire.NewSet(
	impl.NewConverter,
	wire.Bind(new(presenters.ToCreateArticleResponse), new(*impl.Converter)),
	wire.Bind(new(presenters.ToUpdateArticleTitleResponse), new(*impl.Converter)),
)
