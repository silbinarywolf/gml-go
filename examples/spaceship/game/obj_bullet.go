package game

import "github.com/silbinarywolf/gml-go/gml"

type Bullet struct {
	gml.Object
}

func (inst *Bullet) Create() {
	inst.SetSprite(gml.SpriteLoad("Spaceship"))
}

func (inst *Bullet) Destroy() {

}

func (inst *Bullet) Update() {
	inst.Y -= 8
}

func (inst *Bullet) Draw() {
	gml.DrawSelf(&inst.SpriteState, inst.Pos())
}
