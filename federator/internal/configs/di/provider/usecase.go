package provider

import (
	"github.com/google/wire"
	"github.com/miyamo2/blogapi.miyamo.today/federator/internal/app/usecase"
	abstract "github.com/miyamo2/blogapi.miyamo.today/federator/internal/if-adapter/controller/graphql/resolver/usecase"
)

// compatibility check
var (
	_ abstract.Article            = (*usecase.Article)(nil)
	_ abstract.Articles           = (*usecase.Articles)(nil)
	_ abstract.Tag                = (*usecase.Tag)(nil)
	_ abstract.Tags               = (*usecase.Tags)(nil)
	_ abstract.CreateArticle      = (*usecase.CreateArticle)(nil)
	_ abstract.UpdateArticleTitle = (*usecase.UpdateArticleTitle)(nil)
	_ abstract.UpdateArticleBody  = (*usecase.UpdateArticleBody)(nil)
)

var UsecaseSet = wire.NewSet(
	usecase.NewArticle,
	wire.Bind(new(abstract.Article), new(*usecase.Article)),
	usecase.NewArticles,
	wire.Bind(new(abstract.Articles), new(*usecase.Articles)),
	usecase.NewTag,
	wire.Bind(new(abstract.Tag), new(*usecase.Tag)),
	usecase.NewTags,
	wire.Bind(new(abstract.Tags), new(*usecase.Tags)),
	usecase.NewCreateArticle,
	wire.Bind(new(abstract.CreateArticle), new(*usecase.CreateArticle)),
	usecase.NewUpdateArticleTitle,
	wire.Bind(new(abstract.UpdateArticleTitle), new(*usecase.UpdateArticleTitle)),
	usecase.NewUpdateArticleBody,
	wire.Bind(new(abstract.UpdateArticleBody), new(*usecase.UpdateArticleBody)),
)
