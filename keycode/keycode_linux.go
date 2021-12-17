//go:build linux
// +build linux

package keycode

/*
#cgo LDFLAGS: -lX11
#include <X11/Xlib.h>
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

// Maps 数据来源于X11/keysymdef.h
var Maps = map[string]int{
	"backspace": 0xff08,
	"delete":    0xffff,
	"enter":     0xff0d,
	"tab":       0xff09,
	"esc":       0xff1b,
	"up":        0xff52,
	"down":      0xff54,
	"right":     0xff53,
	"left":      0xff51,
	"home":      0xff50,
	"end":       0xff57,
	"pageup":    0xff55,
	"pagedown":  0xff56,
	//
	"f1":  0xffbe,
	"f2":  0xffbf,
	"f3":  0xffc0,
	"f4":  0xffc1,
	"f5":  0xffc2,
	"f6":  0xffc3,
	"f7":  0xffc4,
	"f8":  0xffc5,
	"f9":  0xffc6,
	"f10": 0xffc7,
	"f11": 0xffc8,
	"f12": 0xffc9,
	//
	"cmd":     0xffeb,
	"alt":     0xffe9,
	"control": 0xffe3,
	"shift":   0xffe1,
	"space":   0x0020,
}

// ForChar char key code
func ForChar(k string) int {
	key := C.CString(k)
	code := C.XStringToKeysym(key)
	C.free(unsafe.Pointer(key))
	return int(code)
}
