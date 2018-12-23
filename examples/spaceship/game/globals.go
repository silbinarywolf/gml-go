package game

const (
	WindowTitle  = "Spaceship"
	WindowWidth  = 640
	WindowHeight = 480
	WindowScale  = 1
)

var global Globals

type Globals struct {
	ShipsSighted int
}
