//go:build windows && !bindings

package app

import (
	"github.com/wailsapp/wails/v2/internal/logger"
	"github.com/wailsapp/wails/v2/internal/wv2installer"
	"github.com/wailsapp/wails/v2/pkg/options"
)

func PreflightChecks(options *options.App, logger *logger.Logger) error {

	_ = options

// 处理 webview2 运行时的情况。我们可以通过在 `wails build` 命令中的 `webview2` 标志传递一种策略。
// 这将决定 wv2runtime.Process 在缺乏有效运行时的情况下如何处理。
	installedVersion, err := wv2installer.Process(options)
	if installedVersion != "" {
		logger.Debug("WebView2 Runtime Version '%s' installed. Minimum version required: %s.",
			installedVersion, wv2installer.MinimumRuntimeVersion)
	}
	if err != nil {
		return err
	}

	return nil
}
