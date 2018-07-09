package gml

const (
	DEBUG_COLLISION = false
)

type collisionObject interface {
	BaseObject() *Object
}

func PlaceFree(instType collisionObject, position Vec) bool {
	baseObj := instType.BaseObject()

	var instanceManager *instanceManager
	{
		inst := baseObj

		if room := RoomGetInstance(inst.RoomInstanceIndex()); room == nil {
			instanceManager = gState.globalInstances
		} else {
			instanceManager = &room.instanceManager
		}
	}

	inst := baseObj.Space
	r1Left := position.X
	r1Right := r1Left + float64(inst.Size.X)
	r1Top := position.Y
	r1Bottom := r1Top + float64(inst.Size.Y)

	//var debugString string
	hasCollision := false
	for _, bucket := range instanceManager.spaces.Buckets() {
		for i := 0; i < bucket.Len(); i++ {
			other := bucket.Get(i)
			r2Left := other.X
			r2Right := r2Left + float64(other.Size.X)
			r2Top := other.Y
			r2Bottom := r2Top + float64(other.Size.Y)

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
			if r1Left < r2Right && r1Right > r2Left &&
				r1Top < r2Bottom && r1Bottom > r2Top &&
				inst != other &&
				bucket.IsUsed(i) {
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
