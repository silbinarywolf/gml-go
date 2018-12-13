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
		inst := newRawInstance(objectIndex, 0, 0, 0)
		inst.Create()
		baseObj := inst.BaseObject()

		debugObjectMetaList = append(debugObjectMetaList, debugobj.ObjectMeta{
			Rect:        baseObj.Rect,
			ObjectName:  inst.ObjectName(),
			ObjectIndex: int32(inst.ObjectIndex()),
		})
	}

	debugobj.InitDebugObjectMetaList(debugObjectMetaList)
}
