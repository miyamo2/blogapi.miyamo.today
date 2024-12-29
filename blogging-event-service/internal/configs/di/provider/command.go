package provider

import (
	"blogapi.miyamo.today/blogging-event-service/internal/app/usecase/command"
	"blogapi.miyamo.today/blogging-event-service/internal/infra/dynamo"
	"github.com/google/wire"
)

func BloggingEventCommandService() *dynamo.BloggingEventCommandService {
	return dynamo.NewBloggingEventCommandService(nil)
}

var CommandSet = wire.NewSet(
	BloggingEventCommandService,
	wire.Bind(new(command.BloggingEventService), new(*dynamo.BloggingEventCommandService)),
)
