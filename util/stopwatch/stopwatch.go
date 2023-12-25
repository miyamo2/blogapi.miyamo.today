package stopwatch

import (
	"github.com/Songmu/flextime"
	"time"
)

// StopWatch is a measurer that elapsed time.
type StopWatch struct {
	start time.Time
	stop  time.Time
}

// Start returns a new stopwatch
func Start() StopWatch {
	return StopWatch{start: flextime.Now().UTC().Round(0)}
}

// Stop stops the stopwatch
func (sw *StopWatch) Stop() {
	if !sw.stop.IsZero() {
		return
	}
	sw.stop = flextime.Now().UTC().Round(0)
}

// Duration returns the elapsed time that the stopwatch start/stop.
func (sw *StopWatch) Duration() time.Duration {
	sw.Stop()
	stop := sw.stop
	return stop.Sub(sw.start)
}

// SDuration returns the elapsed time as "15:04:05.000000000" format.
//
// See: https://pkg.go.dev/github.com/miyamo2/blogapi-core/pkg/util/stopwatch/#Duration
func (sw *StopWatch) SDuration() string {
	return time.Unix(0, 0).UTC().Add(sw.Duration()).Format("15:04:05.000000000")
}
