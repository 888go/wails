//go:build windows
// +build windows

package windows

import (
	"github.com/pkg/browser"
)

// BrowserOpenURL 使用默认浏览器打开指定的url

// ff:
// url:
func (f *Frontend) BrowserOpenURL(url string) {
	// 特定方法实现
	_ = browser.OpenURL(url)
}
