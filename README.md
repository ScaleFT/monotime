# monotime

[![GoDoc](https://godoc.org/github.com/ScaleFT/monotime?status.svg)](https://godoc.org/github.com/ScaleFT/monotime)

`monotime` provides a portable monotonic timer.

- Linux: Uses `runtime.nanotime`
- Windows: Uses `runtime.nanotime`
- macOS: Uses `runtime.nanotime`

## About `runtime.nanotime`

The Go standard library provides `time.Now` for getting the current time,
but this timestamp is not monotonic -- it can jump forward or backwards depending
on the operating systems' time synchronization.  The standard library
does not provide a public API for getting the monotonic time from the operating system,
but on some operating systems `runtime.nanotime`, a private method, is available, and
generally is the fastest method to get a monotonic time.  When possible
this library uses `runtime.nanotime`, but on platforms where it is not available
`monotime` may use other methods.

See [Go issue #12914](https://github.com/golang/go/issues/12914) for more information.