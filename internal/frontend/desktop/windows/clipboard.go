//go:build windows
// +build windows

package windows

import (
	"github.com/888go/wails/internal/frontend/desktop/windows/win32"
)


// ff:
func (f *Frontend) ClipboardGetText() (string, error) {
	return win32.GetClipboardText()
}


// ff:
// text:
func (f *Frontend) ClipboardSetText(text string) error {
	return win32.SetClipboardText(text)
}
