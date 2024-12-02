//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/configs/di/provider"
)

func GetDependencies() *Dependencies {
	wire.Build(
		provider.NewRelicSet,
		provider.GormSet,
		provider.QueryServiceSet,
		provider.PresenterSet,
		provider.UsecaseSet,
		provider.ArticleServiceSet,
		provider.GRPCServerSet,
		wire.NewSet(NewDependencies),
	)
	return nil
}
