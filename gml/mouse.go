package gml

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/silbinarywolf/gml-go/gml"
)

var (
	pressingMouseButtonLastFrame [mbSize]bool
)

func MouseCheckButton(button int) bool {
	return ebiten.IsMouseButtonPressed(ebiten.MouseButton(button))
}

func MousePosition() Vec {
	x, y := ebiten.CursorPosition()
	return gml.V(x, y)
}

func MouseX() float64 {
	x, _ := ebiten.CursorPosition()
	return x
}

func MouseY() float64 {
	_, y := ebiten.CursorPosition()
	return y
}

//mouse_check_button_pressed
