package main

import (
	"time"

	"github.com/tinygo-org/tinygo/src/machine"
)

const (
	blinkPin = machine.D13
)

// Blink contains info about a blinking LED.
type Blink struct {
	pin   machine.Pin
	delay int
}

// NewBlink return a new instance of a blinking LED.
func NewBlink(pin machine.Pin, delay int) *Blink {
	blink := &Blink{pin: pin, delay: delay}
	blink.pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	return blink
}

// RunBlink blinks a LED with the current delay.
func (b *Blink) RunBlink() {
	for {
		b.pin.High()
		delay(b.delay, time.Millisecond)
		b.pin.Low()
		delay(b.delay, time.Millisecond)
	}
}

// ControlBlink varies the delay (frequency) of the blinking LED.
func (b *Blink) ControlBlink() {
	for {
		for d := 500; d < 2000; d += 100 {
			b.delay = d
			delay(2*d, time.Millisecond)
		}
	}
}

// Main routine for the blinking LED.
func main() {
	blink := NewBlink(blinkPin, 1000)
	go blink.RunBlink()
	blink.ControlBlink()
}

// delay is short for time.Sleep
func delay(length int, unit time.Duration) {
	time.Sleep(time.Duration(length) * unit)
}

// To get the Arduino to blink:
// tinygo flash -target arduino -scheduler tasks blink.go
