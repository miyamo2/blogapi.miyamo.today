package handler

import (
	"context"

	"blogapi.miyamo.today/read-model-updater/internal/app/usecase"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
)

// ToSyncUsecaseInDtoConverter is an interface for the converter of the usecase.SyncUsecaseInDto
type ToSyncUsecaseInDtoConverter interface {
	ToSyncUsecaseInDto(ctx context.Context, body []byte, eventAt synchro.Time[tz.UTC]) (*usecase.SyncUsecaseInDto, error)
}
