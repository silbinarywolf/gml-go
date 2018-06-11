package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/object"
)

type ObjectIndex = object.ObjectIndex

type ObjectType = object.ObjectType

type Object = object.Object

func ObjectGetIndex(name string) (object.ObjectIndex, bool) {
	nameToID := object.NameToID()
	res, ok := nameToID[name]
	return res, ok
}

func ObjectInitTypes(objTypes []object.ObjectType) {
	object.InitTypes(objTypes)
}
