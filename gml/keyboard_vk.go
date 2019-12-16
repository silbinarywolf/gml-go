package gml

const (
	VkNoKey       Key = 0 + iota // keycode representing that no key is pressed
	VkAnykey                     // keycode representing that any key is pressed
	VkLeft                       // keycode for left arrow key
	VkRight                      // keycode for right arrow key
	VkUp                         // keycode for up arrow key
	VkDown                       // keycode for down arrow key
	VkEnter                      // enter key
	VkEscape                     // escape key
	VkSpace                      // space key
	VkShift                      // either of the shift keys
	VkControl                    // either of the control keys
	VkAlt                        // alt key
	VkBackspace                  // backspace key
	VkTab                        // tab key
	VkHome                       // home key
	VkEnd                        // end key
	VkDelete                     // delete key
	VkInsert                     // insert key
	VkPageUp                     // pageup key
	VkPageDown                   // pagedown key
	VkPause                      // pause/break key
	VkPrintScreen                // printscreen/sysrq key
	VkF1                         // keycode for the function keys F1 to F12
	VkF2
	VkF3
	VkF4
	VkF5
	VkF6
	VkF7
	VkF8
	VkF9
	VkF10
	VkF11
	VkF12
	Vk0
	Vk1
	Vk2
	Vk3
	Vk4
	Vk5
	Vk6
	Vk7
	Vk8
	Vk9
	VkNumpad0
	VkNumpad1
	VkNumpad2
	VkNumpad3
	VkNumpad4
	VkNumpad5
	VkNumpad6
	VkNumpad7
	VkNumpad8
	VkNumpad9
	VkNumpadEnter
	VkMultiply // multiply key on the numeric keypad
	VkDivide   // divide key on the numeric keypad
	VkAdd      // key on the numeric keypad
	VkSubtract // subtract key on the numeric keypad
	VkDecimal  // decimal dot keys on the numeric keypad
	VkA
	VkB
	VkC
	VkD
	VkE
	VkF
	VkG
	VkH
	VkI
	VkJ
	VkK
	VkL
	VkM
	VkN
	VkO
	VkP
	VkQ
	VkR
	VkS
	VkT
	VkU
	VkV
	VkW
	VkX
	VkY
	VkZ
	VkSize
)

var keyToString = []string{
	VkNoKey:       "No Key",
	VkAnykey:      "Any Key",
	VkLeft:        "Left",
	VkRight:       "Right",
	VkUp:          "Up",
	VkDown:        "Down",
	VkEnter:       "Enter",
	VkEscape:      "Escape",
	VkSpace:       "Space",
	VkShift:       "Shift",
	VkControl:     "Control",      // either of the control keys
	VkAlt:         "Alt",          // alt key
	VkBackspace:   "Backspace",    // backspace key
	VkTab:         "Tab",          // tab key
	VkHome:        "Home",         // home key
	VkEnd:         "End",          // end key
	VkDelete:      "Delete",       // delete key
	VkInsert:      "Insert",       // insert key
	VkPageUp:      "Page Up",      // pageup key
	VkPageDown:    "Page Down",    // pagedown key
	VkPause:       "Pause",        // pause/break key
	VkPrintScreen: "Print Screen", // printscreen/sysrq key
	VkF1:          "F1",           // keycode for the function keys F1 to F12
	VkF2:          "F2",
	VkF3:          "F3",
	VkF4:          "F4",
	VkF5:          "F5",
	VkF6:          "F6",
	VkF7:          "F7",
	VkF8:          "F8",
	VkF9:          "F9",
	VkF10:         "F10",
	VkF11:         "F11",
	VkF12:         "F12",
	Vk0:           "0",
	Vk1:           "1",
	Vk2:           "2",
	Vk3:           "3",
	Vk4:           "4",
	Vk5:           "5",
	Vk6:           "6",
	Vk7:           "7",
	Vk8:           "8",
	Vk9:           "9",
	VkNumpad0:     "Numpad 0",
	VkNumpad1:     "Numpad 1",
	VkNumpad2:     "Numpad 2",
	VkNumpad3:     "Numpad 3",
	VkNumpad4:     "Numpad 4",
	VkNumpad5:     "Numpad 5",
	VkNumpad6:     "Numpad 6",
	VkNumpad7:     "Numpad 7",
	VkNumpad8:     "Numpad 8",
	VkNumpad9:     "Numpad 9",
	VkNumpadEnter: "Numpad Enter",
	VkMultiply:    "Numpad Multiply",
	VkDivide:      "Numpad Divide",
	VkAdd:         "Numpad Add",
	VkSubtract:    "Numpad Subtract",
	VkDecimal:     "Numpad Decimal",
	VkA:           "A",
	VkB:           "B",
	VkC:           "C",
	VkD:           "D",
	VkE:           "E",
	VkF:           "F",
	VkG:           "G",
	VkH:           "H",
	VkI:           "I",
	VkJ:           "J",
	VkK:           "K",
	VkL:           "L",
	VkM:           "M",
	VkN:           "N",
	VkO:           "O",
	VkP:           "P",
	VkQ:           "Q",
	VkR:           "R",
	VkS:           "S",
	VkT:           "T",
	VkU:           "U",
	VkV:           "V",
	VkW:           "W",
	VkX:           "X",
	VkY:           "Y",
	VkZ:           "Z",
}

func (key Key) String() string {
	return keyToString[key]
}
