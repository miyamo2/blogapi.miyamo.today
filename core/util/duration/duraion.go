package duration

import (
	"time"

	"github.com/Songmu/flextime"
)

// Watch is a measurer that elapsed time.
type Watch struct {
	start time.Time
	stop  time.Time
}

// Start returns a new duration
func Start() Watch {
	return Watch{start: flextime.Now().UTC().Round(0)}
}

// Stop stops the duration
func (sw *Watch) Stop() {
	if !sw.stop.IsZero() {
		return
	}
	sw.stop = flextime.Now().UTC().Round(0)
}

// Duration returns the elapsed time that the duration start/stop.
func (sw *Watch) Duration() time.Duration {
	sw.Stop()
	stop := sw.stop
	return stop.Sub(sw.start)
}

// SDuration returns the elapsed time as "15:04:05.000000000" format.
//
// See: https://pkg.go.dev/github.com/miyamo2/blogapi.miyamo.today/core/pkg/util/stopwatch/#Duration
func (sw *Watch) SDuration() string {
	return time.Unix(0, 0).UTC().Add(sw.Duration()).Format("15:04:05.000000000")
}
