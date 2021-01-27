package delay

import "time"

// Delay is short for time.Sleep
func WithUnit(length int, unit time.Duration) {
	time.Sleep(time.Duration(length) * unit)
}

// DelayMicro is a delay in microseconds.
func Microseconds(length int) {
	WithUnit(length, time.Microsecond)
}

// DelayMilli is a delay in milliseconds.
func Milliseconds(length int) {
	WithUnit(length, time.Millisecond)
}
