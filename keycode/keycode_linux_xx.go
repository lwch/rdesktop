//go:build linux && (386 || amd64)
// +build linux
// +build 386 amd64

package keycode

/*
#cgo LDFLAGS: -lX11 -lxcb -lpthread -lXau -lXdmcp
#include <X11/Xlib.h>
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

// Maps 数据来源于X11/keysymdef.h
var Maps = map[string]int{
	"backspace": 0xff08, // XK_BackSpace
	"delete":    0xffff, // XK_Delete
	"enter":     0xff0d, // XK_Return
	"tab":       0xff09, // XK_Tab
	"esc":       0xff1b, // XK_Escape
	"up":        0xff52, // XK_Up
	"down":      0xff54, // XK_Down
	"right":     0xff53, // XK_Right
	"left":      0xff51, // XK_Left
	"home":      0xff50, // XK_Home
	"end":       0xff57, // XK_End
	"pageup":    0xff55, // XK_Page_Up
	"pagedown":  0xff56, // XK_Page_Down
	//
	"f1":  0xffbe, // XK_F1
	"f2":  0xffbf, // XK_F2
	"f3":  0xffc0, // XK_F3
	"f4":  0xffc1, // XK_F4
	"f5":  0xffc2, // XK_F5
	"f6":  0xffc3, // XK_F6
	"f7":  0xffc4, // XK_F7
	"f8":  0xffc5, // XK_F8
	"f9":  0xffc6, // XK_F9
	"f10": 0xffc7, // XK_F10
	"f11": 0xffc8, // XK_F11
	"f12": 0xffc9, // XK_F12
	//
	"cmd":     0xffeb, // XK_Super_L
	"alt":     0xffe9, // XK_Alt_L
	"control": 0xffe3, // XK_Control_L
	"shift":   0xffe1, // XK_Shift_L
	"space":   0x0020, // XK_space
}

// ForChar char key code
func ForChar(k string) int {
	key := C.CString(k)
	code := C.XStringToKeysym(key)
	C.free(unsafe.Pointer(key))
	return int(code)
}
