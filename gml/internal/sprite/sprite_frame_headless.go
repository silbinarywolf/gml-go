// +build headless

package sprite

import (
	_ "image/png"
)

type SpriteFrame struct {
	spriteFrameShared
	width, height int
}

func (frame *SpriteFrame) Size() (width int, height int) { return frame.width, frame.height }

func createFrame(frameData spriteAssetFrame) (SpriteFrame, error) {
	r := SpriteFrame{
		width:  int(frameData.Size.X),
		height: int(frameData.Size.Y),
	}
	r.init(frameData)
	return r, nil
}

// NOTE(Jake): 2018-06-17
//
// This is commented out as headless mode doesn't
// draw any images.
//
//func GetRawFrame(spr, index) *SpriteFrame {
//	return nil
//}
