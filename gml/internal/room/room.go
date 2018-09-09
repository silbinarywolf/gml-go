package room

import (
	"github.com/silbinarywolf/gml-go/gml/internal/file"
)

func (room *Room) Filepath() string {
	return file.AssetsDirectory + "/room/" + room.Config.UUID
}

//func (room *Room) LayerCount() int {
//	return len(room.InstanceLayers)
//}
