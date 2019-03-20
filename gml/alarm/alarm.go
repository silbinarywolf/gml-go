package alarm

import (
	"bytes"
	"encoding/binary"

	"github.com/silbinarywolf/gml-go/gml/internal/dt"
)

// alarmNotSet is what an alarm will default to when the timer runs out.
// In Game Maker this is -1 but this clashes with Golang's meaningful zero values.
// I've opted to make it 0 as it makes using an alarm with the "Repeat" method, very
// easy to just drop-in and use.
const alarmNotSet = 0

type Alarm struct {
	internal alarmSerialize
}

type alarmSerialize struct {
	IsTimerSet bool
	TimeLeft   float64
}

func (alarm *Alarm) Get() float64 {
	//if !alarm.isTimerSet {
	//	return 0
	//}
	return alarm.internal.TimeLeft
}

// Set an alarm. This requires you process it every Update with Tick
func (alarm *Alarm) Set(ticks float64) {
	if ticks <= 0 {
		alarm.internal.TimeLeft = alarmNotSet
		alarm.internal.IsTimerSet = false
		return
	}
	alarm.internal.TimeLeft = ticks
	alarm.internal.IsTimerSet = true
}

// IsRunning will return true if the timer is set and running.
// When used with Set, it will remain true until Tick() returns true
// When used with Repeat, it will always remain true after Repeat() is called for the first time.
func (alarm *Alarm) IsRunning() bool {
	return alarm.internal.IsTimerSet
}

// Tick will process the timed event and return true if the timer has expired
func (alarm *Alarm) Tick() bool {
	if !alarm.internal.IsTimerSet {
		return false
	}
	alarm.internal.TimeLeft -= 1.0 * dt.DeltaTime()
	if alarm.internal.TimeLeft <= 0 {
		alarm.internal.TimeLeft = alarmNotSet
		alarm.internal.IsTimerSet = false
		return true
	}
	return false
}

// Repeat is like to be used for events that repeat every N frames and will return true once N frames are processed
func (alarm *Alarm) Repeat(ticks float64) bool {
	if !alarm.internal.IsTimerSet {
		alarm.internal.TimeLeft = ticks
		alarm.internal.IsTimerSet = true
	}
	amount := 1.0 * dt.DeltaTime()
	alarm.internal.TimeLeft -= amount
	if alarm.internal.TimeLeft <= 0 {
		alarm.internal.TimeLeft = alarmNotSet
		alarm.internal.IsTimerSet = false
		return true
	}
	return false
}

func (alarm Alarm) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, alarm.internal); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (alarm *Alarm) UnmarshalBinary(data []byte) error {
	buf := bytes.NewReader(data)
	if err := binary.Read(buf, binary.LittleEndian, &alarm.internal); err != nil {
		return err
	}
	return nil
}
