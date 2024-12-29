//go:build wireinject

package di

import (
	"blogapi.miyamo.today/tag-service/internal/configs/di/provider"
	"github.com/google/wire"
)

func GetDependencies() *Dependencies {
	wire.Build(
		provider.NewRelicSet,
		provider.GormSet,
		provider.QueryServiceSet,
		provider.PresenterSet,
		provider.UsecaseSet,
		provider.TagServiceSet,
		provider.EchoSet,
		wire.NewSet(NewDependencies),
	)
	return nil
}
