# screenshot

golang截屏并生成image.RGBA图像，目前支持以下操作系统

- [x] linux(x11)
- [ ] windows
- [ ] macos

## 使用方法

    cli := screenshot.New()
    img, err := cli.Screenshot()
    // use of img