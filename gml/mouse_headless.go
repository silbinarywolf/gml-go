// +build headless

package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/math"
)

func MouseCheckButton(button int) bool {
	return false
}

func MouseCheckPressed(button int) bool {
	return false
}

func MousePosition() Vec {
	return math.V(0, 0)
}

func mousePosition() math.Vec {
	return math.V(0, 0)
}

func mouseUpdate() {
}
