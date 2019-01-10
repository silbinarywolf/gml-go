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

func (self *Player) Update() {
	if self.enemyCreateAlarm.Repeat(60) {
		// Spawn enemies at the top of the frame, every 60 frames
		roomSize := gml.RoomInstanceSize(self.RoomInstanceIndex())
		gml.InstanceCreate(float64(rand.Intn(int(roomSize.X))), 0, self.RoomInstanceIndex(), ObjEnemyShip)
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
		bullet := gml.InstanceCreate(self.X, self.Y, self.RoomInstanceIndex(), ObjBullet).(*Bullet)
		bullet.Owner = self.InstanceIndex()
	}
}

func (self *Player) Draw() {
	gml.DrawSelf(&self.SpriteState, self.Pos())
	gml.DrawTextF(gml.Vec{0, 32}, "Score: %v", self.Score)
	gml.DrawTextF(gml.Vec{0, 64}, "Ships Sighted: %v", global.ShipsSighted)
}
