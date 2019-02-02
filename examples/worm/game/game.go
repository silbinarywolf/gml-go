package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

const (
	// DesignedMaxTPS states that game logic is designed to simulate at 1/60 of a second
	// ie. alarms, move speed, animation speed
	DesignedMaxTPS = 60
)

func GameStart() {
	gml.DrawSetFont(FntDefault)

	// Setup "kinda" delta time
	gml.SetDesignedTPS(DesignedMaxTPS)
	//gml.SetMaxTPS(80)

	// Setup global variables
	// ...

	// Create new empty room
	roomInstanceIndex := gml.RoomInstanceNew()

	// Create background drawer
	gml.InstanceCreate(0, 0, roomInstanceIndex, ObjBackground)

	// Create menu
	menuInst := gml.InstanceCreate(0, 0, roomInstanceIndex, ObjMenu).(*Menu)

	// Create player in the center of the room
	playerInst := gml.InstanceCreate(0, 0, roomInstanceIndex, ObjWorm).(*Worm)

	menuInst.Player = playerInst.InstanceIndex()

	// Play song
	SndSunnyFields.Play()
}
