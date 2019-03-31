package game

import (
	"github.com/silbinarywolf/gml-go/cmd/gmlgo/testdata/serialization_private_error/game/sub"
	"github.com/silbinarywolf/gml-go/gml"
)

type SerializablePrivate struct {
	gml.Object
	gml.ObjectSerialize
	sub.Embed
}
