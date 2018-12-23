// +build debug

package debugobj

import (
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

var gDebugObjectMetaList []ObjectMeta

type ObjectMeta struct {
	geom.Rect
	ObjectIndex int32
	ObjectName  string
	SpriteIndex sprite.SpriteIndex
}

func InitDebugObjectMetaList(debugObjectMetaList []ObjectMeta) {
	debugObjectMetaList = debugObjectMetaList
}

func DebugObjectMetaList() []ObjectMeta {
	return gDebugObjectMetaList
}

func DebugObjectMetaGet(index int32) *ObjectMeta {
	return &gDebugObjectMetaList[index]
}

func (meta *ObjectMeta) Pos() geom.Vec {
	return meta.Vec
}
