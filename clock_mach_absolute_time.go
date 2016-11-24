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
	return uint64(C.mach_absolute_time())
}

func duration(start uint64, end uint64) time.Duration {
	elapsed := end - start
	elapsedNano := elapsed * uint64(tbinfo.numer) / uint64(tbinfo.denom)
	return time.Duration(elapsedNano)
}
