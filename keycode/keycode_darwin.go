package keycode

/*
#include <Carbon/Carbon.h>
*/
import "C"

// Maps code from /System/Library/Frameworks/Carbon.framework/Versions/A/Frameworks/HIToolbox.framework/Versions/A/Headers/Events.h
var Maps = map[string]int{
	"backspace": C.kVK_Delete,
	"delete":    C.kVK_ForwardDelete,
	"enter":     C.kVK_Return,
	"tab":       C.kVK_Tab,
	"esc":       C.kVK_Escape,
	"up":        C.kVK_UpArrow,
	"down":      C.kVK_DownArrow,
	"right":     C.kVK_RightArrow,
	"left":      C.kVK_LeftArrow,
	"home":      C.kVK_Home,
	"end":       C.kVK_End,
	"pageup":    C.kVK_PageUp,
	"pagedown":  C.kVK_PageDown,
	//
	"f1":  C.kVK_F1,
	"f2":  C.kVK_F2,
	"f3":  C.kVK_F3,
	"f4":  C.kVK_F4,
	"f5":  C.kVK_F5,
	"f6":  C.kVK_F6,
	"f7":  C.kVK_F7,
	"f8":  C.kVK_F8,
	"f9":  C.kVK_F9,
	"f10": C.kVK_F10,
	"f11": C.kVK_F11,
	"f12": C.kVK_F12,
	//
	"cmd":     C.kVK_Command,
	"alt":     C.kVK_Option,
	"control": C.kVK_Control,
	"shift":   C.kVK_Shift,
	"space":   C.kVK_Space,
	//
	"a": C.kVK_ANSI_A,
	"b": C.kVK_ANSI_B,
	"c": C.kVK_ANSI_C,
	"d": C.kVK_ANSI_D,
	"e": C.kVK_ANSI_E,
	"f": C.kVK_ANSI_F,
	"g": C.kVK_ANSI_G,
	"h": C.kVK_ANSI_H,
	"i": C.kVK_ANSI_I,
	"j": C.kVK_ANSI_J,
	"k": C.kVK_ANSI_K,
	"l": C.kVK_ANSI_L,
	"m": C.kVK_ANSI_M,
	"n": C.kVK_ANSI_N,
	"o": C.kVK_ANSI_O,
	"p": C.kVK_ANSI_P,
	"q": C.kVK_ANSI_Q,
	"r": C.kVK_ANSI_R,
	"s": C.kVK_ANSI_S,
	"t": C.kVK_ANSI_T,
	"u": C.kVK_ANSI_U,
	"v": C.kVK_ANSI_V,
	"w": C.kVK_ANSI_W,
	"x": C.kVK_ANSI_X,
	"y": C.kVK_ANSI_Y,
	"z": C.kVK_ANSI_Z,
	//
	"`": C.kVK_ANSI_Grave,
	"0": C.kVK_ANSI_0,
	"1": C.kVK_ANSI_1,
	"2": C.kVK_ANSI_2,
	"3": C.kVK_ANSI_3,
	"4": C.kVK_ANSI_4,
	"5": C.kVK_ANSI_5,
	"6": C.kVK_ANSI_6,
	"7": C.kVK_ANSI_7,
	"8": C.kVK_ANSI_8,
	"9": C.kVK_ANSI_9,
	//
	"-":  C.kVK_ANSI_Minus,
	"=":  C.kVK_ANSI_Equal,
	"[":  C.kVK_ANSI_LeftBracket,
	"]":  C.kVK_ANSI_RightBracket,
	"\\": C.kVK_ANSI_Backslash,
	";":  C.kVK_ANSI_Semicolon,
	"'":  C.kVK_ANSI_Quote,
	",":  C.kVK_ANSI_Comma,
	".":  C.kVK_ANSI_Period,
	"/":  C.kVK_ANSI_Slash,
}

// ForChar char key code
func ForChar(k string) int {
	return Maps[k]
}
