package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

const (
	WallSpeed = 8
)

type Wall struct {
	gml.Object
	DontKillPlayerIfInDirt bool
}

func (self *Wall) Create() {
	self.SetSprite(SprWall)
	self.SetDepth(2)
}

func (self *Wall) Update() {
	self.X -= WallSpeed
	if self.X+self.Size.X < 0 {
		gml.InstanceDestroy(self)
	}
}
