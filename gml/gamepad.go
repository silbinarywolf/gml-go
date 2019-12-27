package gml

import "strconv"

type GamepadButton int32

type GamepadAxis int32

const maxGamepads = 12

const (
	GpButtonNone GamepadButton = 0
)

// WARNING: Do not reorder or change for backwards compatibility.
// This allows these values to be used when storing gamepad settings to persistent storage.
const (
	GpButton0 GamepadButton = 1 + iota
	GpButton1
	GpButton2
	GpButton3
	GpButton4
	GpButton5
	GpButton6
	GpButton7
	GpButton8
	GpButton9
	GpButton10
	GpButton11
	GpButton12
	GpButton13
	GpButton14
	GpButton15
	GpButton16
	GpButton17
	GpButton18
	GpButton19
	GpButton20
	GpButton21
	GpButton22
	GpButton23
	GpButton24
	GpButton25
	GpButton26
	GpButton27
	GpButton28
	GpButton29
	GpButton30
	GpButton31
	GpSize
	GpShoulderLT // Special handling: Left shoulder trigger
	GpShoulderRT // Special handling: Right shoulder trigger
)

// WARNING: Do not reorder or change for backwards compatibility.
// This allows these values to be used when storing gamepad settings to persistent storage.
const (
	GpFace1      GamepadButton = 1 + iota // Top button 1 (this maps to the "A" on an Xbox 360 controller and the cross on a PS controller)
	GpFace2                               // Top button 2 (this maps to the "B" on an Xbox 360 controller and the circle on a PS controller)
	GpFace3                               // Top button 3 (this maps to the "X" on an Xbox 360 controller and the square on a PS controller)
	GpFace4                               // Top button 4 (this maps to the "Y" on an Xbox 360 controller and the triangle on a PS controller)
	GpShoulderLB                          // Left shoulder button
	GpShoulderRB                          // Right shoulder button
	GpSelect
	GpStart
	GpStickLeft  // Left-stick, pressed as a button
	GpStickRight // Right-stick, pressed as a button
	GpPadUp
	GpPadRight
	GpPadDown
	GpPadLeft
)

const (
	GpAxisLH GamepadAxis = 1 + iota
	GpAxisLV
	gpAxisSize
)

func (button GamepadAxis) String() string {
	return "Axis " + strconv.Itoa(int(button))
}

var gamepadToString = []string{
	GpButton0:  "Button 0",
	GpButton1:  "Button 1",
	GpButton2:  "Button 2",
	GpButton3:  "Button 3",
	GpButton4:  "Button 4",
	GpButton5:  "Button 5",
	GpButton6:  "Button 6",
	GpButton7:  "Button 7",
	GpButton8:  "Button 8",
	GpButton9:  "Button 9",
	GpButton10: "Button 10",
	GpButton11: "Button 11",
	GpButton12: "Button 12",
	GpButton13: "Button 13",
	GpButton14: "Button 14",
	GpButton15: "Button 15",
	GpButton16: "Button 16",
	GpButton17: "Button 17",
	GpButton18: "Button 18",
	GpButton19: "Button 19",
	GpButton20: "Button 20",
	GpButton21: "Button 21",
	GpButton22: "Button 22",
	GpButton23: "Button 23",
	GpButton24: "Button 24",
	GpButton25: "Button 25",
	GpButton26: "Button 26",
	GpButton27: "Button 27",
	GpButton28: "Button 28",
	GpButton29: "Button 29",
	GpButton30: "Button 30",
	GpButton31: "Button 31",
}

func (button GamepadButton) String() string {
	return gamepadToString[button]
}
