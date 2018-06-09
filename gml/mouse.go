package gml

import (
	"github.com/hajimehoshi/ebiten"
)

var (
	pressingMouseButtonLastFrame [mbSize]bool
)

func MouseCheckButton(button int) bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButton(button))
}

func MousePosition() Vec {
	x, y := ebiten.CursorPosition()
	return V(float64(x), float64(y))
}

func MouseX() float64 {
	x, _ := ebiten.CursorPosition()
	return float64(x)
}

func MouseY() float64 {
	_, y := ebiten.CursorPosition()
	return float64(y)
}

//mouse_check_button_pressed
