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

// Now returns the current time from a monotonic clock, it must be
// treated as an opaque, platform specific value.
//
// For most purposes, you should subtract two of these values
// by using the Duration() method.
func Now() Monotime {
	return monotime()
}

// Duration returns a time.Duration from two previously captured Now() calls.
func Duration(start Monotime, end Monotime) time.Duration {
	return duration(start, end)
}
