// +build netbsd freebsd solaris
// +build !go1.9

package monotime

import (
	"syscall"
	"time"
	"unsafe"
)

func monotime() Monotime {
	var ts syscall.Timespec

	syscall.Syscall(syscall.SYS_CLOCK_GETTIME, clock_monotonic, uintptr(unsafe.Pointer(&ts)), 0)

	return uint64(ts.Nano())
}

func duration(start Monotime, end Monotime) time.Duration {
	return time.Duration(end - start)
}
