package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

type EnemyShip struct {
	gml.Object
}

func (inst *EnemyShip) Create() {
	inst.SetSprite(SprSpaceship)
	inst.ImageScale.Y = -1
}

func (inst *EnemyShip) Destroy() {

}

func (inst *EnemyShip) Update() {
	inst.Y += 8
}

func (inst *EnemyShip) Draw() {
	gml.DrawSelf(&inst.SpriteState, inst.Pos())
}
