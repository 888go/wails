// Package mac 提供了针对 Wails 应用程序的 MacOS 相关实用功能函数
package mac

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/888go/wails/internal/shell"
	"github.com/leaanthony/slicer"
	"github.com/pkg/errors"
)

// StartAtLogin 根据给定的布尔标志，将此应用程序添加到登录项或从中移除。限制条件是当前运行的应用程序必须位于应用包中。
func StartAtLogin(enabled bool) error {
	exe, err := os.Executable()
	if err != nil {
		return errors.Wrap(err, "Error running os.Executable:")
	}
	binName := filepath.Base(exe)
	if !strings.HasSuffix(exe, "/Contents/MacOS/"+binName) {
		return fmt.Errorf("app needs to be running as package.app file to start at login")
	}
	appPath := strings.TrimSuffix(exe, "/Contents/MacOS/"+binName)
	var command string
	if enabled {
		command = fmt.Sprintf("tell application \"System Events\" to make login item at end with properties {name: \"%s\",path:\"%s\", hidden:false}", binName, appPath)
	} else {
		command = fmt.Sprintf("tell application \"System Events\" to delete login item \"%s\"", binName)
	}
	_, stde, err := shell.RunCommand("/tmp", "osascript", "-e", command)
	if err != nil {
		return errors.Wrap(err, stde)
	}
	return nil
}

// StartsAtLogin 表示此应用程序是否已添加至登录项。
// 限制条件是，当前运行的应用程序必须位于应用包内。
func StartsAtLogin() (bool, error) {
	exe, err := os.Executable()
	if err != nil {
		return false, err
	}
	binName := filepath.Base(exe)
	if !strings.HasSuffix(exe, "/Contents/MacOS/"+binName) {
		return false, fmt.Errorf("app needs to be running as package.app file to start at login")
	}
	results, stde, err := shell.RunCommand("/tmp", "osascript", "-e", `tell application "System Events" to get the name of every login item`)
	if err != nil {
		return false, errors.Wrap(err, stde)
	}
	results = strings.TrimSpace(results)
	startupApps := slicer.String(strings.Split(results, ", "))
	return startupApps.Contains(binName), nil
}
