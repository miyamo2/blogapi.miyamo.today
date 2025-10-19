package converter

import (
	"context"
	"encoding/json"

	"blogapi.miyamo.today/read-model-updater/internal/app/usecase"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/cockroachdb/errors"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// Message represents a DynamoDB stream event message.
type Message struct {
	DynamoDB DynamoDB `json:"dynamodb"`
}

// DynamoDB represents the DynamoDB data in the stream event.
type DynamoDB struct {
	NewImage json.RawMessage `json:"NewImage"`
}

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

func (c *Converter) ToSyncUsecaseInDto(ctx context.Context, body []byte, eventAt synchro.Time[tz.UTC]) (
	*usecase.SyncUsecaseInDto, error,
) {
	nrtx := newrelic.FromContext(ctx)
	seg := nrtx.StartSegment("ToSyncUsecaseInDtoSeq")
	defer seg.End()

	var message Message
	err := json.Unmarshal(body, &message)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal body")
	}

	avm, err := attributevalue.UnmarshalMapJSON(message.DynamoDB.NewImage)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode new image json to attribute value map")
	}

	var dto usecase.SyncUsecaseInDto
	err = attributevalue.UnmarshalMap(avm, &dto)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal attribute value map to dto")
	}
	dto.EventAt = eventAt

	return &dto, nil
}
