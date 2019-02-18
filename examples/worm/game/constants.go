package game

const (
	// DesignedMaxTPS states that game logic is designed to simulate at 1/60 of a second
	// ie. alarms, move speed, animation speed
	DesignedMaxTPS = 60
)

const (
	WindowTitle  = "Worm in the Pipes"
	WindowWidth  = 960
	WindowHeight = 640
	WindowScale  = 1
	CreditText   = "Created by Silbinary Wolf | Art by milkroscope | Music by Magicdweedoo"
)

const (
	DepthBackground = 10
	DepthDirt       = 1
	DepthWormBody   = -5
	DepthWorm       = -10
	DepthMenu       = -15
)
