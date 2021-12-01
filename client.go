package screenshot

import (
	"image"
	"sync"

	"github.com/lwch/logging"
)

// Client screenshot client
type Client struct {
	sync.Mutex
	osBase
	img *image.RGBA
}

// New create client
func New() (*Client, error) {
	cli := &Client{}
	err := cli.init()
	if err != nil {
		return nil, err
	}
	size, err := cli.size()
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
	size, err := cli.size()
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
	cli.img = image.NewRGBA(image.Rect(0, 0, size.X, size.Y))
}
