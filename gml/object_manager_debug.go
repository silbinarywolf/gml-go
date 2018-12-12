// +build debug

package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/debugobj"
)

func debugInitObjectMetaList(objTypes []ObjectType) {
	var debugObjectMetaList []debugobj.ObjectMeta
	for _, record := range objTypes {
		objectIndex := record.ObjectIndex()
		inst := newRawInstance(objectIndex, 0, 0, 0)
		inst.Create()
		baseObj := inst.BaseObject()
		meta := debugobj.ObjectMeta{
			Vec:  baseObj.Pos(),
			Size: baseObj.Size,
		}
		debugObjectMetaList = append(debugObjectMetaList, meta)
	}

	debugobj.InitDebugObjectMetaList(debugObjectMetaList)
}
