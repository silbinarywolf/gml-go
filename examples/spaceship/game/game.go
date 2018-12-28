package game

import "github.com/silbinarywolf/gml-go/gml"

func GameStart() {
	gml.DrawSetFont(FntDefault)

	// Setup global variables
	global.ShipsSighted = 0

	// Create new empty room
	roomInstanceIndex := gml.RoomInstanceNew()
	roomSize := gml.RoomInstanceSize(roomInstanceIndex)

	// Create player in the center of the room
	gml.InstanceCreate(roomSize.X/2, roomSize.Y/2, roomInstanceIndex, ObjPlayer)
}
