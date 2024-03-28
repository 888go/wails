//go:build windows
// +build windows

package windows

import (
	"github.com/pkg/browser"
	"golang.org/x/sys/windows"
)

var fallbackBrowserPaths = []string{
	`\Program Files (x86)\Microsoft\Edge\Application\msedge.exe`,
	`\Program Files\Google\Chrome\Application\chrome.exe`,
	`\Program Files (x86)\Google\Chrome\Application\chrome.exe`,
	`\Program Files\Mozilla Firefox\firefox.exe`,
}

// BrowserOpenURL 使用默认浏览器打开指定的url

// ff:
// url:
func (f *Frontend) BrowserOpenURL(url string) {
	// 特定方法实现
	err := browser.OpenURL(url)
	if err == nil {
		return
	}
	for _, fallback := range fallbackBrowserPaths {
		if err := openBrowser(fallback, url); err == nil {
			return
		}
	}
	f.logger.Error("Unable to open default system browser")
}

func openBrowser(path, url string) error {
	return windows.ShellExecute(0, nil, windows.StringToUTF16Ptr(path), windows.StringToUTF16Ptr(url), nil, windows.SW_SHOWNORMAL)
}
