package game

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/silbinarywolf/gml-go/example/worm/asset"
	"github.com/silbinarywolf/gml-go/gml"
	"github.com/silbinarywolf/gml-go/gml/audio"
)

var Global = new(GameController)

type GameController struct {
	gml.Controller
	PersistentGameData
	PreviousRound GameScore
	CurrentRound  GameScore

	Notification Notification
	Player       gml.InstanceIndex
	MusicPlaying audio.SoundIndex
	Score        int
}

type Medal int

const (
	MedalNone   Medal = 0
	MedalBronze Medal = 1
	MedalSilver Medal = 2
	MedalGold   Medal = 3
)

type GameScore struct {
	MedalWorm Medal
	MedalWing Medal
}

type PersistentGameData struct {
	// todo(Jake): 2019-03-13
	// Maybe reimplement saving for the options / high score system
	SoundDisabled bool
	MusicDisabled bool
}

func (*GameController) HasWormStopped() bool {
	if inst, ok := Global.Player.Get().(*Worm); ok {
		if inst.Dead {
			return true
		}
	}
	return false
}

func (*GameController) GameStart() {
	gml.DrawSetFont(asset.FntDefault)

	// Setup "kinda" delta time
	gml.SetDesignedTPS(DesignedMaxTPS)
	//gml.SetMaxTPS(480)

	// Play song
	if !Global.MusicDisabled {
		Global.MusicPlaying = asset.MusSunnyFields
		Global.MusicPlaying.Play()
	}

	// Setup global variables
	// ...

	// Create new empty room
	roomInstanceIndex := gml.RoomInstanceNew()

	// Create background drawer
	roomInstanceIndex.InstanceCreate(0, 0, ObjBackground)

	// Create menu
	roomInstanceIndex.InstanceCreate(0, 0, ObjMenu)
	//gml.InstanceCreate(0, 0, roomInstanceIndex, ObjMenuGameover)

	// Create player in the center of the room
	playerInst := roomInstanceIndex.InstanceCreate(0, 0, ObjWorm).(*Worm)
	Global.Player = playerInst.InstanceIndex()

	//Global.Notification.SetNotification("You got a wing\n\nEach wing will add an extra jump")
}

func (*GameController) MusicRandomizeTrack() {
	if Global.MusicPlaying != 0 {
		Global.MusicPlaying.Stop()
	}
	if Global.MusicDisabled {
		return
	}

	if Global.MusicPlaying != 0 &&
		Global.MusicPlaying != asset.MusClassicTrack {
		if rand.Int63n(1000) == 1 {
			Global.MusicPlaying = asset.MusClassicTrack
			Global.MusicPlaying.Play()
			return
		}
	}

	// NOTE: Jake: 2019-02-13
	// This doesn't technically cycle through the tracks between games
	// as I expected but this is how the original code worked, so I'm
	// leaving it as is.
	switch Global.MusicPlaying {
	case asset.MusSunnyFields:
		Global.MusicPlaying = asset.MusRacer
	default:
		Global.MusicPlaying = asset.MusSunnyFields
	}
	Global.MusicPlaying.Play()
}

func (*GameController) GamePreUpdate() {
	if Global.MusicPlaying != 0 && !Global.MusicPlaying.IsPlaying() {
		Global.MusicRandomizeTrack()
	}
	Global.Notification.Update()
}

func (*GameController) GameReset() {
	inst, ok := Global.Player.Get().(*Worm)
	if !ok {
		panic("Cannot find Player object to call GameReset")
	}
	inst.Reset()

	// Reset game music if game over
	if Global.MusicPlaying == asset.MusGameover {
		Global.MusicPlaying.Stop()
		Global.MusicPlaying = asset.MusSunnyFields
		Global.MusicPlaying.Play()
	}

	// Make walls from previous playthrough become disabled
	screenSize := gml.CameraGetViewSize(0)
	for _, id := range inst.RoomIndex().WithAll() {
		inst := id.Get()
		switch inst := inst.(type) {
		case *Wall:
			if inst.X+inst.Size.X > screenSize.X {
				// Destroy walls that were spawned off-screen
				gml.InstanceDestroy(inst)
			}
			inst.DontKillPlayer = true
		}
	}

	if inst.Dead {
		inst.Vec = inst.Start
		inst.Y = -140
		inst.Speed.Y = 0
		inst.SetSprite(asset.SprWormHead)
		inst.Dead = false
	}
}

func (*GameController) GamePostDraw() {
	// Draw frame usage
	gml.DrawText(8, 8, gml.DebugFrameUsage(), color.White)

	// Draw score
	if playerInst, ok := Global.Player.Get().(*Worm); ok {
		var scoreIndexes [8]float64

		// Split score into seperate numbers
		i := 0
		score := playerInst.Score
		fontWidth := asset.SprScoreFont.Size().X
		textWidth := 0.0
		for score >= 1 {
			index := math.Mod(score, 10)
			score = math.Floor(score / 10)
			scoreIndexes[i] = index
			i++
			textWidth += fontWidth
		}

		// Draw numbers in correct order
		x := (gml.CameraGetViewSize(0).X / 2) - (textWidth / 2)
		y := 32.0
		for i > 0 {
			i--
			index := scoreIndexes[i]
			gml.DrawSpriteColor(asset.SprScoreFont, index, x-1, y, color.Black)
			gml.DrawSprite(asset.SprScoreFont, index, x, y+1)
			x += fontWidth
		}
	}

	// Draw notification
	Global.Notification.Draw()
}
