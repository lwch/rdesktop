package screenshot

import (
	"image/jpeg"
	"os"
	"testing"
)

func TestScreenshot(t *testing.T) {
	cli, err := New()
	if err != nil {
		t.Fatal(err)
	}
	defer cli.Close()
	cli.ShowCursor(true)
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
