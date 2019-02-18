// +build headless

package test

import (
	"github.com/silbinarywolf/gml-go/examples/spaceship/game"
	"github.com/silbinarywolf/gml-go/gml"

	"testing"
)

func TestGame(t *testing.T) {
	frame := 0
	gml.TestBootstrap(game.Global, gml.GameSettings{
		WindowWidth:  game.WindowWidth,
		WindowHeight: game.WindowHeight,
		WindowTitle:  game.WindowTitle,
	}, gml.TestSettings{
		PostUpdate: func() bool {
			defer func() {
				frame++
			}()
			// Run for approx ~3 seconds (assuming 60 FPS)
			if frame >= gml.DesignedTPS()*3 {
				if _, ok := game.Global.Player.Get().(*game.Player); !ok {
					t.Errorf("Frame %v: Expected Player ship to still be alive.", frame)
				}
				return false
			}
			return true
		},
	})
}
