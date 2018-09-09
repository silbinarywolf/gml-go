package reditor

type Menu int

const (
	MenuNone Menu = 0 + iota
	MenuEntity
	MenuSprite
	MenuBackground
	MenuNewRoom
	MenuLoadRoom
	MenuNewLayer
	MenuSetOrder
)
