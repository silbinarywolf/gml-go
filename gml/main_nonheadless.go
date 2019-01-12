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

func run(gameSettings GameSettings) {
	ebiten.SetRunnableInBackground(true)
	ebiten.Run(updateEbiten, int(gameSettings.WindowWidth), int(gameSettings.WindowHeight), gameSettings.WindowScale, gameSettings.WindowTitle)
}
