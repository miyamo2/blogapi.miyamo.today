package provider

import (
	"github.com/miyamo2/blogapi-article-service/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi-article-service/internal/if-adapter/controller/pb/presenter"
	impl "github.com/miyamo2/blogapi-article-service/internal/if-adapter/presenter/pb"
	"go.uber.org/fx"
)

// compatibility check
var _ presenter.ToGetNextConverter[dto.Tag, dto.Article, *dto.GetNextOutDto] = (*impl.Converter)(nil)
var _ presenter.ToGetAllConverter[dto.Tag, dto.Article, *dto.GetAllOutDto] = (*impl.Converter)(nil)
var _ presenter.ToGetByIdConverter[dto.Tag, *dto.GetByIdOutDto] = (*impl.Converter)(nil)

var Presenter = fx.Options(
	fx.Provide(
		fx.Annotate(
			impl.NewConverter,
			fx.As(new(presenter.ToGetNextConverter[dto.Tag, dto.Article, *dto.GetNextOutDto])),
		),
		fx.Annotate(
			impl.NewConverter,
			fx.As(new(presenter.ToGetAllConverter[dto.Tag, dto.Article, *dto.GetAllOutDto])),
		),
		fx.Annotate(
			impl.NewConverter,
			fx.As(new(presenter.ToGetByIdConverter[dto.Tag, *dto.GetByIdOutDto])),
		),
	),
)
