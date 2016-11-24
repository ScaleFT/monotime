package monotime

import (
	"time"
)

type Timer interface {
	// Returns the elaspsed time since this Timer was created.
	Elapsed() time.Duration
}

// New creates a portable monotonic timer.  If the underlying system
// changes times, this interface will still return an increasing duration.
func New() Timer {
	return newTimer()
}

// Now returns the current time from a monotonic clock, it should be
// treated as an opaque, platform specific value.
//
// For most purposes, you should subtract two of these values
// and using the Duration() method.
func Now() uint64 {
	return monotime()
}

// Duration returns
func Duration(start uint64, end uint64) time.Duration {
	return duration(start, end)
}

type timer struct {
	start uint64
}

func newTimer() *timer {
	return &timer{
		start: monotime(),
	}
}

func (t *timer) Elapsed() time.Duration {
	return duration(t.start, monotime())
}
