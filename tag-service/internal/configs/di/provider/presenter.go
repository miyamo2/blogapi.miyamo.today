package provider

import (
	"blogapi.miyamo.today/tag-service/internal/if-adapter/controller/pb/presenter/convert"
	impl "blogapi.miyamo.today/tag-service/internal/if-adapter/presenter/pb/convert"
	"github.com/google/wire"
)

// compatibility check
var _ convert.ToGetNext = (*impl.GetNextTags)(nil)
var _ convert.ToGetAll = (*impl.GetAllTags)(nil)
var _ convert.ToGetById = (*impl.GetByIdTag)(nil)
var _ convert.ToGetPrev = (*impl.GetPrevTags)(nil)

var PresenterSet = wire.NewSet(
	impl.NewGetNextTags,
	wire.Bind(new(convert.ToGetNext), new(*impl.GetNextTags)),
	impl.NewGetAllTags,
	wire.Bind(new(convert.ToGetAll), new(*impl.GetAllTags)),
	impl.NewGetByIdTag,
	wire.Bind(new(convert.ToGetById), new(*impl.GetByIdTag)),
	impl.NewGetPrevTags,
	wire.Bind(new(convert.ToGetPrev), new(*impl.GetPrevTags)),
)
