//go:build windows
// +build windows

package windows

import (
	"github.com/888go/wails/internal/frontend/desktop/windows/win32"
)

// ff:剪贴板取文本
func (f *Frontend) ClipboardGetText() (string, error) {
	return win32.GetClipboardText()
}

// ff:剪贴板置文本
// text:文本
func (f *Frontend) ClipboardSetText(text string) error {
	return win32.SetClipboardText(text)
}
