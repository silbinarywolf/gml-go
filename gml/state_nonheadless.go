// +build !headless

package gml

// draw the room that the camera is in
func (state *state) draw() {
	for i := 0; i < len(gCameraManager.cameras); i++ {
		view := &gCameraManager.cameras[i]
		if !view.enabled {
			continue
		}
		view.update()
		cameraSetActive(i)

		cameraClear(i)

		if inst := InstanceGet(view.follow); inst != nil {
			// Render instances in same room as instance following
			inst := inst.BaseObject()
			roomInst := roomGetInstance(inst.RoomInstanceIndex())
			if roomInst == nil {
				panic("RoomInstance this object belongs to has been destroyed")
			}
			roomInst.draw()
		} else {
			// If no follower is configured, just render the first active room found
			roomInst := roomGetInstance(1)
			if roomInst == nil {
				panic("Unable to find room instance: 1")
			}
			roomInst.draw()
		}

		// Render camera onto OS-window
		cameraDraw(i)
	}
	cameraClearActive()
}
