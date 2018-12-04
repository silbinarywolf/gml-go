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
	gState.draw()
}

func Run(gameStartFunc func(), updateFunc func(), width int, height int, scale float64, title string) {
	gWindowWidth = width
	gWindowHeight = height
	gWindowScale = scale

	gMainFunctions.gameStart = gameStartFunc
	gMainFunctions.update = updateFunc
	gMainFunctions.gameStart()

	ebiten.SetRunnableInBackground(true)
	ebiten.Run(updateEbiten, width, height, scale, title)
}
