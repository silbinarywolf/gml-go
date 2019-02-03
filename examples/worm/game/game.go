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
	//gml.SetMaxTPS(80)

	// Setup global variables
	// ...

	// Create new empty room
	roomInstanceIndex := gml.RoomInstanceNew()

	// Create background drawer
	gml.InstanceCreate(0, 0, roomInstanceIndex, ObjBackground)

	// Create menu
	gml.InstanceCreate(0, 0, roomInstanceIndex, ObjMenu)

	// Create player in the center of the room
	playerInst := gml.InstanceCreate(0, 0, roomInstanceIndex, ObjWorm).(*Worm)
	Global.Player = playerInst.InstanceIndex()

	// Play song
	Global.MusicPlaying = SndSunnyFields
	Global.MusicPlaying.Play()
}

//func (*GameController) GameRestart() {
//	Global.MusicPlaying.Stop()
//}

func (*GameController) GamePostDraw() {
	// Draw frame usage
	gml.DrawTextF(32, 32, "%s", gml.FrameUsage())

	// Draw score
	if playerInst, ok := Global.Player.Get().(*Worm); ok {
		var scoreIndexes [8]float64

		// Split score into seperate numbers
		i := 0
		score := playerInst.Score
		fontWidth := gml.SpriteSize(SprScoreFont).X
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
