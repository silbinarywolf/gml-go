// +build !headless

package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/math"

	"github.com/hajimehoshi/ebiten"
)

var (
	pressingMouseButtonLastFrame [mbSize]bool
	_mousePos                    math.Vec
)

func MouseCheckButton(button int) bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButton(button))
}

func MouseCheckPressed(button int) bool {
	isHeld := MouseCheckButton(button)
	if !isHeld {
		pressingMouseButtonLastFrame[button] = false
	}
	if pressingMouseButtonLastFrame[button] {
		return false
	}
	if isHeld {
		pressingMouseButtonLastFrame[button] = true
	}
	return isHeld
}

func MousePosition() Vec {
	return mousePosition()
}

/*func MouseX() float64 {
	x, _ := ebiten.CursorPosition()
	return float64(x)
}

func MouseY() float64 {
	_, y := ebiten.CursorPosition()
	return float64(y)
}*/

func mousePosition() math.Vec {
	return _mousePos
}

func mouseUpdate() {
	x, y := ebiten.CursorPosition()
	newPos := math.V(float64(x), float64(y))

	// NOTE(Jake): 2018-06-09
	//
	// We offset the mouse position to the location
	// in the world, like Game Maker.
	//
	// This is probably incorrect to only use the
	// hardcode to camera 0, it should probably account for the
	// camera you're clicking into and then offset.
	//
	// This is future-me's problem though!
	//
	cam := &cameraList[0]
	newPos.X += cam.X
	newPos.Y += cam.Y

	_mousePos = newPos
}

//mouse_check_button_pressed
