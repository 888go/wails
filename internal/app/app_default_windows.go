//go:build !dev && !production && !bindings && windows

package app

import (
	"os/exec"

	"github.com/888go/wails/internal/frontend/desktop/windows/winc/w32"
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
	result := w32.MessageBox(0,
		`没有正确的构建标记，Wails应用程序将无法构建。
请使用“wails build”或按“OK”打开文档，了解如何使用“go build”"`,
		"Error",
		w32.MB_ICONERROR|w32.MB_OKCANCEL)
	if result == 1 {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://wails.io/docs/guides/manual-builds").Start()
	}
	return nil, nil
}
