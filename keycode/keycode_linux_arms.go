//go:build linux && (arm || arm64)
// +build linux
// +build arm arm64

package keycode

var Maps = map[string]int{}

func ForChar(k string) int {
	return 0
}
