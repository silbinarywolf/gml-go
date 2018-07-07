package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

type mainFunctions struct {
	gameStart func()
	update    func()
}

var gMainFunctions *mainFunctions = new(mainFunctions)

var (
	gWidth  int
	gHeight int
)

func update() error {
	sprite.DebugWatch()
	keyboardUpdate()
	keyboardStringUpdate()
	mouseUpdate()
	if EditorIsActive() {
		EditorUpdate()
		EditorDraw()
	} else {
		gMainFunctions.update()
	}
	if g_game.hasGameRestarted {
		gState.globalInstances.reset()
		gMainFunctions.gameStart()
		g_game.hasGameRestarted = false
	}
	return nil
}

func windowWidth() int {
	return gWidth
}

func windowHeight() int {
	return gHeight
}

func Update(animationUpdate bool) {
	cameraSetActive(0)
	defer cameraClearActive()

	gState.update(animationUpdate)
}
