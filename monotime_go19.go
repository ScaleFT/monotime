// +build go1.9

package monotime

import (
	"time"
)

type Monotime = time.Time

type timer struct {
	start Monotime
}

func newTimer() *timer {
	return &timer{
		start: time.Now(),
	}
}

func (t *timer) Elapsed() time.Duration {
	return time.Now().Sub(t.start)
}

func monotime() Monotime {
	return time.Now()
}

func duration(start Monotime, end Monotime) time.Duration {
	return end.Sub(start)
}