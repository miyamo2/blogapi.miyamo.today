package provider

import (
	"blogapi.miyamo.today/article-service/internal/if-adapter/controller/pb/presenter/convert"
	impl "blogapi.miyamo.today/article-service/internal/if-adapter/presenter/pb/convert"
	"github.com/google/wire"
)

// compatibility check
var _ convert.ListAfter = (*impl.ListAfter)(nil)
var _ convert.ListAll = (*impl.ListAll)(nil)
var _ convert.GetByID = (*impl.GetByID)(nil)
var _ convert.ListBefore = (*impl.ListBefore)(nil)

var PresenterSet = wire.NewSet(
	impl.NewListAfter,
	wire.Bind(new(convert.ListAfter), new(*impl.ListAfter)),
	impl.NewListAll,
	wire.Bind(new(convert.ListAll), new(*impl.ListAll)),
	impl.NewGetByID,
	wire.Bind(new(convert.GetByID), new(*impl.GetByID)),
	impl.NewListBefore,
	wire.Bind(new(convert.ListBefore), new(*impl.ListBefore)),
)
