package gml

import (
	"github.com/silbinarywolf/gml-go/gml/internal/geom"
)

const (
	DEBUG_COLLISION = false
)

type collisionObject interface {
	BaseObject() *Object
}

func CollisionRectList(instType collisionObject, position geom.Vec) []ObjectType {
	inst := instType.BaseObject()
	room := roomGetInstance(inst.BaseObject().RoomInstanceIndex())
	if room == nil {
		panic("RoomInstance this object belongs to has been destroyed")
	}

	// Create collision rect at position provided in function
	r1 := inst.Rect
	r1.Vec = position
	r1.Size = inst.Size

	// todo(Jake): 2018-12-01 - #18
	// Consider pooling reusable ObjectType slices to
	// improve performance.
	var list []ObjectType
	for i := 0; i < len(room.instanceLayers); i++ {
		for _, otherT := range room.instanceLayers[i].manager.instances {
			other := otherT.BaseObject()
			if !IsDestroyed(other) &&
				r1.CollisionRectangle(other.Rect) &&
				inst != other {
				list = append(list, otherT)
			}
		}
	}
	if len(list) == 0 {
		return nil
	}
	return list
}

func PlaceFree(instType collisionObject, position geom.Vec) bool {
	inst := instType.BaseObject()
	room := roomGetInstance(inst.BaseObject().RoomInstanceIndex())
	if room == nil {
		panic("RoomInstance this object belongs to has been destroyed")
	}

	// Create collision rect at position provided in function
	r1 := inst.Rect
	r1.Vec = position
	r1.Size = inst.Size

	//var debugString string
	hasCollision := false
	for i := 0; i < len(room.instanceLayers); i++ {
		for _, other := range room.instanceLayers[i].manager.instances {
			other := other.BaseObject()
			if other.Solid() &&
				r1.CollisionRectangle(other.Rect) &&
				inst != other {
				hasCollision = true
			}
		}
		/*spaces := &room.instanceLayers[i].manager.spaces
		for _, bucket := range spaces.Buckets() {
			for i := 0; i < bucket.Len(); i++ {
				other := bucket.Get(i)
				// NOTE(Jake): 2018-07-08
				//
				// For JavaScript performance, we get a 1.2x speedup if we
				// handle as much logic in one if-statement as possible.
				//
				// For native binaries, it doesn't seem to change performance noticeably
				// at all if I add "if inst == other || !instanceManager.spaces.IsUsed(i) { continue; }"
				//
				// ("gjbt" and Chrome 67 Windows were for benchmarking)
				//
				// NOTE(Jake): 2018-08-11
				//
				// Heavily refactored this since the above benchmark. But who cares really. I'll probably
				// need to re-do this collision engine so it supports spatial hashing.
				//
				if other.Solid() &&
					r1.CollisionRectangle(other.Rect) &&
					inst != other &&
					bucket.IsUsed(i) {
					hasCollision = true
				}
			}
		}*/
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
		/*spaces := layer.sprites
		for _, bucket := range spaces.Buckets() {
			for i := 0; i < bucket.Len(); i++ {
				other := bucket.Get(i)
				if r1.CollisionRectangle(other.Rect) &&
					inst != other &&
					bucket.IsUsed(i) {
					hasCollision = true
				}
			}
		}*/
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
