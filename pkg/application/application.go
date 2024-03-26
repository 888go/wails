package application

import (
	"context"
	"sync"

	"github.com/888go/wails/internal/app"
	"github.com/888go/wails/internal/signal"
	"github.com/888go/wails/pkg/menu"
	"github.com/888go/wails/pkg/options"
)

// Application 是 Wails 主应用程序
type Application struct {
	application *app.App
	options     *options.App

	// running flag
	running bool

	shutdown sync.Once
}

// NewWithOptions 使用给定的选项创建一个新的Application
func X创建并按选项(App选项 *options.App) *Application {
	if App选项 == nil {
		return X创建并按默认选项()
	}
	return &Application{
		options: App选项,
	}
}

// New 创建一个使用默认选项的新 Application
func X创建并按默认选项() *Application {
	return &Application{
		options: &options.App{},
	}
}

// 设置应用菜单 将设置应用程序的菜单
func (a *Application) X设置菜单(菜单 *menu.Menu) {
	if a.running {
		a.application.SetApplicationMenu(菜单)
		return
	}

	a.options.X菜单 = 菜单
}

// Run 启动应用程序
func (a *Application) X运行() error {
	err := applicationInit()
	if err != nil {
		return err
	}

	application, err := app.CreateApp(a.options)
	if err != nil {
		return err
	}

	a.application = application

	// Control-C handlers
	signal.OnShutdown(func() {
		a.application.Shutdown()
	})
	signal.Start()

	a.running = true

	err = a.application.Run()
	return err
}

// Quit 将关闭应用程序
func (a *Application) X退出() {
	a.shutdown.Do(func() {
		a.application.Shutdown()
	})
}

// 将给定的结构体绑定到应用程序
func (a *Application) Bind(boundStruct any) {
	a.options.Bind = append(a.options.Bind, boundStruct)
}

func (a *Application) On(eventType EventType, callback func()) {
	c := func(ctx context.Context) {
		callback()
	}

	switch eventType {
	case StartUp:
		a.options.X启动前回调函数 = c
	case ShutDown:
		a.options.X应用退出回调函数 = c
	case DomReady:
		a.options.DOM就绪回调函数 = c
	}
}
