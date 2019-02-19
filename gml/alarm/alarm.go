package alarm

import (
	"github.com/silbinarywolf/gml-go/gml/internal/dt"
)

// alarmNotSet is what an alarm will default to when the timer runs out.
// In Game Maker this is -1 but this clashes with Golang's meaningful zero values.
// I've opted to make it 0 as it makes using an alarm with the "Repeat" method, very
// easy to just drop-in and use.
const alarmNotSet = 0

type Alarm struct {
	isTimerSet bool
	timeLeft   float64
}

func (alarm *Alarm) Get() float64 {
	//if !alarm.isTimerSet {
	//	return 0
	//}
	return alarm.timeLeft
}

// Set an alarm. This requires you process it every Update with Tick
func (alarm *Alarm) Set(ticks float64) {
	if ticks <= 0 {
		alarm.timeLeft = alarmNotSet
		alarm.isTimerSet = false
		return
	}
	alarm.timeLeft = ticks
	alarm.isTimerSet = true
}

// Tick will process the timed event and return true if the timer has expired
func (alarm *Alarm) Tick() bool {
	if !alarm.isTimerSet {
		return false
	}
	alarm.timeLeft -= 1.0 * dt.DeltaTime()
	if alarm.timeLeft <= 0 {
		alarm.timeLeft = alarmNotSet
		alarm.isTimerSet = false
		return true
	}
	return false
}

// Repeat is like to be used for events that repeat every N frames and will return true once N frames are processed
func (alarm *Alarm) Repeat(ticks float64) bool {
	if !alarm.isTimerSet {
		alarm.timeLeft = ticks
		alarm.isTimerSet = true
	}
	amount := 1.0 * dt.DeltaTime()
	alarm.timeLeft -= amount
	if alarm.timeLeft <= 0 {
		alarm.timeLeft = alarmNotSet
		alarm.isTimerSet = false
		return true
	}
	return false
}
