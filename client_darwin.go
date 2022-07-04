package rdesktop

/*
#cgo LDFLAGS: -framework CoreGraphics -framework CoreFoundation
#include <CoreGraphics/CoreGraphics.h>

CGEventRef createWheelEvent(int x, int y) {
	return CGEventCreateScrollWheelEvent(NULL, kCGScrollEventUnitPixel, 2, y, x);
}
*/
import "C"

import (
	"errors"
	"fmt"
	"image"
	"time"
	"unsafe"

	"golang.org/x/image/draw"
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
	defer C.CFRelease(C.CFTypeRef(display))
	raw := C.CGDataProviderCopyData(C.CGImageGetDataProvider(display))
	bits := C.CGImageGetBitsPerPixel(display)
	width := C.CGImageGetWidth(display)
	height := C.CGImageGetHeight(display)
	ptr := unsafe.Pointer(C.CFDataGetBytePtr(raw))
	size := img.Bounds()
	if size.Dx() == int(width) && size.Dy() == int(height) {
		copy(img.Pix, C.GoBytes(ptr, C.int(len(img.Pix))))
	} else {
		var src *image.RGBA
		switch bits {
		case 32:
			src = image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
		default:
			return errors.New("not supported bits")
		}
		copy(src.Pix, C.GoBytes(ptr, C.int(len(src.Pix))))
		draw.NearestNeighbor.Scale(img, img.Bounds(), src, src.Bounds(), draw.Over, nil)
	}
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
