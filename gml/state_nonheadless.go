// +build !headless

package gml

func (state *state) draw() {
	for i := 0; i < len(gCameraManager.cameras); i++ {
		view := &gCameraManager.cameras[i]
		if !view.enabled {
			continue
		}
		view.update()
		cameraSetActive(i)

		cameraClear(i)

		// Render global instances
		//state.globalInstances.draw()

		if view.follow != nil {
			// Render instances in same room as instance following
			inst := view.follow.BaseObject()
			roomInst := roomGetInstance(inst.RoomInstanceIndex())
			if roomInst == nil {
				panic("RoomInstance this object belongs to has been destroyed")
			}
			roomInst.draw()
		} else {
			// Render each instance in each room instance
			for i := 1; i < len(state.roomInstances); i++ {
				roomInst := &state.roomInstances[i]
				if !roomInst.used {
					continue
				}
				roomInst.draw()
			}
		}

		// Render camera onto OS-window
		cameraDraw(i)
	}
	cameraClearActive()
}
