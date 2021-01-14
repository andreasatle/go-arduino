package main

import (
	"machine"
	"time"
)

// ADC seems to use range 0-65536 rather than 0-1023.
const (
	unitTime  = 200
	getToVolt = 5.0 / 65536.0
)

// delay sleeps for n milliseconds.
func delay(n int) {
	time.Sleep(time.Duration(n) * time.Millisecond)
}

// Specify what LED to blink. LED is short for D13.
var adc machine.ADC

// init sets up the global environment (Pulse Wave Modulator)
func init() {
	machine.InitADC()
	adc = machine.ADC{Pin: machine.ADC0}
}

func outputVoltage() {
	for {
		delay(unitTime)
		// Read the voltage from port A0, assuming 5V is active.
		println(float32(adc.Get()) * getToVolt)
	}
}

func main() {

	// Watch serial output:
	// screen /dev/cu.usbmodemXYZUVW (where XYZUVW are some digits)

	outputVoltage()
}
