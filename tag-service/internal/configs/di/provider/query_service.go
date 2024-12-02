package provider

import (
	"github.com/google/wire"
	"github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/app/usecase/query"
	impl "github.com/miyamo2/blogapi.miyamo.today/tag-service/internal/infra/rdb/query"
)

// compatibility check
var _ query.TagService = (*impl.TagService)(nil)

var QueryServiceSet = wire.NewSet(
	impl.NewTagService,
	wire.Bind(new(query.TagService), new(*impl.TagService)),
)
