package game

import "github.com/silbinarywolf/gml-go/gml"

const ObjPlayer = 1

type Player struct {
	gml.Object
}

func (inst *Player) ObjectIndex() gml.ObjectIndex {
	return ObjPlayer
}

// todo(Jake): 2018-11-22
// Make this auto-generated based on the struct name
func (inst *Player) ObjectName() string {
	return "Player"
}

func (inst *Player) Create() {
	inst.SetSprite(gml.SpriteLoad("spaceship"))
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
