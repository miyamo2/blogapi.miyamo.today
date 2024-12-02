package model

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
	"io"
	"time"
	"unsafe"
)

func MarshalTime(t synchro.Time[tz.UTC]) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, fmt.Sprintf(`"%s"`, t.Format(time.RFC3339)))
	})
}

func UnmarshalTime(v interface{}) (synchro.Time[tz.UTC], error) {
	switch v := v.(type) {
	case string:
		t, err := synchro.Parse[tz.UTC](time.RFC3339, v)
		if err != nil {
			return synchro.Time[tz.UTC]{}, err
		}
		return t, nil
	case []byte:
		t, err := synchro.Parse[tz.UTC](time.RFC3339, *(*string)(unsafe.Pointer(&v)))
		if err != nil {
			return synchro.Time[tz.UTC]{}, err
		}
		return t, nil
	default:
		return synchro.Time[tz.UTC]{}, fmt.Errorf("%T is incompatible with synchro.Time[tz.UTC]", v)
	}
}
