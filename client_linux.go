package screenshot

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
}

func (cli *osBase) size() (image.Point, error) {
	return cli.cli.GetSize()
}

func (cli *osBase) screenshot(img *image.RGBA) error {
	return cli.cli.GetImage(img)
}
