// +build debug

package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/debugobj"
)

func debugInitObjectMetaList(objTypes []ObjectType) {
	debugObjectMetaList := make([]debugobj.ObjectMeta, len(objTypes))
	for _, record := range objTypes {
		if record == nil {
			continue
		}
		objectIndex := record.ObjectIndex()
		if objectIndex == 0 {
			continue
		}
		inst := allocateNewInstance(objectIndex)
		// todo(Jake): 2018-12-16:
		// Deprecate calling inst.Create() when determining
		// object size / sprite.
		inst.Create()
		baseObj := inst.BaseObject()
		meta := debugobj.ObjectMeta{
			Rect:        baseObj.Rect,
			SpriteIndex: baseObj.SpriteIndex(),
			ObjectName:  inst.ObjectName(),
			ObjectIndex: int32(inst.ObjectIndex()),
		}

		debugObjectMetaList = append(debugObjectMetaList, meta)
	}

	debugobj.InitDebugObjectMetaList(debugObjectMetaList)
}
