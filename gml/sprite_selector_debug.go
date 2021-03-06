// +build debug

package gml

import (
	"image/color"
	"os"
	"path/filepath"

	"github.com/silbinarywolf/gml-go/gml/internal/file"
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

var (
	debugScratchSpriteList  []sprite.SpriteIndex
	debugSpriteViewerLoaded bool
)

type debugSpriteViewer struct {
}

func (viewer *debugSpriteViewer) lazyLoad() {
	if debugSpriteViewerLoaded {
		return
	}
	debugSpriteViewerLoaded = true
	spritePath := file.AssetDirectory + "/" + sprite.SpriteDirectoryBase
	err := filepath.Walk(spritePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// Skip files
			return nil
		}
		if path == spritePath {
			// Skip self
			return nil
		}
		name := filepath.Base(path)
		sprite.SpriteLoadByName(name)
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func (viewer *debugSpriteViewer) update() (sprite.SpriteIndex, bool) {
	viewer.lazyLoad()

	typingText := KeyboardString()
	spriteMenuFiltered := debugScratchSpriteList[:0]
	for spriteIndex, name := range sprite.SpriteNames() {
		hasMatch := hasFilterMatch(name, typingText)
		if !hasMatch {
			continue
		}
		spriteMenuFiltered = append(spriteMenuFiltered, sprite.SpriteIndex(spriteIndex))
	}

	// Input
	selected := KeyboardCheckPressed(VkEnter) ||
		KeyboardCheckPressed(VkNumpadEnter)
	if selected &&
		len(spriteMenuFiltered) > 0 {
		selectedSpr := spriteMenuFiltered[0]
		ClearKeyboardString()
		return selectedSpr, true
	}

	// Draw
	{
		DrawSetGUI(true)
		// Add black opacity over screen with menu open
		DrawRectangle(0, 0, 2048, 2048, color.RGBA{0, 0, 0, 190})

		ui := geom.Vec{
			X: float64(WindowSize().X) / 2,
			Y: 32,
		}

		{
			searchText := "Search for image (type + press enter)"
			DrawText(ui.X-(StringWidth(searchText)/4), ui.Y, searchText, color.White)
			ui.Y += 24
		}
		{
			typingText := KeyboardString()
			DrawText(ui.X, ui.Y, typingText, color.White)
			DrawText(ui.X+StringWidth(typingText), ui.Y, "|", color.White)
			ui.Y += 24
		}
		previewSize := geom.Vec{32, 32}
		for _, spr := range spriteMenuFiltered {
			if spr == 0 {
				continue
			}
			var pos geom.Vec
			pos.X = ui.X - 40
			pos.Y = ui.Y - (previewSize.Y / 2)
			calcPreviewSize := previewSize
			calcPreviewSize.X /= float64(spr.Size().X)
			calcPreviewSize.Y /= float64(spr.Size().Y)
			DrawSpriteScaled(spr, 0, pos.X, pos.Y, calcPreviewSize)
			name := spr.Name()
			DrawText(ui.X, ui.Y, name, color.White)
			ui.Y += previewSize.Y + 16
		}
	}
	return sprite.SprUndefined, false
}
