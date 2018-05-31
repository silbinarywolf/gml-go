package gml

import (
	"github.com/hajimehoshi/ebiten"
)

type mainFunctions struct {
	gameStart func()
	update    func()
}

var gMainFunctions *mainFunctions = new(mainFunctions)

var gScreen *ebiten.Image

func update(s *ebiten.Image) error {
	gScreen = s
	gMainFunctions.update()
	if g_game.hasGameRestarted {
		gInstanceManager.reset()
		gMainFunctions.gameStart()
		g_game.hasGameRestarted = false
	}
	//ebitenutil.DebugPrint(s, "Hello world!")
	return nil
}

func Init(idToEntityData []ObjectType, nameToID map[string]ObjectIndex) {
	manager := gInstanceManager
	manager.idToEntityData = idToEntityData
	manager.nameToID = nameToID
}

func Update() {
	manager := gInstanceManager
	manager.update()
	manager.draw()
}

func Run(gameStartFunc func(), updateFunc func(), width int, height int, title string) {
	gMainFunctions.gameStart = gameStartFunc
	gMainFunctions.update = updateFunc

	gMainFunctions.gameStart()
	ebiten.SetRunnableInBackground(true)
	ebiten.Run(update, width, height, 2, title)
}
