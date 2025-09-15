package provider

import (
	impl "blogapi.miyamo.today/article-service/internal/app/usecase"
	"blogapi.miyamo.today/article-service/internal/app/usecase/query"
	"blogapi.miyamo.today/article-service/internal/if-adapter/controller/pb/usecase"
	"blogapi.miyamo.today/core/db/gorm"
	"github.com/google/wire"
)

// compatibility check
var (
	_ usecase.GetById = (*impl.GetById)(nil)
	_ usecase.GetAll  = (*impl.GetAll)(nil)
	_ usecase.GetNext = (*impl.GetNext)(nil)
	_ usecase.GetPrev = (*impl.GetPrev)(nil)
)

func GetByIdUsecase(queryService query.ArticleService) *impl.GetById {
	return impl.NewGetById(gorm.Manager(), queryService)
}

func GetAllUsecase(queryService query.ArticleService) *impl.GetAll {
	return impl.NewGetAll(gorm.Manager(), queryService)
}

func GetNextUsecase(queryService query.ArticleService) *impl.GetNext {
	return impl.NewGetNext(gorm.Manager(), queryService)
}

func GetPrevUsecase(queryService query.ArticleService) *impl.GetPrev {
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
