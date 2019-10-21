package game

import (
	"github.com/silbinarywolf/gml-go/gml"
)

type GameObject struct {
	gml.Object
}

type SuperGameObject struct {
	GameObject
}

type GameObjectA struct {
	SuperGameObject
}
