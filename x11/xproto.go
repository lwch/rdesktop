package x11

import (
	"encoding/binary"
	"image"
)

// https://tronche.com/gui/x/xlib/graphics/XGetImage.html
const (
	xyBitmap  = 0 // depth 1
	xyPixmap  = 1 // depth==drawable depth
	zPixmap   = 2 // depth==drawable depth
	allPlanes = 0xffffffff
)

// GetImage draw image from x11
func (cli *Client) GetImage(img *image.RGBA) error {
	screen := cli.info.roots[0]
	var data [20]byte
	data[0] = 73                                              // opcode
	data[1] = zPixmap                                         // format
	binary.BigEndian.PutUint16(data[2:], 5)                   // length
	binary.BigEndian.PutUint32(data[4:], uint32(screen.root)) // drawable
	// x: 2bytes
	// y: 2bytes
	binary.BigEndian.PutUint16(data[12:], screen.widthInPixels)  // width
	binary.BigEndian.PutUint16(data[14:], screen.heightInPixels) // height
	binary.BigEndian.PutUint32(data[16:], allPlanes)
	ret, err := cli.call(data[:])
	if err != nil {
		return err
	}
	err = errCheck(ret)
	if err != nil {
		return err
	}
	_ = ret[1] // depth
	// TODO: depth is not 24
	offset := 0
	for y := 0; y < int(screen.heightInPixels); y++ {
		for x := 0; x < int(screen.widthInPixels); x++ {
			img.Pix[offset+2] = ret[offset+32] // b
			img.Pix[offset+1] = ret[offset+33] // g
			img.Pix[offset] = ret[offset+34]   // r
			offset += 4
		}
	}
	return nil
}
