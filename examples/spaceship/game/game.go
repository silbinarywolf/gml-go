package game

import "github.com/silbinarywolf/gml-go/gml"

func GameStart() {
	gml.DrawSetFont(FntDefault)

	// Setup global variables
	global.ShipsSighted = 0

	// Setup camera
	gml.CameraCreate(0, 0, 0, gml.WindowWidth(), gml.WindowHeight())
	roomInstanceIndex := gml.RoomInstanceNew()

	roomSize := gml.RoomInstanceSize(roomInstanceIndex)
	inst := gml.InstanceCreate(roomSize.X/2, roomSize.Y/2, roomInstanceIndex, ObjPlayer)
	gml.CameraSetViewTarget(0, inst.BaseObject().InstanceIndex())
}

func GameUpdate() {
	gml.Update()
	gml.Draw()
}
