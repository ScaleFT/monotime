// +build darwin
// +build !go1.9

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

func monotime() Monotime {
	return Monotime(C.mach_absolute_time())
}

func duration(start Monotime, end Monotime) time.Duration {
	elapsed := end - start
	elapsedNano := elapsed * uint64(tbinfo.numer) / uint64(tbinfo.denom)
	return time.Duration(elapsedNano)
}
