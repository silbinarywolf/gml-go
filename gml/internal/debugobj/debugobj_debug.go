// +build debug

package debugobj

import (
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

var gDebugObjectMetaList []ObjectMeta

type ObjectMeta struct {
	geom.Vec
	Size        geom.Size
	SpriteIndex sprite.SpriteIndex
}

func InitDebugObjectMetaList(debugObjectMetaList []ObjectMeta) {
	debugObjectMetaList = debugObjectMetaList
}

func DebugObjectMetaList() []ObjectMeta {
	return gDebugObjectMetaList
}
