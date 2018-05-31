package gml

func (m *Map) Run() {
	for _, inst := range m.Entities {
		InstanceCreate(V(float64(inst.X), float64(inst.Y)), ObjectIndex(inst.ObjectIndex))
	}
}
