package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

type Bullet struct {
	gml.Object
	Owner gml.InstanceIndex
}

func (self *Bullet) Create() {
	self.SetSprite(SprBullet)
}

func (self *Bullet) Update() {
	self.Y -= 8

	for _, otherId := range gml.CollisionRectList(self, self.X, self.Y) {
		other, ok := otherId.Get().(*EnemyShip)
		if !ok {
			continue
		}
		owner := self.Owner.Get().(*Player)
		owner.Score += 1
		gml.InstanceDestroy(other)
	}

	if self.Y < 0 {
		gml.InstanceDestroy(self)
	}
}
