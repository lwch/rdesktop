package screenshot

import "testing"

func TestLinux(t *testing.T) {
	cli := New()
	cli.Screenshot()
}
