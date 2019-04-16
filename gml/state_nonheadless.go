// +build !headless

package gml

// draw the room that the each camera is in
func (state *state) draw() {
	for i := 0; i < len(gCameraManager.cameras); i++ {
		view := &gCameraManager.cameras[i]
		if !view.enabled {
			continue
		}
		cameraSetActive(i)
		cameraPreDraw(i)
		cameraClearSurface(i)

		if inst := view.follow.getBaseObject(); inst != nil {
			// Render instances in same room as instance following
			roomInst := roomGetInstance(inst.RoomInstanceIndex())
			if roomInst == nil {
				panic("draw: RoomInstance this object belongs to has been destroyed")
			}
			roomInst.draw()
		} else {
			// If no follower is configured, just render the first active room found
			roomInst := roomLastCreated()
			if roomInst == nil {
				panic("No room exists, you must create a room")
			}
			roomInst.draw()
		}

		// Render camera onto OS-window
		cameraDraw(i)
	}
	//cameraClearActive()
	// NOTE(Jake): 2019-04-15
	// Default to first camera for level editors / animation editor
	// etc.
	cameraSetActive(0)
}
