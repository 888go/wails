package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/888go/wails/pkg/logger"
	"github.com/888go/wails/pkg/options"
	"github.com/888go/wails/pkg/options/mac"
	"github.com/888go/wails/pkg/options/windows"
)

//go:embed frontend/src
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

// 不依赖Wails工具最简单的方式运行, 参考 https://wails.io/zh-Hans/docs/guides/manual-builds
// 对于 dev 构建，最少的命令是：go build -tags dev -gcflags "all=-N -l"
// 对于生产构建，最少的命令是：go build -tags desktop,production -ldflags "-w -s -H windowsgui"

// goland设置方式: 左侧小箭头-->修改运行配置...-->GO工具实参, 设置值: -tags dev -gcflags "all=-N -l"
func main() {
	// 创建一个app结构体的实例
	app := NewApp()

	// 使用选项创建应用程序
	err := wails.X运行(&options.App{
		X标题:             "pure-javascript",
		X宽度:             1024,
		X高度:            768,
		X最小宽度:          1024,
		X最小高度:         768,
		X禁用调整大小:     false,
		X全屏:        false,
		X无边框:         false,
		X启动时隐藏窗口:       false,
		X关闭时隐藏窗口: false,
		X背景颜色:  &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		Assets弃用:            assets,
		X菜单:              nil,
		X日志记录器:            nil,
		X日志级别:          logger.X常量_日志级别_调试,
		X绑定启动前函数:         app.startup,
		X绑定DOM就绪函数:        app.domReady,
		X绑定应用关闭前函数:     app.beforeClose,
		X绑定应用退出函数:        app.shutdown,
		X窗口启动状态:  options.X常量_正常,
		X绑定调用方法: []interface{}{
			app,
		},
		// Windows平台特定的选项
		Windows选项: &windows.Options{
			X开启Webview透明: false,
			X开启窗口半透明:  false,
			X禁用窗口图标:    false,
			// 禁用无边框窗口装饰: false,
			Webview用户数据路径: "",
		},
		// Mac平台特定的选项
		Mac选项: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "pure-javascript",
				Message: "",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
