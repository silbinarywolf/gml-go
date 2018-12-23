package game

import "github.com/silbinarywolf/gml-go/gml"

func GameStart() {
	gml.DrawSetFont(FntDefault)

	// Setup global variables
	global.ShipsSighted = 0

	// Setup camera
	// todo(Jake): 2018-11-24 - #3
	// Change CameraCreate to use geom.Size for w/h
	gml.CameraCreate(0, 0, 0, float64(gml.WindowWidth()), float64(gml.WindowHeight()))
	currentRoomIndex := gml.RoomInstanceNew()

	roomSize := gml.RoomInstanceSize(currentRoomIndex)
	inst := gml.InstanceCreate(float64(roomSize.X/2), float64(roomSize.Y/2), currentRoomIndex, ObjPlayer)
	gml.CameraSetViewTarget(0, inst.BaseObject().InstanceIndex())
}

func GameUpdate() {
	gml.Update()
	gml.Draw()
}
