//go:build windows && !wv2runtime.error && !wv2runtime.browser && !wv2runtime.embed
// +build windows,!wv2runtime.error,!wv2runtime.browser,!wv2runtime.embed

package wv2installer

import (
	"fmt"

	"github.com/888go/wails/internal/webview2runtime"
	"github.com/888go/wails/pkg/options/windows"
)

func doInstallationStrategy(installStatus installationStatus, messages *windows.Messages) error {
	message := messages.WebView2需安装
	if installStatus == needsUpdating {
		message = messages.WebView2需更新
	}
	confirmed, err := webview2runtime.Confirm(message, messages.X缺少必要组件)
	if err != nil {
		return err
	}
	if !confirmed {
		return fmt.Errorf(messages.WebView2未安装)
	}
	installedCorrectly, err := webview2runtime.InstallUsingBootstrapper()
	if err != nil {
		_ = webview2runtime.Error(err.Error(), messages.X出错)
		return err
	}
	if !installedCorrectly {
		err = webview2runtime.Error(messages.X安装失败, messages.X出错)
		return err
	}
	return nil
}
