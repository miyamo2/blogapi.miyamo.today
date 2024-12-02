//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/configs/di/provider"
)

func GetDependencies() *Dependencies {
	wire.Build(
		provider.NewRelicSet,
		provider.GormSet,
		provider.QueryServiceSet,
		provider.PresenterSet,
		provider.UsecaseSet,
		provider.TagServiceSet,
		provider.GRPCServerSet,
		wire.NewSet(NewDependencies),
	)
	return nil
}
