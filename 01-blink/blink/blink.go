package blink

import (
	"github.com/andreasatle/go-arduino/delay"
	"github.com/tinygo-org/tinygo/src/machine"
)

// Blink contains info about a blinking LED.
type Blink struct {
	Pin   machine.Pin
	Delay int
}

// NewBlink return a new instance of a blinking LED.
func NewBlink(pin machine.Pin, delay int) *Blink {
	blink := &Blink{Pin: pin, Delay: delay}
	blink.Pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	return blink
}

// RunBlink blinks a LED with the current delay.
func (b *Blink) Run() {
	for {
		b.Pin.High()
		delay.DelayMilli(b.Delay)
		b.Pin.Low()
		delay.DelayMilli(b.Delay)
	}
}

// ControlBlink varies the delay (frequency) of the blinking LED.
func (b *Blink) Control() {
	for {
		for d := 500; d < 2000; d += 100 {
			b.Delay = d
			delay.DelayMilli(2 * d)
		}
	}
}
