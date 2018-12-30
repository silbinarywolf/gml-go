// +build headless

package gml

import (
	"time"
)

func draw() {
	// no-op
}

func run(gameSettings GameSettings) {
	// Loop
	tick := time.Tick(16 * time.Millisecond)
	for {
		select {
		case <-tick:
			if err := update(); err != nil {
				return
			}

			if gGameSettings.updateCallback != nil &&
				!gGameSettings.updateCallback() {
				return
			}
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
