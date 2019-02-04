package game

import (
	"image/color"
	"math"

	"github.com/silbinarywolf/gml-go/gml"
	"github.com/silbinarywolf/gml-go/gml/audio"
)

var Global = new(GameController)

type GameController struct {
	gml.Controller
	Player        gml.InstanceIndex
	MusicPlaying  audio.SoundIndex
	Score         int
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
	gml.DrawSetFont(FntDefault)

	// Setup "kinda" delta time
	gml.SetDesignedTPS(DesignedMaxTPS)
	//gml.SetMaxTPS(120)

	// Setup global variables
	// ...

	// Create new empty room
	roomInstanceIndex := gml.RoomInstanceNew()

	// Create background drawer
	gml.InstanceCreate(0, 0, roomInstanceIndex, ObjBackground)

	// Create menu
	gml.InstanceCreate(0, 0, roomInstanceIndex, ObjMenu)
	//gml.InstanceCreate(0, 0, roomInstanceIndex, ObjMenuGameover)

	// Create player in the center of the room
	playerInst := gml.InstanceCreate(0, 0, roomInstanceIndex, ObjWorm).(*Worm)
	Global.Player = playerInst.InstanceIndex()

	// Play song
	Global.MusicPlaying = SndSunnyFields
	Global.MusicPlaying.Play()
}

func (*GameController) GameReset() {
	inst, ok := Global.Player.Get().(*Worm)
	if !ok {
		panic("Cannot find Player object to call GameReset")
	}
	inst.WallSpawner.Reset()

	// Reset game music if game over
	if Global.MusicPlaying == SndGameover {
		Global.MusicPlaying.Stop()
		Global.MusicPlaying = SndSunnyFields
		Global.MusicPlaying.Play()
	}

	// Make walls from previous playthrough become disabled
	screenSize := gml.CameraGetViewSize(0)
	for _, id := range gml.WithAll(inst) {
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
		inst.SetSprite(SprWormHead)
		inst.Dead = false
	}
}

func (*GameController) GamePostDraw() {
	// Draw frame usage
	gml.DrawTextF(32, 32, "%s", gml.FrameUsage())

	// Draw score
	if playerInst, ok := Global.Player.Get().(*Worm); ok {
		var scoreIndexes [8]float64

		// Split score into seperate numbers
		i := 0
		score := playerInst.Score
		fontWidth := SprScoreFont.Size().X
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
			gml.DrawSpriteColor(SprScoreFont, index, x-1, y, color.Black)
			gml.DrawSprite(SprScoreFont, index, x, y+1)
			x += fontWidth
		}
	}
}
