//go:build windows && wv2runtime.browser
// +build windows,wv2runtime.browser

package wv2installer

import (
	"fmt"
	"github.com/888go/wails/internal/webview2runtime"
	"github.com/888go/wails/pkg/options/windows"
)

func doInstallationStrategy(installStatus installationStatus, messages *windows.Messages) error {
	confirmed, err := webview2runtime.Confirm(messages.X跳转WebView2下载页面+MinimumRuntimeVersion, messages.X缺少必要组件)
	if err != nil {
		return err
	}
	if confirmed {
		err = webview2runtime.OpenInstallerDownloadWebpage()
		if err != nil {
			return err
		}
	}

	return fmt.Errorf(messages.X安装失败)
}
