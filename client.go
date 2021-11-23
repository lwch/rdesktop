package screenshot

import "image"

type Client struct {
	osBase
}

func New() *Client {
	cli := &Client{}
	cli.init()
	return cli
}

func (cli *Client) Screenshot() (*image.RGBA, error) {
	size := cli.size()
	_ = size
	return nil, nil
}
