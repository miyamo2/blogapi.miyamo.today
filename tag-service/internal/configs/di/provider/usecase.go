package provider

import (
	"github.com/google/wire"
	"github.com/miyamo2/blogapi.miyamo.today/core/db/gorm"
	impl "github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/app/usecase"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/app/usecase/query"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/if-adapter/controller/pb/usecase"
)

// compatibility check
var (
	_ usecase.GetById = (*impl.GetById)(nil)
	_ usecase.GetAll  = (*impl.GetAll)(nil)
	_ usecase.GetNext = (*impl.GetNext)(nil)
	_ usecase.GetPrev = (*impl.GetPrev)(nil)
)

func GetByIdUsecase(queryService query.TagService) *impl.GetById {
	return impl.NewGetById(gorm.Manager(), queryService)
}

func GetAllUsecase(queryService query.TagService) *impl.GetAll {
	return impl.NewGetAll(gorm.Manager(), queryService)
}

func GetNextUsecase(queryService query.TagService) *impl.GetNext {
	return impl.NewGetNext(gorm.Manager(), queryService)
}

func GetPrevUsecase(queryService query.TagService) *impl.GetPrev {
	return impl.NewGetPrev(gorm.Manager(), queryService)
}

var UsecaseSet = wire.NewSet(
	GetByIdUsecase,
	wire.Bind(new(usecase.GetById), new(*impl.GetById)),
	GetAllUsecase,
	wire.Bind(new(usecase.GetAll), new(*impl.GetAll)),
	GetNextUsecase,
	wire.Bind(new(usecase.GetNext), new(*impl.GetNext)),
	GetPrevUsecase,
	wire.Bind(new(usecase.GetPrev), new(*impl.GetPrev)),
)
