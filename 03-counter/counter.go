package main

import (
	"machine"
	"time"
)

const (
	unitTime = 200
)

// delay sleeps for n milliseconds.
func delay(n int) {
	time.Sleep(time.Duration(n) * time.Millisecond)
}

// Specify what LED to blink. LED is short for D13.
var bit1 = machine.D10
var bit2 = machine.D11
var bit4 = machine.D12
var bit8 = machine.D13

// init sets up the global environment
func init() {
	bit1.Configure(machine.PinConfig{Mode: machine.PinOutput})
	bit2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	bit4.Configure(machine.PinConfig{Mode: machine.PinOutput})
	bit8.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

func main() {
	factors := [4]int{1, 2, 3, 2}
	for {
		for q := range factors {
			for i := 0; i < 16; i++ {
				if i&1 == 1 {
					bit1.High()
				}
				if i&2 == 2 {
					bit2.High()
				}
				if i&4 == 4 {
					bit4.High()
				}
				if i&8 == 8 {
					bit8.High()
				}
				delay(q * unitTime)
				if i&1 == 1 {
					bit1.Low()
				}
				if i&2 == 2 {
					bit2.Low()
				}
				if i&4 == 4 {
					bit4.Low()
				}
				if i&8 == 8 {
					bit8.Low()
				}
			}
		}
	}
}
