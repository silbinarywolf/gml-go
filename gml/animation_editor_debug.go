// +build debug

package gml

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"math"
	"strconv"

	"github.com/silbinarywolf/gml-go/gml/internal/file"
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
	"github.com/silbinarywolf/gml-go/gml/internal/sprite"
)

type animMenu int

const (
	animMenuNone animMenu = 0 + iota
	animMenuSprite
	animMenuSpriteBboxLeft
	animMenuSpriteBboxTop
	animMenuSpriteBboxRight
	animMenuSpriteBboxBottom
)

const (
	handleDragNone int = 0 + iota
	handleDragLeftTop
	handleDragRightTop
	handleDragRightBottom
	handleDragLeftBottom
)

var (
	animationEditor *debugAnimationEditor
)

type debugAnimationEditor struct {
	debugSpriteViewer
	menuOpened         animMenu
	handleDragging     int
	handleDragBeginPos geom.Vec
	spriteViewing      SpriteState
	isInPlayback       bool
}

type animationEditorConfig struct {
	SpriteSelected string `json:"SpriteSelected,omitempty"`
}

func (editor *debugAnimationEditor) LazyLoad() {
	if animationEditor != nil {
		return
	}
	animationEditor = new(debugAnimationEditor)
	animationEditor.animationConfigLoad()
	if animationEditor.spriteViewing.SpriteIndex() == 0 {
		animationEditor.spriteViewing.SetSprite(1)
	}
}

func (editor *debugAnimationEditor) animationConfigLoad() {
	configPath := debugConfigPath("animation_editor")
	fileData, err := file.OpenFile(configPath)
	if err == nil {
		bytes, err := ioutil.ReadAll(fileData)
		if err != nil {
			panic("Error loading " + configPath + "\n" + "Error: " + err.Error())
		}
		editorConfig := animationEditorConfig{}
		if err := json.Unmarshal(bytes, &editorConfig); err != nil {
			panic("Error unmarshalling " + configPath + "\n" + "Error: " + err.Error())
		}
		name := editorConfig.SpriteSelected
		editor.spriteViewing = SpriteState{}
		// todo(Jake): 2018-10-28
		// Add function to load a sprite if it exists, we don't want to crash
		// if we remove a sprite that we previously had loaded.
		spr := sprite.SpriteLoadByName(name)
		editor.spriteViewing.SetSprite(spr)
	}
}

func (editor *debugAnimationEditor) animationConfigSave() {
	editorConfig := animationEditorConfig{}
	editorConfig.SpriteSelected = editor.spriteViewing.SpriteIndex().Name()
	json, _ := json.MarshalIndent(editorConfig, "", "\t")
	configPath := debugConfigPath("animation_editor")
	err := ioutil.WriteFile(configPath, json, 0644)
	if err != nil {
		println("Failed to write animation editor config: " + configPath + "\n" + "Error: " + err.Error())
	}
}

func (editor *debugAnimationEditor) animationEditorToggleMenu(menu animMenu) {
	if editor.menuOpened == menu {
		menu = animMenuNone
	}
	spriteIndex := editor.spriteViewing.SpriteIndex()
	imageIndex := int(math.Floor(editor.spriteViewing.ImageIndex()))
	collisionMask := sprite.GetCollisionMask(spriteIndex, imageIndex, 0)
	value, err := strconv.ParseFloat(KeyboardString(), 64)
	if err == nil {
		switch editor.menuOpened {
		case animMenuSpriteBboxLeft:
			diff := value - collisionMask.Rect.X
			collisionMask.Rect.X = value
			collisionMask.Rect.Size.X -= diff
		case animMenuSpriteBboxTop:
			diff := value - collisionMask.Rect.Y
			collisionMask.Rect.Y = value
			collisionMask.Rect.Size.Y -= diff
		case animMenuSpriteBboxRight:
			collisionMask.Rect.Size.X = value - collisionMask.Rect.X
		case animMenuSpriteBboxBottom:
			collisionMask.Rect.Size.Y = value - collisionMask.Rect.Y
		}
	}
	editor.menuOpened = menu
}

func (editor *debugAnimationEditor) Open() {
	editor.LazyLoad()
}

func (editor *debugAnimationEditor) Close() {
}

func (editor *debugAnimationEditor) Update() {
}

func (editor *debugAnimationEditor) Draw() {
	cameraSetActive(0)
	cameraClearSurface(0)
	defer func() {
		cameraDraw(0)
		cameraClearActive()
	}()

	DrawSetGUI(true)

	//
	{
		pos := geom.Vec{
			X: 16,
			Y: 16,
		}
		DrawText(pos.X, pos.Y, "Animation Editor", color.White)
		pos.Y += 24
		DrawText(pos.X, pos.Y, "Space = Play/Pause Animation", color.White)
		pos.Y += 24
		DrawText(pos.X, pos.Y, "CTRL + P = Open Sprite List", color.White)

		if spriteIndex := editor.spriteViewing.SpriteIndex(); spriteIndex != sprite.SprUndefined {
			pos.Y += 24
			DrawText(pos.X, pos.Y, "CTRL + S = Save", color.White)

			if KeyboardCheck(VkControl) && KeyboardCheckPressed(VkS) {
				err := sprite.DebugWriteSpriteConfig(spriteIndex)
				if err != nil {
					panic(err)
				}
			}
		}
	}

	// Shortcut keys
	if KeyboardCheck(VkControl) {
		if KeyboardCheckPressed(VkP) {
			editor.animationEditorToggleMenu(animMenuSprite)
		}
	}
	if KeyboardCheckPressed(VkSpace) {
		editor.isInPlayback = !editor.isInPlayback
	}

	// Change frame viewing
	if spr := editor.spriteViewing.SpriteIndex(); spr != 0 {
		imageIndex := math.Floor(editor.spriteViewing.ImageIndex())
		if KeyboardCheckPressed(VkLeft) {
			imageIndex -= 1
			if imageIndex < 0 {
				imageIndex = editor.spriteViewing.ImageNumber() - 1
			}
			editor.spriteViewing.SetImageIndex(imageIndex)
		}
		if KeyboardCheckPressed(VkRight) {
			imageIndex += 1
			if imageIndex > editor.spriteViewing.ImageNumber() {
				imageIndex = 0
			}
			editor.spriteViewing.SetImageIndex(imageIndex)
		}
	}

	//
	var collisionMask *sprite.CollisionMask
	var inheritCollisionMask *sprite.CollisionMask
	if spriteIndex := editor.spriteViewing.SpriteIndex(); spriteIndex != 0 {
		imageIndex := int(math.Floor(editor.spriteViewing.ImageIndex()))
		collisionMask = sprite.GetCollisionMask(spriteIndex, imageIndex, 0)
		switch collisionMask.Kind {
		case sprite.CollisionMaskInherit:
			for ; imageIndex > 0; imageIndex-- {
				collisionMask = sprite.GetCollisionMask(spriteIndex, imageIndex, 0)
				if collisionMask.Kind != sprite.CollisionMaskInherit {
					break
				}
			}
			if imageIndex == 0 {
				collisionMask = sprite.GetCollisionMask(spriteIndex, imageIndex, 0)
				if collisionMask.Kind == sprite.CollisionMaskInherit {
					collisionMask = &sprite.CollisionMask{
						Kind: sprite.CollisionMaskManual,
						Rect: geom.Rect{
							Size: spriteIndex.Size(),
						},
					}
				}
			}
			inheritCollisionMask = collisionMask
		case sprite.CollisionMaskManual:
			//
		}
	}

	if spriteIndex := editor.spriteViewing.SpriteIndex(); spriteIndex != 0 {
		size := spriteIndex.Size()
		pos := geom.Vec{
			X: float64(WindowSize().X/2) - (float64(size.X) / 2),
			Y: float64(WindowSize().Y/2) - (float64(size.Y) / 2),
		}

		{
			// Draw backdrop
			pos := pos
			DrawRectangle(pos.X, pos.Y, size.X, size.Y, color.RGBA{195, 195, 195, 255})
		}

		// Sprite
		if editor.isInPlayback {
			editor.spriteViewing.ImageUpdate()
		}
		DrawSprite(spriteIndex, editor.spriteViewing.ImageIndex(), pos.X, pos.Y)

		if collisionMask != nil {
			// Draw collision box
			var rect geom.Rect = collisionMask.Rect
			rect.X += pos.X
			rect.Y += pos.Y
			DrawRectangle(rect.X, rect.Y, rect.Size.X, rect.Size.Y, color.RGBA{255, 0, 0, 128})
		}

		if collisionMask != nil &&
			inheritCollisionMask == nil {
			// Draw resize handles
			offset := pos

			// Get distance mouse moved
			var diffX, diffY float64
			{
				mousePos := MouseScreenPosition()
				handleBeginPos := editor.handleDragBeginPos
				diffX = mousePos.X - handleBeginPos.X
				diffY = mousePos.Y - handleBeginPos.Y
				if KeyboardCheck(VkControl) {
					diffX = math.Round(diffX / 4)
					diffY = math.Round(diffY / 4)
				}
			}

			{
				// Top-Left
				rect := geom.Rect{}
				rect.Size = geom.Vec{12, 12}
				rect.X = offset.X + collisionMask.Rect.Left() - float64(rect.Size.X/2)
				rect.Y = offset.Y + collisionMask.Rect.Top() - float64(rect.Size.Y/2)

				// Handle hitbox handles
				if editor.handleDragging == handleDragLeftTop {
					collisionMask.Rect.X += diffX
					collisionMask.Rect.Size.X -= diffX
					collisionMask.Rect.Y += diffY
					collisionMask.Rect.Size.Y -= diffY
				}
				col := color.RGBA{255, 255, 255, 255}
				if debugDrawIsMouseOver(rect.Pos(), rect.Size) {
					if MouseCheckPressed(MbLeft) {
						editor.handleDragging = handleDragLeftTop
					}
					col = color.RGBA{200, 200, 200, 255}
				}
				DrawRectangle(rect.X, rect.Y, rect.Size.X, rect.Size.Y, col)
			}
			{
				// Top-Right
				rect := geom.Rect{}
				rect.Size = geom.Vec{12, 12}
				rect.X = offset.X + collisionMask.Rect.Right() - float64(rect.Size.X/2)
				rect.Y = offset.Y + collisionMask.Rect.Top() - float64(rect.Size.Y/2)

				// Handle hitbox handles
				if editor.handleDragging == handleDragRightTop {
					collisionMask.Rect.Size.X += diffX
					collisionMask.Rect.Y += diffY
					collisionMask.Rect.Size.Y -= diffY
				}
				col := color.RGBA{255, 255, 255, 255}
				if debugDrawIsMouseOver(rect.Pos(), rect.Size) {
					if MouseCheckPressed(MbLeft) {
						editor.handleDragging = handleDragRightTop
					}
					col = color.RGBA{200, 200, 200, 255}
				}
				DrawRectangle(rect.X, rect.Y, rect.Size.X, rect.Size.Y, col)
			}
			{
				// Bottom-Left
				rect := geom.Rect{}
				rect.Size = geom.Vec{12, 12}
				rect.X = offset.X + collisionMask.Rect.Left() - float64(rect.Size.X/2)
				rect.Y = offset.Y + collisionMask.Rect.Bottom() - float64(rect.Size.Y/2)

				// Handle hitbox handles
				if editor.handleDragging == handleDragLeftBottom {
					collisionMask.Rect.X += diffX
					collisionMask.Rect.Size.X -= diffX
					//collisionMask.Rect.Y = diffY
					collisionMask.Rect.Size.Y += diffY
				}
				col := color.RGBA{255, 255, 255, 255}
				if debugDrawIsMouseOver(rect.Pos(), rect.Size) {
					if MouseCheckPressed(MbLeft) {
						editor.handleDragging = handleDragLeftBottom
					}
					col = color.RGBA{200, 200, 200, 255}
				}
				DrawRectangle(rect.X, rect.Y, rect.Size.X, rect.Size.Y, col)
			}
			{
				// Bottom-Right
				rect := geom.Rect{}
				rect.Size = geom.Vec{12, 12}
				rect.X = offset.X + collisionMask.Rect.Right() - float64(rect.Size.X/2)
				rect.Y = offset.Y + collisionMask.Rect.Bottom() - float64(rect.Size.Y/2)

				// Handle hitbox handles
				if editor.handleDragging == handleDragRightBottom {
					collisionMask.Rect.Size.X += diffX
					collisionMask.Rect.Size.Y += diffY
				}
				col := color.RGBA{255, 255, 255, 255}
				if debugDrawIsMouseOver(rect.Pos(), rect.Size) {
					if MouseCheckPressed(MbLeft) {
						editor.handleDragging = handleDragRightBottom
					}
					col = color.RGBA{200, 200, 200, 255}
				}
				DrawRectangle(rect.X, rect.Y, rect.Size.X, rect.Size.Y, col)
			}
			{
				// Update State
				editor.handleDragBeginPos = MouseScreenPosition()
				if !MouseCheckButton(MbLeft) {
					editor.handleDragging = handleDragNone
				}
			}
		}
	}

	if editor.menuOpened != animMenuNone {
		switch editor.menuOpened {
		case animMenuSprite:
			if selectedSpr, ok := animationEditor.debugSpriteViewer.update(); ok {
				editor.spriteViewing = SpriteState{}
				editor.spriteViewing.SetSprite(selectedSpr)
				editor.menuOpened = animMenuNone
				editor.animationConfigSave()
			}
		}
	}

	if spriteIndex := editor.spriteViewing.SpriteIndex(); spriteIndex != 0 {
		basePos := geom.Vec{(float64(WindowSize().X) / 2) - 140, float64(WindowSize().Y)}
		basePos.Y -= 210

		imageIndex := int(math.Floor(editor.spriteViewing.ImageIndex()))
		DrawText(basePos.X, basePos.Y, fmt.Sprintf("Frame: %d", imageIndex), color.White)
		basePos.Y += 24
		if drawButton(basePos, "Kind: Inherit") {
			collisionMask = sprite.GetCollisionMask(spriteIndex, imageIndex, 0)
			collisionMask.Kind = sprite.CollisionMaskInherit
		}
		basePos.Y += 30
		if drawButton(basePos, "Kind: Manual") {
			collisionMask = sprite.GetCollisionMask(spriteIndex, imageIndex, 0)
			if collisionMask.Kind != sprite.CollisionMaskManual {
				collisionMask.Rect = inheritCollisionMask.Rect
				collisionMask.Kind = sprite.CollisionMaskManual
			}
		}
		basePos.Y += 40

		pos := basePos

		//
		drawMask := inheritCollisionMask
		if drawMask == nil {
			drawMask = collisionMask
		}

		if drawMask != nil {
			{
				text := strconv.FormatFloat(drawMask.Rect.Left(), 'f', -1, 64)
				if KeyboardCheck(VkControl) && KeyboardCheckPressed(Vk1) {
					editor.animationEditorToggleMenu(animMenuSpriteBboxLeft)
					if editor.menuOpened == animMenuSpriteBboxLeft {
						SetKeyboardString(text)
					}
				}
				if drawInputText(&pos, "Left (CTRL + 1)", text, editor.menuOpened == animMenuSpriteBboxLeft) {
					editor.animationEditorToggleMenu(animMenuSpriteBboxLeft)
				}
			}
			{
				pos.Y += 24

				text := strconv.FormatFloat(drawMask.Rect.Bottom(), 'f', -1, 64)
				if KeyboardCheck(VkControl) && KeyboardCheckPressed(Vk3) {
					editor.animationEditorToggleMenu(animMenuSpriteBboxBottom)
					if editor.menuOpened == animMenuSpriteBboxBottom {
						SetKeyboardString(text)
					}
				}
				if drawInputText(&pos, "Bottom (CTRL + 3)", text, editor.menuOpened == animMenuSpriteBboxBottom) {
					editor.animationEditorToggleMenu(animMenuSpriteBboxBottom)
				}
			}
			pos = basePos
			pos.X += 160
			{
				text := strconv.FormatFloat(drawMask.Rect.Top(), 'f', -1, 64)
				if KeyboardCheck(VkControl) && KeyboardCheckPressed(Vk2) {
					editor.animationEditorToggleMenu(animMenuSpriteBboxTop)
					if editor.menuOpened == animMenuSpriteBboxTop {
						SetKeyboardString(text)
					}
				}
				if drawInputText(&pos, "Top (CTRL + 2)", text, editor.menuOpened == animMenuSpriteBboxTop) {
					editor.animationEditorToggleMenu(animMenuSpriteBboxTop)
				}
			}
			{
				pos.Y += 24

				text := strconv.FormatFloat(drawMask.Rect.Right(), 'f', -1, 64)
				if KeyboardCheck(VkControl) && KeyboardCheckPressed(Vk4) {
					editor.animationEditorToggleMenu(animMenuSpriteBboxRight)
					if editor.menuOpened == animMenuSpriteBboxRight {
						SetKeyboardString(text)
					}
				}
				if drawInputText(&pos, "Right (CTRL + 4)", text, editor.menuOpened == animMenuSpriteBboxRight) {
					editor.animationEditorToggleMenu(animMenuSpriteBboxRight)
				}
			}
		}
	}
}
