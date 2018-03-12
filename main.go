package gml

import (
	"github.com/hajimehoshi/ebiten"
)

var g_customUpdateFunc func()

type ProgramSettings struct {
	Title         string
	Width, Height int
}

func update(s *ebiten.Image) error {
	// Set "screen"
	g_screen = s
	g_customUpdateFunc()
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

func Run(updateFunc func(), width int, height int, title string) {
	g_customUpdateFunc = updateFunc
	ebiten.Run(update, width, height, 2, title)
}
