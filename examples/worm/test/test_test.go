// +build headless

package test

import (
	"github.com/silbinarywolf/gml-go/examples/worm/game"
	"github.com/silbinarywolf/gml-go/gml"

	"testing"
)

func TestGame(t *testing.T) {
	frames := 0
	gml.TestBootstrap(game.Global, gml.GameSettings{
		WindowWidth:  game.WindowWidth,
		WindowHeight: game.WindowHeight,
		WindowTitle:  game.WindowTitle,
	}, func() bool {
		defer func() {
			frames++
		}()
		// Run for approx ~3 seconds (assuming 60 FPS)
		if frames >= 60*60 {
			// todo(Jake): 2018-12-29 -
			// Replace this hardcoded #1 with the ability to query by
			// object type (just use the first player found / expect 1 player alive)
			//if _, ok := gml.InstanceIndex(1).Get().(*game.Worm); !ok {
			//	t.Errorf("Expected Player ship to still be alive after %d frames.", frames)
			//}
			return false
		}
		return true
	})
}
