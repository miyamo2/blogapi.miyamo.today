package converter

import (
	"context"
	"encoding/json"

	"blogapi.miyamo.today/read-model-updater/internal/app/usecase"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/cockroachdb/errors"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// Message represents a DynamoDB stream event message.
type Message struct {
	DynamoDB DynamoDB `json:"dynamodb"`
}

// DynamoDB represents the DynamoDB data in the stream event.
type DynamoDB struct {
	NewImage usecase.SyncUsecaseInDto `json:"NewImage"`
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
	dto := message.DynamoDB.NewImage
	dto.EventAt = eventAt

	return &dto, nil
}
