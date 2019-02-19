// The alarm tests in this package are built and tested against how
// Game Maker Studio alarms work.
// Version: Game Maker Studio 2
// IDE Version: 2.2.1.375
// Runtime Version: 2 2.2.1.297

package alarm

import (
	"testing"
)

var frameHasAlarm = map[int]bool{
	/*
		This is a print of Game Maker for setting an alarm[0] = 4 in
		the create event and resetting it every time it ticks over
		------------------------------------------------------------
		begin_step: 0
		step: 0
		begin_step: 1
		step: 1
		begin_step: 2
		step: 2
		begin_step: 3
		alarm0: fired
		step: 3
		begin_step: 4
		step: 4
		begin_step: 5
		step: 5
		begin_step: 6
		step: 6
		begin_step: 7
		alarm0: fired
		step: 7
		begin_step: 8
		step: 8
	*/
	3: true,
	7: true,
}

func TestAlarmSet(t *testing.T) {
	timer := Alarm{}
	timer.Set(4)
	for frame := 0; frame < 10; frame++ {
		isSet, _ := frameHasAlarm[frame]
		if timer.Tick() {
			if !isSet {
				t.Errorf("Alarm is not meant to fire on frame %d", frame)
			}
			timer.Set(4)
		} else if isSet {
			t.Errorf("Alarm did not fire when expected on frame %d", frame)
		}
	}
}

func TestAlarmRepeat(t *testing.T) {
	timer := Alarm{}
	for frame := 0; frame < 10; frame++ {
		isSet, _ := frameHasAlarm[frame]
		if timer.Repeat(4) {
			if !isSet {
				t.Errorf("Alarm is not meant to fire on frame %d", frame)
			}
		} else if isSet {
			t.Errorf("Alarm did not fire when expected on frame %d", frame)
		}
	}
}
