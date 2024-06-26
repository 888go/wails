//go:build linux
// +build linux

package linux

import "github.com/pkg/browser"

// BrowserOpenURL 使用默认浏览器打开指定的url

// ff:
// url:
func (f *Frontend) BrowserOpenURL(url string) {
	// 特定方法实现
	if err := browser.OpenURL(url); err != nil {
		f.logger.Error("Unable to open default system browser")
	}
}
