package gml

func (e *Entity) PlaceFree(position Vec) bool {
	manager := g_entityManager
	for _, entity := range manager.entities {
		entity := entity.BaseEntity()
		if e == entity {
			// Skip self
			continue
		}
		// todo: actually detect collision with another entity
	}
	return true
}
