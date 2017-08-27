// +build windows
// +build !go1.9

package monotime

import (
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	modKernel32 = windows.NewLazySystemDLL("kernel32.dll")
)

var (
	procQueryPerformanceCounter   = modKernel32.NewProc("QueryPerformanceCounter")
	procQueryPerformanceFrequency = modKernel32.NewProc("QueryPerformanceFrequency")
)

var counterFreq float64

func init() {
	var freq int64
	ret, _, err := procQueryPerformanceFrequency.Call(uintptr(unsafe.Pointer(&freq)))
	if ret == 0 {
		panic(err.Error())
	}

	counterFreq = float64(freq) / 1e9
}

func monotime() Monotime {
	var t uint64

	ret, _, err := procQueryPerformanceCounter.Call(uintptr(unsafe.Pointer(&t)))
	if ret == 0 {
		panic(err.Error())
	}

	return t
}

func duration(start Monotime, end Monotime) time.Duration {
	return time.Duration(float64((end - start)) / counterFreq)
}
