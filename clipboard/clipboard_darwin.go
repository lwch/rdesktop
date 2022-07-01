package clipboard

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include "clipboard_darwin.h"
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

func init() {
	C.clipboard_init()
}

// Set set text to clipboard
func Set(text string) error {
	str := C.CString(text)
	defer C.free(unsafe.Pointer(str))
	if !C.set_clipboard(str) {
		return errors.New("can not set clipboard data")
	}
	return nil
}

// Get get clipboard text
func Get() (string, error) {
	str := C.get_clipboard()
	return C.GoString(str), nil
}
