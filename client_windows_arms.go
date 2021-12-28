//go:build windows && (arm || arm64)
// +build windows
// +build arm arm64

package rdesktop

import "image"

type osBase struct{}

func (cli *osBase) init() error {
	return ErrUnsupported
}

func (cli *osBase) size() (image.Point, error) {
	return image.Point{}, ErrUnsupported
}

func (cli *osBase) Close() {
}

func (cli *Client) screenshot(img *image.RGBA) error {
	return ErrUnsupported
}

// MouseMove move mouse to x,y
func (cli *osBase) MouseMove(x, y int) error {
	return ErrUnsupported
}

// ToggleMouse toggle mouse button event
func (cli *Client) ToggleMouse(button MouseButton, down bool) error {
	return ErrUnsupported
}

// ToggleKey toggle keyboard event
func (cli *Client) ToggleKey(key string, down bool) error {
	return ErrUnsupported
}
