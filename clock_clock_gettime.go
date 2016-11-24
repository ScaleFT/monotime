// +build netbsd freebsd solaris

package monotime

import (
	"syscall"
	"time"
	"unsafe"
)

func monotime() uint64 {
	var ts syscall.Timespec

	syscall.Syscall(syscall.SYS_CLOCK_GETTIME, clock_monotonic, uintptr(unsafe.Pointer(&ts)), 0)

	return uint64(ts.Nano())
}

func duration(start uint64, end uint64) time.Duration {
	return time.Duration(end - start)
}
