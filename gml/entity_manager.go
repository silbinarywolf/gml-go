package gml

import "reflect"

var g_entityManager entityManager

type entityManagerResettableData struct {
	entities []EntityType
}

type entityManager struct {
	entityManagerResettableData
	idToEntityData []EntityType
}

func Instances() []EntityType {
	return g_entityManager.entities
}

func InstanceCreate(position Vec, entityID int) EntityType {
	if entityID == 0 {
		panic("Cannot pass 0 as 2nd parameter to InstanceCreate(position, entityID)")
	}
	valToCopy := g_entityManager.idToEntityData[entityID]
	e := reflect.New(reflect.ValueOf(valToCopy).Elem().Type()).Interface().(EntityType)
	index := len(g_entityManager.entities)
	g_entityManager.entities = append(g_entityManager.entities, e)
	be := e.BaseEntity()
	be.index = index
	be.init()
	e.Create()
	be.Vec = position
	return e
}

func InstanceDestroy(entity EntityType) {
	be := entity.BaseEntity()

	// Unordered delete
	i := be.index
	manager := &g_entityManager
	lastEntry := manager.entities[len(manager.entities)-1]
	manager.entities[i] = lastEntry
	manager.entities = manager.entities[:len(manager.entities)-1]

	// maybetodo(Jake): 2018-05-27
	//
	// Add func Destroy() to Entity interface and Call e.Destroy()
	//
}

func (manager *entityManager) reset() {
	manager.entityManagerResettableData = entityManagerResettableData{}
}

func (manager *entityManager) update() {
	// Delete entities flagged for deletion
	/*for i, e := range manager.entities {
		be := e.BaseEntity()
		if be.isDestroyed {
			// Unordered delete
			lastEntry := manager.entities[len(manager.entities)-1]
			manager.entities[i] = lastEntry
			manager.entities = manager.entities[:len(manager.entities)-1]
		}
	}*/

	for _, e := range manager.entities {
		e.Update()
	}

	for _, e := range manager.entities {
		if e == nil {
			continue
		}
		be := e.BaseEntity()
		be.SpriteState.imageUpdate()
	}
}

func (manager *entityManager) draw() {
	for _, e := range manager.entities {
		if e == nil {
			continue
		}
		e.Draw()
	}
}
