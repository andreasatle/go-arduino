package main

import (
	"github.com/andreasatle/go-arduino/02-morse/morse"
	"github.com/tinygo-org/tinygo/src/machine"
)

const (
	morsePin = machine.D13
	msg      = "SOS"
)

func main() {
	m := morse.NewMorse(morsePin, 200, msg)
	m.Run()
}
