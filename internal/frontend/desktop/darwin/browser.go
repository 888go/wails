//go:build darwin
// +build darwin

package darwin

import (
	"github.com/pkg/browser"
)

// BrowserOpenURL 使用默认浏览器打开指定的url
func (f *Frontend) BrowserOpenURL(url string) {
	// 特定方法实现
	if err := browser.OpenURL(url); err != nil {
		f.logger.Error("Unable to open default system browser")
	}
}
