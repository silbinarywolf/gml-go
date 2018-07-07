package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/room"
)

type Room = room.Room

type RoomObject = room.RoomObject

func LoadRoom(name string) *Room {
	return room.LoadRoom(name)
}
