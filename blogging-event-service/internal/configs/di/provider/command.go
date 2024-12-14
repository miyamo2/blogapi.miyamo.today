package provider

import (
	"github.com/google/wire"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/app/usecase/command"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/infra/dynamo"
)

func BloggingEventCommandService() *dynamo.BloggingEventCommandService {
	return dynamo.NewBloggingEventCommandService(nil)
}

var CommandSet = wire.NewSet(
	BloggingEventCommandService,
	wire.Bind(new(command.BloggingEventService), new(*dynamo.BloggingEventCommandService)),
)
