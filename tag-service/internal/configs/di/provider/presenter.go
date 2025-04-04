package provider

import (
	"blogapi.miyamo.today/tag-service/internal/if-adapter/controller/pb/presenter"
	impl "blogapi.miyamo.today/tag-service/internal/if-adapter/presenter/pb"
	"github.com/google/wire"
)

// compatibility check
var _ presenter.ToGetNextConverter = (*impl.Converter)(nil)
var _ presenter.ToGetAllConverter = (*impl.Converter)(nil)
var _ presenter.ToGetByIdConverter = (*impl.Converter)(nil)
var _ presenter.ToGetPrevConverter = (*impl.Converter)(nil)

var PresenterSet = wire.NewSet(
	impl.NewConverter,
	wire.Bind(new(presenter.ToGetNextConverter), new(*impl.Converter)),
	wire.Bind(new(presenter.ToGetAllConverter), new(*impl.Converter)),
	wire.Bind(new(presenter.ToGetByIdConverter), new(*impl.Converter)),
	wire.Bind(new(presenter.ToGetPrevConverter), new(*impl.Converter)),
)
