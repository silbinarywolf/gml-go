package game

import "github.com/silbinarywolf/gml-go/gml"

func GameStart() {
	gml.DrawSetFont(FntDefault)

	// Setup global variables
	global.ShipsDestroyed = 0

	// Setup camera
	// todo(Jake): 2018-11-24 - #3
	// Change CameraCreate to use geom.Size for w/h
	gml.CameraCreate(0, 0, 0, float64(gml.WindowWidth()), float64(gml.WindowHeight()))
	currentRoomIndex := gml.RoomInstanceNew()

	// todo(Jake): 2018-12-06 - #38
	// Add function to get RoomSize from RoomInstanceIndex
	// (once gml.RoomInstanceIndex is implemented)
	windowSize := gml.WindowSize().Vec()
	inst := gml.InstanceCreate(windowSize.X/2, windowSize.Y/2, currentRoomIndex, ObjPlayer)
	gml.CameraSetViewTarget(0, inst.BaseObject().InstanceIndex())
}

func GameUpdate() {
	gml.Update(true)
	gml.Draw()
}
