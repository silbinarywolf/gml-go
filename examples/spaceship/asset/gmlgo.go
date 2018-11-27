package asset

import "github.com/silbinarywolf/gml-go/gml"

// todo(Jake): 2018-11-24
// Auto-generate this file

const (
	SprSpaceship gml.SpriteIndex = 1
	SprBullet    gml.SpriteIndex = 2
)

func init() {
	gml.SpriteInitializeIndexToName([]string{
		SprSpaceship: "Spaceship",
		SprBullet:    "Bullet",
	}, map[string]gml.SpriteIndex{
		"Spaceship": SprSpaceship,
		"Bullet":    SprBullet,
	})
}
