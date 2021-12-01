package x11

import (
	"encoding/binary"
	"errors"
	"image"
)

var ErrNoScreens = errors.New("no screens found")

// GetSize get screen size
func (cli *Client) GetSize() (image.Point, error) {
	opcode := cli.opcode("XINERAMA")
	if opcode == 0 {
		return image.Point{}, errors.New("extension XINERAMA not supported")
	}
	var data [4]byte
	data[0] = opcode                        // opcode
	data[1] = 5                             // QueryScreens
	binary.BigEndian.PutUint16(data[2:], 1) // size
	ret, err := cli.call(data[:])
	if err != nil {
		return image.Point{}, err
	}
	err = errCheck(ret)
	if err != nil {
		return image.Point{}, err
	}
	count := binary.BigEndian.Uint32(ret[8:])
	if count == 0 {
		return image.Point{}, ErrNoScreens
	}
	// 默认取第一个屏幕的尺寸
	// 32: x_org
	// 34: y_org
	width := binary.BigEndian.Uint16(ret[36:])
	height := binary.BigEndian.Uint16(ret[38:])
	return image.Point{
		X: int(width),
		Y: int(height),
	}, nil
}
