package x11

import (
	"encoding/binary"
	"errors"
	"io"
	"os"
	"path"
)

// from X11/Xauth.h
type xauth struct {
	family  uint16
	addrLen uint16
	addr    []byte
	numLen  uint16
	num     []byte
	nameLen uint16
	name    []byte
	dataLen uint16
	data    []byte
}

// https://gitlab.freedesktop.org/xorg/app/xauth/-/blob/master/process.c#L410
func readAuth() (xauth, error) {
	var auth xauth
	fname := os.Getenv("XAUTHORITY")
	if len(fname) == 0 {
		home := os.Getenv("HOME")
		if len(home) == 0 {
			return auth, errors.New("can not get env of HOME directory")
		}
		fname = path.Join(home, ".Xauthority")
	}
	f, err := os.Open(fname)
	if err != nil {
		return auth, err
	}
	defer f.Close()
	err = binary.Read(f, binary.BigEndian, &auth.family)
	if err != nil {
		return auth, err
	}
	err = readBytes(f, &auth.addrLen, &auth.addr)
	if err != nil {
		return auth, err
	}
	err = readBytes(f, &auth.numLen, &auth.num)
	if err != nil {
		return auth, err
	}
	err = readBytes(f, &auth.nameLen, &auth.name)
	if err != nil {
		return auth, err
	}
	err = readBytes(f, &auth.dataLen, &auth.data)
	if err != nil {
		return auth, err
	}
	return auth, nil
}

func readBytes(r io.Reader, l *uint16, data *[]byte) error {
	err := binary.Read(r, binary.BigEndian, l)
	if err != nil {
		return err
	}
	*data = make([]byte, *l)
	_, err = io.ReadFull(r, *data)
	return err
}
