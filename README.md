# rdesktop

[![rdesktop](https://github.com/lwch/rdesktop/actions/workflows/build.yml/badge.svg)](https://github.com/lwch/rdesktop/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/lwch/rdesktop.svg)](https://pkg.go.dev/github.com/lwch/rdesktop)
[![Go Report Card](https://goreportcard.com/badge/github.com/lwch/rdesktop)](https://goreportcard.com/report/github.com/lwch/rdesktop)
[![go-mod](https://img.shields.io/github/go-mod/go-version/lwch/rdesktop)](https://github.com/lwch/rdesktop)
[![license](https://img.shields.io/github/license/lwch/rdesktop)](https://opensource.org/licenses/MIT)

golang远程桌面统一封装库，目前已支持功能

- [x] 截屏
- [ ] 键鼠操作
- [ ] 剪切板操作

已支持操作系统

- [x] linux(x11)
- [x] windows
- [ ] macos

## 截屏

    cli, _ := rdesktop.New()
    cli.ShowCursor(true) // 是否绘制鼠标
    img, err := cli.Screenshot()
    // use of img

## 鼠标操作

    cli, _ := rdesktop.New()
    cli.MouseMove(100, 100) // 将鼠标移动到100,100
    cli.ToggleMouse(rdesktop.MouseLeft, true) // 鼠标左键按下
    cli.ToggleMouse(rdesktop.MouseLeft, false) // 鼠标左键弹起
    cli.Scroll(0, -100) // 向下滚动100像素

## 键盘操作

    cli, _ := rdesktop.New()
    cli.ToggleKey("control", true) // 按下ctrl键
    cli.ToggleKey("a", true) // 按下a键
    cli.ToggleKey("control", false) // 弹起ctrl键
    cli.ToggleKey("a", false) // 弹起a键

## 剪切板操作

    cli, _ := rdesktop.New()
    cli.ClipboardSet("hello") // 将剪贴板内容设置为hello
    data, _ := cli.ClipboardGet() // 获取剪贴板内容