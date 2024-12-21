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

// UpdateArticleThumbnail is a use-case for updating an article thumbnail.
type UpdateArticleThumbnail struct {
	bloggingEventCommand command.BloggingEventService
}

// Execute executes the UpdateArticleThumbnail use-case.
func (u *UpdateArticleThumbnail) Execute(ctx context.Context, in *dto.UpdateArticleThumbnailInDto) (*dto.UpdateArticleThumbnailOutDto, error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("Execute").End()

	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
	}
	logger.InfoContext(ctx, "BEGIN")
	defer func() {
		if err != nil {
			logger.WarnContext(ctx, "END",
				slog.Group("return",
					slog.Any("dto.UpdateArticleThumbnail", nil),
					slog.Any("error", err)))
		}
		logger.InfoContext(ctx, "END")
	}()

	command := model.NewUpdateArticleThumbnailEvent(in.ID(), in.ThumbnailUrl())
	commandOut := db.NewSingleStatementResult[*model.BloggingEventKey]()
	err = u.bloggingEventCommand.UpdateArticleThumbnail(ctx, command, commandOut).Execute(ctx)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}

	key := commandOut.StrictGet()
	result := dto.NewUpdateArticleThumbnailOutDto(key.EventID(), key.ArticleID())
	return &result, nil
}

// NewUpdateArticleThumbnail is a constructor for UpdateArticleThumbnail use-case.
func NewUpdateArticleThumbnail(bloggingEventCommand command.BloggingEventService) *UpdateArticleThumbnail {
	return &UpdateArticleThumbnail{bloggingEventCommand: bloggingEventCommand}
}
