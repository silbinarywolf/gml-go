package gml

type gamepadButton int

type gamepadAxis int

const maxGamepads = 12

const (
	GpFace1      gamepadButton = 1 + iota // Top button 1 (this maps to the "A" on an Xbox 360 controller and the cross on a PS controller)
	GpFace2                               // Top button 2 (this maps to the "B" on an Xbox 360 controller and the circle on a PS controller)
	GpFace3                               // Top button 3 (this maps to the "X" on an Xbox 360 controller and the square on a PS controller)
	GpFace4                               // Top button 4 (this maps to the "Y" on an Xbox 360 controller and the triangle on a PS controller)
	GpShoulderLB                          // Left shoulder button
	GpShoulderLT                          // Left shoulder trigger
	GpShoulderRB                          // Right shoulder button
	GpShoulderRT                          // Right shoulder trigger
	GpSelect
	GpStart
	GpStickLeft  // Left-stick, pressed as a button
	GpStickRight // Right-stick, pressed as a button
	GpPadUp
	GpPadDown
	GpPadLeft
	GpPadRight
	/*GamepadButton0
	GamepadButton1
	GamepadButton2
	GamepadButton3
	GamepadButton4
	GamepadButton5
	GamepadButton6
	GamepadButton7
	GamepadButton8
	GamepadButton9
	GamepadButton10
	GamepadButton11
	GamepadButton12
	GamepadButton13
	GamepadButton14
	GamepadButton15
	GamepadButton16
	GamepadButton17
	GamepadButton18
	GamepadButton19
	GamepadButton20
	GamepadButton21
	GamepadButton22
	GamepadButton23
	GamepadButton24
	GamepadButton25
	GamepadButton26
	GamepadButton27
	GamepadButton28
	GamepadButton29
	GamepadButton30
	GamepadButton31*/
	gpSize
)

const (
	GpAxisLH gamepadAxis = 1 + iota
	GpAxisLV
	gpAxisSize
)
