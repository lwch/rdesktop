package x11

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/lwch/logging"
)

// Client x11 client
type Client struct {
	callLock   sync.Mutex
	conn       net.Conn
	info       setupInfo
	extLock    sync.RWMutex
	extensions map[string]byte
	keyMapping []byte
	keyPer     byte
}

// New new client
func New() (*Client, error) {
	disp := os.Getenv("DISPLAY")
	idx := strings.TrimPrefix(disp, ":")
	if len(idx) == 0 {
		idx = "1"
		logging.Info("default to DISPLAY 1")
	}
	conn, err := net.Dial("unix", path.Join("/tmp", ".X11-unix", fmt.Sprintf("X%s", idx)))
	if err != nil {
		return nil, err
	}
	cli := &Client{
		conn:       conn,
		extensions: make(map[string]byte),
	}
	err = cli.handshake()
	if err != nil {
		conn.Close()
		return nil, err
	}
	// 初始化过程详见：https://gitlab.freedesktop.org/xorg/lib/libx11/-/blob/master/src/KeyBind.c#L257
	n := cli.info.maxKeycode - cli.info.minKeycode + 1
	cli.keyMapping, cli.keyPer, err = cli.GetKeyboardMapping(cli.info.minKeycode, n)
	if err != nil {
		conn.Close()
		return nil, err
	}
	return cli, nil
}

// Close close client
func (cli *Client) Close() {
	cli.conn.Close()
}

func pad(n int) int {
	return (n + 3) & ^3
}

func (cli *Client) opcode(name string) byte {
	cli.extLock.RLock()
	code := cli.extensions[name]
	cli.extLock.RUnlock()
	if code == 0 {
		ok, code, err := cli.queryExtension(name)
		if err != nil {
			return 0
		}
		if !ok {
			return 0
		}
		cli.extLock.Lock()
		cli.extensions[name] = code
		cli.extLock.Unlock()
		return code
	}
	return code
}

func (cli *Client) call(data []byte) ([]byte, error) {
	cli.callLock.Lock()
	defer cli.callLock.Unlock()
	_, err := io.Copy(cli.conn, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("send data: %v", err)
	}
	var hdr [32]byte
	_, err = io.ReadFull(cli.conn, hdr[:])
	if err != nil {
		return nil, fmt.Errorf("read header: %v", err)
	}
	size := binary.BigEndian.Uint32(hdr[4:])
	if size == 0 {
		return hdr[:], nil
	}
	data = make([]byte, 32+size*4)
	copy(data, hdr[:])
	_, err = io.ReadFull(cli.conn, data[32:])
	if err != nil {
		return nil, fmt.Errorf("read body: %v", err)
	}
	return data, nil
}

func (cli *Client) callNoResp(data []byte) error {
	cli.callLock.Lock()
	defer cli.callLock.Unlock()
	_, err := io.Copy(cli.conn, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("send data: %v", err)
	}
	return nil
}

func (cli *Client) queryExtension(name string) (bool, byte, error) {
	size := 8 + pad(len(name))
	data := make([]byte, size)
	data[0] = 98 // type
	// pad 1 byte
	binary.BigEndian.PutUint16(data[2:], uint16(size/4))    // body size
	binary.BigEndian.PutUint16(data[4:], uint16(len(name))) // name length
	// pad 2 bytes
	copy(data[8:], name)
	var err error
	data, err = cli.call(data)
	if err != nil {
		return false, 0, err
	}
	err = errCheck(data)
	if err != nil {
		return false, 0, err
	}
	return data[8] != 0, data[9], nil
}

func errCheck(data []byte) error {
	// TODO error parse
	if data[0] == 0 {
		logging.Error("error:\n%s", hex.Dump(data))
		return errors.New("error occurred")
	}
	return nil
}
