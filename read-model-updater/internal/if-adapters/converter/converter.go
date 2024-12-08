package converter

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/app/usecase"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/if-adapters/model"
	"github.com/newrelic/go-agent/v3/newrelic"
	"iter"
)

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

func (c *Converter) ToSyncUsecaseInDtoSeq(ctx context.Context, records []model.Record) iter.Seq2[int, usecase.SyncUsecaseInDto] {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToSyncUsecaseInDtoSeq").End()
	return func(yield func(int, usecase.SyncUsecaseInDto) bool) {
		for i, record := range records {
			var image model.Image
			if err := attributevalue.UnmarshalMap(record.DynamoDB.NewImage, &image); err != nil {
				continue
			}
			if !yield(i, usecase.NewSyncUsecaseInDto(
				image.EventID,
				image.ArticleID,
				image.Title,
				image.Content,
				image.Thumbnail,
				image.AttacheTags,
				image.DetachTags,
				image.Invisible,
			)) {
				return
			}
		}
	}
}
