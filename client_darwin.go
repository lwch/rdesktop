package rdesktop

/*
#cgo LDFLAGS: -framework CoreGraphics -framework CoreFoundation
#include <CoreGraphics/CoreGraphics.h>
*/
import "C"

import (
	"image"
	"unsafe"
)

type osBase struct {
	id C.CGDirectDisplayID
}

func (cli *osBase) init() error {
	cli.id = getDisplayID()
	return nil
}

// Close close client
func (cli *osBase) Close() {
}

func getDisplayID() C.CGDirectDisplayID {
	var id C.CGDirectDisplayID
	if C.CGGetActiveDisplayList(C.uint32_t(1), (*C.CGDirectDisplayID)(unsafe.Pointer(&id)), nil) != C.kCGErrorSuccess {
		return 0
	}
	return id
}

func (cli *osBase) size() (image.Point, error) {
	rect := C.CGDisplayBounds(cli.id)
	return image.Point{
		X: int(rect.size.width),
		Y: int(rect.size.height),
	}, nil
}

func (cli *Client) screenshot(img *image.RGBA) error {
	display := C.CGDisplayCreateImage(cli.id)
	raw := C.CGDataProviderCopyData(C.CGImageGetDataProvider(display))
	ptr := unsafe.Pointer(C.CFDataGetBytePtr(raw))
	copy(img.Pix, C.GoBytes(ptr, C.int(len(img.Pix))))
	// BGR => RGB
	for i := 0; i < len(img.Pix); i += 4 {
		img.Pix[i], img.Pix[i+2] = img.Pix[i+2], img.Pix[i]
	}
	return nil
}

// MouseMove move mouse to x,y
func (cli *osBase) MouseMove(x, y int) error {
	// TODO
	return nil
}

// ToggleMouse toggle mouse button event
func (cli *Client) ToggleMouse(button MouseButton, down bool) error {
	return nil
}

// ToggleKey toggle keyboard event
func (cli *Client) ToggleKey(key string, down bool) error {
	return nil
}

// Scroll mouse scroll
func (cli *Client) Scroll(x, y int) {
}
