// +build !darwin,!windows

package monotime

import (
	"syscall"
	"time"
	"unsafe"
)

const (
	CLOCK_MONOTONIC = 1
)

func monotime() int64 {
	var ts syscall.Timespec

	syscall.Syscall(syscall.SYS_CLOCK_GETTIME, CLOCK_MONOTONIC, uintptr(unsafe.Pointer(&ts)), 0)

	return ts.Nano()
}

type timer struct {
	start int64
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
