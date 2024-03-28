package runtime

import (
	"context"

	"github.com/888go/wails/pkg/options"
)

// WindowSetTitle 设置窗口的标题

// ff:窗口设置标题
// title:标题
// ctx:上下文
func X窗口设置标题(上下文 context.Context, 标题 string) {
	appFrontend := getFrontend(上下文)
	appFrontend.WindowSetTitle(标题)
}

// WindowFullscreen 将窗口设置为全屏模式

// ff:窗口设置全屏
// ctx:上下文
func X窗口设置全屏(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.WindowFullscreen()
}

// WindowUnfullscreen 将窗口设置为未全屏状态

// ff:窗口取消全屏
// ctx:上下文
func X窗口取消全屏(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.WindowUnfullscreen()
}

// WindowCenter 将窗口居中于当前屏幕

// ff:窗口居中
// ctx:上下文
func X窗口居中(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.WindowCenter()
}

// WindowReload 将重新加载窗口内容

// ff:窗口重载
// ctx:上下文
func X窗口重载(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.WindowReload()
}

// WindowReloadApp 将重新加载应用程序

// ff:窗口重载应用程序前端
// ctx:上下文
func X窗口重载应用程序前端(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.WindowReloadApp()
}


// ff:窗口设置系统默认主题
// ctx:上下文
func X窗口设置系统默认主题(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.WindowSetSystemDefaultTheme()
}


// ff:窗口设置浅色主题
// ctx:上下文
func X窗口设置浅色主题(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.WindowSetLightTheme()
}


// ff:窗口设置深色主题
// ctx:上下文
func X窗口设置深色主题(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.WindowSetDarkTheme()
}

// WindowShow 如果窗口被隐藏，则显示窗口

// ff:窗口显示
// ctx:上下文
func X窗口显示(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.WindowShow()
}

// WindowHide the window

// ff:窗口隐藏
// ctx:上下文
func X窗口隐藏(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.WindowHide()
}

// WindowSetSize 设置窗口的大小

// ff:窗口设置尺寸
// height:高
// width:宽
// ctx:上下文
func X窗口设置尺寸(上下文 context.Context, 宽 int, 高 int) {
	appFrontend := getFrontend(上下文)
	appFrontend.WindowSetSize(宽, 高)
}


// ff:窗口取尺寸
// ctx:上下文
func X窗口取尺寸(上下文 context.Context) (int, int) {
	appFrontend := getFrontend(上下文)
	return appFrontend.WindowGetSize()
}

// WindowSetMinSize 设置窗口的最小尺寸

// ff:窗口设置最小尺寸
// height:高
// width:宽
// ctx:上下文
func X窗口设置最小尺寸(上下文 context.Context, 宽 int, 高 int) {
	appFrontend := getFrontend(上下文)
	appFrontend.WindowSetMinSize(宽, 高)
}

// WindowSetMaxSize 设置窗口的最大尺寸

// ff:窗口设置最大尺寸
// height:高
// width:宽
// ctx:上下文
func X窗口设置最大尺寸(上下文 context.Context, 宽 int, 高 int) {
	appFrontend := getFrontend(上下文)
	appFrontend.WindowSetMaxSize(宽, 高)
}

// WindowSetAlwaysOnTop 设置窗口是否总在顶部展示

// ff:窗口设置置顶
// b:置顶
// ctx:上下文
func X窗口设置置顶(上下文 context.Context, 置顶 bool) {
	appFrontend := getFrontend(上下文)
	appFrontend.WindowSetAlwaysOnTop(置顶)
}

// WindowSetPosition 设置窗口的位置

// ff:窗口设置位置
// y:
// x:
// ctx:上下文
func X窗口设置位置(上下文 context.Context, x int, y int) {
	appFrontend := getFrontend(上下文)
	appFrontend.WindowSetPosition(x, y)
}


// ff:窗口取位置
// ctx:上下文
func X窗口取位置(上下文 context.Context) (int, int) {
	appFrontend := getFrontend(上下文)
	return appFrontend.WindowGetPosition()
}

// WindowMaximise 窗口最大化

// ff:窗口最大化
// ctx:上下文
func X窗口最大化(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.WindowMaximise()
}

// WindowToggleMaximise：最大化窗口

// ff:窗口最大化切换
// ctx:上下文
func X窗口最大化切换(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.WindowToggleMaximise()
}

// WindowUnmaximise 将窗口取消最大化

// ff:窗口取消最大化
// ctx:上下文
func X窗口取消最大化(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.WindowUnmaximise()
}

// WindowMinimise 窗口最小化

// ff:窗口最小化
// ctx:上下文
func X窗口最小化(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.WindowMinimise()
}

// WindowUnminimise 将窗口从最小化状态恢复

// ff:窗口取消最小化
// ctx:上下文
func X窗口取消最小化(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.WindowUnminimise()
}

// WindowIsFullscreen 获取窗口状态，判断窗口是否全屏

// ff:窗口是否全屏
// ctx:上下文
func X窗口是否全屏(上下文 context.Context) bool {
	appFrontend := getFrontend(上下文)
	return appFrontend.WindowIsFullscreen()
}

// WindowIsMaximised 获取窗口状态是否为最大化

// ff:窗口是否最大化
// ctx:上下文
func X窗口是否最大化(上下文 context.Context) bool {
	appFrontend := getFrontend(上下文)
	return appFrontend.WindowIsMaximised()
}

// WindowIsMinimised 获取窗口状态是否为最小化

// ff:窗口是否最小化
// ctx:上下文
func X窗口是否最小化(上下文 context.Context) bool {
	appFrontend := getFrontend(上下文)
	return appFrontend.WindowIsMinimised()
}

// WindowIsNormal 获取窗口状态是否为正常窗口

// ff:窗口是否为正常
// ctx:上下文
func X窗口是否为正常(上下文 context.Context) bool {
	appFrontend := getFrontend(上下文)
	return appFrontend.WindowIsNormal()
}

// WindowExecJS在window环境中执行给定的Js代码

// ff:窗口执行JS
// js:js代码
// ctx:上下文
func X窗口执行JS(上下文 context.Context, js代码 string) {
	appFrontend := getFrontend(上下文)
	appFrontend.ExecJS(js代码)
}


// ff:窗口设置背景色
// A:
// B:
// G:
// R:
// ctx:上下文
func X窗口设置背景色(上下文 context.Context, R, G, B, A uint8) {
	appFrontend := getFrontend(上下文)
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
func X窗口打开打印对话框(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.WindowPrint()
}
