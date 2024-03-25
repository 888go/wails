//go:build !dev && !production && !bindings && (linux || darwin)

package app

import (
	"fmt"

	"github.com/888go/wails/pkg/options"
)

// ff:运行
func (a *App) Run() error {
	return nil
}

// CreateApp 创建应用！

// ff:
// _:
func CreateApp(_ *options.App) (*App, error) {
	return nil, fmt.Errorf(`Wails applications will not build without the correct build tags.`)
}
