// +build headless

package gml

import "time"

func Draw() {
	// no-op
}

func Run(gameStartFunc func(), updateFunc func(), width int, height int, title string) {
	gWidth = width
	gHeight = height

	gMainFunctions.gameStart = gameStartFunc
	gMainFunctions.update = updateFunc
	gMainFunctions.gameStart()

	// Loop
	tick := time.Tick(16 * time.Millisecond)
	for {
		select {
		case <-tick:
			updateFunc()
		}
	}
}
