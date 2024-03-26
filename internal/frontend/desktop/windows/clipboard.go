//go:build windows
// +build windows

package windows

import (
	"github.com/888go/wails/internal/frontend/desktop/windows/win32"
)

func (f *Frontend) ClipboardGetText() (string, error) {
	return win32.GetClipboardText()
}

func (f *Frontend) ClipboardSetText(text string) error {
	return win32.SetClipboardText(text)
}
