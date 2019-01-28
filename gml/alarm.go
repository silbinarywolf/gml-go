package gml

import "github.com/silbinarywolf/gml-go/gml/internal/dt"

type Alarm struct {
	isTimerSet bool
	timeLeft   float64
}

// Set an alarm. This requires you process it every Update with Tick
func (alarm *Alarm) Set(ticks float64) {
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
		alarm.isTimerSet = false
		return true
	}
	return false
}

// Repeat is like to be used for events that repeat every N frames and will return true once N frames are processed
func (alarm *Alarm) Repeat(ticks float64) bool {
	// todo(Jake): 2018-12-02: #23
	// I'd like to test this alarm system against Game Maker and
	// see if I can make it feel the same.
	// (ie. you give the same values as Game Maker, you can get the same results)
	if !alarm.isTimerSet {
		alarm.timeLeft = ticks
		alarm.isTimerSet = true
		return false
	}
	amount := 1.0 * dt.DeltaTime()
	alarm.timeLeft -= amount
	if alarm.timeLeft <= 0 {
		alarm.timeLeft += ticks
		alarm.isTimerSet = false
		return true
	}
	return false
}
