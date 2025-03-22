//go:build wireinject

package di

import (
	"blogapi.miyamo.today/federator/internal/configs/di/provider"
	"github.com/google/wire"
)

func GetDependencies() *Dependencies {
	wire.Build(
		provider.NewRelicSet,
		provider.HTTPSet,
		provider.GRPCClientSet,
		provider.PresenterSet,
		provider.UsecaseSet,
		provider.GqlgenSet,
		provider.EchoSet,
		wire.NewSet(NewDependencies),
	)
	return nil
}
