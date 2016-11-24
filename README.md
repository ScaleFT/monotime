# monotime

[![GoDoc](https://godoc.org/github.com/ScaleFT/monotime?status.svg)](https://godoc.org/github.com/ScaleFT/monotime)
[![Build Status](https://travis-ci.org/ScaleFT/monotime.svg?branch=master)](https://travis-ci.org/ScaleFT/monotime)

`monotime` provides a portable monotonic timer.

- Linux: Uses `runtime.nanotime`
- Windows: Uses `QueryPerformanceCounter` [docs](https://msdn.microsoft.com/en-us/library/windows/desktop/ms644904(v=vs.85).aspx)
- macOS: Uses `mach_absolute_time` [docs](https://developer.apple.com/library/content/qa/qa1398/_index.html) (requires CGO)

## About `runtime.nanotime`

The Go standard library provides `time.Now` for getting the current time,
but this timestamp is not monotonic -- it can jump forward or backwards depending
on the operating systems' time synchronization.  The standard library
does not provide a public API for getting the monotonic time from the operating system,
but on some operating systems `runtime.nanotime`, a private method, is available, and
generally is the fastest method to get a monotonic time.  However, on many operating systems
including macOS and Windows, `runtime.nanotime` doesn't actually implement a monotonic timer.

When possible this library uses `runtime.nanotime`, but on platforms where it is not available
`monotime` may use other methods.

See the following `golang/go` issues for more background:

- [#12914: runtime: time: expose monotonic clock source](https://github.com/golang/go/issues/12914)
- [#17610: runtime: use mach_absolute_time for runtime.nanotime](https://github.com/golang/go/issues/17610)
- [#6007: time: use monotonic time on netbsd](https://github.com/golang/go/issues/6007)