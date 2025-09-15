package types

import (
	"fmt"

	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"github.com/goccy/go-json"
)

// UTCTime is a type alias for synchro.Time with UTC timezone.
type UTCTime = synchro.Time[tz.UTC]

type Article struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	Thumbnail string  `json:"thumbnail"`
	CreatedAt UTCTime `json:"created_at"`
	UpdatedAt UTCTime `json:"updated_at"`
}

type Articles []Article

func (a *Articles) Scan(src any) error {
	data, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("unexpected type: %T", src)
	}
	if len(data) == 0 {
		return nil
	}
	return json.Unmarshal(data, a)
}
