package main

import (
	"github.com/silbinarywolf/gml-go/examples/spaceship/game"
	"github.com/silbinarywolf/gml-go/gml"
)

func main() {
	gml.Run(game.GameStart, game.GameUpdate, game.WindowWidth, game.WindowHeight, game.WindowScale, game.WindowTitle)
}
