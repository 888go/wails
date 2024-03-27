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

// startup 在应用程序启动时被调用。将上下文保存下来，
// 以便我们可以调用运行时方法
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet 函数为给定的名字返回一个问候语

// ff:
// name:
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
