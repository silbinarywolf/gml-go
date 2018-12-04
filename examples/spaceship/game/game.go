package game

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
	// todo(Jake): 2018-11-24 - #6
	// Change int to gml.RoomIndex
	CurrentRoomIndex int
}

func GameStart() {
	// todo(Jake): 2018-11-24 - #15
	// - Simplify this so that you can just pass "asset.AlteHaasGroteskRegular"?
	// - Change LoadFont to return FontIndex
	gml.DrawSetFont(gml.LoadFont("AlteHaasGroteskRegular", gml.FontSettings{
		Size: 16, // 12pt == 16px
		DPI:  96,
	}))

	// Setup camera
	// todo(Jake): 2018-11-24 - #3
	// Change CameraCreate to use geom.Size for w/h
	gml.CameraCreate(0, 0, 0, float64(gml.WindowWidth()), float64(gml.WindowHeight()))
	gameWorld.CurrentRoomIndex = gml.RoomInstanceNew()
	gml.InstanceCreateRoom(gml.Vec{float64(gml.WindowWidth()) / 2, float64(gml.WindowHeight()) / 2}, gameWorld.CurrentRoomIndex, ObjPlayer)
}

func GameUpdate() {
	gml.Update(true)
	gml.Draw()
}
