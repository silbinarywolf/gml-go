package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

type EnemyShip struct {
	gml.Object
}

func (self *EnemyShip) Create() {
	self.SetSprite(SprSpaceship)
	self.ImageScale.Y = -1
}

func (self *EnemyShip) Destroy() {

}

func (self *EnemyShip) Update() {
	self.Y += 8
}

func (self *EnemyShip) Draw() {
	gml.DrawSelf(&self.SpriteState, self.Pos())
}
