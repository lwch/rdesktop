# screenshot

golang截屏并生成image.RGBA图像，目前支持以下操作系统

- [x] linux(x11)
- [x] windows
- [ ] macos

## 使用方法

    cli, err := screenshot.New()
    cli.ShowCursor(true) // 是否绘制鼠标
    img, err := cli.Screenshot()
    // use of img