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
	result := update()
	draw()
	return result
}

func draw() {
	if ebiten.IsDrawingSkipped() {
		return
	}
	// Draw
	gController.GamePreDraw()
	gState.draw()
	gController.GamePostDraw()
}

func runBefore() {
	ebiten.SetRunnableInBackground(true)
}

func run(gameSettings GameSettings) {
	runBefore()
	ebiten.Run(updateEbiten, int(gameSettings.WindowWidth), int(gameSettings.WindowHeight), gameSettings.WindowScale, gameSettings.WindowTitle)
}

// runTest is run by TestBootstrap
func runTest(gameSettings GameSettings, testSettings TestSettings) {
	runBefore()
	ebiten.Run(func(s *ebiten.Image) error {
		gScreen = s
		for i := 0; i < testSettings.SpeedMultiplier; i++ {
			if err := runTestUpdateLoop(testSettings); err != nil {
				draw()
				return err
			}
		}
		draw()
		return nil
	}, int(gameSettings.WindowWidth), int(gameSettings.WindowHeight), gameSettings.WindowScale, gameSettings.WindowTitle)
}
