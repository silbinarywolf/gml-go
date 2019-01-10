package gml

import (
	"path/filepath"
	"runtime"

	"github.com/silbinarywolf/gml-go/gml/internal/file"
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
	"github.com/silbinarywolf/gml-go/gml/internal/timegml"
)

type GameSettings struct {
	GameStart    func()
	GameUpdate   func()
	WindowTitle  string
	WindowWidth  float64
	WindowHeight float64
	WindowScale  float64

	updateCallback func() bool
}

var gGameSettings GameSettings

func (gameSettings *GameSettings) setup() {
	// Setup defaults
	if gameSettings.WindowWidth == 0 {
		gameSettings.WindowWidth = 1024
	}
	if gameSettings.WindowHeight == 0 {
		gameSettings.WindowScale = 768
	}
	if gameSettings.WindowScale == 0 {
		gameSettings.WindowScale = 1
	}

	// Copy settings
	gGameSettings = *gameSettings

	// Initialize!
	file.InitAssetDir()
	sprite.LoadAllSprites()
	gCameraManager.reset()

	// Bootup game
	gGameSettings.GameStart()
}

// TestBootstrap the game to give control over continuing / stopping execution per-frame
// this method is for additional control when testing
func TestBootstrap(gameSettings GameSettings, updateCallback func() bool) {
	// Set asset directory relative to the test code file path
	// for `go test` support
	_, filename, _, _ := runtime.Caller(1)
	dir := filepath.Clean(filepath.Dir(filename) + "/../" + file.AssetDirectoryBase)
	file.SetAssetDir(dir)

	gameSettings.updateCallback = updateCallback
	gameSettings.setup()

	// NOTE(Jake): 2018-12-30
	// We currently run the update loop as fast as possible as
	// the simulation is fixed 60 FPS and we don't have a concept of delta-time
	// or anything like that (yet?)
	for {
		if err := update(); err != nil {
			return
		}
		if gGameSettings.updateCallback != nil &&
			!gGameSettings.updateCallback() {
			return
		}
	}
}

// Run
func Run(gameSettings GameSettings) {
	// Setup defaults
	gameSettings.setup()
	run(gameSettings)
}

func update() error {
	frameStartTime := timegml.Now()
	keyboardUpdate()
	keyboardStringUpdate()
	mouseUpdate()

	debugUpdate()

	switch debugMenuID {
	case debugMenuNone:
		if gGameSettings.GameUpdate == nil {
			// Default to simple Update/Draw()
			Update()
			draw()
		} else {
			gGameSettings.GameUpdate()
		}
	case debugMenuRoomEditor,
		debugMenuAnimationEditor:
		cameraSetActive(0)
		cameraClear(0)

		switch debugMenuID {
		case debugMenuRoomEditor:
			editorLazyInit()
			editorUpdate()
		case debugMenuAnimationEditor:
			animationEditorUpdate()
		default:
			panic("Invalid inner debug mode.")
		}

		cameraDraw(0)
		cameraClearActive()
	default:
		panic("Invalid debug mode.")
	}

	// NOTE(Jake): 2018-09-29
	// Ignoring when 0 is reported. This happens on Windows
	// and just makes the frame usage timer completely unhelpful.
	// Not a great workaround but eh.
	frameBudgetUsed := timegml.Now() - frameStartTime
	if frameBudgetUsed > 0 {
		gState.frameBudgetNanosecondsUsed = frameBudgetUsed
	}
	return nil
}

func WindowSize() geom.Vec {
	return geom.Vec{
		X: gGameSettings.WindowWidth,
		Y: gGameSettings.WindowHeight,
	}
}

func WindowWidth() float64 {
	return gGameSettings.WindowWidth
}

func WindowHeight() float64 {
	return gGameSettings.WindowHeight
}

func WindowScale() float64 {
	return gGameSettings.WindowScale
}

// Update runs the game logic, this includes object Update methods, room animation updates
// and more
func Update() {
	cameraSetActive(0)
	defer cameraClearActive()

	gState.update()
}
