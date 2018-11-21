package sprite

import (
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
)

const (
	CollisionMaskInherit CollisionMaskKind = 0 + iota
	CollisionMaskManual
)

type CollisionMaskKind int

type CollisionMask struct {
	Kind CollisionMaskKind
	Rect geom.Rect
}
