// +build headless

package test

import (
	"math"

	"github.com/silbinarywolf/gml-go/examples/worm/game"
	"github.com/silbinarywolf/gml-go/examples/worm/game/input"
	"github.com/silbinarywolf/gml-go/gml"

	"testing"
)

// TestWormJump will test what happens when a user presses jump immediately
// and compares the X/Y positions to what they are in the Game Maker Studio version
// of the game.
func TestWormJump(t *testing.T) {
	frame := 0
	testData := wormJumpData
	gml.TestBootstrap(game.Global, gml.GameSettings{
		WindowWidth:  game.WindowWidth,
		WindowHeight: game.WindowHeight,
		WindowTitle:  game.WindowTitle,
	}, gml.TestSettings{
		PreUpdate: func() {
			input.TestResetJumpPressed()
			if frame >= len(testData) {
				// No data to manipulate
				return
			}
			frameInfo := testData[frame]
			wormInfo := frameInfo[0]
			if wormInfo.HasPressedJump {
				input.TestSetJumpPressed(true)
			}
		},
		PostUpdate: func() bool {
			defer func() {
				frame++
			}()

			if frame >= len(testData) {
				// Finish simulating when out of data
				return false
			}

			frameInfo := testData[frame]
			wormInfo := frameInfo[0]
			bodyParts := frameInfo[1:]

			if inst := game.Global.Player.Get().(*game.Worm); inst != nil {
				// NOTE(Jake): 2019-02-10
				// Game Maker Studio 2 code seems to be more imprecise with
				// floating point values than Go. Since the logic is sound
				// when rounded to 2 decimal places, we simply round the test
				// data and current simulated game to 2 decimal points for comparison
				// to ensure the game is loosely running the same way.
				wormInfoY := math.Round(wormInfo.Y*100) / 100
				wormY := math.Round(inst.Y*100) / 100

				if wormInfo.X != inst.X ||
					wormInfoY != wormY {
					t.Errorf("Frame %v: Not matching test data\n", frame)
					if wormInfo.X != inst.X {
						t.Errorf("- X expected %v but got %v\n", wormInfo.X, inst.X)
					}
					if wormInfoY != wormY {
						t.Errorf("- Y expected %v but got %v\n", wormInfoY, wormY)
					}
				}

				for i, _ := range inst.BodyParts {
					bodyPart := &inst.BodyParts[i]
					if i >= len(bodyParts) {
						if bodyPart.HasSprouted {
							t.Errorf("- Worm Body Part %d - should not exist.", i)
						}
						break
					}
					bodyInfo := bodyParts[i]
					bodyPartY := bodyInfo.Y // math.Floor(bodyInfo.Y*100) / 100
					if bodyInfo.X != bodyPart.X ||
						bodyInfo.Y != bodyPartY {
						t.Errorf("Frame %v: Not matching test data\n", frame)
						if bodyInfo.X != bodyPart.X {
							t.Errorf("- Worm Body Part %d - X expected %v but got %v\n", i, bodyInfo.X, bodyPart.X)
						}
						if bodyInfo.Y != bodyPartY {
							t.Errorf("- Worm Body Part %d - Y expected %v but got %v\n", i, bodyInfo.Y, bodyPartY)
						}
					}
				}
			} else {
				t.Errorf("Frame %v: Worm instance does not exist, expected it to exist.", frame)
			}

			return true
		},
	})
}
