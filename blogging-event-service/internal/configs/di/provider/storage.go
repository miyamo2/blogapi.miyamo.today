package provider

import (
	"blogapi.miyamo.today/blogging-event-service/internal/app/usecase/storage"
	"blogapi.miyamo.today/blogging-event-service/internal/infra/s3"
	"github.com/google/wire"
)

var StorageSet = wire.NewSet(
	s3.NewUploader,
	wire.Bind(new(storage.Uploader), new(*s3.Uploader)),
)
