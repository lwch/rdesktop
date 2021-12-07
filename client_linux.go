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
