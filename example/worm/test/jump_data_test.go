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
		{HasPressedJump: true, X: 304, Y: 507.660003662109375},
		// WormBody
		{X: 264, Y: 528.000000000000000},
		// WormBody
		{X: 224, Y: 528.000000000000000},
		// WormBody
		{X: 184, Y: 528.000000000000000},
		// WormBody
		{X: 144, Y: 528.000000000000000},
		// WormBody
		{X: 104, Y: 528.000000000000000},
		// WormBody
		{X: 64, Y: 528.000000000000000},
		// WormBody
		{X: 24, Y: 528.000000000000000},
		// WormBody
		{X: -16, Y: 528.000000000000000},
		// WormBody
		{X: -56, Y: 528.000000000000000},
		// WormBody
		{X: -96, Y: 528.000000000000000},
	},
	// Frame 2
	{
		// Worm
		{X: 304, Y: 487.980010986328125},
		// WormBody
		{X: 264, Y: 528.000000000000000},
		// WormBody
		{X: 224, Y: 528.000000000000000},
		// WormBody
		{X: 184, Y: 528.000000000000000},
		// WormBody
		{X: 144, Y: 528.000000000000000},
		// WormBody
		{X: 104, Y: 528.000000000000000},
		// WormBody
		{X: 64, Y: 528.000000000000000},
		// WormBody
		{X: 24, Y: 528.000000000000000},
		// WormBody
		{X: -16, Y: 528.000000000000000},
		// WormBody
		{X: -56, Y: 528.000000000000000},
		// WormBody
		{X: -96, Y: 528.000000000000000},
	},
	// Frame 3
	{
		// Worm
		{X: 304, Y: 468.960021972656250},
		// WormBody
		{X: 264, Y: 507.660003662109375},
		// WormBody
		{X: 224, Y: 528.000000000000000},
		// WormBody
		{X: 184, Y: 528.000000000000000},
		// WormBody
		{X: 144, Y: 528.000000000000000},
		// WormBody
		{X: 104, Y: 528.000000000000000},
		// WormBody
		{X: 64, Y: 528.000000000000000},
		// WormBody
		{X: 24, Y: 528.000000000000000},
		// WormBody
		{X: -16, Y: 528.000000000000000},
		// WormBody
		{X: -56, Y: 528.000000000000000},
		// WormBody
		{X: -96, Y: 528.000000000000000},
	},
	// Frame 4
	{
		// Worm
		{X: 304, Y: 450.600036621093750},
		// WormBody
		{X: 264, Y: 507.660003662109375},
		// WormBody
		{X: 224, Y: 528.000000000000000},
		// WormBody
		{X: 184, Y: 528.000000000000000},
		// WormBody
		{X: 144, Y: 528.000000000000000},
		// WormBody
		{X: 104, Y: 528.000000000000000},
		// WormBody
		{X: 64, Y: 528.000000000000000},
		// WormBody
		{X: 24, Y: 528.000000000000000},
		// WormBody
		{X: -16, Y: 528.000000000000000},
		// WormBody
		{X: -56, Y: 528.000000000000000},
		// WormBody
		{X: -96, Y: 528.000000000000000},
	},
	// Frame 5
	{
		// Worm
		{X: 304, Y: 432.900024414062500},
		// WormBody
		{X: 264, Y: 468.960021972656250},
		// WormBody
		{X: 224, Y: 507.660003662109375},
		// WormBody
		{X: 184, Y: 528.000000000000000},
		// WormBody
		{X: 144, Y: 528.000000000000000},
		// WormBody
		{X: 104, Y: 528.000000000000000},
		// WormBody
		{X: 64, Y: 528.000000000000000},
		// WormBody
		{X: 24, Y: 528.000000000000000},
		// WormBody
		{X: -16, Y: 528.000000000000000},
		// WormBody
		{X: -56, Y: 528.000000000000000},
		// WormBody
		{X: -96, Y: 528.000000000000000},
	},
	// Frame 6
	{
		// Worm
		{X: 304, Y: 415.860015869140625},
		// WormBody
		{X: 264, Y: 468.960021972656250},
		// WormBody
		{X: 224, Y: 507.660003662109375},
		// WormBody
		{X: 184, Y: 528.000000000000000},
		// WormBody
		{X: 144, Y: 528.000000000000000},
		// WormBody
		{X: 104, Y: 528.000000000000000},
		// WormBody
		{X: 64, Y: 528.000000000000000},
		// WormBody
		{X: 24, Y: 528.000000000000000},
		// WormBody
		{X: -16, Y: 528.000000000000000},
		// WormBody
		{X: -56, Y: 528.000000000000000},
		// WormBody
		{X: -96, Y: 528.000000000000000},
	},
	// Frame 7
	{
		// Worm
		{X: 304, Y: 399.480010986328125},
		// WormBody
		{X: 264, Y: 432.900024414062500},
		// WormBody
		{X: 224, Y: 468.960021972656250},
		// WormBody
		{X: 184, Y: 507.660003662109375},
		// WormBody
		{X: 144, Y: 528.000000000000000},
		// WormBody
		{X: 104, Y: 528.000000000000000},
		// WormBody
		{X: 64, Y: 528.000000000000000},
		// WormBody
		{X: 24, Y: 528.000000000000000},
		// WormBody
		{X: -16, Y: 528.000000000000000},
		// WormBody
		{X: -56, Y: 528.000000000000000},
		// WormBody
		{X: -96, Y: 528.000000000000000},
	},
	// Frame 8
	{
		// Worm
		{X: 304, Y: 383.760009765625000},
		// WormBody
		{X: 264, Y: 432.900024414062500},
		// WormBody
		{X: 224, Y: 468.960021972656250},
		// WormBody
		{X: 184, Y: 507.660003662109375},
		// WormBody
		{X: 144, Y: 528.000000000000000},
		// WormBody
		{X: 104, Y: 528.000000000000000},
		// WormBody
		{X: 64, Y: 528.000000000000000},
		// WormBody
		{X: 24, Y: 528.000000000000000},
		// WormBody
		{X: -16, Y: 528.000000000000000},
		// WormBody
		{X: -56, Y: 528.000000000000000},
		// WormBody
		{X: -96, Y: 528.000000000000000},
	},
	// Frame 9
	{
		// Worm
		{X: 304, Y: 368.700012207031250},
		// WormBody
		{X: 264, Y: 399.480010986328125},
		// WormBody
		{X: 224, Y: 432.900024414062500},
		// WormBody
		{X: 184, Y: 468.960021972656250},
		// WormBody
		{X: 144, Y: 507.660003662109375},
		// WormBody
		{X: 104, Y: 528.000000000000000},
		// WormBody
		{X: 64, Y: 528.000000000000000},
		// WormBody
		{X: 24, Y: 528.000000000000000},
		// WormBody
		{X: -16, Y: 528.000000000000000},
		// WormBody
		{X: -56, Y: 528.000000000000000},
		// WormBody
		{X: -96, Y: 528.000000000000000},
	},
	// Frame 10
	{
		// Worm
		{X: 304, Y: 354.300018310546875},
		// WormBody
		{X: 264, Y: 399.480010986328125},
		// WormBody
		{X: 224, Y: 432.900024414062500},
		// WormBody
		{X: 184, Y: 468.960021972656250},
		// WormBody
		{X: 144, Y: 507.660003662109375},
		// WormBody
		{X: 104, Y: 528.000000000000000},
		// WormBody
		{X: 64, Y: 528.000000000000000},
		// WormBody
		{X: 24, Y: 528.000000000000000},
		// WormBody
		{X: -16, Y: 528.000000000000000},
		// WormBody
		{X: -56, Y: 528.000000000000000},
		// WormBody
		{X: -96, Y: 528.000000000000000},
	},
	// Frame 11
	{
		// Worm
		{X: 304, Y: 340.560028076171875},
		// WormBody
		{X: 264, Y: 368.700012207031250},
		// WormBody
		{X: 224, Y: 399.480010986328125},
		// WormBody
		{X: 184, Y: 432.900024414062500},
		// WormBody
		{X: 144, Y: 468.960021972656250},
		// WormBody
		{X: 104, Y: 507.660003662109375},
		// WormBody
		{X: 64, Y: 528.000000000000000},
		// WormBody
		{X: 24, Y: 528.000000000000000},
		// WormBody
		{X: -16, Y: 528.000000000000000},
		// WormBody
		{X: -56, Y: 528.000000000000000},
		// WormBody
		{X: -96, Y: 528.000000000000000},
	},
	// Frame 12
	{
		// Worm
		{X: 304, Y: 327.480041503906250},
		// WormBody
		{X: 264, Y: 368.700012207031250},
		// WormBody
		{X: 224, Y: 399.480010986328125},
		// WormBody
		{X: 184, Y: 432.900024414062500},
		// WormBody
		{X: 144, Y: 468.960021972656250},
		// WormBody
		{X: 104, Y: 507.660003662109375},
		// WormBody
		{X: 64, Y: 528.000000000000000},
		// WormBody
		{X: 24, Y: 528.000000000000000},
		// WormBody
		{X: -16, Y: 528.000000000000000},
		// WormBody
		{X: -56, Y: 528.000000000000000},
		// WormBody
		{X: -96, Y: 528.000000000000000},
	},
	// Frame 13
	{
		// Worm
		{X: 304, Y: 315.060028076171875},
		// WormBody
		{X: 264, Y: 340.560028076171875},
		// WormBody
		{X: 224, Y: 368.700012207031250},
		// WormBody
		{X: 184, Y: 399.480010986328125},
		// WormBody
		{X: 144, Y: 432.900024414062500},
		// WormBody
		{X: 104, Y: 468.960021972656250},
		// WormBody
		{X: 64, Y: 507.660003662109375},
		// WormBody
		{X: 24, Y: 528.000000000000000},
		// WormBody
		{X: -16, Y: 528.000000000000000},
		// WormBody
		{X: -56, Y: 528.000000000000000},
		// WormBody
		{X: -96, Y: 528.000000000000000},
	},
	// Frame 14
	{
		// Worm
		{X: 304, Y: 303.300018310546875},
		// WormBody
		{X: 264, Y: 340.560028076171875},
		// WormBody
		{X: 224, Y: 368.700012207031250},
		// WormBody
		{X: 184, Y: 399.480010986328125},
		// WormBody
		{X: 144, Y: 432.900024414062500},
		// WormBody
		{X: 104, Y: 468.960021972656250},
		// WormBody
		{X: 64, Y: 507.660003662109375},
		// WormBody
		{X: 24, Y: 528.000000000000000},
		// WormBody
		{X: -16, Y: 528.000000000000000},
		// WormBody
		{X: -56, Y: 528.000000000000000},
		// WormBody
		{X: -96, Y: 528.000000000000000},
	},
	// Frame 15
	{
		// Worm
		{X: 304, Y: 292.200012207031250},
		// WormBody
		{X: 264, Y: 315.060028076171875},
		// WormBody
		{X: 224, Y: 340.560028076171875},
		// WormBody
		{X: 184, Y: 368.700012207031250},
		// WormBody
		{X: 144, Y: 399.480010986328125},
		// WormBody
		{X: 104, Y: 432.900024414062500},
		// WormBody
		{X: 64, Y: 468.960021972656250},
		// WormBody
		{X: 24, Y: 507.660003662109375},
		// WormBody
		{X: -16, Y: 528.000000000000000},
		// WormBody
		{X: -56, Y: 528.000000000000000},
		// WormBody
		{X: -96, Y: 528.000000000000000},
	},
	// Frame 16
	{
		// Worm
		{X: 304, Y: 281.760009765625000},
		// WormBody
		{X: 264, Y: 315.060028076171875},
		// WormBody
		{X: 224, Y: 340.560028076171875},
		// WormBody
		{X: 184, Y: 368.700012207031250},
		// WormBody
		{X: 144, Y: 399.480010986328125},
		// WormBody
		{X: 104, Y: 432.900024414062500},
		// WormBody
		{X: 64, Y: 468.960021972656250},
		// WormBody
		{X: 24, Y: 507.660003662109375},
		// WormBody
		{X: -16, Y: 528.000000000000000},
		// WormBody
		{X: -56, Y: 528.000000000000000},
		// WormBody
		{X: -96, Y: 528.000000000000000},
	},
	// Frame 17
	{
		// Worm
		{X: 304, Y: 271.980010986328125},
		// WormBody
		{X: 264, Y: 292.200012207031250},
		// WormBody
		{X: 224, Y: 315.060028076171875},
		// WormBody
		{X: 184, Y: 340.560028076171875},
		// WormBody
		{X: 144, Y: 368.700012207031250},
		// WormBody
		{X: 104, Y: 399.480010986328125},
		// WormBody
		{X: 64, Y: 432.900024414062500},
		// WormBody
		{X: 24, Y: 468.960021972656250},
		// WormBody
		{X: -16, Y: 507.660003662109375},
		// WormBody
		{X: -56, Y: 528.000000000000000},
		// WormBody
		{X: -96, Y: 528.000000000000000},
	},
	// Frame 18
	{
		// Worm
		{X: 304, Y: 262.860015869140625},
		// WormBody
		{X: 264, Y: 292.200012207031250},
		// WormBody
		{X: 224, Y: 315.060028076171875},
		// WormBody
		{X: 184, Y: 340.560028076171875},
		// WormBody
		{X: 144, Y: 368.700012207031250},
		// WormBody
		{X: 104, Y: 399.480010986328125},
		// WormBody
		{X: 64, Y: 432.900024414062500},
		// WormBody
		{X: 24, Y: 468.960021972656250},
		// WormBody
		{X: -16, Y: 507.660003662109375},
		// WormBody
		{X: -56, Y: 528.000000000000000},
		// WormBody
		{X: -96, Y: 528.000000000000000},
	},
	// Frame 19
	{
		// Worm
		{X: 304, Y: 254.400009155273438},
		// WormBody
		{X: 264, Y: 271.980010986328125},
		// WormBody
		{X: 224, Y: 292.200012207031250},
		// WormBody
		{X: 184, Y: 315.060028076171875},
		// WormBody
		{X: 144, Y: 340.560028076171875},
		// WormBody
		{X: 104, Y: 368.700012207031250},
		// WormBody
		{X: 64, Y: 399.480010986328125},
		// WormBody
		{X: 24, Y: 432.900024414062500},
		// WormBody
		{X: -16, Y: 468.960021972656250},
		// WormBody
		{X: -56, Y: 507.660003662109375},
		// WormBody
		{X: -96, Y: 528.000000000000000},
	},
	// Frame 20
	{
		// Worm
		{X: 304, Y: 246.600006103515625},
		// WormBody
		{X: 264, Y: 271.980010986328125},
		// WormBody
		{X: 224, Y: 292.200012207031250},
		// WormBody
		{X: 184, Y: 315.060028076171875},
		// WormBody
		{X: 144, Y: 340.560028076171875},
		// WormBody
		{X: 104, Y: 368.700012207031250},
		// WormBody
		{X: 64, Y: 399.480010986328125},
		// WormBody
		{X: 24, Y: 432.900024414062500},
		// WormBody
		{X: -16, Y: 468.960021972656250},
		// WormBody
		{X: -56, Y: 507.660003662109375},
		// WormBody
		{X: -96, Y: 528.000000000000000},
	},
	// Frame 21
	{
		// Worm
		{X: 304, Y: 239.460006713867188},
		// WormBody
		{X: 264, Y: 254.400009155273438},
		// WormBody
		{X: 224, Y: 271.980010986328125},
		// WormBody
		{X: 184, Y: 292.200012207031250},
		// WormBody
		{X: 144, Y: 315.060028076171875},
		// WormBody
		{X: 104, Y: 340.560028076171875},
		// WormBody
		{X: 64, Y: 368.700012207031250},
		// WormBody
		{X: 24, Y: 399.480010986328125},
		// WormBody
		{X: -16, Y: 432.900024414062500},
		// WormBody
		{X: -56, Y: 468.960021972656250},
		// WormBody
		{X: -96, Y: 507.660003662109375},
	},
	// Frame 22
	{
		// Worm
		{X: 304, Y: 232.980010986328125},
		// WormBody
		{X: 264, Y: 254.400009155273438},
		// WormBody
		{X: 224, Y: 271.980010986328125},
		// WormBody
		{X: 184, Y: 292.200012207031250},
		// WormBody
		{X: 144, Y: 315.060028076171875},
		// WormBody
		{X: 104, Y: 340.560028076171875},
		// WormBody
		{X: 64, Y: 368.700012207031250},
		// WormBody
		{X: 24, Y: 399.480010986328125},
		// WormBody
		{X: -16, Y: 432.900024414062500},
		// WormBody
		{X: -56, Y: 468.960021972656250},
		// WormBody
		{X: -96, Y: 507.660003662109375},
	},
	// Frame 23
	{
		// Worm
		{X: 304, Y: 227.160003662109375},
		// WormBody
		{X: 264, Y: 239.460006713867188},
		// WormBody
		{X: 224, Y: 254.400009155273438},
		// WormBody
		{X: 184, Y: 271.980010986328125},
		// WormBody
		{X: 144, Y: 292.200012207031250},
		// WormBody
		{X: 104, Y: 315.060028076171875},
		// WormBody
		{X: 64, Y: 340.560028076171875},
		// WormBody
		{X: 24, Y: 368.700012207031250},
		// WormBody
		{X: -16, Y: 399.480010986328125},
		// WormBody
		{X: -56, Y: 432.900024414062500},
		// WormBody
		{X: -96, Y: 468.960021972656250},
	},
	// Frame 24
	{
		// Worm
		{X: 304, Y: 222.000000000000000},
		// WormBody
		{X: 264, Y: 239.460006713867188},
		// WormBody
		{X: 224, Y: 254.400009155273438},
		// WormBody
		{X: 184, Y: 271.980010986328125},
		// WormBody
		{X: 144, Y: 292.200012207031250},
		// WormBody
		{X: 104, Y: 315.060028076171875},
		// WormBody
		{X: 64, Y: 340.560028076171875},
		// WormBody
		{X: 24, Y: 368.700012207031250},
		// WormBody
		{X: -16, Y: 399.480010986328125},
		// WormBody
		{X: -56, Y: 432.900024414062500},
		// WormBody
		{X: -96, Y: 468.960021972656250},
	},
	// Frame 25
	{
		// Worm
		{X: 304, Y: 217.500000000000000},
		// WormBody
		{X: 264, Y: 227.160003662109375},
		// WormBody
		{X: 224, Y: 239.460006713867188},
		// WormBody
		{X: 184, Y: 254.400009155273438},
		// WormBody
		{X: 144, Y: 271.980010986328125},
		// WormBody
		{X: 104, Y: 292.200012207031250},
		// WormBody
		{X: 64, Y: 315.060028076171875},
		// WormBody
		{X: 24, Y: 340.560028076171875},
		// WormBody
		{X: -16, Y: 368.700012207031250},
		// WormBody
		{X: -56, Y: 399.480010986328125},
		// WormBody
		{X: -96, Y: 432.900024414062500},
	},
	// Frame 26
	{
		// Worm
		{X: 304, Y: 213.660003662109375},
		// WormBody
		{X: 264, Y: 227.160003662109375},
		// WormBody
		{X: 224, Y: 239.460006713867188},
		// WormBody
		{X: 184, Y: 254.400009155273438},
		// WormBody
		{X: 144, Y: 271.980010986328125},
		// WormBody
		{X: 104, Y: 292.200012207031250},
		// WormBody
		{X: 64, Y: 315.060028076171875},
		// WormBody
		{X: 24, Y: 340.560028076171875},
		// WormBody
		{X: -16, Y: 368.700012207031250},
		// WormBody
		{X: -56, Y: 399.480010986328125},
		// WormBody
		{X: -96, Y: 432.900024414062500},
	},
	// Frame 27
	{
		// Worm
		{X: 304, Y: 210.479995727539063},
		// WormBody
		{X: 264, Y: 217.500000000000000},
		// WormBody
		{X: 224, Y: 227.160003662109375},
		// WormBody
		{X: 184, Y: 239.460006713867188},
		// WormBody
		{X: 144, Y: 254.400009155273438},
		// WormBody
		{X: 104, Y: 271.980010986328125},
		// WormBody
		{X: 64, Y: 292.200012207031250},
		// WormBody
		{X: 24, Y: 315.060028076171875},
		// WormBody
		{X: -16, Y: 340.560028076171875},
		// WormBody
		{X: -56, Y: 368.700012207031250},
		// WormBody
		{X: -96, Y: 399.480010986328125},
	},
	// Frame 28
	{
		// Worm
		{X: 304, Y: 207.959991455078125},
		// WormBody
		{X: 264, Y: 217.500000000000000},
		// WormBody
		{X: 224, Y: 227.160003662109375},
		// WormBody
		{X: 184, Y: 239.460006713867188},
		// WormBody
		{X: 144, Y: 254.400009155273438},
		// WormBody
		{X: 104, Y: 271.980010986328125},
		// WormBody
		{X: 64, Y: 292.200012207031250},
		// WormBody
		{X: 24, Y: 315.060028076171875},
		// WormBody
		{X: -16, Y: 340.560028076171875},
		// WormBody
		{X: -56, Y: 368.700012207031250},
		// WormBody
		{X: -96, Y: 399.480010986328125},
	},
	// Frame 29
	{
		// Worm
		{X: 304, Y: 206.099990844726563},
		// WormBody
		{X: 264, Y: 210.479995727539063},
		// WormBody
		{X: 224, Y: 217.500000000000000},
		// WormBody
		{X: 184, Y: 227.160003662109375},
		// WormBody
		{X: 144, Y: 239.460006713867188},
		// WormBody
		{X: 104, Y: 254.400009155273438},
		// WormBody
		{X: 64, Y: 271.980010986328125},
		// WormBody
		{X: 24, Y: 292.200012207031250},
		// WormBody
		{X: -16, Y: 315.060028076171875},
		// WormBody
		{X: -56, Y: 340.560028076171875},
		// WormBody
		{X: -96, Y: 368.700012207031250},
	},
	// Frame 30
	{
		// Worm
		{X: 304, Y: 204.899993896484375},
		// WormBody
		{X: 264, Y: 210.479995727539063},
		// WormBody
		{X: 224, Y: 217.500000000000000},
		// WormBody
		{X: 184, Y: 227.160003662109375},
		// WormBody
		{X: 144, Y: 239.460006713867188},
		// WormBody
		{X: 104, Y: 254.400009155273438},
		// WormBody
		{X: 64, Y: 271.980010986328125},
		// WormBody
		{X: 24, Y: 292.200012207031250},
		// WormBody
		{X: -16, Y: 315.060028076171875},
		// WormBody
		{X: -56, Y: 340.560028076171875},
		// WormBody
		{X: -96, Y: 368.700012207031250},
	},
	// Frame 31
	{
		// Worm
		{X: 304, Y: 204.359985351562500},
		// WormBody
		{X: 264, Y: 206.099990844726563},
		// WormBody
		{X: 224, Y: 210.479995727539063},
		// WormBody
		{X: 184, Y: 217.500000000000000},
		// WormBody
		{X: 144, Y: 227.160003662109375},
		// WormBody
		{X: 104, Y: 239.460006713867188},
		// WormBody
		{X: 64, Y: 254.400009155273438},
		// WormBody
		{X: 24, Y: 271.980010986328125},
		// WormBody
		{X: -16, Y: 292.200012207031250},
		// WormBody
		{X: -56, Y: 315.060028076171875},
		// WormBody
		{X: -96, Y: 340.560028076171875},
	},
	// Frame 32
	{
		// Worm
		{X: 304, Y: 204.479980468750000},
		// WormBody
		{X: 264, Y: 206.099990844726563},
		// WormBody
		{X: 224, Y: 210.479995727539063},
		// WormBody
		{X: 184, Y: 217.500000000000000},
		// WormBody
		{X: 144, Y: 227.160003662109375},
		// WormBody
		{X: 104, Y: 239.460006713867188},
		// WormBody
		{X: 64, Y: 254.400009155273438},
		// WormBody
		{X: 24, Y: 271.980010986328125},
		// WormBody
		{X: -16, Y: 292.200012207031250},
		// WormBody
		{X: -56, Y: 315.060028076171875},
		// WormBody
		{X: -96, Y: 340.560028076171875},
	},
	// Frame 33
	{
		// Worm
		{X: 304, Y: 205.159973144531250},
		// WormBody
		{X: 264, Y: 204.359985351562500},
		// WormBody
		{X: 224, Y: 206.099990844726563},
		// WormBody
		{X: 184, Y: 210.479995727539063},
		// WormBody
		{X: 144, Y: 217.500000000000000},
		// WormBody
		{X: 104, Y: 227.160003662109375},
		// WormBody
		{X: 64, Y: 239.460006713867188},
		// WormBody
		{X: 24, Y: 254.400009155273438},
		// WormBody
		{X: -16, Y: 271.980010986328125},
		// WormBody
		{X: -56, Y: 292.200012207031250},
		// WormBody
		{X: -96, Y: 315.060028076171875},
	},
	// Frame 34
	{
		// Worm
		{X: 304, Y: 206.399963378906250},
		// WormBody
		{X: 264, Y: 204.359985351562500},
		// WormBody
		{X: 224, Y: 206.099990844726563},
		// WormBody
		{X: 184, Y: 210.479995727539063},
		// WormBody
		{X: 144, Y: 217.500000000000000},
		// WormBody
		{X: 104, Y: 227.160003662109375},
		// WormBody
		{X: 64, Y: 239.460006713867188},
		// WormBody
		{X: 24, Y: 254.400009155273438},
		// WormBody
		{X: -16, Y: 271.980010986328125},
		// WormBody
		{X: -56, Y: 292.200012207031250},
		// WormBody
		{X: -96, Y: 315.060028076171875},
	},
	// Frame 35
	{
		// Worm
		{X: 304, Y: 208.199966430664063},
		// WormBody
		{X: 264, Y: 205.159973144531250},
		// WormBody
		{X: 224, Y: 204.359985351562500},
		// WormBody
		{X: 184, Y: 206.099990844726563},
		// WormBody
		{X: 144, Y: 210.479995727539063},
		// WormBody
		{X: 104, Y: 217.500000000000000},
		// WormBody
		{X: 64, Y: 227.160003662109375},
		// WormBody
		{X: 24, Y: 239.460006713867188},
		// WormBody
		{X: -16, Y: 254.400009155273438},
		// WormBody
		{X: -56, Y: 271.980010986328125},
		// WormBody
		{X: -96, Y: 292.200012207031250},
	},
	// Frame 36
	{
		// Worm
		{X: 304, Y: 210.559967041015625},
		// WormBody
		{X: 264, Y: 205.159973144531250},
		// WormBody
		{X: 224, Y: 204.359985351562500},
		// WormBody
		{X: 184, Y: 206.099990844726563},
		// WormBody
		{X: 144, Y: 210.479995727539063},
		// WormBody
		{X: 104, Y: 217.500000000000000},
		// WormBody
		{X: 64, Y: 227.160003662109375},
		// WormBody
		{X: 24, Y: 239.460006713867188},
		// WormBody
		{X: -16, Y: 254.400009155273438},
		// WormBody
		{X: -56, Y: 271.980010986328125},
		// WormBody
		{X: -96, Y: 292.200012207031250},
	},
	// Frame 37
	{
		// Worm
		{X: 304, Y: 213.479965209960938},
		// WormBody
		{X: 264, Y: 208.199966430664063},
		// WormBody
		{X: 224, Y: 205.159973144531250},
		// WormBody
		{X: 184, Y: 204.359985351562500},
		// WormBody
		{X: 144, Y: 206.099990844726563},
		// WormBody
		{X: 104, Y: 210.479995727539063},
		// WormBody
		{X: 64, Y: 217.500000000000000},
		// WormBody
		{X: 24, Y: 227.160003662109375},
		// WormBody
		{X: -16, Y: 239.460006713867188},
		// WormBody
		{X: -56, Y: 254.400009155273438},
		// WormBody
		{X: -96, Y: 271.980010986328125},
	},
	// Frame 38
	{
		// Worm
		{X: 304, Y: 216.959960937500000},
		// WormBody
		{X: 264, Y: 208.199966430664063},
		// WormBody
		{X: 224, Y: 205.159973144531250},
		// WormBody
		{X: 184, Y: 204.359985351562500},
		// WormBody
		{X: 144, Y: 206.099990844726563},
		// WormBody
		{X: 104, Y: 210.479995727539063},
		// WormBody
		{X: 64, Y: 217.500000000000000},
		// WormBody
		{X: 24, Y: 227.160003662109375},
		// WormBody
		{X: -16, Y: 239.460006713867188},
		// WormBody
		{X: -56, Y: 254.400009155273438},
		// WormBody
		{X: -96, Y: 271.980010986328125},
	},
	// Frame 39
	{
		// Worm
		{X: 304, Y: 220.999954223632813},
		// WormBody
		{X: 264, Y: 213.479965209960938},
		// WormBody
		{X: 224, Y: 208.199966430664063},
		// WormBody
		{X: 184, Y: 205.159973144531250},
		// WormBody
		{X: 144, Y: 204.359985351562500},
		// WormBody
		{X: 104, Y: 206.099990844726563},
		// WormBody
		{X: 64, Y: 210.479995727539063},
		// WormBody
		{X: 24, Y: 217.500000000000000},
		// WormBody
		{X: -16, Y: 227.160003662109375},
		// WormBody
		{X: -56, Y: 239.460006713867188},
		// WormBody
		{X: -96, Y: 254.400009155273438},
	},
	// Frame 40
	{
		// Worm
		{X: 304, Y: 225.599945068359375},
		// WormBody
		{X: 264, Y: 213.479965209960938},
		// WormBody
		{X: 224, Y: 208.199966430664063},
		// WormBody
		{X: 184, Y: 205.159973144531250},
		// WormBody
		{X: 144, Y: 204.359985351562500},
		// WormBody
		{X: 104, Y: 206.099990844726563},
		// WormBody
		{X: 64, Y: 210.479995727539063},
		// WormBody
		{X: 24, Y: 217.500000000000000},
		// WormBody
		{X: -16, Y: 227.160003662109375},
		// WormBody
		{X: -56, Y: 239.460006713867188},
		// WormBody
		{X: -96, Y: 254.400009155273438},
	},
	// Frame 41
	{
		// Worm
		{X: 304, Y: 230.759948730468750},
		// WormBody
		{X: 264, Y: 220.999954223632813},
		// WormBody
		{X: 224, Y: 213.479965209960938},
		// WormBody
		{X: 184, Y: 208.199966430664063},
		// WormBody
		{X: 144, Y: 205.159973144531250},
		// WormBody
		{X: 104, Y: 204.359985351562500},
		// WormBody
		{X: 64, Y: 206.099990844726563},
		// WormBody
		{X: 24, Y: 210.479995727539063},
		// WormBody
		{X: -16, Y: 217.500000000000000},
		// WormBody
		{X: -56, Y: 227.160003662109375},
		// WormBody
		{X: -96, Y: 239.460006713867188},
	},
	// Frame 42
	{
		// Worm
		{X: 304, Y: 236.479949951171875},
		// WormBody
		{X: 264, Y: 220.999954223632813},
		// WormBody
		{X: 224, Y: 213.479965209960938},
		// WormBody
		{X: 184, Y: 208.199966430664063},
		// WormBody
		{X: 144, Y: 205.159973144531250},
		// WormBody
		{X: 104, Y: 204.359985351562500},
		// WormBody
		{X: 64, Y: 206.099990844726563},
		// WormBody
		{X: 24, Y: 210.479995727539063},
		// WormBody
		{X: -16, Y: 217.500000000000000},
		// WormBody
		{X: -56, Y: 227.160003662109375},
		// WormBody
		{X: -96, Y: 239.460006713867188},
	},
	// Frame 43
	{
		// Worm
		{X: 304, Y: 242.759948730468750},
		// WormBody
		{X: 264, Y: 230.759948730468750},
		// WormBody
		{X: 224, Y: 220.999954223632813},
		// WormBody
		{X: 184, Y: 213.479965209960938},
		// WormBody
		{X: 144, Y: 208.199966430664063},
		// WormBody
		{X: 104, Y: 205.159973144531250},
		// WormBody
		{X: 64, Y: 204.359985351562500},
		// WormBody
		{X: 24, Y: 206.099990844726563},
		// WormBody
		{X: -16, Y: 210.479995727539063},
		// WormBody
		{X: -56, Y: 217.500000000000000},
		// WormBody
		{X: -96, Y: 227.160003662109375},
	},
	// Frame 44
	{
		// Worm
		{X: 304, Y: 249.599945068359375},
		// WormBody
		{X: 264, Y: 230.759948730468750},
		// WormBody
		{X: 224, Y: 220.999954223632813},
		// WormBody
		{X: 184, Y: 213.479965209960938},
		// WormBody
		{X: 144, Y: 208.199966430664063},
		// WormBody
		{X: 104, Y: 205.159973144531250},
		// WormBody
		{X: 64, Y: 204.359985351562500},
		// WormBody
		{X: 24, Y: 206.099990844726563},
		// WormBody
		{X: -16, Y: 210.479995727539063},
		// WormBody
		{X: -56, Y: 217.500000000000000},
		// WormBody
		{X: -96, Y: 227.160003662109375},
	},
	// Frame 45
	{
		// Worm
		{X: 304, Y: 256.999938964843750},
		// WormBody
		{X: 264, Y: 242.759948730468750},
		// WormBody
		{X: 224, Y: 230.759948730468750},
		// WormBody
		{X: 184, Y: 220.999954223632813},
		// WormBody
		{X: 144, Y: 213.479965209960938},
		// WormBody
		{X: 104, Y: 208.199966430664063},
		// WormBody
		{X: 64, Y: 205.159973144531250},
		// WormBody
		{X: 24, Y: 204.359985351562500},
		// WormBody
		{X: -16, Y: 206.099990844726563},
		// WormBody
		{X: -56, Y: 210.479995727539063},
		// WormBody
		{X: -96, Y: 217.500000000000000},
	},
	// Frame 46
	{
		// Worm
		{X: 304, Y: 264.959930419921875},
		// WormBody
		{X: 264, Y: 242.759948730468750},
		// WormBody
		{X: 224, Y: 230.759948730468750},
		// WormBody
		{X: 184, Y: 220.999954223632813},
		// WormBody
		{X: 144, Y: 213.479965209960938},
		// WormBody
		{X: 104, Y: 208.199966430664063},
		// WormBody
		{X: 64, Y: 205.159973144531250},
		// WormBody
		{X: 24, Y: 204.359985351562500},
		// WormBody
		{X: -16, Y: 206.099990844726563},
		// WormBody
		{X: -56, Y: 210.479995727539063},
		// WormBody
		{X: -96, Y: 217.500000000000000},
	},
	// Frame 47
	{
		// Worm
		{X: 304, Y: 273.479919433593750},
		// WormBody
		{X: 264, Y: 256.999938964843750},
		// WormBody
		{X: 224, Y: 242.759948730468750},
		// WormBody
		{X: 184, Y: 230.759948730468750},
		// WormBody
		{X: 144, Y: 220.999954223632813},
		// WormBody
		{X: 104, Y: 213.479965209960938},
		// WormBody
		{X: 64, Y: 208.199966430664063},
		// WormBody
		{X: 24, Y: 205.159973144531250},
		// WormBody
		{X: -16, Y: 204.359985351562500},
		// WormBody
		{X: -56, Y: 206.099990844726563},
		// WormBody
		{X: -96, Y: 210.479995727539063},
	},
	// Frame 48
	{
		// Worm
		{X: 304, Y: 282.559906005859375},
		// WormBody
		{X: 264, Y: 256.999938964843750},
		// WormBody
		{X: 224, Y: 242.759948730468750},
		// WormBody
		{X: 184, Y: 230.759948730468750},
		// WormBody
		{X: 144, Y: 220.999954223632813},
		// WormBody
		{X: 104, Y: 213.479965209960938},
		// WormBody
		{X: 64, Y: 208.199966430664063},
		// WormBody
		{X: 24, Y: 205.159973144531250},
		// WormBody
		{X: -16, Y: 204.359985351562500},
		// WormBody
		{X: -56, Y: 206.099990844726563},
		// WormBody
		{X: -96, Y: 210.479995727539063},
	},
	// Frame 49
	{
		// Worm
		{X: 304, Y: 292.199890136718750},
		// WormBody
		{X: 264, Y: 273.479919433593750},
		// WormBody
		{X: 224, Y: 256.999938964843750},
		// WormBody
		{X: 184, Y: 242.759948730468750},
		// WormBody
		{X: 144, Y: 230.759948730468750},
		// WormBody
		{X: 104, Y: 220.999954223632813},
		// WormBody
		{X: 64, Y: 213.479965209960938},
		// WormBody
		{X: 24, Y: 208.199966430664063},
		// WormBody
		{X: -16, Y: 205.159973144531250},
		// WormBody
		{X: -56, Y: 204.359985351562500},
		// WormBody
		{X: -96, Y: 206.099990844726563},
	},
	// Frame 50
	{
		// Worm
		{X: 304, Y: 302.399902343750000},
		// WormBody
		{X: 264, Y: 273.479919433593750},
		// WormBody
		{X: 224, Y: 256.999938964843750},
		// WormBody
		{X: 184, Y: 242.759948730468750},
		// WormBody
		{X: 144, Y: 230.759948730468750},
		// WormBody
		{X: 104, Y: 220.999954223632813},
		// WormBody
		{X: 64, Y: 213.479965209960938},
		// WormBody
		{X: 24, Y: 208.199966430664063},
		// WormBody
		{X: -16, Y: 205.159973144531250},
		// WormBody
		{X: -56, Y: 204.359985351562500},
		// WormBody
		{X: -96, Y: 206.099990844726563},
	},
	// Frame 51
	{
		// Worm
		{X: 304, Y: 313.159912109375000},
		// WormBody
		{X: 264, Y: 292.199890136718750},
		// WormBody
		{X: 224, Y: 273.479919433593750},
		// WormBody
		{X: 184, Y: 256.999938964843750},
		// WormBody
		{X: 144, Y: 242.759948730468750},
		// WormBody
		{X: 104, Y: 230.759948730468750},
		// WormBody
		{X: 64, Y: 220.999954223632813},
		// WormBody
		{X: 24, Y: 213.479965209960938},
		// WormBody
		{X: -16, Y: 208.199966430664063},
		// WormBody
		{X: -56, Y: 205.159973144531250},
		// WormBody
		{X: -96, Y: 204.359985351562500},
	},
	// Frame 52
	{
		// Worm
		{X: 304, Y: 324.479919433593750},
		// WormBody
		{X: 264, Y: 292.199890136718750},
		// WormBody
		{X: 224, Y: 273.479919433593750},
		// WormBody
		{X: 184, Y: 256.999938964843750},
		// WormBody
		{X: 144, Y: 242.759948730468750},
		// WormBody
		{X: 104, Y: 230.759948730468750},
		// WormBody
		{X: 64, Y: 220.999954223632813},
		// WormBody
		{X: 24, Y: 213.479965209960938},
		// WormBody
		{X: -16, Y: 208.199966430664063},
		// WormBody
		{X: -56, Y: 205.159973144531250},
		// WormBody
		{X: -96, Y: 204.359985351562500},
	},
	// Frame 53
	{
		// Worm
		{X: 304, Y: 336.359924316406250},
		// WormBody
		{X: 264, Y: 313.159912109375000},
		// WormBody
		{X: 224, Y: 292.199890136718750},
		// WormBody
		{X: 184, Y: 273.479919433593750},
		// WormBody
		{X: 144, Y: 256.999938964843750},
		// WormBody
		{X: 104, Y: 242.759948730468750},
		// WormBody
		{X: 64, Y: 230.759948730468750},
		// WormBody
		{X: 24, Y: 220.999954223632813},
		// WormBody
		{X: -16, Y: 213.479965209960938},
		// WormBody
		{X: -56, Y: 208.199966430664063},
		// WormBody
		{X: -96, Y: 205.159973144531250},
	},
	// Frame 54
	{
		// Worm
		{X: 304, Y: 348.799926757812500},
		// WormBody
		{X: 264, Y: 313.159912109375000},
		// WormBody
		{X: 224, Y: 292.199890136718750},
		// WormBody
		{X: 184, Y: 273.479919433593750},
		// WormBody
		{X: 144, Y: 256.999938964843750},
		// WormBody
		{X: 104, Y: 242.759948730468750},
		// WormBody
		{X: 64, Y: 230.759948730468750},
		// WormBody
		{X: 24, Y: 220.999954223632813},
		// WormBody
		{X: -16, Y: 213.479965209960938},
		// WormBody
		{X: -56, Y: 208.199966430664063},
		// WormBody
		{X: -96, Y: 205.159973144531250},
	},
	// Frame 55
	{
		// Worm
		{X: 304, Y: 361.799926757812500},
		// WormBody
		{X: 264, Y: 336.359924316406250},
		// WormBody
		{X: 224, Y: 313.159912109375000},
		// WormBody
		{X: 184, Y: 292.199890136718750},
		// WormBody
		{X: 144, Y: 273.479919433593750},
		// WormBody
		{X: 104, Y: 256.999938964843750},
		// WormBody
		{X: 64, Y: 242.759948730468750},
		// WormBody
		{X: 24, Y: 230.759948730468750},
		// WormBody
		{X: -16, Y: 220.999954223632813},
		// WormBody
		{X: -56, Y: 213.479965209960938},
		// WormBody
		{X: -96, Y: 208.199966430664063},
	},
	// Frame 56
	{
		// Worm
		{X: 304, Y: 375.359924316406250},
		// WormBody
		{X: 264, Y: 336.359924316406250},
		// WormBody
		{X: 224, Y: 313.159912109375000},
		// WormBody
		{X: 184, Y: 292.199890136718750},
		// WormBody
		{X: 144, Y: 273.479919433593750},
		// WormBody
		{X: 104, Y: 256.999938964843750},
		// WormBody
		{X: 64, Y: 242.759948730468750},
		// WormBody
		{X: 24, Y: 230.759948730468750},
		// WormBody
		{X: -16, Y: 220.999954223632813},
		// WormBody
		{X: -56, Y: 213.479965209960938},
		// WormBody
		{X: -96, Y: 208.199966430664063},
	},
	// Frame 57
	{
		// Worm
		{X: 304, Y: 389.479919433593750},
		// WormBody
		{X: 264, Y: 361.799926757812500},
		// WormBody
		{X: 224, Y: 336.359924316406250},
		// WormBody
		{X: 184, Y: 313.159912109375000},
		// WormBody
		{X: 144, Y: 292.199890136718750},
		// WormBody
		{X: 104, Y: 273.479919433593750},
		// WormBody
		{X: 64, Y: 256.999938964843750},
		// WormBody
		{X: 24, Y: 242.759948730468750},
		// WormBody
		{X: -16, Y: 230.759948730468750},
		// WormBody
		{X: -56, Y: 220.999954223632813},
		// WormBody
		{X: -96, Y: 213.479965209960938},
	},
	// Frame 58
	{
		// Worm
		{X: 304, Y: 404.159912109375000},
		// WormBody
		{X: 264, Y: 361.799926757812500},
		// WormBody
		{X: 224, Y: 336.359924316406250},
		// WormBody
		{X: 184, Y: 313.159912109375000},
		// WormBody
		{X: 144, Y: 292.199890136718750},
		// WormBody
		{X: 104, Y: 273.479919433593750},
		// WormBody
		{X: 64, Y: 256.999938964843750},
		// WormBody
		{X: 24, Y: 242.759948730468750},
		// WormBody
		{X: -16, Y: 230.759948730468750},
		// WormBody
		{X: -56, Y: 220.999954223632813},
		// WormBody
		{X: -96, Y: 213.479965209960938},
	},
	// Frame 59
	{
		// Worm
		{X: 304, Y: 419.399902343750000},
		// WormBody
		{X: 264, Y: 389.479919433593750},
		// WormBody
		{X: 224, Y: 361.799926757812500},
		// WormBody
		{X: 184, Y: 336.359924316406250},
		// WormBody
		{X: 144, Y: 313.159912109375000},
		// WormBody
		{X: 104, Y: 292.199890136718750},
		// WormBody
		{X: 64, Y: 273.479919433593750},
		// WormBody
		{X: 24, Y: 256.999938964843750},
		// WormBody
		{X: -16, Y: 242.759948730468750},
		// WormBody
		{X: -56, Y: 230.759948730468750},
		// WormBody
		{X: -96, Y: 220.999954223632813},
	},
	// Frame 60
	{
		// Worm
		{X: 304, Y: 435.199890136718750},
		// WormBody
		{X: 264, Y: 389.479919433593750},
		// WormBody
		{X: 224, Y: 361.799926757812500},
		// WormBody
		{X: 184, Y: 336.359924316406250},
		// WormBody
		{X: 144, Y: 313.159912109375000},
		// WormBody
		{X: 104, Y: 292.199890136718750},
		// WormBody
		{X: 64, Y: 273.479919433593750},
		// WormBody
		{X: 24, Y: 256.999938964843750},
		// WormBody
		{X: -16, Y: 242.759948730468750},
		// WormBody
		{X: -56, Y: 230.759948730468750},
		// WormBody
		{X: -96, Y: 220.999954223632813},
	},
	// Frame 61
	{
		// Worm
		{X: 304, Y: 451.559906005859375},
		// WormBody
		{X: 264, Y: 419.399902343750000},
		// WormBody
		{X: 224, Y: 389.479919433593750},
		// WormBody
		{X: 184, Y: 361.799926757812500},
		// WormBody
		{X: 144, Y: 336.359924316406250},
		// WormBody
		{X: 104, Y: 313.159912109375000},
		// WormBody
		{X: 64, Y: 292.199890136718750},
		// WormBody
		{X: 24, Y: 273.479919433593750},
		// WormBody
		{X: -16, Y: 256.999938964843750},
		// WormBody
		{X: -56, Y: 242.759948730468750},
		// WormBody
		{X: -96, Y: 230.759948730468750},
	},
	// Frame 62
	{
		// Worm
		{X: 304, Y: 468.479919433593750},
		// WormBody
		{X: 264, Y: 419.399902343750000},
		// WormBody
		{X: 224, Y: 389.479919433593750},
		// WormBody
		{X: 184, Y: 361.799926757812500},
		// WormBody
		{X: 144, Y: 336.359924316406250},
		// WormBody
		{X: 104, Y: 313.159912109375000},
		// WormBody
		{X: 64, Y: 292.199890136718750},
		// WormBody
		{X: 24, Y: 273.479919433593750},
		// WormBody
		{X: -16, Y: 256.999938964843750},
		// WormBody
		{X: -56, Y: 242.759948730468750},
		// WormBody
		{X: -96, Y: 230.759948730468750},
	},
	// Frame 63
	{
		// Worm
		{X: 304, Y: 485.959930419921875},
		// WormBody
		{X: 264, Y: 451.559906005859375},
		// WormBody
		{X: 224, Y: 419.399902343750000},
		// WormBody
		{X: 184, Y: 389.479919433593750},
		// WormBody
		{X: 144, Y: 361.799926757812500},
		// WormBody
		{X: 104, Y: 336.359924316406250},
		// WormBody
		{X: 64, Y: 313.159912109375000},
		// WormBody
		{X: 24, Y: 292.199890136718750},
		// WormBody
		{X: -16, Y: 273.479919433593750},
		// WormBody
		{X: -56, Y: 256.999938964843750},
		// WormBody
		{X: -96, Y: 242.759948730468750},
	},
	// Frame 64
	{
		// Worm
		{X: 304, Y: 503.999938964843750},
		// WormBody
		{X: 264, Y: 451.559906005859375},
		// WormBody
		{X: 224, Y: 419.399902343750000},
		// WormBody
		{X: 184, Y: 389.479919433593750},
		// WormBody
		{X: 144, Y: 361.799926757812500},
		// WormBody
		{X: 104, Y: 336.359924316406250},
		// WormBody
		{X: 64, Y: 313.159912109375000},
		// WormBody
		{X: 24, Y: 292.199890136718750},
		// WormBody
		{X: -16, Y: 273.479919433593750},
		// WormBody
		{X: -56, Y: 256.999938964843750},
		// WormBody
		{X: -96, Y: 242.759948730468750},
	},
	// Frame 65
	{
		// Worm
		{X: 304, Y: 522.599914550781250},
		// WormBody
		{X: 264, Y: 485.959930419921875},
		// WormBody
		{X: 224, Y: 451.559906005859375},
		// WormBody
		{X: 184, Y: 419.399902343750000},
		// WormBody
		{X: 144, Y: 389.479919433593750},
		// WormBody
		{X: 104, Y: 361.799926757812500},
		// WormBody
		{X: 64, Y: 336.359924316406250},
		// WormBody
		{X: 24, Y: 313.159912109375000},
		// WormBody
		{X: -16, Y: 292.199890136718750},
		// WormBody
		{X: -56, Y: 273.479919433593750},
		// WormBody
		{X: -96, Y: 256.999938964843750},
	},
	// Frame 66
	{
		// Worm
		{X: 304, Y: 541.759887695312500},
		// WormBody
		{X: 264, Y: 485.959930419921875},
		// WormBody
		{X: 224, Y: 451.559906005859375},
		// WormBody
		{X: 184, Y: 419.399902343750000},
		// WormBody
		{X: 144, Y: 389.479919433593750},
		// WormBody
		{X: 104, Y: 361.799926757812500},
		// WormBody
		{X: 64, Y: 336.359924316406250},
		// WormBody
		{X: 24, Y: 313.159912109375000},
		// WormBody
		{X: -16, Y: 292.199890136718750},
		// WormBody
		{X: -56, Y: 273.479919433593750},
		// WormBody
		{X: -96, Y: 256.999938964843750},
	},
}
