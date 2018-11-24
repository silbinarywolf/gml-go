package asset

import "github.com/silbinarywolf/gml-go/gml"

// todo(Jake): 2018-11-24
// Auto-generate this file

const (
	_            gml.SpriteIndex = 0
	SprSpaceship                 = 1
)

func init() {
	gml.SpriteInitializeIndexToName([]string{
		SprSpaceship: "Spaceship",
	}, map[string]gml.SpriteIndex{
		"Spaceship": SprSpaceship,
	})
}
