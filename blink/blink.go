package main

import (
	"machine"
	"time"
)

var intervals = []int{100, 200, 300, 400, 300, 200}

func main() {
	unit := time.Millisecond
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		for _, interval := range intervals {
			led.High()
			delay(interval, unit)
			led.Low()
			delay(500-interval, unit)
		}
	}
}

func delay(length int, unit time.Duration) {
	time.Sleep(time.Duration(interval) * unit)
}

// To get the Arduino to blink:
// tinygo flash -target arduino blink.go
