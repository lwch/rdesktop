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

	"github.com/lwch/screenshot/x11"
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
	cli     *x11.Client
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
	var err error
	cli.cli, err = x11.New()
	if err != nil {
		return err
	}
	return nil
}

// Close close client
func (cli *osBase) Close() {
	C.XCloseDisplay(cli.display)
}

func (cli *osBase) size() (image.Point, error) {
	return cli.cli.GetSize()
}

func (cli *osBase) screenshot(img *image.RGBA) error {
	return cli.cli.GetImage(img)
}
