package provider

import (
	"blogapi.miyamo.today/blogging-event-service/internal/if-adapter/controller/pb/presenters"
	impl "blogapi.miyamo.today/blogging-event-service/internal/if-adapter/presenter/pb"
	"github.com/google/wire"
)

// compatibility check
var (
	_ presenters.ToCreateArticleResponse          = (*impl.Converter)(nil)
	_ presenters.ToUpdateArticleTitleResponse     = (*impl.Converter)(nil)
	_ presenters.ToUpdateArticleBodyResponse      = (*impl.Converter)(nil)
	_ presenters.ToUpdateArticleThumbnailResponse = (*impl.Converter)(nil)
	_ presenters.ToAttachTagsResponse             = (*impl.Converter)(nil)
	_ presenters.ToDetachTagsResponse             = (*impl.Converter)(nil)
	_ presenters.ToUploadImageResponse            = (*impl.Converter)(nil)
)

var PresenterSet = wire.NewSet(
	impl.NewConverter,
	wire.Bind(new(presenters.ToCreateArticleResponse), new(*impl.Converter)),
	wire.Bind(new(presenters.ToUpdateArticleTitleResponse), new(*impl.Converter)),
	wire.Bind(new(presenters.ToUpdateArticleBodyResponse), new(*impl.Converter)),
	wire.Bind(new(presenters.ToUpdateArticleThumbnailResponse), new(*impl.Converter)),
	wire.Bind(new(presenters.ToAttachTagsResponse), new(*impl.Converter)),
	wire.Bind(new(presenters.ToDetachTagsResponse), new(*impl.Converter)),
	wire.Bind(new(presenters.ToUploadImageResponse), new(*impl.Converter)),
)
