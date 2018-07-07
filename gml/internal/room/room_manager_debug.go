// +build debug

package room

func (room *Room) DebugWriteDataFile(roomPath string) {
	go func() {
		err := room.writeDataFile(roomPath)
		if err != nil {
			panic("Failed writing " + roomPath + ", error: " + err.Error())
		}
	}()
}
