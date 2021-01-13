package main

import (
	"machine"
	"time"
)

// Constants for morse encoding (from wikipedia, except msgSpace).
const (
	unit        = 200
	dot         = unit
	dash        = 3 * unit
	charSpace   = unit
	letterSpace = 3 * unit
	wordSpace   = 7 * unit
	msgSpace    = 15 * unit
)

// createLetters stores array with morse-codes for capital letters
// (see wikipedia).
func createLetters() {
	letters = [26]string{
		".-",
		"-...",
		"-.-.",
		"-..",
		".",
		"..-.",
		"--.",
		"....",
		"..",
		".---",
		"-.-",
		".-..",
		"--",
		"-.",
		"---",
		".--.",
		"--.-",
		".-.",
		"...",
		"-",
		"..-",
		"...-",
		".--",
		"-..-",
		"-.--",
		"--..",
	}
}

// createNumbers stores array with morse-codes for digits
// (see wikipedia).
func createNumbers() {
	numbers = [10]string{
		"-----",
		".----",
		"..---",
		"...--",
		"....-",
		".....",
		"-....",
		"--...",
		"---..",
		"----.",
	}
}

// isUpperAlpha ckecks if input character is an uppercase letter.
func isUpperAlpha(ch rune) bool {
	return ch >= 'A' && ch <= 'Z'
}

// isNum ckecks if input character is a digit.
func isNum(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

// isSpace ckecks if input character is a space.
func isSpace(ch rune) bool {
	return ch == ' '
}

// isValid rune checks if input character is either an uppercase letter,
// a digit, or a space.
func isValidRune(ch rune) bool {
	return isUpperAlpha(ch) || isNum(ch) || isSpace(ch)
}

// delay sleeps for n milliseconds.
func delay(n int) {
	time.Sleep(time.Duration(n) * time.Millisecond)
}

// press handles a valid input character.
func press(c rune) {
	var n int

	switch c {
	case '.':
		n = dot
	case '-':
		n = dash
	default:
		n = 0
	}

	led.High()
	delay(n)
	led.Low()
}

// outputChar encode the morse-code for a valid character.
func outputChar(c rune) {

	var code string
	if isUpperAlpha(c) {
		code = letters[byte(c)-byte('A')]
	} else if isNum(c) {
		code = numbers[byte(c)-byte('0')]
	}

	press([]rune(code)[0])
	for _, c := range []rune(code)[1:] {
		delay(charSpace)
		press(c)
	}
}

// validMessage checks if the message is encodable.
func validMessage() bool {
	for _, c := range msg {
		if !isValidRune(c) {
			return false
		}
	}
	return true
}

// fastBlink blinks the LED with half the unit speed.
func fastBlink() {
	for {
		led.High()
		delay(unit / 2)
		led.Low()
		delay(unit / 2)
	}

}

var letters [26]string
var numbers [10]string

// Specify what LED to blink. LED is short for D13.
var led = machine.LED

// Specify the message to be encoded.
var msg = "ANTON"

// init sets up the global environment
func init() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	createLetters()
	createNumbers()
}

func main() {
	// Check if message is valid
	if !validMessage() {
		fastBlink()
		return
	}

	// Decode a valid message
	for {
		for i, c := range msg {
			if isSpace(c) {
				delay(wordSpace)
				continue
			}
			outputChar(c)

			if i+1 < len(msg) && !isSpace([]rune(msg)[i+1]) {
				delay(letterSpace)
			}

		}
		delay(msgSpace)
	}
}
