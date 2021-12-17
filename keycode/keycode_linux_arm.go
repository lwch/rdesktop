//go:build (linux && arm) || (linux && arm64)
// +build linux,arm linux,arm64

package keycode

var Maps = map[string]int{}

func ForChar(k string) int {
	return 0
}
