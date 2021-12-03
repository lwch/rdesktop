package screenshot

import "image"

type osBase struct {
}

func (cli *osBase) init() error {
	return nil
}

// Close close client
func (cli *osBase) Close() {
}

func (cli *osBase) size() (image.Point, error) {
	return image.Point{}, nil
}

func (cli *Client) screenshot(img *image.RGBA) error {
	return nil
}
