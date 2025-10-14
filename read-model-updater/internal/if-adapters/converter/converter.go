package converter

import (
	"context"
	"iter"
	"log/slog"
	"time"

	"blogapi.miyamo.today/read-model-updater/internal/app/usecase"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams/types"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

func (c *Converter) ToSyncUsecaseInDtoSeq(
	ctx context.Context, records []types.Record,
) iter.Seq2[int, usecase.SyncUsecaseInDto] {
	nrtx := newrelic.FromContext(ctx)
	defer nrtx.StartSegment("ToSyncUsecaseInDtoSeq").End()
	return func(yield func(int, usecase.SyncUsecaseInDto) bool) {
		for i, record := range records {
			dto, err := toDto(record.Dynamodb.NewImage, record.Dynamodb.ApproximateCreationDateTime)
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

func toDto(in map[string]types.AttributeValue, eventAt *time.Time) (usecase.SyncUsecaseInDto, error) {
	var v usecase.SyncUsecaseInDto
	avm, err := attributevalue.FromDynamoDBStreamsMap(in)
	if err != nil {
		return v, err
	}
	slog.Default().Info("attribute value map", slog.Any("attribute value map", avm))
	err = attributevalue.UnmarshalMap(avm, &v)
	if err != nil {
		return v, err
	}
	if eventAt != nil {
		v.EventAt = synchro.In[tz.UTC](*eventAt)
	}
	slog.Default().Info("dto", slog.Any("dto", v))
	return v, nil
}
