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

func GetByIdUsecase(qs query.TagService) *impl.GetById {
	return impl.NewGetById(gorm.Manager(), qs)
}

func GetAllUsecase(qs query.TagService) *impl.GetAll {
	return impl.NewGetAll(gorm.Manager(), qs)
}

func GetNextUsecase(qs query.TagService) *impl.GetNext {
	return impl.NewGetNext(gorm.Manager(), qs)
}

func GetPrevUsecase(qs query.TagService) *impl.GetPrev {
	return impl.NewGetPrev(gorm.Manager(), qs)
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
