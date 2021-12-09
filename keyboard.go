package rdesktop

// https://github.com/go-vgo/robotgo/blob/master/key/goKey.h#L142
func checkKeycodes(key string) int {
	if len(key) == 1 {
		return keyCodeForChar(key)
	}
	return keyMaps[key]
}
