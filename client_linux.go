package screenshot

/*
#cgo CFLAGS: -DXUTIL_DEFINE_FUNCTIONS
#cgo LDFLAGS: -lX11
#include <X11/Xlib.h>
#include <X11/X.h>
#include <X11/Xutil.h>
*/
import "C"

import (
	"errors"
	"image"
)

// ErrOpenDiaplay can not open display
var ErrOpenDiaplay = errors.New("can not open display")

// ErrGetWindow can not get root window
var ErrGetWindow = errors.New("can not get root window")

// ErrGetImage can not get image
var ErrGetImage = errors.New("can not get image")

type osBase struct {
	display *C.Display
	window  C.Window
}

func (cli *osBase) init() error {
	cli.display = C.XOpenDisplay(nil)
	if cli.display == nil {
		return ErrOpenDiaplay
	}
	cli.window = C.XDefaultRootWindow(cli.display)
	if cli.window == 0 {
		return ErrGetWindow
	}
	return nil
}

// Close close client
func (cli *osBase) Close() {
	C.XCloseDisplay(cli.display)
}

func (cli *osBase) size() image.Point {
	var gwa C.XWindowAttributes
	C.XGetWindowAttributes(cli.display, cli.window, &gwa)
	return image.Point{X: int(gwa.width), Y: int(gwa.height)}
}

func (cli *osBase) screenshot(img *image.RGBA) error {
	size := img.Rect.Max
	cimage := C.XGetImage(cli.display, cli.window, 0, 0, C.uint(size.X), C.uint(size.Y), C.AllPlanes, C.ZPixmap)
	if cimage == nil {
		return ErrGetImage
	}
	defer C.XDestroyImage(cimage)
	redMask := cimage.red_mask
	greenMask := cimage.green_mask
	blueMask := cimage.blue_mask
	offset := 0
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			pix := C.XGetPixel(cimage, C.int(x), C.int(y))
			img.Pix[offset+2] = uint8(pix & blueMask)
			img.Pix[offset+1] = uint8((pix & greenMask) >> 8)
			img.Pix[offset+0] = uint8((pix & redMask) >> 16)
			offset += 4
		}
	}
	return nil
}
