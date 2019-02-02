package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

func WallSpeed() float64 {
	const wallSpeed = 8
	return wallSpeed * gml.DeltaTime()
}

type Wall struct {
	gml.Object
	DontKillPlayerIfInDirt bool
}

func (self *Wall) Create() {
	self.SetSprite(SprWall)
	self.SetDepth(2)
}

func (self *Wall) Update() {
	if Global.HasWormStopped() {
		return
	}

	self.X -= WallSpeed()
	if self.X+self.Size.X < 0 {
		gml.InstanceDestroy(self)
	}
}
