# rdesktop

[![rdesktop](https://github.com/lwch/rdesktop/actions/workflows/build.yml/badge.svg)](https://github.com/lwch/rdesktop/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/lwch/rdesktop.svg)](https://pkg.go.dev/github.com/lwch/rdesktop)
[![Go Report Card](https://goreportcard.com/badge/github.com/lwch/rdesktop)](https://goreportcard.com/report/github.com/lwch/rdesktop)
[![go-mod](https://img.shields.io/github/go-mod/go-version/lwch/rdesktop)](https://github.com/lwch/rdesktop)
[![license](https://img.shields.io/github/license/lwch/rdesktop)](https://opensource.org/licenses/MIT)

golang desktop controller library

- [x] screenshot
- [x] keyboard/mouse events
- [x] scroll events
- [x] clipboard get/set(only supported text data)

supported system

- [x] linux(x11)
- [x] windows
- [x] macos

## screenshot

    cli, _ := rdesktop.New()
    cli.ShowCursor(true) // show the cursor image
    img, err := cli.Screenshot()
    // use of img

## mouse

    cli, _ := rdesktop.New()
    cli.MouseMove(100, 100) // move mouse to 100,100
    cli.ToggleMouse(rdesktop.MouseLeft, true) // mouse left button press down
    cli.ToggleMouse(rdesktop.MouseLeft, false) // mouse left button press up
    cli.Scroll(0, -100) // scroll down 100 pixel

## keyboard

    cli, _ := rdesktop.New()
    cli.ToggleKey("control", true) // press down ctrl
    cli.ToggleKey("a", true) // press down a
    cli.ToggleKey("control", false) // press up ctrl
    cli.ToggleKey("a", false) // press up a

## clipboard

    cli, _ := rdesktop.New()
    cli.ClipboardSet("hello") // set "hello" text to clipboard
    data, _ := cli.ClipboardGet() // get clipboard data