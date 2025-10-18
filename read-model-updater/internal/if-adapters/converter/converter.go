package converter

import (
	"context"
	"fmt"
	"iter"
	"time"

	"blogapi.miyamo.today/read-model-updater/internal/app/usecase"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams/types"
	"github.com/cockroachdb/errors"
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
	seg := nrtx.StartSegment("ToSyncUsecaseInDtoSeq")
	defer seg.End()
	seg.AddAttribute("records", fmt.Sprintf("%+v", records))
	return func(yield func(int, usecase.SyncUsecaseInDto) bool) {
		for i, record := range records {
			if record.EventName != types.OperationTypeInsert {
				continue
			}
			dto, err := toDto(ctx, record.Dynamodb.NewImage, record.Dynamodb.ApproximateCreationDateTime)
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

func toDto(ctx context.Context, in map[string]types.AttributeValue, eventAt *time.Time) (
	usecase.SyncUsecaseInDto, error,
) {
	nrtx := newrelic.FromContext(ctx)
	seg := nrtx.StartSegment("toDto")
	defer seg.End()

	var v usecase.SyncUsecaseInDto
	avm, err := attributevalue.FromDynamoDBStreamsMap(in)
	if err != nil {
		return v, errors.WithStack(err)
	}
	err = attributevalue.UnmarshalMap(avm, &v)
	if err != nil {
		return v, errors.WithStack(err)
	}
	if eventAt != nil {
		v.EventAt = synchro.In[tz.UTC](*eventAt)
	}
	return v, nil
}
