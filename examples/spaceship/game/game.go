package game

import "github.com/silbinarywolf/gml-go/gml"

const (
	WindowTitle  = "Spaceship"
	WindowWidth  = 640
	WindowHeight = 480
	WindowScale  = 1
)

var (
	global Globals
)

type Globals struct {
	// todo(Jake): 2018-11-24 - #6
	// Change int to gml.RoomIndex
	//CurrentRoomIndex int
}

func GameStart() {
	gml.DrawSetFont(FntDefault)

	// Setup camera
	// todo(Jake): 2018-11-24 - #3
	// Change CameraCreate to use geom.Size for w/h
	gml.CameraCreate(0, 0, 0, float64(gml.WindowWidth()), float64(gml.WindowHeight()))
	currentRoomIndex := gml.RoomInstanceNew()

	// todo(Jake): 2018-12-06 - #38
	// Add function to get RoomSize from RoomIndex
	// (once gml.RoomIndex is implemented)
	startPos := gml.WindowSize().Vec()
	startPos.X /= 2
	startPos.Y /= 2
	gml.InstanceCreateRoom(startPos, currentRoomIndex, ObjPlayer)
}

func GameUpdate() {
	gml.Update(true)
	gml.Draw()
}
