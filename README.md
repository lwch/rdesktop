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

    cli, err := rdesktop.New()
    cli.ShowCursor(true) // 是否绘制鼠标
    img, err := cli.Screenshot()
    // use of img

## 键鼠操作

暂未实现

## 剪切板操作

暂未实现