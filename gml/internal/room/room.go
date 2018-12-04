package room

import (
	"github.com/silbinarywolf/gml-go/gml/internal/file"
)

func (room *Room) Filepath() string {
	return file.AssetDirectory + "/" + RoomDirectoryBase + "/" + room.Config.UUID
}

//func (room *Room) LayerCount() int {
//	return len(room.InstanceLayers)
//}
