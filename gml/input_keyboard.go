package gml

import (
	"github.com/hajimehoshi/ebiten"
)

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

	// The following constants can only be used with keyboard_check_direct()
	VkLShift
	VkLControl
	VkLAlt
	VkRShift
	VkRControl
	VkRAlt
)

var g_vkToKey = []ebiten.Key{
	VkNoKey:       -1,                    // keycode representing that no key is pressed
	VkAnykey:      0,                     // keycode representing that any key is pressed
	VkLeft:        ebiten.KeyLeft,        // keycode for left arrow key
	VkRight:       ebiten.KeyRight,       // keycode for right arrow key
	VkUp:          ebiten.KeyUp,          // keycode for up arrow key
	VkDown:        ebiten.KeyDown,        // keycode for down arrow key
	VkEnter:       ebiten.KeyEnter,       // enter key
	VkEscape:      ebiten.KeyEscape,      // escape key
	VkSpace:       ebiten.KeySpace,       // space key
	VkShift:       ebiten.KeyShift,       // either of the shift keys
	VkControl:     ebiten.KeyControl,     // either of the control keys
	VkAlt:         ebiten.KeyAlt,         // alt key
	VkBackspace:   ebiten.KeyBackslash,   // backspace key
	VkTab:         ebiten.KeyTab,         // tab key
	VkHome:        ebiten.KeyHome,        // home key
	VkEnd:         ebiten.KeyEnd,         // end key
	VkDelete:      ebiten.KeyDelete,      // delete key
	VkInsert:      ebiten.KeyInsert,      // insert key
	VkPageUp:      ebiten.KeyPageUp,      // pageup key
	VkPageDown:    ebiten.KeyPageDown,    // pagedown key
	VkPause:       0,                     // pause/break key
	VkPrintScreen: ebiten.KeyLeftBracket, // printscreen/sysrq key
	VkF1:          ebiten.KeyF1,          // keycode for the function keys F1 to F12
	VkF2:          ebiten.KeyF2,
	VkF3:          ebiten.KeyF3,
	VkF4:          ebiten.KeyF4,
	VkF5:          ebiten.KeyF5,
	VkF6:          ebiten.KeyF6,
	VkF7:          ebiten.KeyF7,
	VkF8:          ebiten.KeyF8,
	VkF9:          ebiten.KeyF9,
	VkF10:         ebiten.KeyF10,
	VkF11:         ebiten.KeyF11,
	VkF12:         ebiten.KeyF11,
	VkNumpad0:     ebiten.Key0,
	VkNumpad1:     ebiten.Key1,
	VkNumpad2:     ebiten.Key2,
	VkNumpad3:     ebiten.Key3,
	VkNumpad4:     ebiten.Key4,
	VkNumpad5:     ebiten.Key5,
	VkNumpad6:     ebiten.Key6,
	VkNumpad7:     ebiten.Key7,
	VkNumpad8:     ebiten.Key8,
	VkNumpad9:     ebiten.Key9,
	VkMultiply:    0,               // multiply key on the numeric keypad
	VkDivide:      0,               // divide key on the numeric keypad
	VkAdd:         0,               // key on the numeric keypad
	VkSubtract:    ebiten.KeyMinus, // subtract key on the numeric keypad
	VkDecimal:     0,               // decimal dot keys on

	VkLShift:   0,
	VkLControl: 0,
	VkLAlt:     0,
	VkRShift:   0,
	VkRControl: 0,
	VkRAlt:     0,
}

func KeyboardCheck(key int16) bool {
	return ebiten.IsKeyPressed(g_vkToKey[key])
}
