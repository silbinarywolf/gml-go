package audio

import (
	"io"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/hajimehoshi/ebiten/audio/wav"
	"github.com/silbinarywolf/gml-go/gml/assetman"
)

type SoundIndex int32

const (
	sndUndefined       SoundIndex = 0
	SoundDirectoryBase            = "sound"
)

// Play will play a sound
func (index SoundIndex) Play() {
	if disableAudio {
		return
	}
	if index == sndUndefined {
		panic("Cannot play sound if not set")
	}
	sound := &soundManager.assetList[index]
	if sound.audioPlayer == nil {
		panic("Sound isn't loaded")
	}
	sound.audioPlayer.Rewind()
	sound.audioPlayer.Play()
}

// Stop will stop playing a sound
func (index SoundIndex) Stop() {
	if disableAudio {
		return
	}
	if index == sndUndefined {
		panic("Cannot stop sound if not set")
	}
	sound := &soundManager.assetList[index]
	if sound.audioPlayer == nil {
		panic("Sound isn't loaded")
	}
	sound.audioPlayer.Pause()
}

// IsPlaying will check if a sound is playing
func (index SoundIndex) IsPlaying() bool {
	if disableAudio {
		return true
	}
	if index == sndUndefined {
		panic("Cannot check if unset sound is playing")
	}
	sound := &soundManager.assetList[index]
	return sound.audioPlayer.IsPlaying()
}

// Name of the sound asset
func (index SoundIndex) Name() string {
	sound := &soundManager.assetList[index]
	return sound.name
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
	soundManager = new(soundManagerData)
)

func init() {
	assetman.Register(soundManager)
}

func (manager *soundManagerData) LoadAll() {
	if audioContext != nil {
		// Don't reinitialize. This can be called twice if multiple tests
		// call TestBootstrap()
		return
	}
	assetList := soundManager.assetNameToIndex
	if len(assetList) == 0 {
		// Don't initialize audio if there are no audio assets
		return
	}
	var err error
	audioContext, err = audio.NewContext(sampleRate)
	if err != nil {
		panic(err)
	}
	// Load all sounds
	for _, soundIndex := range assetList {
		loadSound(soundIndex)
	}
}

func (manager *soundManagerData) ManifestJSON() (string, map[string]string) {
	result := make(map[string]string)
	assetList := manager.assetNameToIndex
	for _, soundIndex := range assetList {
		name := soundIndex.Name()
		result[name] = name
	}
	return "audio", result
}

type sound struct {
	audioPlayer *audio.Player
	name        string
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
		name:        name,
	}
}

// InitSoundGeneratedData is used by code generated by gmlgo so you can query a sound by index or name
func InitSoundGeneratedData(indexToName []string, nameToIndex map[string]SoundIndex) {
	soundManager.assetIndexToName = indexToName
	soundManager.assetNameToIndex = nameToIndex
	soundManager.assetList = make([]sound, len(soundManager.assetIndexToName))
}
