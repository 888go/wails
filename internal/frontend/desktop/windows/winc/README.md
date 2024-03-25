# # winc

**这是一个仅出于与[Wails](https://github.com/wailsapp/wails)集成目的而从[tadvi/winc](https://github.com/tadvi/winc)分支出来的项目。请注意，此存储库不提供任何支持 **

Windows操作系统上Go GUI应用的通用库。它仅适用于Windows操作系统，因此相比于其他一些用于Go的UI库，该库更小巧。

设计目标：极简主义和简单性。
## # 依赖

除了 Go 标准库之外，没有其他依赖项。
## # 构建

如果你想将图标文件和其他资源打包到二进制文件中，建议使用 **rsrc** 工具：

```bash
rsrc -manifest app.manifest -ico=app.ico,application_edit.ico,application_error.ico -o rsrc.syso
```

这里的 `app.manifest` 是一个格式为 XML 的文件：

```xml
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0">
    <assemblyIdentity version="1.0.0.0" processorArchitecture="*" name="App" type="win32"/>
    <dependency>
        <dependentAssembly>
            <assemblyIdentity type="win32" name="Microsoft.Windows.Common-Controls" version="6.0.0.0" processorArchitecture="*" publicKeyToken="6595b64144ccf1df" language="*"/>
        </dependentAssembly>
    </dependency>
</assembly>
```

大多数 Windows 应用程序不会显示命令提示符。在构建 Go 项目时，使用标志来指示它是一个 Windows GUI 可执行文件：

```bash
go build -ldflags="-H windowsgui"
```
## # 最佳学习如何使用该库的方法是查看所包含的**示例**项目。
## # 配置

1. 确保您已安装并设置了有效的 Go 环境。详情请参阅下方页面：
   http://golang.org/doc/install

2. 使用以下命令获取 Wails 库的 Windows 桌面前端内部依赖包：
   go get github.com/wailsapp/wails/v2/internal/frontend/desktop/windows/winc
## # 图标

当使用 rsrc 将图标打包到二进制文件时，它会显示已打包图标的 ID。

```
rsrc -manifest app.manifest -ico=app.ico,lightning.ico,edit.ico,application_error.ico -o rsrc.syso
清单 ID: 1
图标 app.ico ID: 10
图标 lightning.ico ID: 13
图标 edit.ico ID: 16
图标 application_error.ico ID: 19
```

使用 ID 引用已打包的图标。

```
const myIcon = 13

btn.SetResIcon(myIcon) // 设置按钮上的图标。
```

附带的示例源代码通过 `release.bat` 文件进行基本构建。请注意，图标 ID 的顺序相关。因此，如果您更改 -ico 标志中的顺序，则图标 ID 会不同。如果要保持相同的顺序，只需将新图标添加到 -ico 逗号分隔列表的末尾即可。
## # 布局管理器

SimpleDock 是默认的布局管理器。

当前的停靠和拆分视图设计允许构建简单的应用程序，但如果需要在多个方向上有多个拆分视图，你可能需要创建自己的布局管理器。

重要的一点是要在 SimpleDock 中有一个控件设置为 **Fill** 模式停靠。未设置任何停靠方式的控件会通过 SetPos() 函数进行定位。因此，你可以将一个面板设置为停靠在顶部，然后在该面板内使用另一个停靠来排列控件，或者使用 SetPos() 将控件放置在固定位置。

![示例布局：带有两个工具栏和状态栏](dock_topbottom.png)

这是一个基本布局。在这里，工具栏和状态栏的位置，你可以替换为可调整大小的面板或其他控件。面板内部可以拥有自己的 Dock，用于排列其中的其他控件。

![示例布局：带有两个工具栏和左侧导航](dock_topleft.png)

这是一个在左侧有额外控件的布局。通常左侧是树形视图或列表视图。

规则很简单：你要么使用 SimpleDock 来停靠控件，要么使用 SetPos() 将它们设置在固定位置。就是这样。

在未来某个时间点，**winc** 可能会获得更复杂的布局管理器。
## # 对话框屏幕

对话框屏幕不基于Windows资源文件(.rc)。它们只是在固定坐标位置放置了控件的窗口。这种方式适用于包含10至14个控件的对话框屏幕，效果良好。
# # 最小示例

```go
package main

import (
	"github.com/wailsapp/wails/v2/internal/frontend/desktop/windows/winc"
)

func main() {
	// 创建主窗口
	mainWindow := winc.NewForm(nil)
	mainWindow.SetSize(400, 300) // (宽度, 高度)
	mainWindow.SetText("Hello World 示例")

	// 创建编辑框
	edt := winc.NewEdit(mainWindow)
	edt.SetPos(10, 20)
	// 大多数控件默认都有大小，除非调用 SetSize
	edt.SetText("编辑文本")

	// 创建按钮
	btn := winc.NewPushButton(mainWindow)
	btn.SetText("显示或隐藏")
	btn.SetPos(40, 50) // (x, y)
	btn.SetSize(100, 40) // (宽度, 高度)
	btn.OnClick().Bind(func(e *winc.Event) {
		if edt.Visible() {
			edt.Hide()
		} else {
			edt.Show()
		}
	})

	// 主窗口居中并显示
	mainWindow.Center()
	mainWindow.Show()
	mainWindow.OnClose().Bind(wndOnClose)

	// 启动事件循环
	winc.RunMainLoop()
}

func wndOnClose(arg *winc.Event) {
	winc.Exit()
}
```

![](examples/hello.png)

运行 `sample_minimal` 的结果。
## # 创建自己的控件

基于现有的结构和事件模型创建您自己的控件是一个很好的实践。库中包含了一些以这种方式构建的控件，例如：IconButton（button.go）、ErrorPanel（panel.go）、MultiEdit（edit.go）等。在构建自己的控件之前，请先参考已有的控件作为示例。

在设计自定义控件时，请记住，类型需要在 Go 语言与 Win32 API 之间进行转换，这通常涉及到字符串 UTF8 和 UTF16 的转换，但还有其他类型的转换需要注意。

在开发自定义控件时，您可能还需要：

```go
import "github.com/wailsapp/wails/v2/internal/frontend/desktop/windows/winc/w32"
```

w32 包含了 Win32 API 的底层常量和函数。

请参阅 **sample_control** 示例，了解自定义窗口的构建方式。
## # 伙伴包

[适用于 Windows 系统托盘图标的 Go 语言包，包含菜单和通知功能](https://github.com/tadvi/systray)
## # 致谢

该库基于以下项目构建：

[AllenDang/gform：适用于 Go 的 Windows GUI 框架](https://github.com/AllenDang/gform)

**winc** 从 **gform** 中采纳了大部分设计决策，并在此基础上添加了更多控件和代码示例。
