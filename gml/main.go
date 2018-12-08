package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
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
)

func update() error {
	frameStartTime := timegml.Now()
	keyboardUpdate()
	keyboardStringUpdate()
	mouseUpdate()

	debugUpdate()

	switch debugMenuID {
	case debugMenuNone:
		gMainFunctions.update()
	case debugMenuRoomEditor:
		cameraSetActive(0)
		cameraClear(0)

		editorLazyInit()
		editorUpdate()

		cameraDraw(0)
		cameraClearActive()
	case debugMenuAnimationEditor:
		cameraSetActive(0)
		cameraClear(0)

		animationEditorUpdate()

		cameraDraw(0)
		cameraClearActive()
	default:
		panic("Invalid debug mode.")
	}
	if g_game.hasGameRestarted {
		panic("todo: Fix / test this. I assume its broken")
		gState.globalInstances.reset()
		gMainFunctions.gameStart()
		g_game.hasGameRestarted = false
	}

	// NOTE(Jake): 2018-09-29
	// Ignoring when 0 is reported. This happens on Windows
	// and just makes the frame usage timer completely helpful.
	// Not a good workaround.
	frameBudgetUsed := timegml.Now() - frameStartTime
	if frameBudgetUsed > 0 {
		gState.frameBudgetNanosecondsUsed = frameBudgetUsed
	}
	return nil
}

func WindowSize() geom.Size {
	return geom.Size{
		X: int32(gWindowWidth),
		Y: int32(gWindowHeight),
	}
}

func WindowWidth() int {
	return gWindowWidth
}

func WindowHeight() int {
	return gWindowHeight
}

func WindowScale() float64 {
	return gWindowScale
}

func Update(animationUpdate bool) {
	cameraSetActive(0)
	defer cameraClearActive()

	gState.update(animationUpdate)
}
