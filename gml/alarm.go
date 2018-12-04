package gml

type Alarm struct {
	timeSet  int
	timeLeft int
}

// todo(Jake): 2018-12-02: #23
// I'd like to test this alarm system against Game Maker and
// see if I can make it feel the same.
// (ie. you give the same values as Game Maker, you can get the same results)
func (alarm *Alarm) Update(frames int) bool {
	if alarm.timeSet == 0 {
		alarm.timeSet = frames
		alarm.timeLeft = frames
	}
	alarm.timeLeft -= 1
	if alarm.timeLeft <= 0 {
		alarm.timeSet = 0
		return true
	}
	return false
}
