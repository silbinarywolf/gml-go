// +build !headless

package gml

import (
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
)

func WindowCursorVisible() bool {
	return ebiten.IsCursorVisible()
}

func WindowSetCursorVisible(visible bool) {
	ebiten.SetCursorVisible(visible)
}

// WindowSize will get the size of the current game window
func WindowSize() geom.Vec {
	return geom.Vec{
		X: gGameSettings.WindowWidth,
		Y: gGameSettings.WindowHeight,
	}
}

func WindowGetFullscreen() bool {
	return ebiten.IsFullscreen()
}

func WindowSetFullscreen(value bool) {
	ebiten.SetFullscreen(value)
}

func WindowSetSize(width, height float64) {
	gGameSettings.WindowWidth = width
	gGameSettings.WindowHeight = height
	ebiten.SetScreenSize(int(math.Floor(width)), int(math.Floor(height)))
}

func WindowSetScale(scale float64) {
	gGameSettings.WindowScale = scale
	ebiten.SetScreenScale(scale)
}

func WindowScale() float64 {
	return gGameSettings.WindowScale
}

// IsBrowser will return true if the game is playing in a web browser
func IsBrowser() bool {
	return isBrowser
}
