package rdesktop

import (
	"fmt"
	"image"
	"math"

	"github.com/lwch/rdesktop/x11"
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

func (cli *osBase) Size() (image.Point, error) {
	return cli.cli.GetSize()
}

func (cli *Client) screenshot(img *image.RGBA) error {
	err := cli.cli.GetImage(img)
	if err != nil {
		return err
	}
	if cli.showCursor {
		cli.cli.DrawCursor(img)
	}
	return nil
}

func (cli *Client) GetCursor() (*image.RGBA, error) {
	return cli.cli.GetCursor()
}

// MouseMove move mouse to x,y
func (cli *Client) MouseMove(x, y int) error {
	return cli.cli.WarpPointer(uint16(x), uint16(y))
}

// ToggleMouse toggle mouse button event, https://www.x.org/releases/X11R7.7/doc/xextproto/xtest.html
func (cli *Client) ToggleMouse(button MouseButton, down bool) error {
	t := 4 // button down
	if !down {
		t = 5 // button up
	}
	return cli.cli.TestFakeInput(byte(t), byte(button)+1)
}

// ToggleKey toggle keyboard event
func (cli *Client) ToggleKey(key string, down bool) error {
	code := checkKeycodes(key)
	if code == 0 {
		return fmt.Errorf("key not found: %s", key)
	}
	t := 2 // key down
	if !down {
		t = 3 // key up
	}
	n := cli.cli.KeysymToKeycode(code)
	return cli.cli.TestFakeInput(byte(t), n)
}

// Scroll https://github.com/go-vgo/robotgo/blob/master/mouse/mouse_c.h#L313
func (cli *Client) Scroll(x, y int) {
	run := func(dir byte, cnt int) {
		for i := 0; i < cnt; i++ {
			// https://gitlab.freedesktop.org/xorg/lib/libxtst/-/blob/master/src/XTest.c#L181
			// transform press to 4 and up to 5
			cli.cli.TestFakeInput(4, dir)
			cli.cli.TestFakeInput(5, dir)
		}
	}
	if x != 0 {
		dir := 6 // up
		if x < 0 {
			dir = 7 // down
		}
		run(byte(dir), int(math.Abs(float64(x))))
	}
	if y != 0 {
		dir := 4 // up
		if y < 0 {
			dir = 5 // down
		}
		run(byte(dir), int(math.Abs(float64(y))))
	}
}
