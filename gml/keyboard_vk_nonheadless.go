// +build !headless

package gml

import (
	"github.com/hajimehoshi/ebiten"
)

var keyboardVkToEbiten = []ebiten.Key{
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
	VkPause:       ebiten.KeyPause,       // pause/break key
	VkPrintScreen: ebiten.KeyPrintScreen, // printscreen/sysrq key
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
	VkNumpad0:     ebiten.KeyKP0,
	VkNumpad1:     ebiten.KeyKP1,
	VkNumpad2:     ebiten.KeyKP2,
	VkNumpad3:     ebiten.KeyKP3,
	VkNumpad4:     ebiten.KeyKP4,
	VkNumpad5:     ebiten.KeyKP5,
	VkNumpad6:     ebiten.KeyKP6,
	VkNumpad7:     ebiten.KeyKP7,
	VkNumpad8:     ebiten.KeyKP8,
	VkNumpad9:     ebiten.KeyKP9,
	VkNumpadEnter: ebiten.KeyKPEnter,
	VkMultiply:    0,               // multiply key on the numeric keypad
	VkDivide:      0,               // divide key on the numeric keypad
	VkAdd:         0,               // key on the numeric keypad
	VkSubtract:    ebiten.KeyMinus, // subtract key on the numeric keypad
	VkDecimal:     0,               // decimal dot keys on

	// The following in GML are letters keys, handled like: keyboard_check(ord("R")), however we'll just use VK
	VkA: ebiten.KeyA,
	VkB: ebiten.KeyB,
	VkC: ebiten.KeyC,
	VkD: ebiten.KeyD,
	VkE: ebiten.KeyE,
	VkF: ebiten.KeyF,
	VkG: ebiten.KeyG,
	VkH: ebiten.KeyH,
	VkI: ebiten.KeyI,
	VkJ: ebiten.KeyJ,
	VkK: ebiten.KeyK,
	VkL: ebiten.KeyL,
	VkM: ebiten.KeyM,
	VkN: ebiten.KeyN,
	VkO: ebiten.KeyO,
	VkP: ebiten.KeyP,
	VkQ: ebiten.KeyQ,
	VkR: ebiten.KeyR,
	VkS: ebiten.KeyS,
	VkT: ebiten.KeyT,
	VkU: ebiten.KeyU,
	VkV: ebiten.KeyV,
	VkW: ebiten.KeyW,
	VkX: ebiten.KeyX,
	VkY: ebiten.KeyY,
	VkZ: ebiten.KeyZ,

	// Not supported by Game Maker outside of Windows platform.
	//VkLShift:   0,
	//VkLControl: 0,
	//VkLAlt:     0,
	//VkRShift:   0,
	//VkRControl: 0,
	//VkRAlt:     0,
}
