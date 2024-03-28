package main

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp 创建一个新的 App 应用程序结构体

// ff:
func NewApp() *App {
	return &App{}
}

// startup 在应用程序启动时被调用
func (a *App) startup(ctx context.Context) {
	// 在此处执行你的初始化设置
	a.ctx = ctx
}

// domReady 在前端资源加载完成后被调用
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose 在应用程序即将退出时调用，
// 这种情况可能是通过点击窗口关闭按钮或调用 runtime.Quit 函数。
// 返回 true 将使应用程序继续运行，返回 false 则会按照正常流程继续关闭操作。
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown 在应用程序终止时被调用
func (a *App) shutdown(ctx context.Context) {
	// 在这里执行你的清理工作
}

// Greet 函数为给定的名字返回一个问候语

// ff:
// name:
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
