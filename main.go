package main

import (
	"embed"
	"log"

	"github.com/888go/wails/pkg/logger"
	"github.com/888go/wails/pkg/options"
	"github.com/888go/wails/pkg/options/windows"
)

//go:embed frontend/src
var assets embed.FS

var icon []byte

// 工具实参  go build -tags dev -gcflags "all=-N -l"

// 对于 dev 构建，最少的命令是：go build -tags dev -gcflags "all=-N -l"
// 对于生产构建，最少的命令是：go build -tags desktop,production -ldflags "-w -s -H windowsgui"
func main() {
	// 创建一个app结构体的实例
	app := NewApp()

	// 使用选项创建应用程序
	err := Run(&options.App{
		Title:             "demo2",
		Width:             1024,
		Height:            768,
		MinWidth:          1024,
		MinHeight:         768,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		Assets:            assets,
		Menu:              nil,
		Logger:            nil,
		LogLevel:          logger.DEBUG,
		OnStartup:         app.startup,
		OnDomReady:        app.domReady,
		OnBeforeClose:     app.beforeClose,
		OnShutdown:        app.shutdown,
		WindowStartState:  options.Normal,
		Bind: []interface{}{
			app,
		},
		// Windows平台特定的选项
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
			// 禁用无边框窗口装饰: false,
			WebviewUserDataPath: "",
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
