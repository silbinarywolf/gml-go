package game

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/silbinarywolf/gml-go/examples/spaceship/asset"
	"github.com/silbinarywolf/gml-go/gml"
	"github.com/silbinarywolf/gml-go/gml/alarm"
)

type Player struct {
	gml.Object
	enemyCreateAlarm alarm.Alarm
	Score            int
}

func (self *Player) Create() {
	self.SetSprite(asset.SprSpaceship)
}

func (self *Player) Update() {
	if self.enemyCreateAlarm.Repeat(60) {
		// Spawn enemies at the top of the frame, every 60 frames
		roomSize := self.RoomInstanceIndex().Size()
		self.RoomInstanceIndex().InstanceCreate(float64(rand.Intn(int(roomSize.X))), 0, ObjEnemyShip)
	}

	if gml.KeyboardCheck(gml.VkLeft) {
		self.X -= 8
	}
	if gml.KeyboardCheck(gml.VkRight) {
		self.X += 8
	}
	if gml.KeyboardCheck(gml.VkUp) {
		self.Y -= 8
	}
	if gml.KeyboardCheck(gml.VkDown) {
		self.Y += 8
	}
	if gml.KeyboardCheckPressed(gml.VkSpace) {
		bullet := self.RoomInstanceIndex().InstanceCreate(self.X, self.Y, ObjBullet).(*Bullet)
		bullet.Owner = self.InstanceIndex()
	}
}

func (self *Player) Draw() {
	self.Object.Draw()
	gml.DrawText(16, 16, fmt.Sprintf("Score: %v\n\nShips Sighted: %v", self.Score, Global.ShipsSighted), color.White)
}
