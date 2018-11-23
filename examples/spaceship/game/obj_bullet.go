package game

import "github.com/silbinarywolf/gml-go/gml"

const ObjBullet = 2

type Bullet struct {
	gml.Object
}

func (inst *Bullet) ObjectIndex() gml.ObjectIndex {
	return ObjBullet
}

func (inst *Bullet) ObjectName() string {
	return "Bullet"
}

func (inst *Bullet) Create() {
	inst.SetSprite(gml.SpriteLoad("spaceship"))
}

func (inst *Bullet) Destroy() {

}

func (inst *Bullet) Update() {
	inst.Y -= 8
}

func (inst *Bullet) Draw() {
	gml.DrawSelf(&inst.SpriteState, inst.Pos())
}
