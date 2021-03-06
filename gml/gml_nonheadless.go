// +build !headless

package gml

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/silbinarywolf/gml-go/gml/internal/dt"
	"github.com/silbinarywolf/gml-go/gml/monotime"
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
	frameStartTime := monotime.Now()

	// Update/Draw
	var err error
	{
		gScreen = s
		err = update()
		draw()
	}

	gGameGlobals.frameUpdateBudgetNanosecondsUsed = monotime.Now() - frameStartTime
	gGameGlobals.tickCount++
	return err
}

func draw() {
	if !ebiten.IsDrawingSkipped() {
		context := contextUpdate()
		context.Draw()
		gGameGlobals.frameCount++
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
