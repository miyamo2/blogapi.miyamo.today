package provider

import (
	"github.com/miyamo2/blogapi-tag-service/internal/app/usecase/query"
	impl "github.com/miyamo2/blogapi-tag-service/internal/infra/rdb/query"
	"github.com/miyamo2/blogapi-tag-service/internal/infra/rdb/query/model"
	"go.uber.org/fx"
)

// compatibility check
var _ query.TagService[model.Article, *model.Tag] = (*impl.TagService)(nil)

var QueryService = fx.Options(
	fx.Provide(
		fx.Annotate(
			impl.NewTagService,
			fx.As(new(query.TagService[model.Article, *model.Tag])),
		),
	),
)
