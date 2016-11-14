package monotime

import (
	"time"
)

type Timer interface {
	Elapsed() time.Duration
}

func New() Timer {
	return newTimer()
}
