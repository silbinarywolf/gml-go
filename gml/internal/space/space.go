package space

import (
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
)

type Space struct {
	solid bool
	geom.Rect
}

func (record *Space) SetSolid(isSolid bool) {
	record.solid = isSolid
}

func (record *Space) Solid() bool { return record.solid }
