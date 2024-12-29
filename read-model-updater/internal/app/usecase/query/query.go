package query

import (
	"blogapi.miyamo.today/core/db"
	"blogapi.miyamo.today/read-model-updater/internal/domain/model"
	"context"
)

// BloggingEventService is a query service interface for the BloggingEvent.
type BloggingEventService interface {
	// AllEventsWithArticleID returns all blogging events by article id.
	AllEventsWithArticleID(ctx context.Context, articleID string, out *db.MultipleStatementResult[model.BloggingEvent]) db.Statement
}
