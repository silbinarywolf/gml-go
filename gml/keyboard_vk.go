package gml

const (
	VkNoKey       = 0 + iota // keycode representing that no key is pressed
	VkAnykey                 // keycode representing that any key is pressed
	VkLeft                   // keycode for left arrow key
	VkRight                  // keycode for right arrow key
	VkUp                     // keycode for up arrow key
	VkDown                   // keycode for down arrow key
	VkEnter                  // enter key
	VkEscape                 // escape key
	VkSpace                  // space key
	VkShift                  // either of the shift keys
	VkControl                // either of the control keys
	VkAlt                    // alt key
	VkBackspace              // backspace key
	VkTab                    // tab key
	VkHome                   // home key
	VkEnd                    // end key
	VkDelete                 // delete key
	VkInsert                 // insert key
	VkPageUp                 // pageup key
	VkPageDown               // pagedown key
	VkPause                  // pause/break key
	VkPrintScreen            // printscreen/sysrq key
	VkF1                     // keycode for the function keys F1 to F12
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
	VkMultiply // multiply key on the numeric keypad
	VkDivide   // divide key on the numeric keypad
	VkAdd      // key on the numeric keypad
	VkSubtract // subtract key on the numeric keypad
	VkDecimal  // decimal dot keys on the numeric keypad

	// The following in GML are letters keys, handled like: keyboard_check(ord("R")), however we'll just use VK
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

	// The following constants can only be used with keyboard_check_direct()
	// Therefore, they are not supported.
	//VkLShift
	//VkLControl
	//VkLAlt
	//VkRShift
	//VkRControl
	//VkRAlt

	vkSize
)
