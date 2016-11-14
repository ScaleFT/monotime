package monotime

import (
	"github.com/winlabs/gowin32"
	"time"
)

type timer struct {
	start uint64
}

func newTimer() *timer {
	return &timer{
		start: gowin32.GetTimeCounter(),
	}
}

func (t *timer) Elapsed() time.Duration {
	now := gowin32.GetTimeCounter()
	elapsed := now - t.start
	return time.Duration(elapsed)
}
