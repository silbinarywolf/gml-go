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
	err := update()
	//if gGameSettings.updateCallback != nil &&
	//	!gGameSettings.updateCallback() {
	//	return errors.New("todo")
	//}
	return err
}

func draw() {
	if ebiten.IsDrawingSkipped() {
		//log.Printf("Warning: Rendering is slow, skipping render this frame\n")
		return
	}
	// Draw
	gController.GamePreDraw()
	gState.draw()
	gController.GamePostDraw()
}

func run(gameSettings GameSettings) {
	ebiten.SetRunnableInBackground(true)
	ebiten.Run(updateEbiten, int(gameSettings.WindowWidth), int(gameSettings.WindowHeight), gameSettings.WindowScale, gameSettings.WindowTitle)
}
