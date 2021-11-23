package screenshot

import (
	"image/jpeg"
	"os"
	"testing"
)

func TestLinux(t *testing.T) {
	cli := New()
	defer cli.Close()
	img, err := cli.Screenshot()
	if err != nil {
		t.Fatal(err)
	}
	f, err := os.Create("screenshot.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	err = jpeg.Encode(f, img, nil)
	if err != nil {
		t.Fatal(err)
	}
}
