package provider

import (
	"github.com/miyamo2/blogapi/internal/app/usecase/dto"
	abstract "github.com/miyamo2/blogapi/internal/if-adapter/controller/graphql/resolver/presenter/converter"

	"github.com/miyamo2/blogapi/internal/if-adapter/presenters/graphql/converter"
	"go.uber.org/fx"
)

// compatibility check
var (
	_ abstract.ArticleConverter[dto.Tag, dto.ArticleTag, dto.ArticleOutDto]   = (*converter.Converter)(nil)
	_ abstract.ArticlesConverter[dto.Tag, dto.ArticleTag, dto.ArticlesOutDto] = (*converter.Converter)(nil)
	_ abstract.TagConverter[dto.Article, dto.TagArticle, dto.TagOutDto]       = (*converter.Converter)(nil)
	_ abstract.TagsConverter[dto.Article, dto.TagArticle, dto.TagsOutDto]     = (*converter.Converter)(nil)
)

var Presenter = fx.Options(
	fx.Provide(
		fx.Annotate(
			converter.NewConverter, fx.As(new(abstract.ArticleConverter[dto.Tag, dto.ArticleTag, dto.ArticleOutDto])),
		),
	),
	fx.Provide(
		fx.Annotate(
			converter.NewConverter, fx.As(new(abstract.ArticlesConverter[dto.Tag, dto.ArticleTag, dto.ArticlesOutDto])),
		),
	),
	fx.Provide(
		fx.Annotate(
			converter.NewConverter, fx.As(new(abstract.TagConverter[dto.Article, dto.TagArticle, dto.TagOutDto])),
		),
	),
	fx.Provide(
		fx.Annotate(
			converter.NewConverter, fx.As(new(abstract.TagsConverter[dto.Article, dto.TagArticle, dto.TagsOutDto])),
		),
	),
)
