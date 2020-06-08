package gml

import (
	"errors"
	"path/filepath"
	"runtime"

	"github.com/silbinarywolf/gml-go/gml/assetman"
	"github.com/silbinarywolf/gml-go/gml/internal/dt"
	"github.com/silbinarywolf/gml-go/gml/internal/file"
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
	"github.com/silbinarywolf/gml-go/gml/internal/paniclog"
	"github.com/silbinarywolf/gml-go/gml/monotime"
)

type DefaultContext struct{}

func (context *DefaultContext) Open() {
}

func (context *DefaultContext) Close() {
}

func (context *DefaultContext) Update() {
	// Update
	gController.GamePreUpdate()
	gState.update()
	gController.GamePostUpdate()

	// Remove deleted entities at safe point
	// ie. not while executing user-code / at the end of the frame
	gState.removeDeletedEntities()

	// NOTE: Jake: 2019-02-24
	// `cameraUpdate` should run after all update logic so that it snaps
	// to the object being followed. If a user needs custom camera behaviour,
	// they can leverage CameraSetUpdateFunction()
	cameraUpdate()
}

func (context *DefaultContext) Draw() {
	// NOTE(Jake): 2020-06-08
	// Default to first camera for level editors / animation editor
	// etc.
	cameraSetActive(0)

	// PreDraw
	gController.GamePreDraw()

	// Draw
	for i := 0; i < len(gCameraManager.cameras); i++ {
		view := &gCameraManager.cameras[i]
		if !view.enabled {
			continue
		}
		cameraSetActive(i)
		cameraMaybeAllocSurface(i)
		cameraClearSurface(i)

		if inst := view.follow.getBaseObject(); inst != nil {
			// Render instances in same room as instance following
			roomInst := roomGetInstance(inst.RoomInstanceIndex())
			if roomInst == nil {
				panic("draw: RoomInstance this object belongs to has been destroyed")
			}
			roomInst.draw()
		} else {
			// If no follower is configured, just try render the first active room found
			roomInst := roomLastCreated()
			if roomInst != nil {
				// NOTE(Jae): 2020-03-29
				// We used to throw an error if no room existed but doing that
				// isn't great when you're not in a room, like a title screen
				roomInst.draw()
			}
		}

		// Render camera onto OS-window
		cameraDraw(i)
	}

	// NOTE(Jake): 2019-04-15
	// Default to first camera for level editors / animation editor
	// etc.
	cameraSetActive(0)

	// PostDraw
	gController.GamePostDraw()

	// Clear the camera so that calls to CameraGetActive() will cause
	// a panic. (ie. not allowed!)
	cameraClearActive()
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

// testInitAssetDir is designed to be called from a gmlgo generated method inside the game projects "asset"
// directory, this allows us to determine where asset directory exists when running in circumstances
// where the executables location is nowhere near the assets. (ie. "go test")
func testInitAssetDir() {
	if !file.IsAssetDirSet() {
		_, filename, _, _ := runtime.Caller(2)
		dir := filepath.Clean(filepath.Dir(filename))
		// NOTE(Jae): 2020-05-30
		// In release builds, this logic would fail because the folder
		// being referenced would be where the code was built on the machine
		// that compiled it.
		// We might want to make this logic still work for debug mode or something
		// so we can catch dev. mistakes if/when this logic needs to change.
		//if _, err := os.Stat(dir); os.IsNotExist(err) {
		//	panic("Cannot find asset directory: " + dir + ", integration tests must go inside \"{project}/test\"")
		//}
		file.SetAssetDir(dir)
	}
}

// TestBootstrap the game to give control over continuing / stopping execution per-frame
// this method is for additional control when testing
func TestBootstrap(controller gameController, gameSettings GameSettings, testSettings TestSettings) {
	setup(controller, &gameSettings)
	runTest(gameSettings, testSettings)
}

// Run
func Run(controller gameController, gameSettings GameSettings) {
	// We only want panic/crashes to be redirected to log files for release
	// builds. By placing the logic here, other commands like "go test" wont end
	// up logging relative to the current working directory or similar.
	//
	// A small concern is that with this approach, if a crash occurs during bootup
	// it wont be caught. (ie. in a Golang init() function) but I think I can live
	// with that.
	paniclog.Init()

	// NOTE(Jae): 2020-05-30
	// We setup the asset directory here. For tests this program would've already
	// executed "testInitAssetDir()" to determine the asset directory
	// to support other execution modes like "go test".
	// However, that folder location would be incorrect for release builds
	// and we want it to be relative to the executables location.
	file.InitAssetDir()

	setup(controller, &gameSettings)
	run(gameSettings)
}

// GameEnd will close the game
func GameEnd() {
	gGameGlobals.hasGameEnded = true
}

type contextUpdateLoop interface {
	Open()
	Close()
	Update()
	Draw()
}

type contextUpdateLoopItem struct {
	windowCursorVisible bool
	windowSize          geom.Vec
	WindowScale         geom.Vec
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
	currentContext.Close()
	gUpdateContextStack = gUpdateContextStack[:len(gUpdateContextStack)-1]
}

// PushUpdateContext allows you to override the state of the game with
// a special behaviour interface
func ContextUpdatePush(context contextUpdateLoop) {
	current := gUpdateContextStack[len(gUpdateContextStack)-1]
	if current == context {
		panic("Cannot push current context again")
	}
	context.Open()
	gUpdateContextStack = append(gUpdateContextStack, context)
}

func updateFrameTimerAndTickCount(frameStartTime int64) {
	timeTaken := monotime.Now() - frameStartTime
	if gGameGlobals.frameUpdateBudgetNanosecondsUsed == 0 {
		gGameGlobals.frameUpdateBudgetNanosecondsUsed = timeTaken
	} else {
		// Exponentially smoothed moving average
		gGameGlobals.frameUpdateBudgetNanosecondsUsed = int64(float64(gGameGlobals.frameUpdateBudgetNanosecondsUsed) + (0.10 * float64(timeTaken-gGameGlobals.frameUpdateBudgetNanosecondsUsed)))
	}
	gGameGlobals.tickCount++
}

func update() error {

	// update inputs
	keyboardUpdate()
	gamepadUpdate()
	keyboardStringUpdate()
	mouseUpdate()

	// debugUpdate will do things like live-asset reloading
	debugUpdate()

	// run game loop, or debug animation editor or other update-loop context
	// depending on the stack
	context := contextUpdate()
	context.Update()

	if gGameGlobals.hasGameEnded {
		return errGameEnd
	}
	return nil
}
