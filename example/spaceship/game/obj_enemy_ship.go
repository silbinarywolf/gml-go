package game

import (
	"github.com/silbinarywolf/gml-go/example/spaceship/asset"
	"github.com/silbinarywolf/gml-go/gml"
)

type EnemyShip struct {
	gml.Object
}

func (self *EnemyShip) Create() {
	self.SetSprite(asset.SprSpaceship)
	self.ImageScale.Y = -1
}

func (self *EnemyShip) Destroy() {
	Global.ShipsSighted += 1
}

func (self *EnemyShip) Update() {
	self.Y += 8

	// todo(Jake): 2019-01-23
	// Add some code here to destroy the player ship?

	if self.Y > self.RoomIndex().Size().Y {
		gml.InstanceDestroy(self)
		return
	}
}
