package types

import (
	"github.com/Code-Hex/synchro"
	"github.com/Code-Hex/synchro/tz"
)

// UTCTime is a type alias for synchro.Time with UTC timezone.
type UTCTime = synchro.Time[tz.UTC]
