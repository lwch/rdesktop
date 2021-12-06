package screenshot

import (
	"errors"
	"fmt"
	"image"
	"syscall"
	"unsafe"

	"github.com/lwch/logging"
	"github.com/lwch/screenshot/windef"
)

type osBase struct {
	hwnd   uintptr
	hdc    uintptr
	bits   uintptr
	width  uintptr
	height uintptr
	buffer uintptr
}

func (cli *osBase) init() error {
	return cli.getHandle()
}

func (cli *osBase) getHandle() error {
	hwnd, _, err := syscall.Syscall(windef.FuncGetDesktopWindow, 0, 0, 0, 0)
	if hwnd == 0 {
		return fmt.Errorf("get desktop window: %v", err)
	}
	hdc, _, err := syscall.Syscall(windef.FuncGetDC, 1, cli.hwnd, 0, 0)
	if hdc == 0 {
		return fmt.Errorf("get dc: %v", err)
	}
	if cli.hdc != 0 {
		syscall.Syscall(windef.FuncReleaseDC, 2, cli.hwnd, cli.hdc, 0)
	}
	cli.hwnd = hwnd
	cli.hdc = hdc
	return nil
}

func (cli *osBase) size() (image.Point, error) {
	err := cli.getHandle()
	if err != nil {
		return image.Point{}, err
	}
	bits, _, err := syscall.Syscall(windef.FuncGetDeviceCaps, 2, cli.hdc, windef.BITSPIXEL, 0)
	if bits == 0 {
		return image.Point{}, fmt.Errorf("get device caps(bits): %v", err)
	}
	if cli.bits != 32 {
		cli.bits = 32
		logging.Info("reset bits to 32")
	}
	width, _, err := syscall.Syscall(windef.FuncGetDeviceCaps, 2, cli.hdc, windef.HORZRES, 0)
	if width == 0 {
		return image.Point{}, fmt.Errorf("get device caps(width): %v", err)
	}
	height, _, err := syscall.Syscall(windef.FuncGetDeviceCaps, 2, cli.hdc, windef.VERTRES, 0)
	if height == 0 {
		return image.Point{}, fmt.Errorf("get device caps(height): %v", err)
	}
	if cli.width != width || cli.height != height {
		err := cli.resizeBuffer(int(bits), int(width), int(height))
		if err != nil {
			return image.Point{}, err
		}
	}
	cli.bits = bits
	cli.width = width
	cli.height = height
	return image.Point{
		X: int(width),
		Y: int(height),
	}, nil
}

func (cli *osBase) resizeBuffer(bits, width, height int) error {
	addr, _, err := syscall.Syscall(windef.FuncGlobalAlloc, 2, windef.GMEMFIXED, uintptr(bits*width*height/8), 0)
	if addr == 0 {
		return fmt.Errorf("global alloc: %v", err)
	}
	if cli.buffer != 0 {
		syscall.Syscall(windef.FuncGlobalFree, 1, cli.buffer, 0, 0)
	}
	cli.buffer = addr
	return nil
}

func (cli *osBase) Close() {
	if cli.hwnd != 0 && cli.hdc != 0 {
		syscall.Syscall(windef.FuncReleaseDC, 2, cli.hwnd, cli.hdc, 0)
	}
}

func (cli *Client) screenshot(img *image.RGBA) error {
	memDC, bitmap, free, err := cli.bitblt(img.Rect.Max.X, img.Rect.Max.Y)
	if err != nil {
		return err
	}
	defer free()
	defer cli.copyImageData(cli.hdc, bitmap, img)
	if cli.showCursor {
		cli.drawCursor(memDC)
	}
	return nil
}

func (cli *osBase) bitblt(width, height int) (uintptr, uintptr, func(), error) {
	memDC, _, err := syscall.Syscall(windef.FuncCreateCompatibleDC, 1, cli.hdc, 0, 0)
	if memDC == 0 {
		return 0, 0, nil, errors.New("create dc: " + err.Error())
	}
	bitmap, _, err := syscall.Syscall(windef.FuncCreateCompatibleBitmap, 3, cli.hdc,
		uintptr(width), uintptr(height))
	if bitmap == 0 {
		syscall.Syscall(windef.FuncDeleteDC, 1, memDC, 0, 0)
		return 0, 0, nil, errors.New("create bitmap: " + err.Error())
	}
	oldDC, _, err := syscall.Syscall(windef.FuncSelectObject, 2, memDC, bitmap, 0)
	if oldDC == 0 {
		syscall.Syscall(windef.FuncDeleteObject, 1, bitmap, 0, 0)
		syscall.Syscall(windef.FuncDeleteDC, 1, memDC, 0, 0)
		return 0, 0, nil, errors.New("select object: " + err.Error())
	}
	ok, _, err := syscall.Syscall9(windef.FuncBitBlt, 9, memDC, 0, 0,
		uintptr(width), uintptr(height), cli.hdc, 0, 0, windef.SRCCOPY)
	if ok == 0 {
		syscall.Syscall(windef.FuncSelectObject, 2, memDC, oldDC, 0)
		syscall.Syscall(windef.FuncDeleteObject, 1, bitmap, 0, 0)
		syscall.Syscall(windef.FuncDeleteDC, 1, memDC, 0, 0)
		return 0, 0, nil, errors.New("bitblt: " + err.Error())
	}
	return memDC, bitmap, func() {
		syscall.Syscall(windef.FuncSelectObject, 2, memDC, oldDC, 0)
		syscall.Syscall(windef.FuncDeleteObject, 1, bitmap, 0, 0)
		syscall.Syscall(windef.FuncDeleteDC, 1, memDC, 0, 0)
	}, nil
}

// BITMAPINFOHEADER https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-bitmapinfoheader
type BITMAPINFOHEADER struct {
	BiSize          uint32
	BiWidth         int32
	BiHeight        int32
	BiPlanes        uint16
	BiBitCount      uint16
	BiCompression   uint32
	BiSizeImage     uint32
	BiXPelsPerMeter int32
	BiYPelsPerMeter int32
	BiClrUsed       uint32
	BiClrImportant  uint32
}

func (cli *osBase) copyImageData(hdc, bitmap uintptr, img *image.RGBA) {
	var hdr BITMAPINFOHEADER
	hdr.BiSize = uint32(unsafe.Sizeof(hdr))
	hdr.BiPlanes = 1
	hdr.BiBitCount = uint16(cli.bits)
	hdr.BiWidth = int32(img.Rect.Max.X)
	hdr.BiHeight = int32(-img.Rect.Max.Y)
	hdr.BiCompression = windef.BIRGB
	hdr.BiSizeImage = 0
	lines, _, err := syscall.Syscall9(windef.FuncGetDIBits, 7, hdc, bitmap, 0, uintptr(img.Rect.Max.Y),
		cli.buffer, uintptr(unsafe.Pointer(&hdr)), windef.DIBRGBCOLORS, 0, 0)
	if lines == 0 {
		logging.Error("get bits: %v", err)
	}
	// TODO: support difference of 32 bits
	for i := 0; i < len(img.Pix); i++ {
		img.Pix[i] = *(*uint8)(unsafe.Pointer(cli.buffer + uintptr(i)))
	}
	// BGR => RGB
	for i := 0; i < len(img.Pix); i += int(cli.bits / 8) {
		img.Pix[i], img.Pix[i+2] = img.Pix[i+2], img.Pix[i]
	}
}

func (cli *osBase) drawCursor(memDC uintptr) error {
	var curInfo windef.CURSORINFO
	curInfo.CbSize = windef.DWORD(unsafe.Sizeof(curInfo))
	ok, _, err := syscall.Syscall(windef.FuncGetCursorInfo, 1, uintptr(unsafe.Pointer(&curInfo)), 0, 0)
	if ok == 0 {
		logging.Error("get cursor info: %v", err)
		return nil
	}
	if curInfo.Flags == windef.CURSORSHOWING {
		var info windef.ICONINFO
		ok, _, err = syscall.Syscall(windef.FuncGetIconInfo, 2, uintptr(curInfo.HCursor), uintptr(unsafe.Pointer(&info)), 0)
		if ok == 0 {
			logging.Error("get icon info: %v", err)
			return nil
		}
		x := curInfo.PTScreenPos.X - windef.LONG(info.XHotspot)
		y := curInfo.PTScreenPos.Y - windef.LONG(info.YHotspot)
		ok, _, err = syscall.Syscall6(windef.FuncDrawIcon, 4, memDC, uintptr(x), uintptr(y), uintptr(curInfo.HCursor), 0, 0)
		if ok == 0 {
			logging.Error("draw icon: %v", err)
			return nil
		}
	}
	return nil
}
