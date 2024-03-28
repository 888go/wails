//go:build windows && wv2runtime.error
// +build windows,wv2runtime.error

package wv2installer

import (
	"fmt"
	"github.com/888go/wails/internal/webview2runtime"
	"github.com/888go/wails/pkg/options/windows"
)

func doInstallationStrategy(installStatus installationStatus, messages *windows.Messages) error {
	_ = webview2runtime.Error(messages.X联系管理员, messages.X出错)
	return fmt.Errorf(messages.WebView2未安装)
}
