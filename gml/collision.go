package gml

import (
	"runtime"
	"strconv"
	"strings"
)

var (
	DEBUG_COLLISION = false
)

func (inst *Object) PlaceFree(position Vec) bool {
	entities := gInstanceManager.entities

	r1Left := position.X
	r1Right := r1Left + inst.Size.X
	r1Top := position.Y
	r1Bottom := r1Top + inst.Size.Y

	hasCollision := false
	var debugString string
	for _, other := range entities {
		other := other.BaseObject()
		if inst == other {
			// Skip self
			continue
		}
		r2Left := other.X
		r2Right := r2Left + other.Size.X
		r2Top := other.Y
		r2Bottom := r2Top + other.Size.Y

		if r1Left < r2Right && r1Right > r2Left &&
			r1Top < r2Bottom && r1Bottom > r2Top {
			hasCollision = true
			// Debug
			if DEBUG_COLLISION {
				debugString += "- " + other.Sprite().name + "\n"
			}
		}
	}
	if len(debugString) > 0 {
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
		//fmt.Printf("PlaceFree: collision between %s:\n%s%s\n\n", e.Sprite().name, debugString, message)
	}
	//fmt.Printf("EndPlaceFree\n\n")
	return !hasCollision
}
