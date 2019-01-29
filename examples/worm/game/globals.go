package game

const (
	WindowTitle  = "Worm in the Pipes"
	WindowWidth  = 960
	WindowHeight = 640
	WindowScale  = 1
)

const (
	DepthMenu       = -15
	DepthWorm       = -10
	DepthWormBody   = -5
	DepthDirt       = 1
	DepthBackground = 10
)

var global Globals

type Globals struct {
	Score         int
	SoundDisabled bool
	MusicDisabled bool
}
