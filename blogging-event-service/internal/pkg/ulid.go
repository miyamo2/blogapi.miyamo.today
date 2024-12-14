package pkg

import "github.com/oklog/ulid/v2"

type ULIDGenerator func() ulid.ULID
