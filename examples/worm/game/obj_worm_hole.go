package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

type WormHole struct {
	gml.Object
}

func (self *WormHole) Create() {
	self.SetSprite(SprWormHole)
	self.SetDepth(1)
}

func (self *WormHole) Update() {
	self.X -= WallSpeed
	if self.X+self.Size.X < 0 {
		gml.InstanceDestroy(self)
	}
}
