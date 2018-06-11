package gml

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
