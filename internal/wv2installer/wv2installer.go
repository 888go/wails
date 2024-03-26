//go:build windows

package wv2installer

import (
	"fmt"

	"github.com/wailsapp/go-webview2/webviewloader"
	"github.com/888go/wails/pkg/options"
	"github.com/888go/wails/pkg/options/windows"
)

const MinimumRuntimeVersion string = "94.0.992.31" // WebView2 SDK 1.0.992.28
// （注：此代码无具体功能实现，仅为版本号注释）
// 此处注释表明使用的是WebView2 SDK（软件开发工具包）的1.0.992.28版本

type installationStatus int

const (
	needsInstalling installationStatus = iota
	needsUpdating
)

func Process(appoptions *options.App) (string, error) {
	messages := windows.X运行时默认提示()
	if appoptions.Windows选项 != nil && appoptions.Windows选项.X用户消息 != nil {
		messages = appoptions.Windows选项.X用户消息
	}

	installStatus := needsInstalling

	// 如果存在手动指定的webview路径，则覆盖版本检查
	var webviewPath = ""
	if opts := appoptions.Windows选项; opts != nil && opts.Webview浏览器路径 != "" {
		webviewPath = opts.Webview浏览器路径
	}

	installedVersion, err := webviewloader.GetAvailableCoreWebView2BrowserVersionString(webviewPath)
	if err != nil {
		return "", err
	}

	if installedVersion != "" {
		installStatus = needsUpdating
		compareResult, err := webviewloader.CompareBrowserVersions(installedVersion, MinimumRuntimeVersion)
		if err != nil {
			return "", err
		}
		updateRequired := compareResult < 0
		// 已安装且无需更新
		if !updateRequired {
			return installedVersion, nil
		}
	}

	// 如果手动指定了webview，则强制采用错误处理策略
	if webviewPath != "" {
		return installedVersion, fmt.Errorf(messages.InvalidFixedWebview2)
	}

	return installedVersion, doInstallationStrategy(installStatus, messages)
}
