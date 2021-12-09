package x11

import (
	"encoding/binary"
)

// KeysymToKeycode https://gitlab.freedesktop.org/xorg/lib/libx11/-/blob/master/src/KeyBind.c#L131
func (cli *Client) KeysymToKeycode(code int) byte {
	for j := byte(0); j < cli.keyPer; j++ {
		for i := cli.info.minKeycode; i <= cli.info.maxKeycode; i++ {
			if cli.KeyCodetoKeySym(i, int(j)) == code {
				return byte(i)
			}
		}
	}
	return 0
}

func intValue(data []byte, offset int) int {
	return int(binary.BigEndian.Uint32(data[offset*4:]))
}

// KeyCodetoKeySym https://gitlab.freedesktop.org/xorg/lib/libx11/-/blob/master/src/KeyBind.c#L85
func (cli *Client) KeyCodetoKeySym(code keyCode, col int) int {
	if col < 0 || (col >= int(cli.keyPer) && col > 3) ||
		code < cli.info.minKeycode || code > cli.info.maxKeycode {
		return 0
	}
	per := cli.keyPer
	offset := (int(code) - int(cli.info.minKeycode)) * int(cli.keyPer)
	if col < 4 {
		if col > 1 {
			for {
				if per > 2 && intValue(cli.keyMapping, offset+int(per)-1) == 0 {
					per--
					continue
				}
				break
			}
			if per < 3 {
				col -= 2
			}
		}
		if int(per) <= (col|1) || intValue(cli.keyMapping, offset+(col|1)) == 0 {
			n := intValue(cli.keyMapping, offset+(col&(^1)))
			return n & 0xffff
		}
	}
	return intValue(cli.keyMapping, offset+col)
}
