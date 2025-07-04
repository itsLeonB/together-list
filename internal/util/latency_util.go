package util

import "time"

func MeasureLatency(f func()) time.Duration {
	start := time.Now()
	f()
	return time.Since(start)
}
