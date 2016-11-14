# monotime

[![GoDoc](https://godoc.org/github.com/ScaleFT/monotime?status.svg)](https://godoc.org/github.com/ScaleFT/monotime)

`monotime` provides a portable monolithic timer.

- Linux: Uses `clock_gettime(CLOCK_MONOTONIC, ...)`
- macOS: Uses `mach_absolute_time` [docs](https://developer.apple.com/library/content/qa/qa1398/_index.html)

