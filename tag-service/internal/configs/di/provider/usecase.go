package provider

import (
	impl "blogapi.miyamo.today/tag-service/internal/app/usecase"
	"blogapi.miyamo.today/tag-service/internal/if-adapter/controller/pb/usecase"
	"github.com/google/wire"
)

// compatibility check
var (
	_ usecase.GetById    = (*impl.GetById)(nil)
	_ usecase.ListAll    = (*impl.ListAll)(nil)
	_ usecase.ListAfter  = (*impl.ListAfter)(nil)
	_ usecase.ListBefore = (*impl.ListBefore)(nil)
)

var UsecaseSet = wire.NewSet(
	impl.NewGetById,
	wire.Bind(new(usecase.GetById), new(*impl.GetById)),
	impl.NewListAll,
	wire.Bind(new(usecase.ListAll), new(*impl.ListAll)),
	impl.NewListAfter,
	wire.Bind(new(usecase.ListAfter), new(*impl.ListAfter)),
	impl.NewListBefore,
	wire.Bind(new(usecase.ListBefore), new(*impl.ListBefore)),
)
