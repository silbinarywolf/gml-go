package sprite

import (
	"bytes"
	"encoding/gob"
	"errors"
	"io/ioutil"
	"sync"

	_ "image/png"

	"github.com/silbinarywolf/gml-go/gml/assetman"
	"github.com/silbinarywolf/gml-go/gml/internal/file"
)

const (
	SpriteDirectoryBase = "sprite"
)

var (
	gSpriteManager = newSpriteManager()
)

func init() {
	assetman.Register(gSpriteManager)
}

type spriteManager struct {
	assetList        []Sprite
	assetNameToIndex map[string]SpriteIndex
	assetIndexToName []string
	assetIndexToPath []string
}

func newSpriteManager() *spriteManager {
	manager := &spriteManager{}
	return manager
}

// LoadAll sprite assets
func (manager *spriteManager) LoadAll() {
	/*start := timeprec.Now()
	defer func() {
		const NanosecondsPerMillisecond = 1000000
		taken := timeprec.Now() - start
		budget := float64(taken) / NanosecondsPerMillisecond
		fmt.Printf("DEBUG: Time spent loading sprites: %vns (%vms)\n", taken, budget)
	}()*/

	// NOTE(Jake): 2019-06-03
	// Split up asset loading of sprites into 4 chunks
	// this is so the WASM build will load 4 remote sprite files
	// at a time.
	// Quick perf. checks indiciate it's ~2x faster to load them all. (2.5ms vs 5ms)
	// So this logic can go away if it causes problems.
	assetList := gSpriteManager.assetIndexToPath
	var wg sync.WaitGroup
	const StartOffset = 1
	chunkCount := 4
	chunkSize := (len(assetList) + chunkCount - 1) / chunkCount
	for start := StartOffset; start < len(assetList); start += chunkSize {
		wg.Add(1)
		end := start + chunkSize

		if end > len(assetList) {
			end = len(assetList)
		}

		go func(i int, end int) {
			for i < end {
				spriteIndex := SpriteIndex(i)
				SpriteLoad(spriteIndex)
				i++
			}
			wg.Done()
		}(start, end)
	}
	wg.Wait()
}

// InitSpriteGeneratedData is used by code generated by gmlgo so you can query a sprite by index or name
func InitSpriteGeneratedData(indexToName []string, nameToIndex map[string]SpriteIndex, indexToPath []string) {
	gSpriteManager.assetIndexToName = indexToName
	gSpriteManager.assetNameToIndex = nameToIndex
	gSpriteManager.assetIndexToPath = indexToPath
	gSpriteManager.assetList = make([]Sprite, len(gSpriteManager.assetIndexToName))
}

func SpriteNames() []string {
	return gSpriteManager.assetIndexToName
}

// SpriteLoadByName is used internally by the room editor, animation editor,
// live-sprite reloading watcher and more
func SpriteLoadByName(name string) SpriteIndex {
	index, ok := gSpriteManager.assetNameToIndex[name]
	if !ok {
		return SprUndefined
	}
	SpriteLoad(index)
	return index
}

func SpriteLoad(index SpriteIndex) {
	manager := gSpriteManager
	sprite := &manager.assetList[index]
	if sprite.isLoaded {
		return
	}
	sprite.loadSprite(index)
}

func loadSpriteFromData(name string) *spriteAsset {
	path := file.AssetDirectory + "/" + SpriteDirectoryBase + "/" + name + ".data"
	fileData, err := file.OpenFile(path)
	if err != nil {
		return nil
	}
	bytesData, err := ioutil.ReadAll(fileData)
	fileData.Close()
	if err != nil {
		panic(errors.New("Unable to read bytes from image: " + path))
	}
	buf := bytes.NewReader(bytesData)
	asset := new(spriteAsset)
	gob.NewDecoder(buf).Decode(asset)
	return asset
}

func (sprite *Sprite) loadSprite(index SpriteIndex) {
	name := gSpriteManager.assetIndexToName[index]
	path := gSpriteManager.assetIndexToPath[index]

	// If debug mode, write out the sprite to *.data file
	debugWriteSprite(name, path)

	// Load from *.data
	spriteAsset := loadSpriteFromData(name)
	if spriteAsset == nil {
		panic("Unable to load sprite from data file: " + name)
	}
	frames := make([]SpriteFrame, len(spriteAsset.Frames))
	for i := 0; i < len(spriteAsset.Frames); i++ {
		frameAsset := spriteAsset.Frames[i]
		frame, err := createFrame(frameAsset)
		if err != nil {
			panic("Sprite frame load error for \"" + name + "\": " + err.Error())
		}
		frames[i] = frame
	}

	// Create sprite
	*sprite = newSprite(spriteAsset.Name, frames, spriteConfig{
		ImageSpeed: spriteAsset.ImageSpeed,
	})
}
