package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

type Player struct {
	gml.Object
}

func (inst *Player) Create() {
	inst.SetSprite(SprSpaceship)
}

func (inst *Player) Destroy() {
}

func (inst *Player) Update() {
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
		gml.InstanceCreateRoom(inst.Pos(), gameWorld.CurrentRoomIndex, ObjBullet)
	}
}

func (inst *Player) Draw() {
	gml.DrawSelf(&inst.SpriteState, inst.Pos())
}
