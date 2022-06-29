package rdesktop

/*
#cgo LDFLAGS: -framework CoreGraphics -framework CoreFoundation
#include <CoreGraphics/CoreGraphics.h>
*/
import "C"

import (
	"fmt"
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
	C.CGDisplayRelease(cli.id)
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
	if cli.showCursor {
		C.CGDisplayShowCursor(cli.id)
	} else {
		C.CGDisplayHideCursor(cli.id)
	}
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
	pt := C.CGPointMake(C.double(x), C.double(y))
	err := C.CGDisplayMoveCursorToPoint(cli.id, pt)
	if err != 0 {
		return fmt.Errorf("can not move: %d", err)
	}
	return nil
}

func getMousePosition() C.CGPoint {
	event := C.CGEventCreate(C.CGEventSourceRef(0))
	defer C.CFRelease(C.CFTypeRef(event))
	return C.CGEventGetLocation(event)
}

// ToggleMouse toggle mouse button event
func (cli *Client) ToggleMouse(button MouseButton, down bool) error {
	var t C.CGEventType
	var btn C.CGMouseButton
	switch button {
	case MouseLeft:
		if down {
			t = C.kCGEventLeftMouseDown
		} else {
			t = C.kCGEventLeftMouseUp
		}
		btn = 0
	case MouseMiddle:
		if down {
			t = C.kCGEventOtherMouseDown
		} else {
			t = C.kCGEventOtherMouseUp
		}
		btn = 2
	case MouseRight:
		if down {
			t = C.kCGEventRightMouseDown
		} else {
			t = C.kCGEventRightMouseUp
		}
		btn = 1
	}
	event := C.CGEventCreateMouseEvent(C.CGEventSourceRef(0), t, getMousePosition(), btn)
	defer C.CFRelease(C.CFTypeRef(event))
	C.CGEventPost(C.kCGSessionEventTap, event)
	return nil
}

// ToggleKey toggle keyboard event
func (cli *Client) ToggleKey(key string, down bool) error {
	return nil
}

// Scroll mouse scroll
func (cli *Client) Scroll(x, y int) {
}
