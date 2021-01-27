package main

import (
	"github.com/andreasatle/go-arduino/01-blink/blink"
	"github.com/tinygo-org/tinygo/src/machine"
)

const (
	blinkPin     = machine.D13
	initialDelay = 1000
)

// Main routine for the blinking LED.
func main() {
	b := blink.NewBlink(blinkPin, initialDelay)
	go b.Run()
	b.Control()
}

// To get the Arduino to blink:
// tinygo flash -target arduino -scheduler tasks blink.go
