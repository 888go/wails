//go:build windows

package windows

import (
	"github.com/888go/wails/internal/frontend/desktop/windows/win32"
	"github.com/888go/wails/pkg/options/windows"
)

func (w *Window) UpdateTheme() {

	// 如果主题没有变化，则不要重绘
	if !w.themeChanged {
		return
	}
	w.themeChanged = false

	if win32.IsCurrentlyHighContrastMode() {
		return
	}

	if !win32.SupportsThemes() {
		return
	}

	var isDarkMode bool
	switch w.theme {
	case windows.X常量_win主题_默认:
		isDarkMode = win32.IsCurrentlyDarkMode()
	case windows.X常量_win主题_暗黑:
		isDarkMode = true
	case windows.X常量_win主题_浅色:
		isDarkMode = false
	}
	win32.SetTheme(w.Handle(), isDarkMode)

	// 自定义主题处理
	winOptions := w.frontendOptions.Windows选项
	var customTheme *windows.ThemeSettings
	if winOptions != nil {
		customTheme = winOptions.X自定义主题
	}
	// Custom theme
	if win32.SupportsCustomThemes() && customTheme != nil {
		if w.isActive {
			if isDarkMode {
				win32.SetTitleBarColour(w.Handle(), customTheme.DarkModeTitleBar)
				win32.SetTitleTextColour(w.Handle(), customTheme.DarkModeTitleText)
				win32.SetBorderColour(w.Handle(), customTheme.DarkModeBorder)
			} else {
				win32.SetTitleBarColour(w.Handle(), customTheme.LightModeTitleBar)
				win32.SetTitleTextColour(w.Handle(), customTheme.LightModeTitleText)
				win32.SetBorderColour(w.Handle(), customTheme.LightModeBorder)
			}
		} else {
			if isDarkMode {
				win32.SetTitleBarColour(w.Handle(), customTheme.DarkModeTitleBarInactive)
				win32.SetTitleTextColour(w.Handle(), customTheme.DarkModeTitleTextInactive)
				win32.SetBorderColour(w.Handle(), customTheme.DarkModeBorderInactive)
			} else {
				win32.SetTitleBarColour(w.Handle(), customTheme.LightModeTitleBarInactive)
				win32.SetTitleTextColour(w.Handle(), customTheme.LightModeTitleTextInactive)
				win32.SetBorderColour(w.Handle(), customTheme.LightModeBorderInactive)
			}
		}
	}
}
