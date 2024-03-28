package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/888go/wails/pkg/options"
	"github.com/888go/wails/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

// 不依赖Wails工具最简单的方式运行, 参考 https://wails.io/zh-Hans/docs/guides/manual-builds
// 对于 dev 构建，最少的命令是：go build -tags dev -gcflags "all=-N -l"
// 对于生产构建，最少的命令是：go build -tags desktop,production -ldflags "-w -s -H windowsgui"

// goland设置方式: 左侧小箭头-->修改运行配置...-->GO工具实参, 设置值: -tags dev -gcflags "all=-N -l"
func main() {
	// 创建一个app结构体的实例
	app := NewApp()

	// 使用选项创建应用程序
	err := wails.X运行(&options.App{
		X标题:  "svelte",
		X宽度:  1024,
		X高度: 768,
		X绑定http请求: &assetserver.Options{
			X静态资源: assets,
		},
		X背景颜色: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		X绑定启动前函数:        app.startup,
		X绑定调用方法: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
