package gml

import (
	"path/filepath"
	"runtime"

	"github.com/hajimehoshi/ebiten"
	"github.com/silbinarywolf/gml-go/gml/audio"
	"github.com/silbinarywolf/gml-go/gml/internal/dt"
	"github.com/silbinarywolf/gml-go/gml/internal/file"
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
	"github.com/silbinarywolf/gml-go/gml/internal/timegml"
)

type GameSettings struct {
	WindowTitle  string
	WindowWidth  float64
	WindowHeight float64
	WindowScale  float64
}

type TestSettings struct {
	PreUpdate  func()
	PostUpdate func() bool
}

const (
	defaultWindowWidth  = 1024
	defaultWindowHeight = 768
)

var (
	gController   gameController
	gGameSettings GameSettings
)

func setup(controller gameController, gameSettings *GameSettings) {
	// Setup defaults
	if gameSettings.WindowWidth == 0 {
		gameSettings.WindowWidth = defaultWindowWidth
	}
	if gameSettings.WindowHeight == 0 {
		gameSettings.WindowHeight = defaultWindowHeight
	}
	if gameSettings.WindowScale == 0 {
		gameSettings.WindowScale = 1
	}

	// Copy controller and settings
	gController = controller
	gGameSettings = *gameSettings

	// Initialize
	file.InitAssetDir()
	gCameraManager.reset()

	// Load all assets
	audio.InitAndLoadAllSounds()
	sprite.LoadAllSprites()

	// Setup TPS
	SetDesignedTPS(dt.DefaultMaxTPS)
	SetMaxTPS(dt.DefaultMaxTPS)

	// Bootup game
	controller.GameStart()
}

// MaxTPS returns the current maximum TPS.
func MaxTPS() int {
	return ebiten.MaxTPS()
}

// SetMaxTPS sets the maximum TPS (ticks per second), that represents how many updating function is called per second. The initial value is 60.
//
// If tps is UncappedTPS, TPS is uncapped and the game is updated per frame. If tps is negative but not UncappedTPS, SetMaxTPS panics.
func SetMaxTPS(tps int) {
	ebiten.SetMaxTPS(tps)
	dt.SetMaxTPS(tps)
}

// SetDesignedTPS is the ticks-per-second the game was initially designed to run at. ie. 30tps, 60tps, etc
//
// For example, if you're porting a Game Maker game that ran at 30 frames per second, you'd want this to be 30 so
// that translation of alarm logic works seamlessly.
func SetDesignedTPS(tps int) {
	dt.SetDesignedTPS(tps)
}

// DesignedTPS() returns the ticks-per-second that the game was designed to run at.
//
// Default is 60
func DesignedTPS() int {
	return dt.DesignedTPS()
}

// DeltaTime gets the fixed delta time based on the designed TPS divided by max TPS.
func DeltaTime() float64 {
	return dt.DeltaTime()
}

// TestBootstrap the game to give control over continuing / stopping execution per-frame
// this method is for additional control when testing
func TestBootstrap(controller gameController, gameSettings GameSettings, testSettings TestSettings) {
	// Set asset directory relative to the test code file path
	// for `go test` support
	_, filename, _, _ := runtime.Caller(1)
	dir := filepath.Clean(filepath.Dir(filename) + "/../" + file.AssetDirectoryBase)
	file.SetAssetDir(dir)

	setup(controller, &gameSettings)

	// NOTE(Jake): 2018-12-30
	// We currently run the update loop as fast as possible as
	// the simulation is fixed 60 FPS and we don't have a concept of delta-time
	// or anything like that (yet?)
	for {
		if testSettings.PreUpdate != nil {
			testSettings.PreUpdate()
		}
		if err := update(); err != nil {
			break
		}
		if testSettings.PostUpdate != nil &&
			!testSettings.PostUpdate() {
			break
		}
	}
}

// Run
func Run(controller gameController, gameSettings GameSettings) {
	// Setup defaults
	setup(controller, &gameSettings)
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
		// Update
		gController.GamePreUpdate()
		gState.update()
		gController.GamePostUpdate()

		// NOTE: Jake: 2019-02-24
		// `cameraUpdate` should run after all update logic so that it snaps
		// to the object being followed. If a user needs to update the camera after
		// that,
		cameraUpdate()

		// Draw
		draw()
	case debugMenuAnimationEditor:
		//debugMenuRoomEditor,
		cameraSetActive(0)
		cameraClearSurface(0)

		switch debugMenuID {
		//case debugMenuRoomEditor:
		//editorLazyInit()
		//editorUpdate()
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

	// NOTE(Jake): 2019-01-26
	// Swapped to high precision timer on Windows.
	// So this should be accurate.
	frameBudgetUsed := timegml.Now() - frameStartTime
	gState.frameBudgetNanosecondsUsed = frameBudgetUsed
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
