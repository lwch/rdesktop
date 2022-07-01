package rdesktop

import (
	"image"
	"sync"

	"github.com/lwch/logging"
	"github.com/lwch/rdesktop/clipboard"
)

// MouseButton button of mouse
type MouseButton byte

const (
	// MouseLeft left button for mouse
	MouseLeft MouseButton = iota
	// MouseMiddle middle button for mouse
	MouseMiddle
	// MouseRight right button for mouse
	MouseRight
)

// Client screenshot client
type Client struct {
	sync.Mutex
	osBase
	img        *image.RGBA
	showCursor bool
}

// New create client
func New() (*Client, error) {
	cli := &Client{}
	err := cli.init()
	if err != nil {
		return nil, err
	}
	size, err := cli.Size()
	if err != nil {
		cli.Close()
		return nil, err
	}
	logging.Info("initialize screenshot client, screen_size=(%d, %d)", size.X, size.Y)
	cli.resize(size)
	return cli, nil
}

// Screenshot screenshot
func (cli *Client) Screenshot() (*image.RGBA, error) {
	size, err := cli.Size()
	if err != nil {
		return nil, err
	}
	cli.resize(size)
	return cli.img, cli.screenshot(cli.img)
}

func (cli *Client) resize(size image.Point) {
	if size.X == 0 && size.Y == 0 {
		return
	}
	cli.Lock()
	defer cli.Unlock()
	if cli.img != nil {
		src := cli.img.Rect.Max
		if src.X == size.X && src.Y == size.Y {
			return
		}
	}
	logging.Info("resize screenshot client image, size=(%d, %d)", size.X, size.Y)
	cli.img = image.NewRGBA(image.Rect(0, 0, size.X, size.Y))
}

// ShowCursor set draw cursor
func (cli *Client) ShowCursor(v bool) {
	cli.showCursor = v
}

// ClipboardSet set text to clipboard
func (cli *Client) ClipboardSet(text string) error {
	return clipboard.Set(text)
}

// ClipboardGet get text from clipboard
func (cli *Client) ClipboardGet() (string, error) {
	return clipboard.Get()
}
