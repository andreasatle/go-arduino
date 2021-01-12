package main

import (
	"machine"
	"time"
)

const (
	unit        = 200
	dot         = unit
	dash        = 3 * unit
	charSpace   = unit
	letterSpace = 3 * unit
	wordSpace   = 7 * unit
	msgSpace    = 15 * unit
)

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

func isUpperAlpha(ch rune) bool {
	return ch >= 'A' && ch <= 'Z'
}

func isNum(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func isSpace(ch rune) bool {
	return ch == ' '
}

func isValidRune(ch rune) bool {
	return isUpperAlpha(ch) || isNum(ch) || isSpace(ch)
}

func delay(n int) {
	time.Sleep(time.Duration(n) * time.Millisecond)
}

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

func validMessage() bool {
	for _, c := range msg {
		if !isValidRune(c) {
			return false
		}
	}
	return true
}

var letters [26]string
var numbers [10]string

var led = machine.LED

var msg = "MY NAME IS ANDREAS"

func init() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	createLetters()
	createNumbers()
}

func main() {
	if !validMessage() {
		for {
			led.High()
			delay(unit / 2)
			led.Low()
			delay(unit / 2)
		}
	}
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
