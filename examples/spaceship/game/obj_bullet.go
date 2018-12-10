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

func (self *Bullet) Create() {
	self.SetSprite(SprBullet)
}

func (self *Bullet) Destroy() {

}

func (self *Bullet) Update() {
	self.Y -= 8

	for _, other := range gml.CollisionRectList(self, self.Pos()) {
		other, ok := other.(*EnemyShip)
		if !ok {
			continue
		}
		owner := self.Owner.(*Player)
		owner.Score += 1
		gml.InstanceDestroy(other)
	}
}

func (self *Bullet) Draw() {
	gml.DrawSelf(&self.SpriteState, self.Pos())
}
