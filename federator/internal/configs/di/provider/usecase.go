package provider

import (
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase/dto"
	abstract "github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/controller/graphql/resolver/usecase"
	"go.uber.org/fx"
)

// compatibility check
var (
	_ abstract.Article[dto.ArticleInDto, dto.Tag, dto.ArticleTag, dto.ArticleOutDto]    = (*usecase.Article)(nil)
	_ abstract.Articles[dto.ArticlesInDto, dto.Tag, dto.ArticleTag, dto.ArticlesOutDto] = (*usecase.Articles)(nil)
	_ abstract.Tag[dto.TagInDto, dto.Article, dto.TagArticle, dto.TagOutDto]            = (*usecase.Tag)(nil)
	_ abstract.Tags[dto.TagsInDto, dto.Article, dto.TagArticle, dto.TagsOutDto]         = (*usecase.Tags)(nil)
)

var Usecase = fx.Options(
	fx.Provide(
		fx.Annotate(usecase.NewArticle, fx.As(new(abstract.Article[dto.ArticleInDto, dto.Tag, dto.ArticleTag, dto.ArticleOutDto]))),
	),
	fx.Provide(
		fx.Annotate(usecase.NewArticles, fx.As(new(abstract.Articles[dto.ArticlesInDto, dto.Tag, dto.ArticleTag, dto.ArticlesOutDto]))),
	),
	fx.Provide(
		fx.Annotate(usecase.NewTag, fx.As(new(abstract.Tag[dto.TagInDto, dto.Article, dto.TagArticle, dto.TagOutDto]))),
	),
	fx.Provide(
		fx.Annotate(usecase.NewTags, fx.As(new(abstract.Tags[dto.TagsInDto, dto.Article, dto.TagArticle, dto.TagsOutDto]))),
	),
)
