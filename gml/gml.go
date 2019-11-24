package gml

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"

	"github.com/silbinarywolf/gml-go/gml/assetman"
	"github.com/silbinarywolf/gml-go/gml/internal/dt"
	"github.com/silbinarywolf/gml-go/gml/internal/file"
	_ "github.com/silbinarywolf/gml-go/gml/internal/paniccatch"
	"github.com/silbinarywolf/gml-go/gml/monotime"
)

type DefaultContext struct{}

func (context *DefaultContext) Update() {
	// Update
	gController.GamePreUpdate()
	gState.update()
	gController.GamePostUpdate()

	// Remove deleted entities at safe point
	// ie. not while executign user-code / at the end of the frame
	gState.removeDeletedEntities()

	// NOTE: Jake: 2019-02-24
	// `cameraUpdate` should run after all update logic so that it snaps
	// to the object being followed. If a user needs custom camera behaviour,
	// they can leverage CameraSetUpdateFunction()
	cameraUpdate()
}

type GameSettings struct {
	WindowTitle  string
	WindowWidth  float64
	WindowHeight float64
	WindowScale  float64
}

type TestSettings struct {
	PreUpdate  func()
	PostUpdate func() bool
	// SpeedMultiplier is the number of times Update() methods are called
	// per frame when running tests in headed mode. If set to 0, it will default to 1.
	SpeedMultiplier int
}

const (
	defaultWindowWidth  = 1024
	defaultWindowHeight = 768
)

var (
	gController         gameController
	gGameSettings       GameSettings
	gUpdateContextStack []contextUpdateLoop
	errGameEnd          = errors.New("Game ended")
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
	initDraw()

	// Load all assets
	assetman.UnsafeLoadAll()

	// Setup TPS
	SetDesignedTPS(dt.DefaultMaxTPS)
	SetMaxTPS(dt.DefaultMaxTPS)

	// Setup default context
	// note: we dont call ContextUpdatePush() as it checks the current context first
	//		 which isn't pushed / doesn't exist yet.
	gUpdateContextStack = append(gUpdateContextStack, &DefaultContext{})

	// Bootup game
	controller.GameStart()
}

func runTestUpdateLoop(testSettings TestSettings) error {
	if testSettings.PreUpdate != nil {
		testSettings.PreUpdate()
	}
	if err := update(); err != nil {
		return err
	}
	if testSettings.PostUpdate != nil &&
		!testSettings.PostUpdate() {
		return errors.New("Test exited")
	}
	return nil
}

// MaxTPS returns the current maximum TPS.
func MaxTPS() int {
	return dt.MaxTPS()
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
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		panic("Cannot find asset directory: " + dir + ", integration tests must go inside \"{project}/test\"")
	}
	file.SetAssetDir(dir)

	setup(controller, &gameSettings)
	runTest(gameSettings, testSettings)
}

// Run
func Run(controller gameController, gameSettings GameSettings) {
	// Setup defaults
	setup(controller, &gameSettings)
	run(gameSettings)
}

// GameEnd will close the game
func GameEnd() {
	if gState.hasGameEnded {
		return
	}
	gState.hasGameEnded = true
}

type contextUpdateLoop interface {
	Update()
}

// UpdateContext is the current context being utilized by the game
// ie. game state mode, animation editor mode, map editor mode
func contextUpdate() contextUpdateLoop {
	current := gUpdateContextStack[len(gUpdateContextStack)-1]
	return current
}

func ContextUpdatePop(currentContext contextUpdateLoop) {
	current := gUpdateContextStack[len(gUpdateContextStack)-1]
	if current != currentContext {
		panic("Can only pop context if you can provide a reference to the current context")
	}
	gUpdateContextStack = gUpdateContextStack[:len(gUpdateContextStack)-1]
}

// PushUpdateContext allows you to override the state of the game with
// a special behaviour interface
func ContextUpdatePush(context contextUpdateLoop) {
	current := gUpdateContextStack[len(gUpdateContextStack)-1]
	if current == context {
		panic("Cannot push current context again")
	}
	gUpdateContextStack = append(gUpdateContextStack, context)
}

func update() error {
	frameStartTime := monotime.Now()

	// update inputs
	keyboardUpdate()
	keyboardStringUpdate()
	mouseUpdate()

	// debugUpdate will do things like live-asset reloading
	debugUpdate()

	// run game loop, or debug animation editor or other update-loop context
	// depending on the stack
	context := contextUpdate()
	context.Update()

	/*switch debugMenuID {
	case debugMenuNone:
		// Update
		gController.GamePreUpdate()
		gState.update()
		gController.GamePostUpdate()

		// Remove deleted entities at safe point
		// ie. not while executign user-code / at the end of the frame
		gState.removeDeletedEntities()

		// NOTE: Jake: 2019-02-24
		// `cameraUpdate` should run after all update logic so that it snaps
		// to the object being followed. If a user needs custom camera behaviour,
		// they can leverage CameraSetUpdateFunction()
		cameraUpdate()
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
		panic("Invalid debug mode id")
	}*/

	gState.frameBudgetNanosecondsUsed = monotime.Now() - frameStartTime
	if gState.hasGameEnded {
		return errGameEnd
	}
	return nil
}
