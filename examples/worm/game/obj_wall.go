package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

const (
	WallSpeed = 7
)

type Wall struct {
	gml.Object
}

func (self *Wall) Create() {
	self.SetSprite(SprWall)
}

func (self *Wall) Update() {
	self.X -= WallSpeed
}
