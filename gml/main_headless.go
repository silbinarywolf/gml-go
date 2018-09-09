// +build headless

package gml

import "time"

func Draw() {
	// no-op
}

func Run(gameStartFunc func(), updateFunc func(), width int, height int, scale float64, title string) {
	gWindowWidth = width
	gWindowHeight = height
	gWindowScale = scale

	gMainFunctions.gameStart = gameStartFunc
	gMainFunctions.update = updateFunc
	gMainFunctions.gameStart()

	// Loop
	tick := time.Tick(16 * time.Millisecond)
	for {
		select {
		case <-tick:
			updateFunc()
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
