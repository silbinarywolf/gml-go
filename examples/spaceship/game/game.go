package game

import (
	"github.com/silbinarywolf/gml-go/examples/spaceship/asset"
	"github.com/silbinarywolf/gml-go/gml"
)

var Global = new(GameController)

type GameController struct {
	gml.Controller
	Player       gml.InstanceIndex
	ShipsSighted int
}

func (_ *GameController) GameStart() {
	gml.DrawSetFont(asset.FntDefault)

	// Setup global variables
	Global.ShipsSighted = 0

	// Create new empty room
	roomInstanceIndex := gml.RoomInstanceNew()
	roomSize := roomInstanceIndex.Size()

	// Create player in the center of the room
	Global.Player = roomInstanceIndex.InstanceCreate(roomSize.X/2, roomSize.Y/2, ObjPlayer).BaseObject().InstanceIndex()
}
