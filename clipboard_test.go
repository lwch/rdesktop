package rdesktop

import (
	"testing"
)

func TestClipboard(t *testing.T) {
	cli, err := New()
	if err != nil {
		t.Fatal(err)
	}
	defer cli.Close()
	err = cli.ClipboardSet("hello")
	if err != nil {
		t.Fatal(err)
	}
	data, err := cli.ClipboardGet()
	if err != nil {
		t.Fatal(err)
	}
	if data != "hello" {
		t.Fatalf("unexpected clipboard data")
	}
}
