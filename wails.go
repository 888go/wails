// Package wails 是 Wails 项目的主包。
// 它被客户端应用程序使用。
package wails

import (
	_ "github.com/wailsapp/wails/v2/internal/goversion" // 添加编译时版本检查，确保最低Go版本
	"github.com/wailsapp/wails/v2/pkg/application"
	"github.com/wailsapp/wails/v2/pkg/options"
)

// Run 根据给定的配置创建应用程序并执行它
func Run(options *options.App) error {
	mainApp := application.NewWithOptions(options)
	return mainApp.Run()
}
