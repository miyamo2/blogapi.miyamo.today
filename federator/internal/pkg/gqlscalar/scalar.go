package gqlscalar

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"io"
	"net/url"
	"time"
)

type URL url.URL

var (
	_ graphql.Marshaler   = (*URL)(nil)
	_ graphql.Unmarshaler = (*URL)(nil)
)

// String implements the fmt.Stringer interface
func (u URL) String() string {
	v := url.URL(u)
	return v.String()
}

// MarshalGQL implements the graphql.Marshaler interface
func (u URL) MarshalGQL(w io.Writer) {
	io.WriteString(w, fmt.Sprintf(`"%s"`, u.String()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (u *URL) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case string:
		result, err := url.Parse(v)
		if err != nil {
			return err
		}
		*u = URL(*result)
		return nil
	case []byte:
		result := &url.URL{}
		if err := result.UnmarshalBinary(v); err != nil {
			return err
		}
		*u = URL(*result)
		return nil
	default:
		return fmt.Errorf("%T is incompatible with URL", v)
	}
}

var (
	_ graphql.Marshaler   = (*UTC)(nil)
	_ graphql.Unmarshaler = (*UTC)(nil)
)

type UTC synchro.Time[tz.UTC]

func (u UTC) ToSynchroTime() synchro.Time[tz.UTC] {
	return synchro.Time[tz.UTC](u)
}

func (u UTC) MarshalGQL(w io.Writer) {
	io.WriteString(w, fmt.Sprintf(`"%s"`, synchro.Time[tz.UTC](u).Format(time.RFC3339)))
}

func (u *UTC) UnmarshalGQL(v any) error {
	switch v := v.(type) {
	case string:
		t, err := synchro.Parse[tz.UTC](time.RFC3339Nano, v)
		if err != nil {
			return err
		}
		*u = UTC(t)
		return nil
	case []byte:
		t, err := synchro.Parse[tz.UTC](time.RFC3339Nano, string(v))
		if err != nil {
			return err
		}
		*u = UTC(t)
		return nil
	default:
		return fmt.Errorf("%T is incompatible with UTC", v)
	}
}
