# TODO

## Vesrion 1
- x Add icon
- x Build for Intel
- x Build for M
- x Package it
- x Release it

## Vesrion 2 
- Adjust initial window size
- Fix Theme (Font size and color, x Alignment, etc...)
- Change backround color maybe???
- Add Break 5/20 functionality
- 60 seconds?!?
- Pressing any "Start" button should not "Pause/Stop" current counter.
- Add the v2 into the app icon
- Add Applicaiton icon (now showing the script editor icon)
- Edit title of the app when running Pomodoro / Pause
- Get a Pause UTF code in button / Title
- Fix Segfault:
```sh
mcourtois@moe-home-imac go-pomodoro % go run . -v
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x10 pc=0x10043c7e9]

goroutine 18 [running, locked to thread]:
fyne.io/fyne/v2/app.(*settings).Theme(0xffffffffffffffff?)
        /Users/mcourtois/go/pkg/mod/fyne.io/fyne/v2@v2.4.4/app/settings.go:66 +0x29
fyne.io/fyne/v2/theme.current()
        /Users/mcourtois/go/pkg/mod/fyne.io/fyne/v2@v2.4.4/theme/theme.go:179 +0x31
fyne.io/fyne/v2/theme.Padding()
        /Users/mcourtois/go/pkg/mod/fyne.io/fyne/v2@v2.4.4/theme/size.go:117 +0x13
fyne.io/fyne/v2/layout.(*boxLayout).MinSize(0xc000112034, {0xc000124f80, 0x5, 0x0?})
        /Users/mcourtois/go/pkg/mod/fyne.io/fyne/v2@v2.4.4/layout/boxlayout.go:125 +0x26
fyne.io/fyne/v2.(*Container).MinSize(0x0?)
        /Users/mcourtois/go/pkg/mod/fyne.io/fyne/v2@v2.4.4/container.go:90 +0x58
fyne.io/fyne/v2/layout.(*stackLayout).MinSize(0x0?, {0xc000035870?, 0x1, 0x0?})
        /Users/mcourtois/go/pkg/mod/fyne.io/fyne/v2@v2.4.4/layout/stacklayout.go:48 +0x9d
fyne.io/fyne/v2.(*Container).MinSize(0xc001c5dd28?)
        /Users/mcourtois/go/pkg/mod/fyne.io/fyne/v2@v2.4.4/container.go:90 +0x58
fyne.io/fyne/v2/internal/driver/glfw.(*glCanvas).MinSize(0xc000128000)
        /Users/mcourtois/go/pkg/mod/fyne.io/fyne/v2@v2.4.4/internal/driver/glfw/canvas.go:72 +0x96
fyne.io/fyne/v2/internal/driver/common.(*Canvas).EnsureMinSize(0xc000128000)
        /Users/mcourtois/go/pkg/mod/fyne.io/fyne/v2@v2.4.4/internal/driver/common/canvas.go:92 +0x89
fyne.io/fyne/v2/internal/driver/glfw.(*gLDriver).repaintWindow.func1()
        /Users/mcourtois/go/pkg/mod/fyne.io/fyne/v2@v2.4.4/internal/driver/glfw/loop.go:215 +0x2b
fyne.io/fyne/v2/internal/driver/glfw.(*window).RunWithContext(0xc0002d7f10?, 0xc001c5de48)
        /Users/mcourtois/go/pkg/mod/fyne.io/fyne/v2@v2.4.4/internal/driver/glfw/window.go:923 +0x43
fyne.io/fyne/v2/internal/driver/glfw.(*gLDriver).repaintWindow(0xc0000821c0?, 0xc001c5df38?)
        /Users/mcourtois/go/pkg/mod/fyne.io/fyne/v2@v2.4.4/internal/driver/glfw/loop.go:214 +0x45
fyne.io/fyne/v2/internal/driver/glfw.(*gLDriver).drawSingleFrame(0xc001c5df90?)
        /Users/mcourtois/go/pkg/mod/fyne.io/fyne/v2@v2.4.4/internal/driver/glfw/loop.go:102 +0x156
fyne.io/fyne/v2/internal/driver/glfw.(*gLDriver).startDrawThread.func1()
        /Users/mcourtois/go/pkg/mod/fyne.io/fyne/v2@v2.4.4/internal/driver/glfw/loop.go:270 +0x1d5
created by fyne.io/fyne/v2/internal/driver/glfw.(*gLDriver).startDrawThread in goroutine 1
        /Users/mcourtois/go/pkg/mod/fyne.io/fyne/v2@v2.4.4/internal/driver/glfw/loop.go:246 +0xbb
exit status 2
```


## Version 3
- Try to make it work for iOS
- Add Clock Animation
- Add sound effect
