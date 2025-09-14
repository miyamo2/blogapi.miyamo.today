package usecase

import (
	"iter"
	"slices"

	"blogapi.miyamo.today/tag-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/tag-service/internal/infra/rdb/types"
)

// articleDtoFromQueryModel converts query.Article to dto.Article
func articleDtoFromQueryModel(articles types.Articles) []dto.Article {
	return slices.Collect(
		func(yield func(dto.Article) bool) {
			for _, v := range articles {
				if !yield(
					dto.NewArticle(
						v.ID,
						v.Title,
						v.Thumbnail,
						v.CreatedAt,
						v.UpdatedAt,
					),
				) {
					return
				}
			}
		},
	)
}

func getPage[T any](rows []T, size, number int) iter.Seq[T] {
	i := 1
	for page := range slices.Chunk(rows, size) {
		if i > number {
			break
		}
		if i < number {
			i++
			continue
		}
		return func(yield func(T) bool) {
			for _, v := range page {
				if !yield(v) {
					return
				}
			}
		}
	}
	return func(_ func(T) bool) {
		return
	}
}
