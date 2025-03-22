//go:build wireinject

package di

import (
	"blogapi.miyamo.today/blogging-event-service/internal/configs/di/provider"
	"github.com/google/wire"
)

func GetDependencies() *Dependencies {
	wire.Build(
		provider.AWSSet,
		provider.NewRelicSet,
		provider.StorageSet,
		provider.GormSet,
		provider.CommandSet,
		provider.PresenterSet,
		provider.UsecaseSet,
		provider.BloggingEventServiceServerSet,
		provider.EchoSet,
		wire.NewSet(NewDependencies),
	)
	return nil
}
