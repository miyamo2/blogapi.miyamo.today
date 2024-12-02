package provider

import (
	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/if-adapter/controller/pb/presenter"
	impl "github.com/miyamo2/blogapi.miyamo.today/article-service/internal/if-adapter/presenter/pb"
	"go.uber.org/fx"
)

// compatibility check
var _ presenter.ToGetNextConverter = (*impl.Converter)(nil)
var _ presenter.ToGetAllConverter = (*impl.Converter)(nil)
var _ presenter.ToGetByIdConverter = (*impl.Converter)(nil)
var _ presenter.ToGetPrevConverter = (*impl.Converter)(nil)

var Presenter = fx.Options(
	fx.Provide(
		fx.Annotate(
			impl.NewConverter,
			fx.As(new(presenter.ToGetNextConverter)),
		),
		fx.Annotate(
			impl.NewConverter,
			fx.As(new(presenter.ToGetAllConverter)),
		),
		fx.Annotate(
			impl.NewConverter,
			fx.As(new(presenter.ToGetByIdConverter)),
		),
		fx.Annotate(
			impl.NewConverter,
			fx.As(new(presenter.ToGetPrevConverter)),
		)),
)
