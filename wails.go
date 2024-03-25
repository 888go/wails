// Package wails 是 Wails 项目的主包。
// 它被客户端应用程序使用。
package main

import (
	_ "github.com/888go/wails/internal/goversion" // 添加编译时版本检查，确保最低Go版本
	"github.com/888go/wails/pkg/application"
	"github.com/888go/wails/pkg/options"
)

// Run 根据给定的配置创建应用程序并执行它
func X运行(配置 *options.App) error {
	mainApp := application.X创建并按选项(配置)
	return mainApp.X运行()
}
