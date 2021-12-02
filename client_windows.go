package screenshot

import "image"

type osBase struct {
}

func (cli *osBase) init() error {
	return nil
}

func (cli *osBase) size() (image.Point, error) {
	return image.Point{}, nil
}

func (cli *osBase) Close() {
}

func (cli *osBase) screenshot(img *image.RGBA) error {
	return nil
}
