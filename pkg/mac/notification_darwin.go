// Package mac 提供了针对 Wails 应用程序的 MacOS 相关实用功能函数
package mac

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/wailsapp/wails/v2/internal/shell"
)

// StartAtLogin 根据给定的布尔标志，将此应用程序添加到登录项或从中移除。限制条件是当前运行的应用程序必须位于应用包中。

// ff:
// sound:
// message:
// subtitle:
// title:
func ShowNotification(title string, subtitle string, message string, sound string) error {
	command := fmt.Sprintf("display notification \"%s\"", message)
	if len(title) > 0 {
		command += fmt.Sprintf(" with title \"%s\"", title)
	}
	if len(subtitle) > 0 {
		command += fmt.Sprintf(" subtitle \"%s\"", subtitle)
	}
	if len(sound) > 0 {
		command += fmt.Sprintf(" sound name \"%s\"", sound)
	}
	_, stde, err := shell.RunCommand("/tmp", "osascript", "-e", command)
	if err != nil {
		return errors.Wrap(err, stde)
	}
	return nil
}
