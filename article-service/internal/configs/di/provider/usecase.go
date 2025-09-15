package provider

import (
	impl "blogapi.miyamo.today/article-service/internal/app/usecase"
	"blogapi.miyamo.today/article-service/internal/if-adapter/controller/pb/usecase"
	"github.com/google/wire"
)

// compatibility check
var (
	_ usecase.GetByID    = (*impl.GetByID)(nil)
	_ usecase.ListAll    = (*impl.ListAll)(nil)
	_ usecase.ListAfter  = (*impl.ListAfter)(nil)
	_ usecase.ListBefore = (*impl.ListBefore)(nil)
)

var UsecaseSet = wire.NewSet(
	impl.NewGetByID,
	wire.Bind(new(usecase.GetByID), new(*impl.GetByID)),
	impl.NewListAll,
	wire.Bind(new(usecase.ListAll), new(*impl.ListAll)),
	impl.NewListAfter,
	wire.Bind(new(usecase.ListAfter), new(*impl.ListAfter)),
	impl.NewListBefore,
	wire.Bind(new(usecase.ListBefore), new(*impl.ListBefore)),
)
