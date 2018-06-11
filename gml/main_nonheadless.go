// +build !headless

package gml

import (
	"github.com/hajimehoshi/ebiten"
)

var (
	gScreen *ebiten.Image
)

func updateEbiten(s *ebiten.Image) error {
	gScreen = s
	return update()
}

func Draw() {
	manager := gInstanceManager
	manager.draw()
	for i := 1; i < len(roomInstances); i++ {
		roomInst := &roomInstances[i]
		roomInst.draw()
	}
}

func Run(gameStartFunc func(), updateFunc func(), width int, height int, title string) {
	gWidth = width
	gHeight = height

	gMainFunctions.gameStart = gameStartFunc
	gMainFunctions.update = updateFunc
	gMainFunctions.gameStart()

	ebiten.SetRunnableInBackground(true)
	ebiten.Run(updateEbiten, width, height, 2, title)
}
