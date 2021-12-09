package rdesktop

import (
	"image"

	"github.com/lwch/rdesktop/x11"
)

type mouseButton byte

const (
	MouseLeft   mouseButton = 1
	MouseMiddle mouseButton = 2
	MouseRight  mouseButton = 3
)

type osBase struct {
	cli *x11.Client
}

func (cli *osBase) init() error {
	var err error
	cli.cli, err = x11.New()
	if err != nil {
		return err
	}
	return nil
}

// Close close client
func (cli *osBase) Close() {
	if cli.cli != nil {
		cli.cli.Close()
	}
}

func (cli *osBase) size() (image.Point, error) {
	return cli.cli.GetSize()
}

func (cli *Client) screenshot(img *image.RGBA) error {
	err := cli.cli.GetImage(img)
	if err != nil {
		return err
	}
	if cli.showCursor {
		cli.cli.GetCursorImage(img)
	}
	return nil
}

// MouseMove move mouse to x,y
func (cli *Client) MouseMove(x, y int) error {
	return cli.cli.WarpPointer(uint16(x), uint16(y))
}

// ToggleMouse toggle mouse button event, https://www.x.org/releases/X11R7.7/doc/xextproto/xtest.html
func (cli *Client) ToggleMouse(button mouseButton, down bool) error {
	t := 4 // button down
	if !down {
		t = 5
	}
	return cli.cli.TestFakeInput(byte(t), byte(button), 0, 0)
}
