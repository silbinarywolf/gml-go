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
	debugScratchSpriteList  []*sprite.Sprite
	debugSpriteViewerLoaded bool
)

type debugSpriteViewer struct {
}

func (viewer *debugSpriteViewer) lazyLoad() {
	if debugSpriteViewerLoaded {
		return
	}
	debugSpriteViewerLoaded = true
	spritePath := file.AssetsDirectory + "/sprites"
	err := filepath.Walk(spritePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			println("prevent panic by handling failure accessing a path " + path + ": " + err.Error())
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
		sprite.LoadSprite(name)
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func (viewer *debugSpriteViewer) Update() (*sprite.Sprite, bool) {
	viewer.lazyLoad()

	typingText := KeyboardString()
	spriteMenuFiltered := debugScratchSpriteList[:0]
	for _, spr := range sprite.SpriteList() {
		hasMatch := hasFilterMatch(spr.Name(), typingText)
		if !hasMatch {
			continue
		}
		spriteMenuFiltered = append(spriteMenuFiltered, spr)
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
		DrawRectangle(geom.Vec{0, 0}, geom.Vec{2048, 2048}, color.RGBA{0, 0, 0, 190})

		ui := geom.Vec{
			X: float64(windowWidth()) / 2,
			Y: 32,
		}

		{
			searchText := "Search for image (type + press enter)"
			DrawText(geom.Vec{ui.X - (StringWidth(searchText) / 4), ui.Y}, searchText)
			ui.Y += 24
		}
		{
			typingText := KeyboardString()
			DrawText(geom.Vec{ui.X, ui.Y}, typingText)
			DrawText(geom.Vec{ui.X + StringWidth(typingText), ui.Y}, "|")
			ui.Y += 24
		}
		previewSize := geom.Vec{32, 32}
		for _, spr := range spriteMenuFiltered {
			var pos geom.Vec
			pos.X = ui.X - 40
			pos.Y = ui.Y - (previewSize.Y / 2)
			calcPreviewSize := previewSize
			calcPreviewSize.X /= float64(spr.Size().X)
			calcPreviewSize.Y /= float64(spr.Size().Y)
			DrawSpriteScaled(spr, 0, pos, calcPreviewSize)
			name := spr.Name()
			DrawText(geom.Vec{ui.X, ui.Y}, name)
			ui.Y += previewSize.Y + 16
		}
	}
	return nil, false
}
