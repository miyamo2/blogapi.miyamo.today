package provider

import (
	"github.com/miyamo2/blogapi-article-service/internal/app/usecase/query"
	impl "github.com/miyamo2/blogapi-article-service/internal/infra/rdb/query"
	"go.uber.org/fx"
)

// compatibility check
var _ query.ArticleService[impl.Tag, *impl.Article] = (*impl.ArticleService)(nil)

var QueryService = fx.Options(
	fx.Provide(
		fx.Annotate(
			impl.NewArticleService,
			fx.As(new(query.ArticleService[impl.Tag, *impl.Article])),
		),
	),
)
