package main

import (
	"machine"
	"time"
)

// PWM seems to use range 0-65535 rather than 0-255
const (
	unitTime = 200
	high     = 1<<16 - 1
	low      = 0
)

// delay sleeps for n milliseconds.
func delay(n int) {
	time.Sleep(time.Duration(n) * time.Millisecond)
}

// Specify what LED to blink. LED is short for D13.
var bit1, bit2, bit4, bit8 machine.PWM

// init sets up the global environment (Pulse Wave Modulator)
func init() {
	machine.InitPWM()
	bit1 = machine.PWM{Pin: machine.D6}
	bit2 = machine.PWM{Pin: machine.D9}
	bit4 = machine.PWM{Pin: machine.D10}
	bit8 = machine.PWM{Pin: machine.D11}
	bit1.Configure()
	bit2.Configure()
	bit4.Configure()
	bit8.Configure()
}

func main() {
	factors := [4]int{1, 2, 3, 2}
	for {
		for q := range factors {
			_ = q
			for i := 0; i < 16; i++ {
				high := uint16(1<<i - 1)
				if i&1 == 1 {
					bit1.Set(high)
				}
				if i&2 == 2 {
					bit2.Set(high)
				}
				if i&4 == 4 {
					bit4.Set(high)
				}
				if i&8 == 8 {
					bit8.Set(high)
				}
				delay(unitTime)
				if i&1 == 1 {
					bit1.Set(low)
				}
				if i&2 == 2 {
					bit2.Set(low)
				}
				if i&4 == 4 {
					bit4.Set(low)
				}
				if i&8 == 8 {
					bit8.Set(low)
				}
			}
		}
	}
}
