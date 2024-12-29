//go:build wireinject

package di

import (
	"blogapi.miyamo.today/article-service/internal/configs/di/provider"
	"github.com/google/wire"
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
