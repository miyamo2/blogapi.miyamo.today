package usecase

import (
	"blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/tag-service/internal/infra/rdb/query/model"
)

// tagDtoFromQueryModel converts query.Tag to dto.Tag
func tagDtoFromQueryModel(tag model.Tag) dto.Tag {
	return dto.NewTag(
		tag.ID(),
		tag.Name(),
		func() []dto.Article {
			articles := make([]dto.Article, 0, len(tag.Articles()))
			for _, article := range tag.Articles() {
				articles = append(articles, dto.NewArticle(
					article.ID(),
					article.Title(),
					article.Thumbnail(),
					article.CreatedAt(),
					article.UpdatedAt(),
				))
			}
			return articles
		}())
}
