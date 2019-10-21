package game

// NOTE(Jake): 2019-04-27
// This test confirms that sin has different outputs in Go when compared to Game Maker Studio.

// TestSinCounter compares outputs of a pattern used by Worm against a dataset retrieved from
// how Game Maker Studio 2 outputs the same values each frame.
/*func TestSinCounter(t *testing.T) {
	sinTimer := alarm.Alarm{}
	for frame := 0; frame < len(sinCounterTestData); frame++ {
		data := &sinCounterTestData[frame]
		sinTimer.Repeat(WormSinCounterStart)

		alarm := sinTimer.Get()
		//sine := float64(math32.Sin(float32((alarm * 0.15))))
		sine := math.Sin(alarm * 0.15)
		sinCounter := math.Round(sine * 21)
		if data.Alarm != alarm ||
			data.Sine != sine ||
			data.SinCounter != sinCounter {
			t.Errorf("Frame %v: Not matching test data\n", frame)
			if data.Alarm != alarm {
				t.Errorf("- Alarm expected %v but got %v\n", int(data.Alarm), int(alarm))
			}
			if data.Sine != sine {
				t.Errorf("- Sine expected %v but got %v\n", data.Sine, sine)
			}
			if data.SinCounter != sinCounter {
				t.Errorf("- SinCounter expected %v but got %v\n", data.SinCounter, sinCounter)
			}
		}
	}
}*/

type wormSinCounterData struct {
	Alarm      float64
	Sine       float64
	SinCounter float64
}

// sinCounterTestData is taken from a Game Maker Studio 2 project
var sinCounterTestData = []wormSinCounterData{
	{
		Alarm:      9999998,
		Sine:       0.71,
		SinCounter: 15,
	},
	{
		Alarm:      9999997,
		Sine:       0.86,
		SinCounter: 18,
	},
	{
		Alarm:      9999996,
		Sine:       0.92,
		SinCounter: 19,
	},
	{
		Alarm:      9999995,
		Sine:       0.96,
		SinCounter: 20,
	},
	{
		Alarm:      9999994,
		Sine:       0.99,
		SinCounter: 21,
	},
	{
		Alarm:      9999993,
		Sine:       1.00,
		SinCounter: 21,
	},
	{
		Alarm:      9999992,
		Sine:       0.98,
		SinCounter: 21,
	},
	{
		Alarm:      9999991,
		Sine:       0.94,
		SinCounter: 20,
	},
	{
		Alarm:      9999990,
		Sine:       0.89,
		SinCounter: 19,
	},
	{
		Alarm:      9999989,
		Sine:       0.83,
		SinCounter: 17,
	},
	{
		Alarm:      9999988,
		Sine:       0.75,
		SinCounter: 16,
	},
	{
		Alarm:      9999987,
		Sine:       0.57,
		SinCounter: 12,
	},
	{
		Alarm:      9999986,
		Sine:       0.46,
		SinCounter: 10,
	},
	{
		Alarm:      9999985,
		Sine:       0.35,
		SinCounter: 7,
	},
	{
		Alarm:      9999984,
		Sine:       0.23,
		SinCounter: 5,
	},
	{
		Alarm:      9999983,
		Sine:       0.11,
		SinCounter: 2,
	},
	{
		Alarm:      9999982,
		Sine:       -0.14,
		SinCounter: -3,
	},
	{
		Alarm:      9999981,
		Sine:       -0.27,
		SinCounter: -6,
	},
	{
		Alarm:      9999980,
		Sine:       -0.38,
		SinCounter: -8,
	},
	{
		Alarm:      9999979,
		Sine:       -0.50,
		SinCounter: -10,
	},
	{
		Alarm:      9999978,
		Sine:       -0.60,
		SinCounter: -13,
	},
	{
		Alarm:      9999977,
		Sine:       -0.78,
		SinCounter: -16,
	},
	{
		Alarm:      9999976,
		Sine:       -0.85,
		SinCounter: -18,
	},
	{
		Alarm:      9999975,
		Sine:       -0.91,
		SinCounter: -19,
	},
	{
		Alarm:      9999974,
		Sine:       -0.96,
		SinCounter: -20,
	},
	{
		Alarm:      9999973,
		Sine:       -0.98,
		SinCounter: -21,
	},
	{
		Alarm:      9999972,
		Sine:       -1.00,
		SinCounter: -21,
	},
}
