package gml

func (room *Room) Run() {
	for _, inst := range room.Instances {
		InstanceCreate(V(float64(inst.X), float64(inst.Y)), ObjectIndex(inst.ObjectIndex))
	}
}
