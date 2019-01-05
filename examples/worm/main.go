package main

import (
	"github.com/silbinarywolf/gml-go/examples/worm/game"
	"github.com/silbinarywolf/gml-go/gml"
)

func main() {
	gml.Run(gml.GameSettings{
		GameStart:    game.GameStart,
		WindowWidth:  game.WindowWidth,
		WindowHeight: game.WindowHeight,
		WindowTitle:  game.WindowTitle,
	})
}
