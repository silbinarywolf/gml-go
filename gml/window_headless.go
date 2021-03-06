// +build headless

package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
)

func WindowCursorVisible() bool {
	return false
}

func WindowSetCursorVisible(visible bool) {
}

func WindowGetFullscreen() bool {
	return false
}

func WindowSetFullscreen(value bool) {
}

// WindowSize will get the size of the current game window
func WindowSize() geom.Vec {
	return geom.Vec{
		X: gGameSettings.WindowWidth,
		Y: gGameSettings.WindowHeight,
	}
}

func WindowSetSize(width, height float64) {
	gGameSettings.WindowWidth = width
	gGameSettings.WindowHeight = height
}

func WindowSetScale(scale float64) {
	gGameSettings.WindowScale = scale
}

func WindowScale() float64 {
	return gGameSettings.WindowScale
}

// IsBrowser will return true if the game is playing in a web browser
func IsBrowser() bool {
	return isBrowser
}
