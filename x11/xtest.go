package x11

import (
	"encoding/binary"
	"errors"
)

// TestFakeInput mouse input or keyboard input
func (cli *Client) TestFakeInput(t, detail byte, x, y uint16) error {
	opcode := cli.opcode("XTEST")
	if opcode == 0 {
		return errors.New("extension XTEST not supported")
	}
	screen := cli.info.roots[0]
	var data [36]byte
	data[0] = opcode                        // opcode
	data[1] = 2                             // XTestFakeInput
	binary.BigEndian.PutUint16(data[2:], 9) // size
	data[4] = t                             // type
	data[5] = detail                        // detail
	// pad 2 bytes
	binary.BigEndian.PutUint32(data[8:], 0)                    // time
	binary.BigEndian.PutUint32(data[12:], uint32(screen.root)) // window
	// pad 4 bytes
	// pad 4 bytes
	binary.BigEndian.PutUint16(data[24:], x) // root_x
	binary.BigEndian.PutUint16(data[26:], y) // root_y
	// pad 4 bytes
	// pad 2 bytes
	// pad 1 byte
	// device_id 1 byte
	return cli.callNoResp(data[:])
}
