package delay

import "time"

// Delay is short for time.Sleep
func Delay(length int, unit time.Duration) {
	time.Sleep(time.Duration(length) * unit)
}

// DelayMicro is a delay in microseconds.
func DelayMicro(length int) {
	Delay(length, time.Microsecond)
}

// DelayMilli is a delay in milliseconds.
func DelayMilli(length int) {
	Delay(length, time.Millisecond)
}
