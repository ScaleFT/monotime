// +build darwin

package monotime

// #include <mach/mach_time.h>
import "C"

import (
	"time"
)

var tbinfo C.struct_mach_timebase_info

func init() {
	C.mach_timebase_info(&tbinfo)
}

func monotime() uint64 {
	t := C.mach_absolute_time()

	return uint64(t)
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
	now := monotime()
	elapsed := now - t.start
	elapsedNano := elapsed * uint64(tbinfo.numer) / uint64(tbinfo.denom)
	return time.Duration(elapsedNano)
}
