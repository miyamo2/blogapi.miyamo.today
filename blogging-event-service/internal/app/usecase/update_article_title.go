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

// UpdateArticleTitle is a use-case for creating an article.
type UpdateArticleTitle struct {
	bloggingEventCommand command.BloggingEventService
}

// Execute executes the UpdateArticleTitle use-case.
func (u *UpdateArticleTitle) Execute(ctx context.Context, in *dto.UpdateArticleTitleInDto) (_ *dto.UpdateArticleTitleOutDto, err error) {
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
					slog.Any("dto.UpdateArticleTitleOutDto", nil),
					slog.Any("error", err)))
			return
		}
		logger.InfoContext(ctx, "END")
	}()

	command := model.NewUpdateArticleTitleEvent(in.ID(), in.Title())
	commandOut := db.NewSingleStatementResult[*model.BloggingEventKey]()
	err = u.bloggingEventCommand.UpdateArticleTitle(ctx, command, commandOut).Execute(ctx)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}

	key := commandOut.StrictGet()
	result := dto.NewUpdateArticleTitleOutDto(key.EventID(), key.ArticleID())
	return &result, nil
}

// NewUpdateArticleTitle is a constructor for UpdateArticleTitle use-case.
func NewUpdateArticleTitle(bloggingEventCommand command.BloggingEventService) *UpdateArticleTitle {
	return &UpdateArticleTitle{bloggingEventCommand: bloggingEventCommand}
}
