package pb

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/miyamo2/altnrslog"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/app/usecase/dto"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/infra/grpc"
	"github.com/miyamo2/blogapi.miyamo.today/core/log"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
	"log/slog"
)

type Converter struct{}

func (c Converter) ToCreateArticleArticleResponse(ctx context.Context, from *dto.CreateArticleOutDto) (response *grpc.BloggingEventResponse, err error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToCreateArticleArticleResponse").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
		err = nil
	}
	logger.InfoContext(ctx, "BEGIN", slog.Group("patameters", slog.Any("from", *from)))
	defer func() {
		logger.InfoContext(ctx, "END", slog.Group("return", slog.Any("response", *response)))
	}()
	response = &grpc.BloggingEventResponse{
		EventId:   from.EventID(),
		ArticleId: from.ArticleID(),
	}
	return
}

func (c Converter) ToUpdateArticleTitleResponse(ctx context.Context, from *dto.UpdateArticleTitleOutDto) (response *grpc.BloggingEventResponse, err error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToUpdateArticleTitleResponse").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
		err = nil
	}
	logger.InfoContext(ctx, "BEGIN", slog.Group("patameters", slog.Any("from", *from)))
	defer func() {
		logger.InfoContext(ctx, "END", slog.Group("return", slog.Any("response", *response)))
	}()
	response = &grpc.BloggingEventResponse{
		EventId:   from.EventID(),
		ArticleId: from.ArticleID(),
	}
	return
}

func (c Converter) ToUpdateArticleBodyResponse(ctx context.Context, from *dto.UpdateArticleBodyOutDto) (response *grpc.BloggingEventResponse, err error) {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToUpdateArticleBodyResponse").End()
	logger, err := altnrslog.FromContext(ctx)
	if err != nil {
		err = errors.WithStack(err)
		nrtx.NoticeError(nrpkgerrors.Wrap(err))
		logger = log.DefaultLogger()
		err = nil
	}
	logger.InfoContext(ctx, "BEGIN", slog.Group("parameters", slog.Any("from", *from)))
	defer func() {
		logger.InfoContext(ctx, "END", slog.Group("return", slog.Any("response", *response)))
	}()
	response = &grpc.BloggingEventResponse{
		EventId:   from.EventID(),
		ArticleId: from.ArticleID(),
	}
	return
}

func NewConverter() *Converter {
	return &Converter{}
}
