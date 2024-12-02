//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/configs/di/provider"
)

func GetDependencies() *Dependencies {
	wire.Build(
		provider.NewRelicSet,
		provider.GRPCClientSet,
		provider.PresenterSet,
		provider.UsecaseSet,
		provider.GqlgenSet,
		provider.EchoSet,
		wire.NewSet(NewDependencies),
	)
	return nil
}
