package game

import (
	"math/rand"

	"github.com/silbinarywolf/gml-go/examples/worm/game/wall"
	"github.com/silbinarywolf/gml-go/gml"
)

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

func GameSpawnWall(roomInstanceIndex gml.RoomInstanceIndex) {
	const WallX = 976
	wallSets := wall.WallSets()
	wallSet := wallSets[rand.Intn(len(wallSets))]
	wallInfo := wallSet[rand.Intn(len(wallSet))]
	for _, y := range wallInfo.WallYList {
		gml.InstanceCreate(WallX, y, roomInstanceIndex, ObjWall)
	}
}
