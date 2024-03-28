package runtime

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/options"
)

// WindowSetTitle 设置窗口的标题
func WindowSetTitle(ctx context.Context, title string) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowSetTitle(title)
}

// WindowFullscreen 将窗口设置为全屏模式
func WindowFullscreen(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowFullscreen()
}

// WindowUnfullscreen 将窗口设置为未全屏状态
func WindowUnfullscreen(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowUnfullscreen()
}

// WindowCenter 将窗口居中于当前屏幕
func WindowCenter(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowCenter()
}

// WindowReload 将重新加载窗口内容
func WindowReload(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowReload()
}

// WindowReloadApp 将重新加载应用程序
func WindowReloadApp(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowReloadApp()
}

func WindowSetSystemDefaultTheme(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowSetSystemDefaultTheme()
}

func WindowSetLightTheme(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowSetLightTheme()
}

func WindowSetDarkTheme(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowSetDarkTheme()
}

// WindowShow 如果窗口被隐藏，则显示窗口
func WindowShow(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowShow()
}

// WindowHide the window
func WindowHide(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowHide()
}

// WindowSetSize 设置窗口的大小
func WindowSetSize(ctx context.Context, width int, height int) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowSetSize(width, height)
}

func WindowGetSize(ctx context.Context) (int, int) {
	appFrontend := getFrontend(ctx)
	return appFrontend.WindowGetSize()
}

// WindowSetMinSize 设置窗口的最小尺寸
func WindowSetMinSize(ctx context.Context, width int, height int) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowSetMinSize(width, height)
}

// WindowSetMaxSize 设置窗口的最大尺寸
func WindowSetMaxSize(ctx context.Context, width int, height int) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowSetMaxSize(width, height)
}

// WindowSetAlwaysOnTop 设置窗口是否总在顶部展示
func WindowSetAlwaysOnTop(ctx context.Context, b bool) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowSetAlwaysOnTop(b)
}

// WindowSetPosition 设置窗口的位置
func WindowSetPosition(ctx context.Context, x int, y int) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowSetPosition(x, y)
}

func WindowGetPosition(ctx context.Context) (int, int) {
	appFrontend := getFrontend(ctx)
	return appFrontend.WindowGetPosition()
}

// WindowMaximise 窗口最大化
func WindowMaximise(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowMaximise()
}

// WindowToggleMaximise：最大化窗口
func WindowToggleMaximise(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowToggleMaximise()
}

// WindowUnmaximise 将窗口取消最大化
func WindowUnmaximise(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowUnmaximise()
}

// WindowMinimise 窗口最小化
func WindowMinimise(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowMinimise()
}

// WindowUnminimise 将窗口从最小化状态恢复
func WindowUnminimise(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowUnminimise()
}

// WindowIsFullscreen 获取窗口状态，判断窗口是否全屏
func WindowIsFullscreen(ctx context.Context) bool {
	appFrontend := getFrontend(ctx)
	return appFrontend.WindowIsFullscreen()
}

// WindowIsMaximised 获取窗口状态是否为最大化
func WindowIsMaximised(ctx context.Context) bool {
	appFrontend := getFrontend(ctx)
	return appFrontend.WindowIsMaximised()
}

// WindowIsMinimised 获取窗口状态是否为最小化
func WindowIsMinimised(ctx context.Context) bool {
	appFrontend := getFrontend(ctx)
	return appFrontend.WindowIsMinimised()
}

// WindowIsNormal 获取窗口状态是否为正常窗口
func WindowIsNormal(ctx context.Context) bool {
	appFrontend := getFrontend(ctx)
	return appFrontend.WindowIsNormal()
}

// WindowExecJS在window环境中执行给定的Js代码
func WindowExecJS(ctx context.Context, js string) {
	appFrontend := getFrontend(ctx)
	appFrontend.ExecJS(js)
}

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

func WindowPrint(ctx context.Context) {
	appFrontend := getFrontend(ctx)
	appFrontend.WindowPrint()
}
