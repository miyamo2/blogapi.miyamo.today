package provider

import (
	"github.com/miyamo2/blogapi-tag-service/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi-tag-service/internal/if-adapter/controller/pb/presenter"
	impl "github.com/miyamo2/blogapi-tag-service/internal/if-adapter/presenter/pb"
	"go.uber.org/fx"
)

// compatibility check
var _ presenter.ToGetByIdConverter[dto.Article, *dto.GetByIdOutDto] = (*impl.Converter)(nil)
var _ presenter.ToGetAllConverter[dto.Article, dto.Tag, *dto.GetAllOutDto] = (*impl.Converter)(nil)
var _ presenter.ToGetNextConverter[dto.Article, dto.Tag, *dto.GetNextOutDto] = (*impl.Converter)(nil)
var _ presenter.ToGetPrevConverter[dto.Article, dto.Tag, *dto.GetPrevOutDto] = (*impl.Converter)(nil)

var Presenter = fx.Options(
	fx.Provide(
		fx.Annotate(
			impl.NewConverter,
			fx.As(new(presenter.ToGetByIdConverter[dto.Article, *dto.GetByIdOutDto])),
		),
		fx.Annotate(
			impl.NewConverter,
			fx.As(new(presenter.ToGetAllConverter[dto.Article, dto.Tag, *dto.GetAllOutDto])),
		),
		fx.Annotate(
			impl.NewConverter,
			fx.As(new(presenter.ToGetNextConverter[dto.Article, dto.Tag, *dto.GetNextOutDto])),
		),
		fx.Annotate(
			impl.NewConverter,
			fx.As(new(presenter.ToGetPrevConverter[dto.Article, dto.Tag, *dto.GetPrevOutDto])),
		),
	),
)
