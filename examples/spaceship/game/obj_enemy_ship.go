package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

type EnemyShip struct {
	gml.Object
}

func (self *EnemyShip) Create() {
	self.SetSprite(SprSpaceship)
	self.ImageScale.Y = -1
}

func (self *EnemyShip) Destroy() {
	Global.ShipsSighted += 1
}

func (self *EnemyShip) Update() {
	self.Y += 8

	// todo(Jake): 2019-01-23
	// Add some code here to destroy the player ship?

	// todo(Jake): 2018-12-06 - #38
	// Add function to get RoomSize from RoomInstanceIndex
	// (once gml.RoomInstanceIndex is implemented)
	if self.Y > gml.RoomInstanceSize(self.RoomInstanceIndex()).Y {
		gml.InstanceDestroy(self)
		return
	}
}
