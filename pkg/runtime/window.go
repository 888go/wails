package runtime

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/options"
)

// WindowSetTitle 设置窗口的标题

// ff:窗口设置标题
// title:标题
// ctx:上下文
func WindowSetTitle(ctx context.Context, title string) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowSetTitle(title)
}

// WindowFullscreen 将窗口设置为全屏模式

// ff:窗口设置全屏
// ctx:上下文
func WindowFullscreen(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowFullscreen()
}

// WindowUnfullscreen 将窗口设置为未全屏状态

// ff:窗口取消全屏
// ctx:上下文
func WindowUnfullscreen(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowUnfullscreen()
}

// WindowCenter 将窗口居中于当前屏幕

// ff:窗口居中
// ctx:上下文
func WindowCenter(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowCenter()
}

// WindowReload 将重新加载窗口内容

// ff:窗口重载
// ctx:上下文
func WindowReload(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowReload()
}

// WindowReloadApp 将重新加载应用程序

// ff:窗口重载应用程序前端
// ctx:上下文
func WindowReloadApp(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowReloadApp()
}


// ff:窗口设置系统默认主题
// ctx:上下文
func WindowSetSystemDefaultTheme(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowSetSystemDefaultTheme()
}


// ff:窗口设置浅色主题
// ctx:上下文
func WindowSetLightTheme(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowSetLightTheme()
}


// ff:窗口设置深色主题
// ctx:上下文
func WindowSetDarkTheme(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowSetDarkTheme()
}

// WindowShow 如果窗口被隐藏，则显示窗口

// ff:窗口显示
// ctx:上下文
func WindowShow(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowShow()
}

// WindowHide the window

// ff:窗口隐藏
// ctx:上下文
func WindowHide(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowHide()
}

// WindowSetSize 设置窗口的大小

// ff:窗口设置尺寸
// height:高
// width:宽
// ctx:上下文
func WindowSetSize(ctx context.Context, width int, height int) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowSetSize(width, height)
}


// ff:窗口取尺寸
// ctx:上下文
func WindowGetSize(ctx context.Context) (int, int) {
	appFrontend := getFrontend(ctx)
	return appFrontend.WindowGetSize()
}

// WindowSetMinSize 设置窗口的最小尺寸

// ff:窗口设置最小尺寸
// height:高
// width:宽
// ctx:上下文
func WindowSetMinSize(ctx context.Context, width int, height int) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowSetMinSize(width, height)
}

// WindowSetMaxSize 设置窗口的最大尺寸

// ff:窗口设置最大尺寸
// height:高
// width:宽
// ctx:上下文
func WindowSetMaxSize(ctx context.Context, width int, height int) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowSetMaxSize(width, height)
}

// WindowSetAlwaysOnTop 设置窗口是否总在顶部展示

// ff:窗口设置置顶
// b:置顶
// ctx:上下文
func WindowSetAlwaysOnTop(ctx context.Context, b bool) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowSetAlwaysOnTop(b)
}

// WindowSetPosition 设置窗口的位置

// ff:窗口设置位置
// y:
// x:
// ctx:上下文
func WindowSetPosition(ctx context.Context, x int, y int) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowSetPosition(x, y)
}


// ff:窗口取位置
// ctx:上下文
func WindowGetPosition(ctx context.Context) (int, int) {
	appFrontend := getFrontend(ctx)
	return appFrontend.WindowGetPosition()
}

// WindowMaximise 窗口最大化

// ff:窗口最大化
// ctx:上下文
func WindowMaximise(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowMaximise()
}

// WindowToggleMaximise：最大化窗口

// ff:窗口最大化切换
// ctx:上下文
func WindowToggleMaximise(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowToggleMaximise()
}

// WindowUnmaximise 将窗口取消最大化

// ff:窗口取消最大化
// ctx:上下文
func WindowUnmaximise(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowUnmaximise()
}

// WindowMinimise 窗口最小化

// ff:窗口最小化
// ctx:上下文
func WindowMinimise(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowMinimise()
}

// WindowUnminimise 将窗口从最小化状态恢复

// ff:窗口取消最小化
// ctx:上下文
func WindowUnminimise(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowUnminimise()
}

// WindowIsFullscreen 获取窗口状态，判断窗口是否全屏

// ff:窗口是否全屏
// ctx:上下文
func WindowIsFullscreen(ctx context.Context) bool {
	appFrontend := getFrontend(ctx)
	return appFrontend.WindowIsFullscreen()
}

// WindowIsMaximised 获取窗口状态是否为最大化

// ff:窗口是否最大化
// ctx:上下文
func WindowIsMaximised(ctx context.Context) bool {
	appFrontend := getFrontend(ctx)
	return appFrontend.WindowIsMaximised()
}

// WindowIsMinimised 获取窗口状态是否为最小化

// ff:窗口是否最小化
// ctx:上下文
func WindowIsMinimised(ctx context.Context) bool {
	appFrontend := getFrontend(ctx)
	return appFrontend.WindowIsMinimised()
}

// WindowIsNormal 获取窗口状态是否为正常窗口

// ff:窗口是否为正常
// ctx:上下文
func WindowIsNormal(ctx context.Context) bool {
	appFrontend := getFrontend(ctx)
	return appFrontend.WindowIsNormal()
}

// WindowExecJS在window环境中执行给定的Js代码

// ff:窗口执行JS
// js:js代码
// ctx:上下文
func WindowExecJS(ctx context.Context, js string) {
	appFrontend := getFrontend(ctx)
	appFrontend.ExecJS(js)
}


// ff:窗口设置背景色
// A:
// B:
// G:
// R:
// ctx:上下文
func WindowSetBackgroundColour(ctx context.Context, R, G, B, A uint8) {
	appFrontend := getFrontend(ctx)
	col := &options.RGBA{
		R: R,
		G: G,
		B: B,
		A: A,
	}
	appFrontend.WindowSetBackgroundColour(col)
}


// ff:窗口打开打印对话框
// ctx:上下文
func WindowPrint(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowPrint()
}
