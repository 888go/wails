//go:build production

package app

import (
	"context"

	"github.com/888go/wails/internal/binding"
	"github.com/888go/wails/internal/frontend/desktop"
	"github.com/888go/wails/internal/frontend/dispatcher"
	"github.com/888go/wails/internal/frontend/runtime"
	"github.com/888go/wails/internal/logger"
	"github.com/888go/wails/internal/menumanager"
	"github.com/888go/wails/pkg/options"
)


// ff:
func (a *App) Run() error {
	err := a.frontend.Run(a.ctx)
	a.frontend.RunMainLoop()
	a.frontend.WindowClose()
	if a.shutdownCallback != nil {
		a.shutdownCallback(a.ctx)
	}
	return err
}

// CreateApp 创建应用！

// ff:
// appoptions:
func CreateApp(appoptions *options.App) (*App, error) {
	var err error

	ctx := context.Background()

	// Merge default options
	options.MergeDefaults(appoptions)

	debug := IsDebug()
	devtoolsEnabled := IsDevtoolsEnabled()
	ctx = context.WithValue(ctx, "debug", debug)
	ctx = context.WithValue(ctx, "devtoolsEnabled", devtoolsEnabled)

	// Set up logger
	myLogger := logger.New(appoptions.X日志记录器)
	if IsDebug() {
		myLogger.SetLogLevel(appoptions.X日志级别)
	} else {
		myLogger.SetLogLevel(appoptions.LogLevelProduction)
	}
	ctx = context.WithValue(ctx, "logger", myLogger)
	ctx = context.WithValue(ctx, "obfuscated", IsObfuscated())

	// Preflight Checks
	err = PreflightChecks(appoptions, myLogger)
	if err != nil {
		return nil, err
	}

	// 创建菜单管理器
	menuManager := menumanager.NewManager()

	// 处理应用程序菜单
	if appoptions.X菜单 != nil {
		err = menuManager.SetApplicationMenu(appoptions.X菜单)
		if err != nil {
			return nil, err
		}
	}

	// 创建绑定豁免 - 丑陋的解决方案。肯定有更优的方法
	bindingExemptions := []interface{}{
		appoptions.X启动前回调函数,
		appoptions.X应用退出回调函数,
		appoptions.DOM就绪回调函数,
		appoptions.X应用关闭前回调函数,
	}
	appBindings := binding.NewBindings(myLogger, appoptions.Bind, bindingExemptions, IsObfuscated(), appoptions.EnumBind)
	eventHandler := runtime.NewEvents(myLogger)
	ctx = context.WithValue(ctx, "events", eventHandler)
	// 将日志器附加到上下文中
	if debug {
		ctx = context.WithValue(ctx, "buildtype", "debug")
	} else {
		ctx = context.WithValue(ctx, "buildtype", "production")
	}

	messageDispatcher := dispatcher.NewDispatcher(ctx, myLogger, appBindings, eventHandler, appoptions.ErrorFormatter)
	appFrontend := desktop.NewFrontend(ctx, appoptions, myLogger, appBindings, messageDispatcher)
	eventHandler.AddFrontend(appFrontend)

	ctx = context.WithValue(ctx, "frontend", appFrontend)
	result := &App{
		ctx:              ctx,
		frontend:         appFrontend,
		logger:           myLogger,
		menuManager:      menuManager,
		startupCallback:  appoptions.X启动前回调函数,
		shutdownCallback: appoptions.X应用退出回调函数,
		debug:            debug,
		devtoolsEnabled:  devtoolsEnabled,
		options:          appoptions,
	}

	return result, nil

}
