package rdesktop

/*
#cgo LDFLAGS: -framework CoreGraphics -framework CoreFoundation
#include <CoreGraphics/CoreGraphics.h>

CGEventRef createWheelEvent(int x, int y) {
	return CGEventCreateScrollWheelEvent(NULL, kCGScrollEventUnitPixel, 2, y, x);
}

void get_cursor_size(int *width, int *height);
void cursor_copy(unsigned char* pixels, int width, int height);
void screenshot(unsigned char *pixels, int width, int height, bool show_cursor);
*/
import "C"

import (
	"fmt"
	"image"
	"time"
	"unsafe"
)

type osBase struct {
	id        C.CGDirectDisplayID
	ctrlDown  bool
	altDown   bool
	shiftDown bool
	cmdDown   bool
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

func (cli *osBase) Size() (image.Point, error) {
	display := C.CGDisplayCreateImage(cli.id)
	defer C.CFRelease(C.CFTypeRef(display))
	width := C.CGImageGetWidth(display)
	height := C.CGImageGetHeight(display)
	return image.Point{
		X: int(width),
		Y: int(height),
	}, nil
}

func (cli *Client) screenshot(img *image.RGBA) error {
	C.screenshot((*C.uchar)(unsafe.Pointer(&img.Pix[0])), C.int(img.Rect.Dx()), C.int(img.Rect.Dy()), C.bool(cli.showCursor))
	return nil
}

// GetCursor get cursor image
func (cli *osBase) GetCursor() (*image.RGBA, error) {
	var width, height C.int
	C.get_cursor_size(&width, &height)
	if width == 0 || height == 0 {
		return nil, fmt.Errorf("can not get cursor size")
	}
	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	C.cursor_copy((*C.uchar)(unsafe.Pointer(&img.Pix[0])), width, height)
	return img, nil
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
	code := checkKeycodes(key)
	event := C.CGEventCreateKeyboardEvent(C.CGEventSourceRef(0), C.CGKeyCode(code), true)
	if event == 0 {
		return nil
	}
	defer C.CFRelease(C.CFTypeRef(event))

	if down {
		C.CGEventSetType(event, C.kCGEventKeyDown)
	} else {
		C.CGEventSetType(event, C.kCGEventKeyUp)
	}

	flag := 0
	if cli.ctrlDown {
		flag |= C.kCGEventFlagMaskControl
	}
	if cli.altDown {
		flag |= C.kCGEventFlagMaskAlternate
	}
	if cli.cmdDown {
		flag |= C.kCGEventFlagMaskCommand
	}
	if cli.shiftDown {
		flag |= C.kCGEventFlagMaskShift
	}
	if flag != 0 {
		C.CGEventSetFlags(event, C.CGEventFlags(flag))
	}

	C.CGEventPost(C.kCGSessionEventTap, event)

	switch key {
	case "cmd":
		cli.cmdDown = down
	case "alt":
		cli.altDown = down
	case "control":
		cli.ctrlDown = down
	case "shift":
		cli.shiftDown = down
	}

	time.Sleep(0)
	return nil
}

// Scroll mouse scroll
func (cli *Client) Scroll(x, y int) {
	event := C.createWheelEvent(C.int(x), C.int(y))
	defer C.CFRelease(C.CFTypeRef(event))
	C.CGEventPost(C.kCGHIDEventTap, event)
}
