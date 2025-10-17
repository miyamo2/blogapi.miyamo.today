package dynamo

import (
	"context"
	"fmt"
	"os"
	"slices"

	"blogapi.miyamo.today/read-model-updater/internal/domain/model"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/sqldav"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/oklog/ulid/v2"
)

type bloggingEvent struct {
	EventID    string             `db:"event_id"`
	ArticleID  string             `db:"article_id"`
	Title      *string            `db:"title"`
	Content    *string            `db:"content"`
	Thumbnail  *string            `db:"thumbnail"`
	Tags       sqldav.Set[string] `db:"tags"`
	AttachTags sqldav.Set[string] `db:"attach_tags"`
	DetachTags sqldav.Set[string] `db:"detach_tags"`
	Invisible  *bool              `db:"invisible"`
}

var listEventsByArticleID = fmt.Sprintf(
	`SELECT "event_id",
		"article_id", 
		"title", 
		"content", 
		"thumbnail", 
		"tags", 
		"attach_tags", 
		"detach_tags", 
		"invisible" 
FROM %s."article_id_event_id-Index" 
WHERE "article_id" = $1
`, os.Getenv("BLOGGING_EVENTS_TABLE_NAME"),
)

type DB interface {
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type BloggingEventQueryService struct {
	db DB
}

func (s *BloggingEventQueryService) ListEventsByArticleID(
	ctx context.Context,
	articleID string,
) ([]model.BloggingEvent, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("BloggingEventQueryService#AllEventsWithArticleID").End()

	rows := make([]bloggingEvent, 0)
	err := s.db.SelectContext(ctx, &rows, listEventsByArticleID, articleID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	slices.SortFunc(
		rows, func(i, j bloggingEvent) int {
			var iid ulid.ULID
			iid, err = ulid.Parse(i.EventID)
			if err != nil {
				return 0
			}
			var jid ulid.ULID
			jid, err = ulid.Parse(j.EventID)
			if err != nil {
				return 0
			}
			return iid.Compare(jid)
		},
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	result := make([]model.BloggingEvent, 0, len(rows))
	for _, r := range rows {
		result = append(
			result,
			model.NewBloggingEvent(
				r.EventID,
				r.ArticleID,
				r.Title,
				r.Content,
				r.Thumbnail,
				r.Tags,
				r.AttachTags,
				r.DetachTags,
				r.Invisible,
			),
		)
	}
	return result, nil
}

func NewBloggingEventQueryService(db DB) *BloggingEventQueryService {
	return &BloggingEventQueryService{
		db: db,
	}
}
