package usecase

import (
	"blogapi.miyamo.today/blogging-event-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/blogging-event-service/internal/app/usecase/storage"
	"blogapi.miyamo.today/core/log"
	"context"
	"github.com/miyamo2/altnrslog"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
	"log/slog"
)

type UploadImage struct {
	uploader storage.Uploader
}

func (u *UploadImage) Execute(ctx context.Context, in *dto.UploadImageInDto) (*dto.UploadImageOutDto, error) {
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
					slog.Any("dto.UploadImageOutDto", nil),
					slog.Any("error", err)))
			return
		}
		logger.InfoContext(ctx, "END")
	}()

	uri, err := u.uploader.Upload(ctx, in.Name(), in.Bytes(), in.ContentType())
	if err != nil {
		return nil, err
	}
	result := dto.NewUploadImageOutDto(*uri)
	return &result, nil
}

func NewUploadImage(uploader storage.Uploader) *UploadImage {
	return &UploadImage{uploader: uploader}
}
