package gml

type gameState struct {
	hasGameRestarted bool
}

var g_game gameState

func GameRestart() {
	g_game.hasGameRestarted = true
}
