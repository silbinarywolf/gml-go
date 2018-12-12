package game

import (
	"math/rand"

	"github.com/silbinarywolf/gml-go/gml"
)

type Player struct {
	gml.Object
	enemyCreateAlarm gml.Alarm
	Score            int
}

func (self *Player) Create() {
	self.SetSprite(SprSpaceship)
}

func (self *Player) Destroy() {
}

func (self *Player) Update() {
	if self.enemyCreateAlarm.Update(60) {
		// Spawn enemies at the top of the frame, every 60 frames
		gml.InstanceCreateRoom(gml.Vec{float64(rand.Intn(gml.WindowWidth())), 0}, self.RoomInstanceIndex(), ObjEnemyShip)
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
		bullet := gml.InstanceCreateRoom(self.Pos(), self.RoomInstanceIndex(), ObjBullet).(*Bullet)
		bullet.Owner = self
	}
}

func (self *Player) Draw() {
	gml.DrawSelf(&self.SpriteState, self.Pos())

	gml.DrawTextF(gml.Vec{0, 32}, "Score: %d", self.Score)
}
