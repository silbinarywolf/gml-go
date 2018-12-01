package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

type Bullet struct {
	gml.Object
}

func (inst *Bullet) Create() {
	inst.SetSprite(SprBullet)
}

func (inst *Bullet) Destroy() {

}

func (inst *Bullet) Update() {
	inst.Y -= 8

	for _, other := range gml.CollisionRectList(inst, inst.Pos()) {
		other, ok := other.(*Player)
		if !ok {
			continue
		}
		inst.X += 8
		other.X += 1
		//gml.InstanceDestroy(other)
	}
}

func (inst *Bullet) Draw() {
	gml.DrawSelf(&inst.SpriteState, inst.Pos())
}
