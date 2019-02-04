package game

import "github.com/silbinarywolf/gml-go/gml"

const (
	SproutLerp      = 0.1
	SproutLerpSpeed = 0.1
	SeperationWidth = 40 // Taken from game logic: (sprite_get_width(sprite_index) >> 1) + (sprite_get_width(sprite_index) >> 2);
)

type WormBody struct {
	gml.Vec
	WormLag
	HasSprouted bool
}

func (self *WormBody) SeperationWidth() float64 {
	return SeperationWidth
}
