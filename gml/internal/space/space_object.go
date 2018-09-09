package space

type SpaceObject struct {
	*Space
	spaceIndex int
}

func (record *SpaceObject) Init(space *Space, spaceIndex int) {
	if record.Space != nil {
		panic("Can only initialize SpaceObject once.")
	}
	record.Space = space
	record.spaceIndex = spaceIndex
}

func (space *SpaceObject) SpaceIndex() int {
	return space.spaceIndex
}
