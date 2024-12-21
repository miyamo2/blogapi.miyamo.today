package usecase

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/app/usecase/command"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/domain/model"
	"github.com/miyamo2/blogapi.miyamo.today/core/db"
	"github.com/miyamo2/blogapi.miyamo.today/core/log"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
	"log/slog"
)

// AttachTags is a use-case for creating an article.
type AttachTags struct {
	bloggingEventCommand command.BloggingEventService
}

// Execute executes the AttachTags use-case.
func (u *AttachTags) Execute(ctx context.Context, in *dto.AttachTagsInDto) (_ *dto.AttachTagsOutDto, err error) {
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
					slog.Any("dto.AttachTagsOutDto", nil),
					slog.Any("error", err)))
			return
		}
		logger.InfoContext(ctx, "END")
	}()

	command := model.NewAttachTagsEvent(in.ID(), in.TagNames())
	commandOut := db.NewSingleStatementResult[*model.BloggingEventKey]()
	err = u.bloggingEventCommand.AttachTags(ctx, command, commandOut).Execute(ctx)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}

	key := commandOut.StrictGet()
	result := dto.NewAttachTagsOutDto(key.EventID(), key.ArticleID())
	return &result, nil
}

// NewAttachTags is a constructor for AttachTags use-case.
func NewAttachTags(bloggingEventCommand command.BloggingEventService) *AttachTags {
	return &AttachTags{bloggingEventCommand: bloggingEventCommand}
}
