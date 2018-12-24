// +build !headless

package gml

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
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

func Run(gameStartFunc func(), updateFunc func(), width, height float64, scale float64, title string) {
	gWindowSize = geom.Vec{
		X: width,
		Y: height,
	}
	gWindowScale = scale

	gMainFunctions.gameStart = gameStartFunc
	gMainFunctions.update = updateFunc
	gMainFunctions.gameStart()

	ebiten.SetRunnableInBackground(true)
	ebiten.Run(updateEbiten, int(width), int(height), scale, title)
}
