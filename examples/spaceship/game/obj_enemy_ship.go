package game

import (
	"fmt"

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
	if self.Y > WindowHeight {
		fmt.Printf("todo: Uncomment InstanceDestroy() in obj_enemy_ship, fix InstanceDestroy() method\n")
		//gml.InstanceDestroy(self)
		return
	}
}

func (self *EnemyShip) Draw() {
	gml.DrawSelf(&self.SpriteState, self.Pos())
}
