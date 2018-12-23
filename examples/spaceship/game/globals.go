package game

const (
	WindowTitle  = "Spaceship"
	WindowWidth  = 640
	WindowHeight = 480
	WindowScale  = 1
)

var global Globals

// Globals is a structure where you can define all your global variables
type Globals struct {
	ShipsSighted int
}
