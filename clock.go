// +build !darwin

package monotime

import (
	"time"

	gm "github.com/aristanetworks/goarista/monotime"
)

type timer struct {
	start uint64
}

func monotime() uint64 {
	return gm.Now()
}

func newTimer() *timer {
	return &timer{
		start: monotime(),
	}
}

func (t *timer) Elapsed() time.Duration {
	now := monotime()
	return time.Duration(now - t.start)
}
