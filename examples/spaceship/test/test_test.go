// +build headless

package test

import (
	"github.com/silbinarywolf/gml-go/examples/spaceship/game"
	"github.com/silbinarywolf/gml-go/gml"

	"testing"
)

func TestGame(t *testing.T) {
	frames := 0
	gml.TestBootstrap(gml.GameSettings{
		GameStart:    game.GameStart,
		WindowWidth:  game.WindowWidth,
		WindowHeight: game.WindowHeight,
		WindowTitle:  game.WindowTitle,
	}, func() bool {
		defer func() {
			frames++
		}()
		// Run for approx ~3 seconds (assuming 60 FPS)
		if frames > 383272 {
			switch gml.InstanceGet(1).(type) {
			case *game.Player:
				return false
			default:
				t.Errorf("Expected Player ship to still be alive after %d frames.", frames)
				return false
			}
		}
		return true
	})
}
