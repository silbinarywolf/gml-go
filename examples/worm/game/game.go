package game

import "github.com/silbinarywolf/gml-go/gml"

func GameStart() {
	gml.DrawSetFont(FntDefault)

	// Setup global variables
	// ...

	// Create new empty room
	roomInstanceIndex := gml.RoomInstanceNew()

	// Create background drawer
	gml.InstanceCreate(0, 0, roomInstanceIndex, ObjBackground)

	// Create player in the center of the room
	gml.InstanceCreate(0, 0, roomInstanceIndex, ObjWorm)
}
