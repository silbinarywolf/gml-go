package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/math"
)

type Vec = math.Vec

type Size = math.Size

func V(x float64, y float64) math.Vec {
	return math.Vec{X: x, Y: y}
}
