package dt

// NOTE(Jake): 2019-01-28
// I want to share this code throughout the "internal" packages.
// So the easiest way to do that is to have this logic here.

const (
	DefaultMaxTPS = 60
)

var (
	fixedDeltaTime float64 = 1 // Default to 1 so tests that use it still work
	designedTPS    int     = DefaultMaxTPS
	maxTPS         int     = DefaultMaxTPS
)

// SetDesignedTPS is the ticks-per-second the game was initially designed to run at. ie. 30tps, 60tps, etc
// for example, if you're porting a Game Maker game that ran at 30 frames per second, you'd want this to be 30 so
// that translation of alarm logic works seamlessly.
func SetDesignedTPS(tps int) {
	designedTPS = tps
	fixedDeltaTime = float64(designedTPS) / float64(maxTPS)
}

// DesignedTPS() returns the ticks-per-second that the game was designed to run at.
//
// Default is 60
func DesignedTPS() int {
	return designedTPS
}

// SetMaxTPS is the ticks-per-second the game is trying to run at. ie. 240tps, 480tps
func SetMaxTPS(tps int) {
	maxTPS = tps
	fixedDeltaTime = float64(designedTPS) / float64(maxTPS)
}

// DeltaTime gets the fixed delta time based on the designed TPS divided by max TPS.
func DeltaTime() float64 {
	return fixedDeltaTime
}

// TestSetDeltaTime allows test code to override the delta time value manually
//func TestSetDeltaTime(dt float64) {
//	fixedDeltaTime = dt
//}
