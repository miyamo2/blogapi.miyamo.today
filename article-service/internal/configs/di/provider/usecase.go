package provider

import (
	impl "github.com/miyamo2/blogapi.miyamo.today/article-service/internal/app/usecase"
	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/if-adapter/controller/pb/usecase"
	"go.uber.org/fx"
)

// compatibility check
var (
	_ usecase.GetById = (*impl.GetById)(nil)
	_ usecase.GetAll  = (*impl.GetAll)(nil)
	_ usecase.GetNext = (*impl.GetNext)(nil)
	_ usecase.GetPrev = (*impl.GetPrev)(nil)
)

var Usecase = fx.Options(
	fx.Provide(
		fx.Annotate(
			impl.NewGetById,
			fx.As(new(usecase.GetById)),
		),
		fx.Annotate(
			impl.NewGetAll,
			fx.As(new(usecase.GetAll)),
		),
		fx.Annotate(
			impl.NewGetNext,
			fx.As(new(usecase.GetNext)),
		),
		fx.Annotate(
			impl.NewGetPrev,
			fx.As(new(usecase.GetPrev)),
		),
	),
)
