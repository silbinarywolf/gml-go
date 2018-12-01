package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/room"
)

// todo(Jake): 2018-12-01
// Remove this, change LoadRoom functions to return gml.RoomIndex
type Room = room.Room

func LoadRoom(name string) *Room {
	return room.LoadRoom(name)
}
