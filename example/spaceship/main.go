package main

import (
	"github.com/silbinarywolf/gml-go/example/spaceship/game"
	"github.com/silbinarywolf/gml-go/gml"
)

func main() {
	gml.Run(game.Global, gml.GameSettings{
		WindowWidth:  game.WindowWidth,
		WindowHeight: game.WindowHeight,
		WindowTitle:  game.WindowTitle,
	})
}
