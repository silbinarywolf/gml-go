package gml

import (
	"github.com/hajimehoshi/ebiten"
)

type mainFunctions struct {
	gameStart func()
	update    func()
}

var gMainFunctions *mainFunctions = new(mainFunctions)

var (
	gScreen *ebiten.Image
	gWidth  int
	gHeight int
)

func update(s *ebiten.Image) error {
	gScreen = s
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
		gInstanceManager.reset()
		gMainFunctions.gameStart()
		g_game.hasGameRestarted = false
	}
	//ebitenutil.DebugPrint(s, "Hello world!")
	return nil
}

func windowWidth() int {
	return gWidth
}

func windowHeight() int {
	return gHeight
}

func Update() {
	cameraSetActive(0)
	defer cameraClearActive()

	manager := gInstanceManager
	manager.update()
	for i := 1; i < len(roomInstances); i++ {
		roomInst := &roomInstances[i]
		roomInst.update()
	}
}

func Draw() {
	manager := gInstanceManager
	manager.draw()
	for i := 1; i < len(roomInstances); i++ {
		roomInst := &roomInstances[i]
		roomInst.draw()
	}
}

func Run(gameStartFunc func(), updateFunc func(), width int, height int, title string) {
	gMainFunctions.gameStart = gameStartFunc
	gMainFunctions.update = updateFunc

	gMainFunctions.gameStart()
	ebiten.SetRunnableInBackground(true)
	gWidth = width
	gHeight = height
	ebiten.Run(update, width, height, 2, title)
}
