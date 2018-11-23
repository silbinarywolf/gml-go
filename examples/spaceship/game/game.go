package game

//go:generate gmlgo_gen

import "github.com/silbinarywolf/gml-go/gml"

const (
	WindowTitle  = "Spaceship"
	WindowWidth  = 640
	WindowHeight = 480
	WindowScale  = 1
)

var (
	gameWorld GameWorld
)

type GameWorld struct {
	// todo(Jake): 2018-11-28 - #6
	// Change int to gml.RoomIndex
	CurrentRoomIndex int
}

func GameStart() {
	// Setup camera
	// todo(Jake): 2018-11-28 - #3
	// Change CameraCreate to use geom.Size for w/h
	gml.CameraCreate(0, 0, 0, float64(gml.WindowWidth()), float64(gml.WindowHeight()))
	gameWorld.CurrentRoomIndex = gml.RoomInstanceNew()
	gml.InstanceCreateRoom(gml.Vec{32, 32}, gameWorld.CurrentRoomIndex, ObjPlayer)
}

func GameUpdate() {
	gml.Update(true)
	gml.Draw()
}
