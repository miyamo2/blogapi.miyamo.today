package types

import (
	"database/sql"
	"fmt"

	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/goccy/go-json"
)

// UTCTime is a type alias for synchro.Time with UTC timezone.
type UTCTime = synchro.Time[tz.UTC]

type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var _ sql.Scanner = (*Tags)(nil)

type Tags []Tag

func (t *Tags) Scan(src interface{}) error {
	data, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("unexpected type: %T", src)
	}
	return json.Unmarshal(data, t)
}
