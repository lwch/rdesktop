package rdesktop

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

// MouseMove move mouse to x,y
func (cli *osBase) MouseMove(x, y int) error {
	// TODO
	return nil
}

// ToggleMouse toggle mouse button event
func (cli *Client) ToggleMouse(button MouseButton, down bool) error {
	return nil
}

// ToggleKey toggle keyboard event
func (cli *Client) ToggleKey(key string, down bool) error {
	return nil
}

// Scroll mouse scroll
func (cli *Client) Scroll(x, y int) {
}
