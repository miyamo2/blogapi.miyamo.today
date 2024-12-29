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

// UpdateArticleBody is a use-case for updating an article body.
type UpdateArticleBody struct {
	bloggingEventCommand command.BloggingEventService
}

// Execute executes the UpdateArticleBody use-case.
func (u *UpdateArticleBody) Execute(ctx context.Context, in *dto.UpdateArticleBodyInDto) (*dto.UpdateArticleBodyOutDto, error) {
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
					slog.Any("dto.UpdateArticleBody", nil),
					slog.Any("error", err)))
			return
		}
		logger.InfoContext(ctx, "END")
	}()

	command := model.NewUpdateArticleBodyEvent(in.ID(), in.Body())
	commandOut := db.NewSingleStatementResult[*model.BloggingEventKey]()
	err = u.bloggingEventCommand.UpdateArticleBody(ctx, command, commandOut).Execute(ctx)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}

	key := commandOut.StrictGet()
	result := dto.NewUpdateArticleBodyOutDto(key.EventID(), key.ArticleID())
	return &result, nil
}

// NewUpdateArticleBody is a constructor for UpdateArticleBody use-case.
func NewUpdateArticleBody(bloggingEventCommand command.BloggingEventService) *UpdateArticleBody {
	return &UpdateArticleBody{bloggingEventCommand: bloggingEventCommand}
}
