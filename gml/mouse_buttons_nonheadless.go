// +build !headless

package gml

import "github.com/hajimehoshi/ebiten"

const (
	MbLeft   int = iota + int(ebiten.MouseButtonLeft)
	MbRight      = int(ebiten.MouseButtonRight)
	MbMiddle     = int(ebiten.MouseButtonMiddle)
	MbSize       = int(ebiten.MouseButtonMiddle) + 1
)
