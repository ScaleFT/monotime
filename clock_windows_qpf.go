package monotime

import (
	"golang.org/x/sys/windows"
	"time"
	"unsafe"
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

func monotime() uint64 {
	var t uint64

	ret, _, err := procQueryPerformanceCounter.Call(uintptr(unsafe.Pointer(&t)))
	if ret == 0 {
		panic(err.Error())
	}

	return t
}

func duration(start uint64, end uint64) time.Duration {
	return time.Duration(float64((end - start)) / counterFreq)
}
