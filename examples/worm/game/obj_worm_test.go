package game

import (
	"math"
	"testing"

	"github.com/silbinarywolf/gml-go/gml"
)

// TestSinCounter compares outputs of a pattern used by Worm against a dataset retrieved from
// how Game Maker Studio 2 outputs the same values each frame.
func TestSinCounter(t *testing.T) {
	sinTimer := gml.Alarm{}
	for frame := 0; frame < len(sinCounterTestData); frame++ {
		data := &sinCounterTestData[frame]
		sinTimer.Repeat(WormSinCounterStart)

		alarm := sinTimer.Get()
		sinCounter := math.Round(math.Sin(sinTimer.Get()*0.15) * 21)
		if data.SinCounter != sinCounter ||
			data.Alarm != alarm {
			t.Errorf("Frame %v: Not matching test data\n", frame)
			if data.Alarm != alarm {
				t.Errorf("- Alarm expected %v but got %v\n", int(data.Alarm), int(alarm))
			}
			if data.SinCounter != sinCounter {
				t.Errorf("- SinCounter expected %v but got %v\n", data.SinCounter, sinCounter)
			}
		}
	}
}

type wormSinCounterData struct {
	Alarm      float64
	SinCounter float64
}

// sinCounterTestData is taken from a Game Maker Studio 2 project
var sinCounterTestData = []wormSinCounterData{
	{
		Alarm:      9999998,
		SinCounter: 15,
	},
	{
		Alarm:      9999997,
		SinCounter: 18,
	},
	{
		Alarm:      9999996,
		SinCounter: 19,
	},
	{
		Alarm:      9999995,
		SinCounter: 20,
	},
	{
		Alarm:      9999994,
		SinCounter: 21,
	},
	{
		Alarm:      9999993,
		SinCounter: 21,
	},
	{
		Alarm:      9999992,
		SinCounter: 21,
	},
	{
		Alarm:      9999991,
		SinCounter: 20,
	},
	{
		Alarm:      9999990,
		SinCounter: 19,
	},
	{
		Alarm:      9999989,
		SinCounter: 17,
	},
	{
		Alarm:      9999988,
		SinCounter: 16,
	},
	{
		Alarm:      9999987,
		SinCounter: 12,
	},
	{
		Alarm:      9999986,
		SinCounter: 10,
	},
	{
		Alarm:      9999985,
		SinCounter: 7,
	},
	{
		Alarm:      9999984,
		SinCounter: 5,
	},
	{
		Alarm:      9999983,
		SinCounter: 2,
	},
	{
		Alarm:      9999982,
		SinCounter: -3,
	},
	{
		Alarm:      9999981,
		SinCounter: -6,
	},
	{
		Alarm:      9999980,
		SinCounter: -8,
	},
	{
		Alarm:      9999979,
		SinCounter: -10,
	},
	{
		Alarm:      9999978,
		SinCounter: -13,
	},
	{
		Alarm:      9999977,
		SinCounter: -16,
	},
	{
		Alarm:      9999976,
		SinCounter: -18,
	},
	{
		Alarm:      9999975,
		SinCounter: -19,
	},
	{
		Alarm:      9999974,
		SinCounter: -20,
	},
	{
		Alarm:      9999973,
		SinCounter: -21,
	},
	{
		Alarm:      9999972,
		SinCounter: -21,
	},
	{
		Alarm:      9999971,
		SinCounter: -21,
	},
	{
		Alarm:      9999970,
		SinCounter: -20,
	},
}
