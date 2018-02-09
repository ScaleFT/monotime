// +build !go1.9

package monotime

import "time"

type Monotime uint64

type timer struct {
	start Monotime
}

func newTimer() *timer {
	return &timer{
		start: monotime(),
	}
}

func (t *timer) Elapsed() time.Duration {
	return duration(t.start, monotime())
}
