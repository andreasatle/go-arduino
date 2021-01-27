/*
Package morse contains a simple implementation of a morse-code generator.

The morse-code mappings are inside the Morse-type, but the length of the
different intervals are set as constants.

Since the mappings are inside the Morse-type, we would duplicate the mappings in the unlikely
event of having two instances of the Morse-type.

Since we are targeting the arduino-uno micro-controller, one can argue that the space
is very limited, and it would be good to have as many variables as possible defined as constants.
This would be a trade-off between generality and efficiency.

The main purpose of this code is to play around and learn how to program micro-controllers.
*/
package morse

import (
	"github.com/andreasatle/go-arduino/delay"
	"github.com/tinygo-org/tinygo/src/machine"
)

// Morse contains data for a morse-code generator.
type Morse struct {
	pin         machine.Pin
	message     string
	unit        int
	dot         int
	dash        int
	charSpace   int
	letterSpace int
	wordSpace   int
	msgSpace    int
}

// Letters contains the morse-code for uppercase letters A-Z.
type Letters [26]string

// Numbers contains the digits 0-9.
type Numbers [10]string

var letters *Letters
var numbers *Numbers

func init() {
	letters = createLetters()
	numbers = createNumbers()
}

// NewMorse creates a new instance of Morse.
func NewMorse(pin machine.Pin, unit int, message string) *Morse {
	m := &Morse{
		pin:         pin,
		message:     message,
		unit:        unit,
		dot:         unit,
		dash:        unit,
		charSpace:   unit,
		letterSpace: 3 * unit,
		wordSpace:   7 * unit,
		msgSpace:    15 * unit,
	}
	m.pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	return m
}

// Run defines the behaviour of the morse-code generator.
func (m *Morse) Run() {
	// Check if message is valid
	if !m.validMessage() {
		m.fastBlink()
		return
	}

	// Decode a valid message
	for {
		m.Encode()
	}
}

// fastBlink blinks the LED with half the unit speed.
func (m *Morse) fastBlink() {
	for {
		m.pin.High()
		delay.DelayMilli(m.unit / 2)
		m.pin.Low()
		delay.DelayMilli(m.unit / 2)
	}
}

func (m *Morse) Encode() {
	for i, c := range m.message {
		if isSpace(c) {
			delay.DelayMilli(m.wordSpace)
			continue
		}
		m.outputChar(c)

		if i+1 < len(m.message) && !isSpace([]rune(m.message)[i+1]) {
			delay.DelayMilli(m.letterSpace)
		}

	}
	delay.DelayMilli(m.msgSpace)
}

// createLetters stores array with morse-codes for capital letters
// (see wikipedia).
func createLetters() *Letters {
	letters := &Letters{
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
	return letters
}

// createNumbers stores array with morse-codes for digits
// (see wikipedia).
func createNumbers() *Numbers {
	numbers := &Numbers{
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
	return numbers
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

// press handles a valid input character.
func (m *Morse) press(c rune) {
	var n int

	switch c {
	case '.':
		n = m.dot
	case '-':
		n = m.dash
	default:
		n = 0
	}

	m.pin.High()
	delay.DelayMilli(n)
	m.pin.Low()
}

// outputChar encode the morse-code for a valid character.
func (m *Morse) outputChar(c rune) {

	var code string
	if isUpperAlpha(c) {
		code = letters[byte(c)-byte('A')]
	} else if isNum(c) {
		code = numbers[byte(c)-byte('0')]
	}

	m.press([]rune(code)[0])
	for _, c := range []rune(code)[1:] {
		delay.DelayMilli(m.charSpace)
		m.press(c)
	}
}

// validMessage checks if the message is encodable.
func (m *Morse) validMessage() bool {
	for _, c := range m.message {
		if !isValidRune(c) {
			return false
		}
	}
	return true
}
