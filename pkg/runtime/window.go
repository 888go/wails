package runtime

import (
	"context"

	"github.com/888go/wails/pkg/options"
)

// WindowSetTitle 设置窗口的标题
func X窗口设置标题(上下文 context.Context, 标题 string) {
	appFrontend := getFrontend(上下文)
	appFrontend.X窗口设置标题(标题)
}

// WindowFullscreen 将窗口设置为全屏模式
func X窗口设置全屏(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.X窗口设置全屏()
}

// WindowUnfullscreen 将窗口设置为未全屏状态
func X窗口取消全屏(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.X窗口取消全屏()
}

// WindowCenter 将窗口居中于当前屏幕
func X窗口居中(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.X窗口居中()
}

// WindowReload 将重新加载窗口内容
func X窗口重载(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.X窗口重载()
}

// WindowReloadApp 将重新加载应用程序
func X窗口重载应用程序前端(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.X窗口重载应用程序前端()
}

func X窗口设置系统默认主题(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.X窗口设置系统默认主题()
}

func X窗口设置浅色主题(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.X窗口设置浅色主题()
}

func X窗口设置深色主题(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.X窗口设置深色主题()
}

// WindowShow 如果窗口被隐藏，则显示窗口
func X窗口显示(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.X窗口显示()
}

// WindowHide the window
func X窗口隐藏(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.X窗口隐藏()
}

// WindowSetSize 设置窗口的大小
func X窗口设置尺寸(上下文 context.Context, 宽 int, 高 int) {
	appFrontend := getFrontend(上下文)
	appFrontend.X窗口设置尺寸(宽, 高)
}

func X窗口取尺寸(上下文 context.Context) (int, int) {
	appFrontend := getFrontend(上下文)
	return appFrontend.X窗口取尺寸()
}

// WindowSetMinSize 设置窗口的最小尺寸
func X窗口设置最小尺寸(上下文 context.Context, 宽 int, 高 int) {
	appFrontend := getFrontend(上下文)
	appFrontend.X窗口设置最小尺寸(宽, 高)
}

// WindowSetMaxSize 设置窗口的最大尺寸
func X窗口设置最大尺寸(上下文 context.Context, 宽 int, 高 int) {
	appFrontend := getFrontend(上下文)
	appFrontend.X窗口设置最大尺寸(宽, 高)
}

// WindowSetAlwaysOnTop 设置窗口是否总在顶部展示
func X窗口设置置顶(上下文 context.Context, 置顶 bool) {
	appFrontend := getFrontend(上下文)
	appFrontend.X窗口设置置顶(置顶)
}

// WindowSetPosition 设置窗口的位置

// y:
// x:
func X窗口设置位置(上下文 context.Context, x int, y int) {
	appFrontend := getFrontend(上下文)
	appFrontend.X窗口设置位置(x, y)
}

func X窗口取位置(上下文 context.Context) (int, int) {
	appFrontend := getFrontend(上下文)
	return appFrontend.X窗口取位置()
}

// WindowMaximise 窗口最大化
func X窗口最大化(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.X窗口最大化()
}

// WindowToggleMaximise：最大化窗口
func X窗口最大化切换(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.X窗口最大化切换()
}

// WindowUnmaximise 将窗口取消最大化
func X窗口取消最大化(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.X窗口取消最大化()
}

// WindowMinimise 窗口最小化
func X窗口最小化(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.X窗口最小化()
}

// WindowUnminimise 将窗口从最小化状态恢复
func X窗口取消最小化(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.X窗口取消最小化()
}

// WindowIsFullscreen 获取窗口状态，判断窗口是否全屏
func X窗口是否全屏(上下文 context.Context) bool {
	appFrontend := getFrontend(上下文)
	return appFrontend.X窗口是否全屏()
}

// WindowIsMaximised 获取窗口状态是否为最大化
func X窗口是否最大化(上下文 context.Context) bool {
	appFrontend := getFrontend(上下文)
	return appFrontend.X窗口是否最大化()
}

// WindowIsMinimised 获取窗口状态是否为最小化
func X窗口是否最小化(上下文 context.Context) bool {
	appFrontend := getFrontend(上下文)
	return appFrontend.X窗口是否最小化()
}

// WindowIsNormal 获取窗口状态是否为正常窗口
func X窗口是否为正常(上下文 context.Context) bool {
	appFrontend := getFrontend(上下文)
	return appFrontend.X窗口是否为正常()
}

// WindowExecJS在window环境中执行给定的Js代码
func X窗口执行JS(上下文 context.Context, js代码 string) {
	appFrontend := getFrontend(上下文)
	appFrontend.X窗口执行JS(js代码)
}


// A:
// B:
// G:
// R:
func X窗口设置背景色(上下文 context.Context, R, G, B, A uint8) {
	appFrontend := getFrontend(上下文)
	col := &options.RGBA{
		R: R,
		G: G,
		B: B,
		A: A,
	}
	appFrontend.X窗口设置背景色(col)
}

func X窗口打开打印对话框(上下文 context.Context) {
	appFrontend := getFrontend(上下文)
	appFrontend.X窗口打开打印对话框()
}
