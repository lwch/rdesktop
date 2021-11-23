package screenshot

import (
	"image"
	"sync"
)

// Client screenshot client
type Client struct {
	sync.Mutex
	osBase
	img *image.RGBA
}

// New create client
func New() *Client {
	cli := &Client{}
	cli.init()
	size := cli.size()
	cli.resize(size)
	return cli
}

// Screenshot screenshot
func (cli *Client) Screenshot() (*image.RGBA, error) {
	size := cli.size()
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
