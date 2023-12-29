package duration

import (
	"github.com/Songmu/flextime"
	"reflect"
	"testing"
	"time"
)

var (
	start = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).Round(0)
	stop  = time.Date(2020, 1, 1, 0, 0, 1, 0, time.UTC).Round(0)
)

type WatchDurationTestCase struct {
	target        Watch
	want          time.Duration
	beforeProcess func()
}

func Test_Watch_Duration(t *testing.T) {
	tests := map[string]WatchDurationTestCase{
		"happy_path/already_stopped": {
			target: Watch{start: start, stop: stop},
			want:   stop.Sub(start),
		},
		"happy_path/not_stopped_yet": {
			target: func() Watch {
				flextime.Fix(start)
				return Start()
			}(),
			want: stop.Sub(start),
			beforeProcess: func() {
				flextime.Fix(stop)
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			target := tt.target
			want := tt.want
			beforeProcess := tt.beforeProcess
			if beforeProcess != nil {
				beforeProcess()
			}
			if got := target.Duration(); got != want {
				t.Errorf("Duration() = %v, want %v", got, want)
			}
		})
	}
}

type WatchSDurationTestCase struct {
	target        Watch
	want          string
	beforeProcess func()
}

func Test_Watch_SDuration(t *testing.T) {
	tests := map[string]WatchSDurationTestCase{
		"happy_path": {
			target: func() Watch {
				flextime.Fix(start)
				return Start()
			}(),
			want: "00:00:01.000000000",
			beforeProcess: func() {
				flextime.Fix(stop)
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			target := tt.target
			want := tt.want
			beforeProcess := tt.beforeProcess
			if beforeProcess != nil {
				beforeProcess()
			}
			if got := target.SDuration(); got != want {
				t.Errorf("SDuration() = %v, want %v", got, want)
			}
		})
	}
}

type WatchStopTestCase struct {
	target        Watch
	expect        Watch
	beforeProcess func()
}

func Test_Watch_Stop(t *testing.T) {
	tests := map[string]WatchStopTestCase{
		"happy_path": {
			target: func() Watch {
				flextime.Fix(start)
				return Start()
			}(),
			expect: Watch{start: start, stop: stop},
			beforeProcess: func() {
				flextime.Fix(stop)
			},
		},
		"unhappy_path/already_stopped": {
			target: Watch{start: start, stop: stop},
			expect: Watch{start: start, stop: stop},
			beforeProcess: func() {
				flextime.Fix(time.Date(2020, 12, 31, 23, 59, 59, 999999999, time.UTC))
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			target := tt.target
			expect := tt.expect
			beforeProcess := tt.beforeProcess
			if beforeProcess != nil {
				beforeProcess()
			}
			target.Stop()
			if !reflect.DeepEqual(target, expect) {
				t.Errorf("actual = %v, expected %v", target, expect)
			}
		})
	}
}
