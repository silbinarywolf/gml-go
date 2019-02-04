package audio

import (
	"io"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/hajimehoshi/ebiten/audio/wav"
)

type SoundIndex int32

const (
	SoundDirectoryBase = "sound"
)

// Play will play a sound
func (index SoundIndex) Play() {
	sound := &soundManager.assetList[index]
	if sound.audioPlayer == nil {
		panic("Sound isn't loaded")
	}
	sound.audioPlayer.Rewind()
	sound.audioPlayer.Play()
}

// Stop will stop playing a sound
func (index SoundIndex) Stop() {
	sound := &soundManager.assetList[index]
	if sound.audioPlayer == nil {
		panic("Sound isn't loaded")
	}
	sound.audioPlayer.Pause()
}

const (
	sampleRate = 48000
)

type soundKind int32

const (
	soundKindWAV soundKind = 1
	soundKindMP3 soundKind = 2
)

var (
	audioContext *audio.Context
)

type soundManagerData struct {
	assetList        []sound
	assetNameToIndex map[string]SoundIndex
	assetIndexToName []string
}

var (
	soundManager = &soundManagerData{}
)

type sound struct {
	audioPlayer *audio.Player
}

type soundAsset struct {
	Kind soundKind
	Data []byte
}

func loadSound(index SoundIndex) {
	name := soundManager.assetIndexToName[index]

	soundAsset := debugLoadAndWriteSoundAsset(name)
	if soundAsset == nil {
		panic("missing sound: " + name)
	}

	var soundSteam io.ReadCloser
	soundData := audio.BytesReadSeekCloser(soundAsset.Data)
	switch soundAsset.Kind {
	case soundKindWAV:
		d, err := wav.Decode(audioContext, soundData)
		if err != nil {
			panic(err)
		}
		soundSteam = d
	case soundKindMP3:
		d, err := mp3.Decode(audioContext, soundData)
		if err != nil {
			panic(err)
		}
		soundSteam = d
	default:
		panic("invalid sound type")
	}
	if soundSteam == nil {
		panic("soundStream is nil")
	}
	var err error
	audioPlayer, err := audio.NewPlayer(audioContext, soundSteam)
	if err != nil {
		panic(err)
	}
	soundManager.assetList[index] = sound{
		audioPlayer: audioPlayer,
	}
}

// InitAndLoadAllSounds is used by gmlgo when initializing the engine
func InitAndLoadAllSounds() error {
	var err error
	audioContext, err = audio.NewContext(sampleRate)
	if err != nil {
		return err
	}
	// Load all sounds
	for _, soundIndex := range soundManager.assetNameToIndex {
		loadSound(soundIndex)
	}
	return nil
}

// InitSoundGeneratedData is used by code generated by gmlgo so you can query a sound by index or name
func InitSoundGeneratedData(indexToName []string, nameToIndex map[string]SoundIndex) {
	soundManager.assetIndexToName = indexToName
	soundManager.assetNameToIndex = nameToIndex
	soundManager.assetList = make([]sound, len(soundManager.assetIndexToName))
}