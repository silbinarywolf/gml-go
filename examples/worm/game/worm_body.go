package game

import "github.com/silbinarywolf/gml-go/gml"

const (
	SproutLerp      = 0.1
	SproutLerpSpeed = 0.1
)

type WormBody struct {
	gml.Vec
	WormLag
	HasSprouted bool
}

func (self *WormBody) SeperationWidth() float64 {
	size := SprWormHead.Size()
	return size.X/2 + size.X/6
}
