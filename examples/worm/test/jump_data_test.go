// +build headless

package test

type jumpData struct {
	X, Y           float64
	HasPressedJump bool
}

var wormJumpData = [][]jumpData{
	// Frame 1
	{
		// Worm
		{HasPressedJump: true, X: 304, Y: 507.66},
		// WormBody
		{X: 264, Y: 528},
	},
	// Frame 2
	{
		// Worm
		{X: 304, Y: 487.98},
		// WormBody
		{X: 264, Y: 528},
	},
	// Frame 3
	{
		// Worm
		{X: 304, Y: 468.96},
		// WormBody
		{X: 264, Y: 507.66},
	},
	// Frame 4
	{
		// Worm
		{X: 304, Y: 450.60},
		// WormBody
		{X: 264, Y: 507.66},
	},
	// Frame 5
	{
		// Worm
		{X: 304, Y: 432.90},
		// WormBody
		{X: 264, Y: 468.96},
	},
	// Frame 6
	{
		// Worm
		{X: 304, Y: 415.86},
		// WormBody
		{X: 264, Y: 468.96},
	},
	// Frame 7
	{
		// Worm
		{X: 304, Y: 399.48},
		// WormBody
		{X: 264, Y: 432.90},
	},
	// Frame 8
	{
		// Worm
		{X: 304, Y: 383.76},
		// WormBody
		{X: 264, Y: 432.90},
	},
	// Frame 9
	{
		// Worm
		{X: 304, Y: 368.70},
		// WormBody
		{X: 264, Y: 399.48},
	},
	// Frame 10
	{
		// Worm
		{X: 304, Y: 354.30},
		// WormBody
		{X: 264, Y: 399.48},
	},
	// Frame 11
	{
		// Worm
		{X: 304, Y: 340.56},
		// WormBody
		{X: 264, Y: 368.70},
	},
	// Frame 12
	{
		// Worm
		{X: 304, Y: 327.48},
		// WormBody
		{X: 264, Y: 368.70},
	},
	// Frame 13
	{
		// Worm
		{X: 304, Y: 315.06},
		// WormBody
		{X: 264, Y: 340.56},
	},
	// Frame 14
	{
		// Worm
		{X: 304, Y: 303.30},
		// WormBody
		{X: 264, Y: 340.56},
	},
	// Frame 15
	{
		// Worm
		{X: 304, Y: 292.20},
		// WormBody
		{X: 264, Y: 315.06},
	},
	// Frame 16
	{
		// Worm
		{X: 304, Y: 281.76},
		// WormBody
		{X: 264, Y: 315.06},
	},
	// Frame 17
	{
		// Worm
		{X: 304, Y: 271.98},
		// WormBody
		{X: 264, Y: 292.20},
	},
	// Frame 18
	{
		// Worm
		{X: 304, Y: 262.86},
		// WormBody
		{X: 264, Y: 292.20},
	},
	// Frame 19
	{
		// Worm
		{X: 304, Y: 254.40},
		// WormBody
		{X: 264, Y: 271.98},
	},
	// Frame 20
	{
		// Worm
		{X: 304, Y: 246.60},
		// WormBody
		{X: 264, Y: 271.98},
	},
	// Frame 21
	{
		// Worm
		{X: 304, Y: 239.46},
		// WormBody
		{X: 264, Y: 254.40},
	},
	// Frame 22
	{
		// Worm
		{X: 304, Y: 232.98},
		// WormBody
		{X: 264, Y: 254.40},
	},
	// Frame 23
	{
		// Worm
		{X: 304, Y: 227.16},
		// WormBody
		{X: 264, Y: 239.46},
	},
	// Frame 24
	{
		// Worm
		{X: 304, Y: 222},
		// WormBody
		{X: 264, Y: 239.46},
	},
	// Frame 25
	{
		// Worm
		{X: 304, Y: 217.50},
		// WormBody
		{X: 264, Y: 227.16},
	},
	// Frame 26
	{
		// Worm
		{X: 304, Y: 213.66},
		// WormBody
		{X: 264, Y: 227.16},
	},
	// Frame 27
	{
		// Worm
		{X: 304, Y: 210.48},
		// WormBody
		{X: 264, Y: 217.50},
	},
	// Frame 28
	{
		// Worm
		{X: 304, Y: 207.96},
		// WormBody
		{X: 264, Y: 217.50},
	},
	// Frame 29
	{
		// Worm
		{X: 304, Y: 206.10},
		// WormBody
		{X: 264, Y: 210.48},
	},
	// Frame 30
	{
		// Worm
		{X: 304, Y: 204.90},
		// WormBody
		{X: 264, Y: 210.48},
	},
	// Frame 31
	{
		// Worm
		{X: 304, Y: 204.36},
		// WormBody
		{X: 264, Y: 206.10},
	},
	// Frame 32
	{
		// Worm
		{X: 304, Y: 204.48},
		// WormBody
		{X: 264, Y: 206.10},
	},
	// Frame 33
	{
		// Worm
		{X: 304, Y: 205.16},
		// WormBody
		{X: 264, Y: 204.36},
	},
	// Frame 34
	{
		// Worm
		{X: 304, Y: 206.40},
		// WormBody
		{X: 264, Y: 204.36},
	},
	// Frame 35
	{
		// Worm
		{X: 304, Y: 208.20},
		// WormBody
		{X: 264, Y: 205.16},
	},
	// Frame 36
	{
		// Worm
		{X: 304, Y: 210.56},
		// WormBody
		{X: 264, Y: 205.16},
	},
	// Frame 37
	{
		// Worm
		{X: 304, Y: 213.48},
		// WormBody
		{X: 264, Y: 208.20},
	},
	// Frame 38
	{
		// Worm
		{X: 304, Y: 216.96},
		// WormBody
		{X: 264, Y: 208.20},
	},
	// Frame 39
	{
		// Worm
		{X: 304, Y: 221.00},
		// WormBody
		{X: 264, Y: 213.48},
	},
	// Frame 40
	{
		// Worm
		{X: 304, Y: 225.60},
		// WormBody
		{X: 264, Y: 213.48},
	},
	// Frame 41
	{
		// Worm
		{X: 304, Y: 230.76},
		// WormBody
		{X: 264, Y: 221.00},
	},
	// Frame 42
	{
		// Worm
		{X: 304, Y: 236.48},
		// WormBody
		{X: 264, Y: 221.00},
	},
	// Frame 43
	{
		// Worm
		{X: 304, Y: 242.76},
		// WormBody
		{X: 264, Y: 230.76},
	},
	// Frame 44
	{
		// Worm
		{X: 304, Y: 249.60},
		// WormBody
		{X: 264, Y: 230.76},
	},
	// Frame 45
	{
		// Worm
		{X: 304, Y: 257.00},
		// WormBody
		{X: 264, Y: 242.76},
	},
	// Frame 46
	{
		// Worm
		{X: 304, Y: 264.96},
		// WormBody
		{X: 264, Y: 242.76},
	},
	// Frame 47
	{
		// Worm
		{X: 304, Y: 273.48},
		// WormBody
		{X: 264, Y: 257.00},
	},
	// Frame 48
	{
		// Worm
		{X: 304, Y: 282.56},
		// WormBody
		{X: 264, Y: 257.00},
	},
	// Frame 49
	{
		// Worm
		{X: 304, Y: 292.20},
		// WormBody
		{X: 264, Y: 273.48},
	},
	// Frame 50
	{
		// Worm
		{X: 304, Y: 302.40},
		// WormBody
		{X: 264, Y: 273.48},
	},
	// Frame 51
	{
		// Worm
		{X: 304, Y: 313.16},
		// WormBody
		{X: 264, Y: 292.20},
	},
	// Frame 52
	{
		// Worm
		{X: 304, Y: 324.48},
		// WormBody
		{X: 264, Y: 292.20},
	},
	// Frame 53
	{
		// Worm
		{X: 304, Y: 336.36},
		// WormBody
		{X: 264, Y: 313.16},
	},
	// Frame 54
	{
		// Worm
		{X: 304, Y: 348.80},
		// WormBody
		{X: 264, Y: 313.16},
	},
	// Frame 55
	{
		// Worm
		{X: 304, Y: 361.80},
		// WormBody
		{X: 264, Y: 336.36},
	},
	// Frame 56
	{
		// Worm
		{X: 304, Y: 375.36},
		// WormBody
		{X: 264, Y: 336.36},
	},
	// Frame 57
	{
		// Worm
		{X: 304, Y: 389.48},
		// WormBody
		{X: 264, Y: 361.80},
	},
	// Frame 58
	{
		// Worm
		{X: 304, Y: 404.16},
		// WormBody
		{X: 264, Y: 361.80},
	},
	// Frame 59
	{
		// Worm
		{X: 304, Y: 419.40},
		// WormBody
		{X: 264, Y: 389.48},
	},
	// Frame 60
	{
		// Worm
		{X: 304, Y: 435.20},
		// WormBody
		{X: 264, Y: 389.48},
	},
	// Frame 61
	{
		// Worm
		{X: 304, Y: 451.56},
		// WormBody
		{X: 264, Y: 419.40},
	},
	// Frame 62
	{
		// Worm
		{X: 304, Y: 468.48},
		// WormBody
		{X: 264, Y: 419.40},
	},
	// Frame 63
	{
		// Worm
		{X: 304, Y: 485.96},
		// WormBody
		{X: 264, Y: 451.56},
	},
	// Frame 64
	{
		// Worm
		{X: 304, Y: 504.00},
		// WormBody
		{X: 264, Y: 451.56},
	},
	// Frame 65
	{
		// Worm
		{X: 304, Y: 522.60},
		// WormBody
		{X: 264, Y: 485.96},
	},
}
