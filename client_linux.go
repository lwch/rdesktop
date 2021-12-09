package rdesktop

import (
	"image"

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
	err := cli.cli.WarpPointer(uint16(x), uint16(y))
	if err != nil {
		return err
	}
	return nil
}
