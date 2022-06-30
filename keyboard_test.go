package rdesktop

import "testing"

func TestKeyboardInput(t *testing.T) {
	cli, err := New()
	if err != nil {
		t.Fatal(err)
	}
	defer cli.Close()
	cli.ToggleKey("cmd", true)
	cli.ToggleKey("a", true)
	cli.ToggleKey("cmd", false)
	cli.ToggleKey("a", false)
}
