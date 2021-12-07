package x11

import (
	"encoding/binary"
	"errors"
	"image"
)

// GetCursorImage draw cursor image from x11
func (cli *Client) GetCursorImage(img *image.RGBA) error {
	opcode := cli.opcode("XFIXES")
	if opcode == 0 {
		return errors.New("extension XFIXES not supported")
	}
	var data [4]byte
	data[0] = opcode                        // opcode
	data[1] = 4                             // GetCursorImage
	binary.BigEndian.PutUint16(data[2:], 1) // size
	ret, err := cli.call(data[:])
	if err != nil {
		return err
	}
	err = errCheck(ret)
	if err != nil {
		return err
	}
	x := binary.BigEndian.Uint16(ret[8:])
	y := binary.BigEndian.Uint16(ret[10:])
	width := binary.BigEndian.Uint16(ret[12:])
	height := binary.BigEndian.Uint16(ret[14:])
	xHot := binary.BigEndian.Uint16(ret[16:])
	yHot := binary.BigEndian.Uint16(ret[18:])
	if x >= xHot {
		x -= xHot
	}
	if y >= yHot {
		y -= yHot
	}
	dOffset := 32
	offset := int(y)*img.Rect.Max.X*4 + int(x)*4
	for dy := 0; dy < int(height); dy++ {
		next := offset + img.Rect.Max.X*4
		for dx := 0; dx < int(width); dx++ {
			if ret[dOffset+3] > 0 { // a
				img.Pix[offset+2] = ret[dOffset]   // b
				img.Pix[offset+1] = ret[dOffset+1] // g
				img.Pix[offset] = ret[dOffset+2]   // r
			}
			offset += 4
			dOffset += 4
		}
		offset = next
	}
	return nil
}
