package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
	"github.com/silbinarywolf/gml-go/gml/internal/timegml"
)

type mainFunctions struct {
	gameStart func()
	update    func()
}

var gMainFunctions *mainFunctions = new(mainFunctions)

var (
	gWindowWidth  int
	gWindowHeight int
	gWindowScale  float64 // Window scale
	//lastFrameTime int64
)

func update() error {
	frameStartTime := timegml.Now()
	//frameOffset := timegml.Now() - lastFrameTime
	sprite.DebugWatch()
	keyboardUpdate()
	keyboardStringUpdate()
	mouseUpdate()
	if EditorIsActive() {
		cameraSetActive(0)
		editorUpdate()
		cameraClearActive()
	} else {
		gMainFunctions.update()
	}
	if g_game.hasGameRestarted {
		gState.globalInstances.reset()
		gMainFunctions.gameStart()
		g_game.hasGameRestarted = false
	}
	gState.frameBudgetNanosecondsUsed = timegml.Now() - frameStartTime
	//gState.frameBudgetNanosecondsUsed += frameOffset
	//lastFrameTime = timegml.Now()
	return nil
}

func windowWidth() int {
	return gWindowWidth
}

func windowHeight() int {
	return gWindowHeight
}

func windowScale() float64 {
	return gWindowScale
}

func Update(animationUpdate bool) {
	cameraSetActive(0)
	defer cameraClearActive()

	gState.update(animationUpdate)
}
