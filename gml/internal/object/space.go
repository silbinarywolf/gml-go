package object

import (
	m "github.com/silbinarywolf/gml-go/gml/internal/math"
)

type Space struct {
	m.Vec       // Position (contains X,Y)
	Size  m.Vec // Size (X,Y)
}
