package gml

import (
	"github.com/hajimehoshi/ebiten"
)

type customFunctions struct {
	gameStart func()
	update    func()
}

var g_customFunc customFunctions = customFunctions{}

type ProgramSettings struct {
	Title         string
	Width, Height int
}

func update(s *ebiten.Image) error {
	// Set "screen"
	g_screen = s
	g_customFunc.update()
	if g_game.hasGameRestarted {
		g_entityManager.reset()
		g_customFunc.gameStart()
		g_game.hasGameRestarted = false
	}
	//ebitenutil.DebugPrint(s, "Hello world!")
	return nil
}

func Init(idToEntityData []EntityType) {
	g_entityManager.idToEntityData = idToEntityData
}

func Update() {
	g_entityManager.update()
	g_entityManager.draw()
}

func Run(gameStartFunc func(), updateFunc func(), width int, height int, title string) {
	g_customFunc.gameStart = gameStartFunc
	g_customFunc.update = updateFunc

	g_customFunc.gameStart()
	ebiten.Run(update, width, height, 2, title)
}
