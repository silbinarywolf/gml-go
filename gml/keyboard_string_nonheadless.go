// +build !headless

package gml

import (
	"bytes"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

var (
	keyboardString bytes.Buffer
)

func ClearKeyboardString() {
	keyboardString.Reset()
}

func SetKeyboardString(text string) {
	keyboardString.Reset()
	keyboardString.WriteString(text)
}

// KeyboardString returns the last letters typed by the user since ClearKeyboardString was called
// this is useful for input boxes and easily getting input from a user.
func KeyboardString() string {
	return keyboardString.String()
}

func keyboardStringUpdate() {
	// Update keyboard string
	inputChars := ebiten.InputChars()
	for _, char := range inputChars {
		keyboardString.WriteRune(char)
	}

	// NOTE(Jake): 2018-06-02
	//
	// We don't do this as it renders an ugly square character by default.
	// Also we just don't need to retain the newline for our use-case.
	//
	// If the enter key is pressed, add a line break.
	//if repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyKPEnter) {
	//	keyboardString.WriteByte('\n')
	//}

	// If the backspace key is pressed, remove one character.
	if repeatingKeyPressed(ebiten.KeyBackspace) {
		if keyboardString.Len() >= 1 {
			var lastPos int = -1
			for pos, _ := range keyboardString.String() {
				lastPos = pos
			}
			if lastPos != -1 {
				keyboardString.Truncate(lastPos)
			}
		}
	}
}

// NOTE(Jake): 2018-06-02
//
// Taken as-is from Ebiten Typewriter example
//
// repeatingKeyPressed return true when key is pressed considering the repeat state.
func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}
