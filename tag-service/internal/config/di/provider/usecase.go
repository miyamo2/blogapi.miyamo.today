package provider

import (
	impl "github.com/miyamo2/api.miyamo.today/tag-service/internal/app/usecase"
	"github.com/miyamo2/api.miyamo.today/tag-service/internal/app/usecase/dto"
	"github.com/miyamo2/api.miyamo.today/tag-service/internal/if-adapter/controller/pb/usecase"
	"go.uber.org/fx"
)

// compatibility check
var _ usecase.GetById[dto.GetByIdInDto, dto.Article, *dto.GetByIdOutDto] = (*impl.GetById)(nil)
var _ usecase.GetAll[dto.Article, dto.Tag, *dto.GetAllOutDto] = (*impl.GetAll)(nil)
var _ usecase.GetNext[dto.GetNextInDto, dto.Article, dto.Tag, *dto.GetNextOutDto] = (*impl.GetNext)(nil)
var _ usecase.GetPrev[dto.GetPrevInDto, dto.Article, dto.Tag, *dto.GetPrevOutDto] = (*impl.GetPrev)(nil)

var Usecase = fx.Options(
	fx.Provide(
		fx.Annotate(
			impl.NewGetById,
			fx.As(new(usecase.GetById[dto.GetByIdInDto, dto.Article, *dto.GetByIdOutDto])),
		),
		fx.Annotate(
			impl.NewGetAll,
			fx.As(new(usecase.GetAll[dto.Article, dto.Tag, *dto.GetAllOutDto])),
		),
		fx.Annotate(
			impl.NewGetNext,
			fx.As(new(usecase.GetNext[dto.GetNextInDto, dto.Article, dto.Tag, *dto.GetNextOutDto])),
		),
		fx.Annotate(
			impl.NewGetPrev,
			fx.As(new(usecase.GetPrev[dto.GetPrevInDto, dto.Article, dto.Tag, *dto.GetPrevOutDto])),
		),
	),
)
