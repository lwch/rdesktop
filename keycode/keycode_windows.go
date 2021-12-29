package keycode

import (
	"syscall"

	"github.com/lwch/rdesktop/windef"
)

// Maps https://docs.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes
var Maps = map[string]int{
	"backspace": 0x08, // VK_BACK
	"delete":    0x2E, // VK_DELETE
	"enter":     0x0D, // VK_RETURN
	"tab":       0x09, // VK_TAB
	"esc":       0x1B, // VK_ESCAPE
	"up":        0x26, // VK_UP
	"down":      0x28, // VK_DOWN
	"right":     0x27, // VK_RIGHT
	"left":      0x25, // VK_LEFT
	"home":      0x24, // VK_HOME
	"end":       0x23, // VK_END
	"pageup":    0x21, // VK_PRIOR
	"pagedown":  0x22, // VK_NEXT
	//
	"f1":  0x70, // VK_F1
	"f2":  0x71, // VK_F2
	"f3":  0x72, // VK_F3
	"f4":  0x73, // VK_F4
	"f5":  0x74, // VK_F5
	"f6":  0x75, // VK_F6
	"f7":  0x76, // VK_F7
	"f8":  0x77, // VK_F8
	"f9":  0x78, // VK_F9
	"f10": 0x79, // VK_F10
	"f11": 0x7A, // VK_F11
	"f12": 0x7B, // VK_F12
	//
	"cmd":     0x5B, // VK_LWIN
	"alt":     0x12, // VK_MENU
	"control": 0x11, // VK_CONTROL
	"shift":   0xA0, // VK_LSHIFT
	"space":   0x20, // VK_SPACE
}

// ForChar char key code
func ForChar(k string) int {
	code, _, _ := syscall.Syscall(windef.FuncVkKeyScan, 1, uintptr(k[0]), 0, 0)
	return int(code)
}
