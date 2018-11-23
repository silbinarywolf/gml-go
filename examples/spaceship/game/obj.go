package game

import "github.com/silbinarywolf/gml-go/gml"

func init() {
	gml.ObjectInitTypes([]gml.ObjectType{
		// This is used by gml.InstanceCreate to clone new instances by ObjectIndex
		ObjPlayer: new(Player),
		ObjBullet: new(Bullet),
	})
}
