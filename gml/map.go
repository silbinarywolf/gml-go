package gml

func (m *Map) Run() {
	for _, e := range m.Entities {
		InstanceCreate(V(float64(e.X), float64(e.Y)), EntityID(e.EntityID))
	}
}
