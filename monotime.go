package monotime

import (
	"time"
)

type Timer interface {
	// Returns the elaspsed time since this Timer was created.
	Elapsed() time.Duration
}

// New creates a portable monolithic timer.  If the underlying system
// changes times, this interface will still return an increasing duration.
func New() Timer {
	return newTimer()
}
