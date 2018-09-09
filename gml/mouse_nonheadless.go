// +build !headless

package gml

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
)

var (
	mouseButtonPress             [MbSize]int // this array is reset every frame
	pressingMouseButtonLastFrame [MbSize]bool
	_mousePos                    geom.Vec
)

func MouseCheckButton(button int) bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButton(button))
}

func MouseCheckPressed(button int) bool {
	return mouseButtonPress[button] == 1
	/*isHeld := MouseCheckButton(button)
	if !isHeld {
		pressingMouseButtonLastFrame[button] = false
	}
	if pressingMouseButtonLastFrame[button] {
		return false
	}
	if isHeld {
		pressingMouseButtonLastFrame[button] = true
	}
	return isHeld*/
}

func MousePosition() geom.Vec {
	return _mousePos
}

// Get the mouse position relative to the window
func mouseScreenPosition() geom.Vec {
	x, y := ebiten.CursorPosition()
	return geom.Vec{float64(x), float64(y)}
}

//
// NOTE(Jake): 2018-07-10
//
// Ebiten doesn't have mouseWheel() support on a stable version yet and
// it doesn't support browser mouse wheel.
// - https://github.com/hajimehoshi/ebiten/issues/630
//
// I'll look into this later!
//
/*func mouseWheel() geom.Vec {
	xoff, yoff := ebiten.MouseWheel()
	return geom.V(xoff, yoff)
}*/

func mouseUpdate() {
	x, y := ebiten.CursorPosition()
	newPos := geom.Vec{float64(x), float64(y)}

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
	viewPos := CameraGetViewPos(0)
	newPos.X += viewPos.X
	newPos.Y += viewPos.Y

	_mousePos = newPos

	// Add code to check mouse inputs
	for btn := MbLeft; btn < MbSize; btn++ {
		if MouseCheckButton(btn) {
			mouseButtonPress[btn]++
		} else {
			mouseButtonPress[btn] = 0
		}
	}
}

//mouse_check_button_pressed
