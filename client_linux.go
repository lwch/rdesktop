package screenshot

/*
#cgo LDFLAGS: -lX11
#include <X11/Xlib.h>
#include <X11/X.h>
#include <X11/Xutil.h>
*/
import "C"

import (
	"fmt"
	"image"
)

type osBase struct {
}

func (cli *osBase) init() {
	display := C.XOpenDisplay(nil)
	// root := C.XDefaultRootWindow(display)
	fmt.Println(display)
}

func (cli *osBase) size() image.Rectangle {
	return image.Rect(0, 0, 0, 0)
}
