package usecase

import (
	"blogapi.miyamo.today/blogging-event-service/internal/app/usecase/command"
	"blogapi.miyamo.today/blogging-event-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/blogging-event-service/internal/domain/model"
	"blogapi.miyamo.today/core/db"
	"blogapi.miyamo.today/core/log"
	"context"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/altnrslog"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
	"log/slog"
)

// DetachTags is a use-case for creating an article.
type DetachTags struct {
	bloggingEventCommand command.BloggingEventService
}

// Execute executes the DetachTags use-case.
func (u *DetachTags) Execute(ctx context.Context, in *dto.DetachTagsInDto) (_ *dto.DetachTagsOutDto, err error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Execute").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN")
	defer func() {
		if err != nil {
			logger.WarnContext(ctx, "END",
				slog.Group("return",
					slog.Any("dto.DetachTagsOutDto", nil),
					slog.Any("error", err)))
			return
		}
		logger.InfoContext(ctx, "END")
	}()

	command := model.NewDetachTagsEvent(in.ID(), in.TagNames())
	commandOut := db.NewSingleStatementResult[*model.BloggingEventKey]()
	err = u.bloggingEventCommand.DetachTags(ctx, command, commandOut).Execute(ctx)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}

	key := commandOut.StrictGet()
	result := dto.NewDetachTagsOutDto(key.EventID(), key.ArticleID())
	return &result, nil
}

// NewDetachTags is a constructor for DetachTags use-case.
func NewDetachTags(bloggingEventCommand command.BloggingEventService) *DetachTags {
	return &DetachTags{bloggingEventCommand: bloggingEventCommand}
}
