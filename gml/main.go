package gml

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/silbinarywolf/gml-go/gml/internal/object"
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
	gMainFunctions.update()
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

func Init(idToEntityData []object.ObjectType, nameToID map[string]object.ObjectIndex) {
	object.Init(idToEntityData, nameToID)
}

func Update() {
	manager := gInstanceManager
	manager.update()
	for _, roomInst := range roomInstances {
		roomInst.update()
	}
	manager.draw()
	for _, roomInst := range roomInstances {
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
