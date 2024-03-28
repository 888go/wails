package app

import (
	"context"

	"github.com/wailsapp/wails/v2/internal/frontend"
	"github.com/wailsapp/wails/v2/internal/logger"
	"github.com/wailsapp/wails/v2/internal/menumanager"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/options"
)

// App 定义了一个Wails应用程序结构
type App struct {
	frontend frontend.Frontend
	logger   *logger.Logger
	options  *options.App

	menuManager *menumanager.Manager

	// 表示应用程序是否处于调试模式
	debug bool

	// 表示是否启用了开发者工具
	devtoolsEnabled bool

	// OnStartup/OnShutdown
	startupCallback  func(ctx context.Context)
	shutdownCallback func(ctx context.Context)
	ctx              context.Context
}

// 关闭应用程序
func (a *App) Shutdown() {
	if a.frontend != nil {
		a.frontend.Quit()
	}
}

// 设置应用菜单 将设置应用程序的菜单
func (a *App) SetApplicationMenu(menu *menu.Menu) {
	if a.frontend != nil {
		a.frontend.MenuSetApplicationMenu(menu)
	}
}
