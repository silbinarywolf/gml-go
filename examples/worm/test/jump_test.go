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

			if inst := game.Global.Player.Get().(*game.Worm); inst != nil {
				// NOTE(Jake): 2019-02-10
				// The Game Maker Studio 2 outputs of the float values to 2 decimal places.
				// So we manipulate our Y position to match 2 decimal places.
				wormY := math.Floor(inst.Y*100) / 100

				if wormInfo.X != inst.X ||
					wormInfo.Y != wormY {
					t.Errorf("Frame %v: Not matching test data\n", frame)
					if wormInfo.X != inst.X {
						t.Errorf("- X expected %v but got %v\n", wormInfo.X, inst.X)
					}
					if wormInfo.Y != wormY {
						t.Errorf("- Y expected %v but got %v\n", wormInfo.Y, wormY)
					}
				}

				for i, _ := range inst.BodyParts {
					bodyPart := &inst.BodyParts[i]
					if i+1 >= len(frameInfo) {
						if bodyPart.HasSprouted {
							t.Errorf("- Worm Body Part %d - should not exist.", i)
						}
						break
					}
					bodyInfo := frameInfo[1+i]
					bodyPartY := math.Floor(bodyInfo.Y*100) / 100
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
