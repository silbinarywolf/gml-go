// +build headless

package gml

import (
	"time"

	"github.com/silbinarywolf/gml-go/gml/internal/dt"
	"github.com/silbinarywolf/gml-go/gml/monotime"
)

// SetMaxTPS sets the maximum TPS (ticks per second), that represents how many updating function is called per second. The initial value is 60.
//
// If tps is UncappedTPS, TPS is uncapped and the game is updated per frame. If tps is negative but not UncappedTPS, SetMaxTPS panics.
func SetMaxTPS(tps int) {
	dt.SetMaxTPS(tps)
}

func draw() {
	// no-op
}

func run(gameSettings GameSettings) {
	// Loop
	tick := time.Tick(16 * time.Millisecond)
	for {
		select {
		case <-tick:
			frameStartTime := monotime.Now()
			if err := update(); err != nil {
				return
			}
			gGameGlobals.frameUpdateBudgetNanosecondsUsed = monotime.Now() - frameStartTime
			gGameGlobals.tickCount++
			// todo(Jake): 2018-07-10
			//
			// Should improve this to be more robust!
			// - https://trello.com/c/1RUkMGOx/55-improve-clock-for-headless-mode
			//
			// However, I'm deferring this effort as the way Ebiten works might change
			// in the future:
			// - https://github.com/hajimehoshi/ebiten/issues/605
			//
			// time.Sleep(time.Second / 60)
		}
	}
}

// runTest is run by TestBootstrap
func runTest(gameSettings GameSettings, testSettings TestSettings) {
	// NOTE(Jake): 2018-12-30
	// We currently run the update loop as fast as possible as
	// the simulation is fixed 60 FPS and we don't have a concept of delta-time
	// or anything like that (yet?)
	for {
		if err := runTestUpdateLoop(testSettings); err != nil {
			return
		}
	}
}
