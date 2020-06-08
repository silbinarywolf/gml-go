package gml

const (
	DEBUG_COLLISION = false
)

type collisionObject interface {
	BaseObject() *Object
}

//var list []InstanceIndex

func CollisionRectList(instType collisionObject, x, y float64) []InstanceIndex {
	inst := instType.BaseObject()
	room := roomGetInstance(inst.BaseObject().RoomIndex())
	if room == nil {
		panic("RoomInstance this object belongs to has been destroyed")
	}

	// Create collision rect at position provided in function
	r1 := inst.bboxAt(x, y)

	// todo(Jake): 2018-12-01 - #18
	// Consider pooling reusable InstanceIndex slices to
	// improve performance.
	var list []InstanceIndex
	//list = list[:0]
	for _, otherIndex := range room.instances {
		other := otherIndex.getBaseObject()
		if other == nil {
			continue
		}
		if r1.CollisionRectangle(other.Bbox()) &&
			!other.internal.IsDestroyed &&
			inst != other {
			list = append(list, otherIndex)
		}
	}

	if len(list) == 0 {
		return nil
	}
	return list
}

func PlaceFree(instType collisionObject, x, y float64) bool {
	inst := instType.BaseObject()
	room := roomGetInstance(inst.BaseObject().RoomIndex())
	if room == nil {
		panic("RoomInstance this object belongs to has been destroyed")
	}

	// Create collision rect at position provided in function
	r1 := inst.bboxAt(x, y)

	hasCollision := false
	for _, otherIndex := range room.instances {
		other := otherIndex.getBaseObject()
		if other == nil {
			continue
		}
		if other.Solid() &&
			r1.CollisionRectangle(other.Bbox()) &&
			inst != other {
			hasCollision = true
		}
	}
	return !hasCollision
}
