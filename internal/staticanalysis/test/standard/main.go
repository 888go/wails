package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/888go/wails/pkg/options"
	"github.com/888go/wails/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// 创建一个app结构体的实例
	app := NewApp()

	// 使用选项创建应用程序
	err := wails.Run(&options.App{
		X标题:  "staticanalysis",
		X宽度:  1024,
		X高度: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		X背景颜色: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		X启动前回调函数:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
