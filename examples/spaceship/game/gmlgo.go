// Code generated ;0.1.0; DO NOT EDIT.

package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

const (
	ObjBullet gml.ObjectIndex = 1
	ObjPlayer gml.ObjectIndex = 2
)

func (inst *Bullet) ObjectIndex() gml.ObjectIndex { return ObjBullet }
func (inst *Bullet) ObjectName() string           { return "Bullet" }

func (inst *Player) ObjectIndex() gml.ObjectIndex { return ObjPlayer }
func (inst *Player) ObjectName() string           { return "Player" }

func init() {
	gml.ObjectInitTypes([]gml.ObjectType{
		ObjBullet: new(Bullet),
		ObjPlayer: new(Player),
	})
}
