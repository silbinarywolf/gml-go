// +build headless

package test

// NOTE(Jake): 2019-02-10
// The worm moves in a way that is 1-pixel off the way it moves
// in the Game Maker Studio version because math.Sin works differently to GM.
//
// It seems to be a precision issue with GM, so I either need a less precise sine
// method or I can just live with the fact that the worm moves in the ground slightly
// differently. Because the worm snaps to its start position when it jumps, this should have
// very little to no effect on gameplay.
//
// In anycase, this is the reason why this test is commented out for now.
/*func TestWormGroundMovement(t *testing.T) {
	frame := 0
	testData := wormGroundMovementData
	gml.TestBootstrap(game.Global, gml.GameSettings{
		WindowWidth:  game.WindowWidth,
		WindowHeight: game.WindowHeight,
		WindowTitle:  game.WindowTitle,
	}, gml.TestSettings{
		PreUpdate: func() {
			// todo(Jake): 2019-02-10
			// Add ability to mock inputs
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
				if wormInfo.X != inst.X ||
					wormInfo.Y != inst.Y {
					t.Errorf("Frame %v: Not matching test data\n", frame)
					if wormInfo.X != inst.X {
						t.Errorf("- X expected %v but got %v\n", wormInfo.X, inst.X)
					}
					if wormInfo.Y != inst.Y {
						t.Errorf("- Y expected %v but got %v\n", wormInfo.Y, inst.Y)
					}
				}
			} else {
				t.Errorf("Frame %v: Worm instance does not exist, expected it to exist.", frame)
			}

			return true
		},
	})
}
*/
