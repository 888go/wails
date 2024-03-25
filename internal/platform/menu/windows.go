//go:build windows

package menu

import "github.com/888go/wails/internal/platform/win32"

type Menu struct {
	menu win32.HMENU
}
