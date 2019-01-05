package game

const (
	WindowTitle  = "Worm in the Pipes"
	WindowWidth  = 960
	WindowHeight = 640
	WindowScale  = 1
)

var global Globals

type Globals struct {
	Score         int
	SoundDisabled bool
	MusicDisabled bool
}
