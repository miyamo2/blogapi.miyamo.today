//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/configs/di/provider"
)

func GetDependencies() *Dependencies {
	wire.Build(
		provider.AWSSet,
		provider.NewRelicSet,
		provider.GormSet,
		provider.CommandSet,
		provider.PresenterSet,
		provider.UsecaseSet,
		provider.BloggingEventServiceServerSet,
		provider.GRPCServerSet,
		wire.NewSet(NewDependencies),
	)
	return nil
}
