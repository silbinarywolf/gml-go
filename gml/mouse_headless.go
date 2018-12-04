// +build headless

package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
)

func MouseCheckButton(button int) bool {
	return false
}

func MouseCheckPressed(button int) bool {
	return false
}

func MousePosition() Vec {
	return geom.Vec{}
}

func mousePosition() geom.Vec {
	return geom.Vec{}
}

func mouseScreenPosition() Vec {
	return geom.Vec{}
}

func mouseUpdate() {
}
