package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

type Bullet struct {
	gml.Object
	// todo(Jake): 2018-12-02 - #24
	// Swap this to gml.InstanceIndex when ready
	// Maybe also add "vet" functionality - #25
	Owner gml.ObjectType
}

func (inst *Bullet) Create() {
	inst.SetSprite(SprBullet)
}

func (inst *Bullet) Destroy() {

}

func (inst *Bullet) Update() {
	inst.Y -= 8

	for _, other := range gml.CollisionRectList(inst, inst.Pos()) {
		other, ok := other.(*EnemyShip)
		if !ok {
			continue
		}
		owner := inst.Owner.(*Player)
		owner.Score += 1
		gml.InstanceDestroy(other)
	}
}

func (inst *Bullet) Draw() {
	gml.DrawSelf(&inst.SpriteState, inst.Pos())
}
