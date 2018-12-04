package gml

import (
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
	// and just makes the frame usage timer annoying.
	frameBudgetUsed := timegml.Now() - frameStartTime
	if frameBudgetUsed > 0 {
		gState.frameBudgetNanosecondsUsed = frameBudgetUsed
	}
	return nil
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

// todo: replace windowWidth() with WindowWidth()
func windowWidth() int {
	return gWindowWidth
}

// todo: replace windowHeight() with WindowHeight()
func windowHeight() int {
	return gWindowHeight
}

// todo: replace windowScale() with WindowScale()
func windowScale() float64 {
	return gWindowScale
}

func Update(animationUpdate bool) {
	cameraSetActive(0)
	defer cameraClearActive()

	gState.update(animationUpdate)
}
