package converter

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/cockroachdb/errors"
	"github.com/goccy/go-json"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/app/usecase"
	"github.com/miyamo2/blogapi.miyamo.today/read-model-updater/internal/if-adapters/model"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
	"iter"
)

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

func (c *Converter) ToSyncUsecaseInDtoSeq(ctx context.Context, records []events.DynamoDBEventRecord) iter.Seq2[int, usecase.SyncUsecaseInDto] {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToSyncUsecaseInDtoSeq").End()
	return func(yield func(int, usecase.SyncUsecaseInDto) bool) {
		for i, record := range records {
			dto, err := toDto(record.Change.NewImage)
			if err != nil {
				nrtx.NoticeError(nrpkgerrors.Wrap(err))
				continue
			}
			if !yield(i, dto) {
				return
			}
		}
	}
}

func toDto(in map[string]events.DynamoDBAttributeValue) (usecase.SyncUsecaseInDto, error) {
	attributevalueMap := make(map[string]*dynamodb.AttributeValue)
	var image model.Image
	for k, v := range in {
		var av dynamodb.AttributeValue
		bytes, err := v.MarshalJSON()
		if err != nil {
			err = errors.WithStack(err)
			return usecase.SyncUsecaseInDto{}, err
		}

		if err := json.Unmarshal(bytes, &av); err != nil {
			err = errors.WithStack(err)
			return usecase.SyncUsecaseInDto{}, err
		}

		attributevalueMap[k] = &av
	}

	if err := dynamodbattribute.UnmarshalMap(attributevalueMap, &image); err != nil {
		err = errors.WithStack(err)
		return usecase.SyncUsecaseInDto{}, err
	}
	return usecase.NewSyncUsecaseInDto(
		image.EventID,
		image.ArticleID,
		image.Title,
		image.Content,
		image.Thumbnail,
		image.AttacheTags,
		image.DetachTags,
		image.Invisible,
	), nil
}
