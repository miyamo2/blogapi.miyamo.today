package provider

import (
	"github.com/google/wire"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/app/usecase/storage"
	"github.com/miyamo2/blogapi.miyamo.today/blogging-event-service/internal/infra/s3"
)

var StorageSet = wire.NewSet(
	s3.NewUploader,
	wire.Bind(new(storage.Uploader), new(*s3.Uploader)),
)
