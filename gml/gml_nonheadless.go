// +build !headless

package gml

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/silbinarywolf/gml-go/gml/internal/dt"
)

var (
	gScreen *ebiten.Image
)

// SetMaxTPS sets the maximum TPS (ticks per second), that represents how many updating function is called per second. The initial value is 60.
//
// If tps is UncappedTPS, TPS is uncapped and the game is updated per frame. If tps is negative but not UncappedTPS, SetMaxTPS panics.
func SetMaxTPS(tps int) {
	ebiten.SetMaxTPS(tps)
	dt.SetMaxTPS(tps)
}

func updateAndDrawEbiten(s *ebiten.Image) error {
	gScreen = s
	result := update()
	draw()
	return result
}

func draw() {
	if !ebiten.IsDrawingSkipped() {
		context := contextUpdate()
		context.Draw()
	}
}

func runBefore() {
	ebiten.SetRunnableInBackground(true)
}

func run(gameSettings GameSettings) {
	runBefore()
	ebiten.Run(updateAndDrawEbiten, int(gameSettings.WindowWidth), int(gameSettings.WindowHeight), gameSettings.WindowScale, gameSettings.WindowTitle)
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
