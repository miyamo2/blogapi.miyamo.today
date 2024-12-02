package provider

import (
	"github.com/miyamo2/blogapi.miyamo.today/article-service/internal/app/usecase/query"
	impl "github.com/miyamo2/blogapi.miyamo.today/article-service/internal/infra/rdb/query"
	"go.uber.org/fx"
)

// compatibility check
var _ query.ArticleService = (*impl.ArticleService)(nil)

var QueryService = fx.Options(
	fx.Provide(
		fx.Annotate(
			impl.NewArticleService,
			fx.As(new(query.ArticleService)),
		),
	),
)
