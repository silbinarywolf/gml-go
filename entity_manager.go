package gml

import "reflect"

var g_entityManager entityManager

type entityManager struct {
	entities       []EntityType
	idToEntityData []EntityType
}

func InstanceCreate(position Vec, entityID int) {
	if entityID == 0 {
		panic("Cannot pass 0 as 2nd parameter to InstanceCreate(position, entityID)")
	}
	valToCopy := g_entityManager.idToEntityData[entityID]
	e := reflect.New(reflect.ValueOf(valToCopy).Elem().Type()).Interface().(EntityType)
	g_entityManager.entities = append(g_entityManager.entities, e)
	be := e.BaseEntity()
	be.init()
	e.Create()
	be.Vec = position
}

func (manager *entityManager) update() {
	for _, e := range manager.entities {
		e.Update()
	}

	for _, e := range manager.entities {
		be := e.BaseEntity()
		be.SpriteState.imageUpdate()
	}
}

func (manager *entityManager) draw() {
	for _, e := range manager.entities {
		e.Draw()
	}
}
