package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/object"
)

type ObjectIndex = object.ObjectIndex

type ObjectType = object.ObjectType

type Object = object.Object

func ObjectGetIndex(name string) (object.ObjectIndex, bool) {
	res, ok := object.ObjectGetIndex(name)
	return res, ok
}

// ObjectInitTypes is required to be called so the engine can create game objects
func ObjectInitTypes(objectTypeToData []object.ObjectType, objectIndexList []object.ObjectIndex) {
	object.InitTypes(objectTypeToData, objectIndexList)
}
