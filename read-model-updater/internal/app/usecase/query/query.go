package query

import (
	"context"

	"blogapi.miyamo.today/read-model-updater/internal/domain/model"
)

// BloggingEventService is a query service interface for the BloggingEvent.
type BloggingEventService interface {
	// ListEventsByArticleID returns all blogging events by article id.
	ListEventsByArticleID(ctx context.Context, articleID string) []model.BloggingEvent
}
