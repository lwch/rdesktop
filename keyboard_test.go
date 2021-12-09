package rdesktop

import "testing"

func TestKeyboardInput(t *testing.T) {
	cli, err := New()
	if err != nil {
		t.Fatal(err)
	}
	defer cli.Close()
	cli.ToggleKey("control", true)
	cli.ToggleKey("a", true)
	cli.ToggleKey("control", false)
	cli.ToggleKey("a", false)
}
