package gml

// NOTE(Jake): 2018-09-09
//
// This test sucks and is out of date.
// the server still crashes, probably during a for-loop
// of every instance
//
/*func TestInstanceDestroyStability(t *testing.T) {
	roomInstance := RoomInstanceEmptyCreate()
	roomInstances := make([]object.ObjectType, 1024)
	for i := 0; i < len(roomInstances); i++ {
		roomInstances[i] = roomInstance.InstanceCreate(V(0, 0), ObjDummyPlayer)
	}

	// NOTE(Jake): 2018-07-08
	// Delete instance #128 and create new instance (should use that slot)
	{
		ExpectedSpaceIndex := 128
		roomInstance.InstanceDestroy(roomInstances[ExpectedSpaceIndex])
		inst := roomInstance.InstanceCreate(V(0, 0), ObjDummyPlayer)
		baseObj := inst.BaseObject()
		if baseObj.SpaceIndex() != ExpectedSpaceIndex {
			t.Errorf("SpaceIndex not equal to %d. When instance #%d was deleted, we expect that \"slot\" to be free to use when the next instance is created.", ExpectedSpaceIndex, ExpectedSpaceIndex)
		}
	}

	// NOTE(Jake): 2018-07-08
	// Create new instance
	{
		ExpectedSpaceIndex := 1024
		inst := roomInstance.InstanceCreate(V(0, 0), ObjDummyPlayer)
		baseObj := inst.BaseObject()
		if baseObj.SpaceIndex() != ExpectedSpaceIndex {
			t.Errorf("SpaceIndex not equal to %d. We expected this instance to use a new slot from the end of the array.", ExpectedSpaceIndex)
		}
	}

	// NOTE(Jake): 2018-07-08
	// Delete instance #810 and create new instance. (Using 810 as we have buckets of ~256 items, so we're testing that bucket functionality works)
	{
		ExpectedSpaceIndex := 810
		roomInstance.InstanceDestroy(roomInstances[ExpectedSpaceIndex])
		inst := roomInstance.InstanceCreate(V(0, 0), ObjDummyPlayer)
		baseObj := inst.BaseObject()
		if baseObj.SpaceIndex() != ExpectedSpaceIndex {
			t.Errorf("SpaceIndex not equal to %d. When instance #%d was deleted, we expect that \"slot\" to be free to use when the next instance is created.", ExpectedSpaceIndex, ExpectedSpaceIndex)
		}
	}
}
*/
