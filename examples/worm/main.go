package main

import (
	"math/rand"
	"time"

	"github.com/silbinarywolf/gml-go/examples/worm/game"
	"github.com/silbinarywolf/gml-go/gml"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	gml.Run(gml.GameSettings{
		GameStart:    game.GameStart,
		WindowWidth:  game.WindowWidth,
		WindowHeight: game.WindowHeight,
		WindowTitle:  game.WindowTitle,
	})
}
