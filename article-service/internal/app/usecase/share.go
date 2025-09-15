package usecase

import (
	"iter"
	"slices"

	"blogapi.miyamo.today/article-service/internal/app/usecase/dto"
	"blogapi.miyamo.today/article-service/internal/infra/rdb/types"
)

// tagDtoFromQueryModel converts type.Tag to dto.Tag
func tagDtoFromQueryModel(tags []types.Tag) []dto.Tag {
	return slices.Collect(
		func(yield func(dto.Tag) bool) {
			for _, v := range tags {
				if !yield(dto.NewTag(v.ID, v.Name)) {
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
	return func(_ func(T) bool) {}
}
