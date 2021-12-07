# rdesktop

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