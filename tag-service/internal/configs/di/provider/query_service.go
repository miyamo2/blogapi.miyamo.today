package provider

import (
	"blogapi.miyamo.today/tag-service/internal/app/usecase/query"
	impl "blogapi.miyamo.today/tag-service/internal/infra/rdb/query"
	"github.com/google/wire"
)

// compatibility check
var _ query.TagService = (*impl.TagService)(nil)

var QueryServiceSet = wire.NewSet(
	impl.NewTagService,
	wire.Bind(new(query.TagService), new(*impl.TagService)),
)
