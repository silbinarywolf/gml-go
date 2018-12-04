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

func (inst *Player) Create() {
	inst.SetSprite(SprSpaceship)
}

func (inst *Player) Destroy() {
}

func (inst *Player) Update() {
	if inst.enemyCreateAlarm.Update(60) {
		// Spawn enemies at the top of the frame, every 60 frames
		gml.InstanceCreateRoom(gml.Vec{float64(rand.Intn(gml.WindowWidth())), 0}, gameWorld.CurrentRoomIndex, ObjEnemyShip)
	}

	if gml.KeyboardCheck(gml.VkLeft) {
		inst.X -= 8
	}
	if gml.KeyboardCheck(gml.VkRight) {
		inst.X += 8
	}
	if gml.KeyboardCheck(gml.VkUp) {
		inst.Y -= 8
	}
	if gml.KeyboardCheck(gml.VkDown) {
		inst.Y += 8
	}
	if gml.KeyboardCheckPressed(gml.VkSpace) {
		bullet := gml.InstanceCreateRoom(inst.Pos(), gameWorld.CurrentRoomIndex, ObjBullet).(*Bullet)
		bullet.Owner = inst
	}
}

func (inst *Player) Draw() {
	gml.DrawSelf(&inst.SpriteState, inst.Pos())

	gml.DrawTextF(gml.Vec{0, 32}, "Score: %d", inst.Score)
}
