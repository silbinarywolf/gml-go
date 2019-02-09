// +build headless

package test

import (
	"github.com/silbinarywolf/gml-go/examples/worm/game"
	"github.com/silbinarywolf/gml-go/gml"

	"testing"
)

func TestWormGame(t *testing.T) {
	frames := 0
	gml.TestBootstrap(game.Global, gml.GameSettings{
		WindowWidth:  game.WindowWidth,
		WindowHeight: game.WindowHeight,
		WindowTitle:  game.WindowTitle,
	}, func() bool {
		defer func() {
			frames++
		}()

		if frames >= len(gameSessionTestData) {
			// Finish simulating when out of data
			return false
		}

		frameInfo := gameSessionTestData[frames]
		wormInfo := frameInfo[0]

		if inst := game.Global.Player.Get().(*game.Worm); inst != nil {
			if wormInfo.X != inst.X ||
				wormInfo.Y != inst.Y {
				t.Errorf("Frame %v: Worm not matching test data, expected {X: %v, Y: %v} but got {X: %v, Y: %v}", frames, wormInfo.X, wormInfo.Y, inst.X, inst.Y)
			}
		} else {
			t.Errorf("Frame %v: Worm missing.", frames)
		}

		return true
	})
}
