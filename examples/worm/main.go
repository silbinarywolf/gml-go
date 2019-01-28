package main

import (
	"math/rand"
	"time"

	//"net/http"
	//_ "net/http/pprof"

	"github.com/silbinarywolf/gml-go/examples/worm/game"
	"github.com/silbinarywolf/gml-go/gml"
)

func main() {
	//trace.Start(os.Stderr)
	//defer trace.Stop()
	//go func() {
	//	log.Println(http.ListenAndServe("localhost:6060", nil))
	//  }()

	// NOTE(Jake): 2018-01-06
	// Set here so that if we add tests, they're deterministic.
	rand.Seed(time.Now().UTC().UnixNano())

	gml.Run(gml.GameSettings{
		GameStart:    game.GameStart,
		WindowWidth:  game.WindowWidth,
		WindowHeight: game.WindowHeight,
		WindowTitle:  game.WindowTitle,
	})
}
