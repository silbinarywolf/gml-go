package input

import (
	"github.com/silbinarywolf/gml-go/gml"
)

func JumpPressed() bool {
	return gml.MouseCheckPressed(gml.MbLeft) ||
		gml.KeyboardCheckPressed(gml.VkSpace) ||
		gml.KeyboardCheckPressed(gml.VkEnter)
}
