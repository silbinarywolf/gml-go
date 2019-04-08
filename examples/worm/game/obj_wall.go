package game

import (
	"github.com/silbinarywolf/gml-go/examples/worm/asset"
	"github.com/silbinarywolf/gml-go/gml"
)

func WallSpeed() float64 {
	const wallSpeed = 8
	return wallSpeed * gml.DeltaTime()
}

type Wall struct {
	gml.Object

	// Special flag where wall is jutting into the ground
	// but not enough that the player should die.
	DontKillPlayerIfInDirt bool

	// Special flag for when you reset the game, walls that
	// existed from the previous game will still render on-screen
	// but they won't kill you
	DontKillPlayer bool
}

func (self *Wall) Create() {
	self.SetSprite(asset.SprWall)
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
