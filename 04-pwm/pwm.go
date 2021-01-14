package main

import (
	"machine"
	"time"
)

// PWM seems to use range 0-65535 rather than 0-255.
const (
	unitTime = 200
	high     = uint16(1<<16 - 1)
	low      = uint16(0)
)

// delay sleeps for n milliseconds.
func delay(n int) {
	time.Sleep(time.Duration(n) * time.Millisecond)
}

// Specify what LED to blink. LED is short for D13.
var bit [4]machine.PWM

// init sets up the global environment (Pulse Wave Modulator)
func init() {
	machine.InitPWM()
	bit[0] = machine.PWM{Pin: machine.D6}
	bit[1] = machine.PWM{Pin: machine.D9}
	bit[2] = machine.PWM{Pin: machine.D10}
	bit[3] = machine.PWM{Pin: machine.D11}
	bit[0].Configure()
	bit[1].Configure()
	bit[2].Configure()
	bit[3].Configure()
}

func set(i, pwm uint16) {
	for j := 0; j < 4; j++ {
		mask := uint16(1 << j)
		if i&mask == mask {
			bit[j].Set(pwm)
		}
	}
}

func binaryCounter() {
	factors := [4]int{1, 2, 3, 2}
	for {
		for q := range factors {
			_ = q
			for i := uint16(0); i < 16; i++ {
				brightness := uint16(1<<i - 1)
				set(i, brightness)
				delay(unitTime)
				set(i, low)
			}
		}
	}
}

func main() {
	binaryCounter()
}
