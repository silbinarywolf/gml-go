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
	room := roomGetInstance(inst.BaseObject().RoomInstanceIndex())
	if room == nil {
		panic("RoomInstance this object belongs to has been destroyed")
	}

	// Create collision rect at position provided in function
	r1 := inst.Rect
	r1.X = x
	r1.Y = y
	r1.Size = inst.Size

	// todo(Jake): 2018-12-01 - #18
	// Consider pooling reusable InstanceIndex slices to
	// improve performance.
	var list []InstanceIndex
	//list = list[:0]
	for i := 0; i < len(room.instanceLayers); i++ {
		for _, otherIndex := range room.instanceLayers[i].instances {
			other := instanceGetBaseObject(otherIndex)
			if other == nil {
				continue
			}
			if r1.CollisionRectangle(other.Rect) &&
				!other.isDestroyed &&
				inst != other {
				list = append(list, otherIndex)
			}
		}
	}
	if len(list) == 0 {
		return nil
	}
	return list
}

func PlaceFree(instType collisionObject, x, y float64) bool {
	inst := instType.BaseObject()
	room := roomGetInstance(inst.BaseObject().RoomInstanceIndex())
	if room == nil {
		panic("RoomInstance this object belongs to has been destroyed")
	}

	// Create collision rect at position provided in function
	r1 := inst.Rect
	r1.X = x
	r1.Y = y
	r1.Size = inst.Size

	//var debugString string
	hasCollision := false
	for i := 0; i < len(room.instanceLayers); i++ {
		for _, otherIndex := range room.instanceLayers[i].instances {
			other := instanceGetBaseObject(otherIndex)
			if other == nil {
				continue
			}
			if other.Solid() &&
				r1.CollisionRectangle(other.Rect) &&
				inst != other {
				hasCollision = true
			}
		}
	}
	for i := 0; i < len(room.spriteLayers); i++ {
		layer := &room.spriteLayers[i]
		if !layer.hasCollision {
			continue
		}
		for _, other := range layer.sprites {
			if r1.CollisionRectangle(other.Rect()) {
				hasCollision = true
			}
		}
	}

	/*if DEBUG_COLLISION &&
		len(debugString) > 0 {
		// Get calling function name / line
		var message string
		{
			callIndex := 1
			for i := 0; i < 1; i++ {
				_, file, line, ok := runtime.Caller(callIndex)

				if ok {
					// Reduce full filepath to just the scope of the game
					fileParts := strings.Split(file, "/")
					if len(fileParts) >= 3 {
						file = fileParts[len(fileParts)-3] + "/" + fileParts[len(fileParts)-2] + "/" + fileParts[len(fileParts)-1]
					}
					message = message + file + "(" + strconv.Itoa(line) + ")"
				}
				callIndex++
			}
		}
		fmt.Printf("PlaceFree: collision between %s:\n%s%s\n\n", e.Sprite().name, debugString, message)
	}
	fmt.Printf("EndPlaceFree\n\n")*/
	return !hasCollision
}
