// code from https://github.com/go-vgo/robotgo/blob/master/clipboard/clipboard_unix.go

package clipboard

import (
	"os/exec"
	"strings"
)

const (
	xsel  = "xsel"
	xclip = "xclip"
)

var (
	copyArgs, pasteArgs []string

	xselCopyArgs  = []string{xsel, "--input", "--clipboard"}
	xselPasteArgs = []string{xsel, "--output", "--clipboard"}

	xclipCopyArgs  = []string{xclip, "-in", "-selection", "clipboard"}
	xclipPasteArgs = []string{xclip, "-out", "-selection", "clipboard"}
)

func init() {
	if _, err := exec.LookPath(xclip); err == nil {
		copyArgs = xclipCopyArgs
		pasteArgs = xclipPasteArgs
		return
	}
	if _, err := exec.LookPath(xsel); err == nil {
		copyArgs = xselCopyArgs
		pasteArgs = xselPasteArgs
		return
	}
	Unsupported = true
}

// Set set text to clipboard
func Set(text string) error {
	if Unsupported {
		return ErrUnsupport
	}
	cmd := exec.Command(copyArgs[0], copyArgs[1:]...)
	cmd.Stdin = strings.NewReader(text)
	return cmd.Run()
}

// Get get clipboard text
func Get() (string, error) {
	if Unsupported {
		return "", ErrUnsupport
	}
	cmd := exec.Command(pasteArgs[0], pasteArgs[1:]...)
	data, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(data), nil
}
