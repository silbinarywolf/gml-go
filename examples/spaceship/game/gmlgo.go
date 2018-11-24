package game

import "github.com/silbinarywolf/gml-go/gml"

// todo(Jake): 2018-11-24 - Github Issue #9
// Make this file auto-generated

const (
	_         gml.ObjectIndex = 0
	ObjPlayer                 = 1
	ObjBullet                 = 2
)

func (inst *Player) ObjectIndex() gml.ObjectIndex { return ObjPlayer }
func (inst *Player) ObjectName() string           { return "Player" }

func (inst *Bullet) ObjectIndex() gml.ObjectIndex { return ObjBullet }
func (inst *Bullet) ObjectName() string           { return "Bullet" }

func init() {
	gml.ObjectInitTypes([]gml.ObjectType{
		// This is used by gml.InstanceCreate to clone new instances by ObjectIndex
		ObjPlayer: new(Player),
		ObjBullet: new(Bullet),
	})
}
